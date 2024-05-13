## Install migrate;
https://github.com/golang-migrate/migrate

## Create init_schema: 
```bash
migrate create -exc sql -dir .\internal\database\sqlite\migrations\ -seq init_schema 
```

## Migrate init_schema: 
```bash
migrate -path ./internal/database/sqlite/migrations -database "postgresql://leo:Goldenhand76@localhost:5432/immo?sslmode=disable" -verbose up
```

