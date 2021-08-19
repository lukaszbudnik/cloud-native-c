package main

// I need to explicitly reference stdio and stdlib
// further I need them one after another, followed by import "C"
// if I move "C" inside the import block
// cgo will not be able to find symbols from stdio/stdlib/response header files

// #include <stdio.h>
// #include <stdlib.h>
// #include "response.h"
import "C"

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"unsafe"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {}

//export PutObject
func PutObject(filenameC, bucketC *C.char) int {
	filename := C.GoString(filenameC)
	bucket := C.GoString(bucketC)

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Fprintf(os.Stderr, "configuration error: %v\n", err.Error())
		return C.CONFIGURATION_ERROR
	}

	client := s3.NewFromConfig(cfg)

	file, err := os.Open(filename)

	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to open file: %v\n", filename)
		return C.UNABLE_TO_OPEN_FILE_ERROR
	}

	defer file.Close()

	input := &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   file,
	}

	output, err := client.PutObject(context.TODO(), input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error uploading file: %v\n", err.Error())
		return C.FILE_UPLOAD_ERROR
	}

	fmt.Printf("File uploaded correctly, version id: %v\n", *output.VersionId)

	return C.SUCCESS
}

//export GetObject
func GetObject(keyC, bucketC *C.char) *C.char {
	key := C.GoString(keyC)
	bucket := C.GoString(bucketC)

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Fprintf(os.Stderr, "configuration error: %v\n", err.Error())
		return nil
	}

	client := s3.NewFromConfig(cfg)

	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	output, err := client.GetObject(context.TODO(), input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error getting file: %v\n", err.Error())
		return nil
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(output.Body)
	content := buf.String()

	fmt.Printf("File version id: %v\n", *output.VersionId)

	return C.CString(content)
}

//export Free
func Free(ptr *C.char) {
	C.free(unsafe.Pointer(ptr))
}
