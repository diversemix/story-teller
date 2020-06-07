.PHONY: test build run

test:
	cd lib/markovchain ; go test

build: test
	go mod tidy
	go build app.go

run: test
	go run app.go