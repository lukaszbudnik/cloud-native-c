all: app

app: storage app.c
	gcc storage.a app.c -framework CoreFoundation -framework Security -o app

storage: storage.go
	go build -buildmode=c-archive storage.go

clean:
	go clean
	rm -f app
