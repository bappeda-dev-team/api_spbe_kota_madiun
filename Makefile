GO_VERSION :=1.22

.PHONY: install-go init-go

setup: install-go init-go

# ONLY MAC with brew
install-go:
	echo 'installing go...'

init-go:
	echo 'initialize go in your pc..'

build:
	go build -o api cmd/main.go

tidy:
	go mod tidy

format:
	go fmt ./...

run:
	go run cmd/main.go

install-migrate:
	brew install golang-migrate

db-migrate:
	migrate -path db/migrations -database "mysql://root@tcp(localhost:3306)/db_spbe" up
