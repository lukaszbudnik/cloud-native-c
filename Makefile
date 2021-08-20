OS:= $(shell uname -s)

ifeq ($(OS), Linux)
	build_cmd = gcc -pthread -o app app.c storage.a
endif
ifeq ($(OS), Darwin)
	build_cmd = gcc storage.a app.c -framework CoreFoundation -framework Security -o app
endif

all: app

app: storage app.c
	$(build_cmd)

storage: storage.go
	go build -buildmode=c-archive storage.go

clean:
	go clean
	rm -f app
