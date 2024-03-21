all: build

build: main.go
	go build -o out/omg main.go

run: main.go
	go run main.go

fmt: main.go
	go fmt ./...
