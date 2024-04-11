build:
	GOOS=windows GOARCH=386 go build .\cmd\web\main.go

build-docker:
	docker build -f .\build\package\Dockerfile -t goldenhand/immortals:0.0.1 .

unit-tests:
	go test ./...

run:
	.\main.exe

all: build run