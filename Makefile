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
