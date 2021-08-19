# Cloud native C

This is a sample project which shows how to integrate C and Go.

It contains:

- Go code which uploads and downloads objects to/from AWS S3
- C code which calls Go code to handle upload/download logic
- Makefile

To build and run the project (assuming gcc and go are installed) execute:

```sh
$ make all
$ ./app app.c my-bucket-name
```

If you want to learn more about this project, c, go, and cgo please checkout out my article: [Cloud native C](https://dev.to/lukaszbudnik/cloud-native-c-kik-temp-slug-3139541).
