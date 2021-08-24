# Cloud native C

This is a sample project which shows how to integrate C and Go.

It contains:

- Go code which uploads and downloads objects to/from AWS S3
- C code which calls Go code to handle upload/download logic
- Makefile

Supported platforms are:

- MacOS 10.15
- Ubuntu 20.04
- Windows 2019 with MinGW

See [build.yml](.github/workflows/build.yml)) for details.

To build and run the project (assuming make, gcc and go are installed) simply execute:

```sh
$ make all
$ ./app app.c my-bucket-name
```

If you want to learn more about this project, c, go, and cgo please checkout out my article: [Cloud native C](https://dev.to/lukaszbudnik/cloud-native-c-48m).
