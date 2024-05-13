postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root immo

dropdb: 
	docker exec -it postgres dropdb immo

migrateup:
	migrate -path ./internal/database/sqlite/migrations -database "postgresql://leo:Goldenhand76@localhost:5432/immo?sslmode=disable" -verbose up

migratedown:
	migrate -path ./internal/database/sqlite/migrations -database "postgresql://leo:Goldenhand76@localhost:5432/immo?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONEY: postgres createdb dropdb migrateup migratedown sqlc