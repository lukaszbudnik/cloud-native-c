ifeq ($(OS),Windows_NT)
	build_cmd = gcc -pthread -o app app.c storage.a
	rm_cmd = rm -f app.exe
else
	UNAME := $(shell uname -s)
	ifeq ($(UNAME), Linux)
		build_cmd = gcc -pthread -o app app.c storage.a
	endif
	ifeq ($(UNAME), Darwin)
		build_cmd = gcc storage.a app.c -framework CoreFoundation -framework Security -o app
	endif
	rm_cmd = rm -f app
endif

all: app

app: storage app.c
	$(build_cmd)

storage: storage.go
	go build -buildmode=c-archive storage.go

clean:
	go clean
	$(rm_cmd)
