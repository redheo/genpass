build:
	go vet .
	go build -o genpass main.go
