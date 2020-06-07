.PHONY: test build run get_alice

get_alice:
	wget -O story.txt http://www.gutenberg.org/files/11/11-0.txt

test:
	cd lib/markovchain ; go test

build: test
	go mod tidy
	go build app.go

run: test
	go run app.go