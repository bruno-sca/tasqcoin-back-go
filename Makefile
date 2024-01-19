migrate:
	go run cmd/migrate/main.go

run:
	go run cmd/api/main.go

dev:
	air -c .air.toml