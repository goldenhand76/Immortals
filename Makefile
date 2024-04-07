build:
	GOOS=windows GOARCH=386 go build .\cmd\web\main.go

run:
	.\main.exe

all: build run