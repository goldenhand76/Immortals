build:
	GOOS=windows GOARCH=amd64 go build .\cmd\web\main.go

build-docker:
	docker build -f .\build\package\Dockerfile -t goldenhand/immortals:0.0.1 .

migrateforce: 
	migrate -path ./internal/database/sqlite/migrations -database "postgresql://leo:Goldenhand76@localhost:5432/immo?sslmode=disable" force 1

migrateup:
	migrate -path ./internal/database/sqlite/migrations -database "postgresql://leo:Goldenhand76@localhost:5432/immo?sslmode=disable" -verbose up

migratedown:
	migrate -path ./internal/database/sqlite/migrations -database "postgresql://leo:Goldenhand76@localhost:5432/immo?sslmode=disable" -verbose down

test:
	go test -v -cover ./...
run:
	.\main.exe

all: build run