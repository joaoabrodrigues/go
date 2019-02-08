package main

import "fmt"

func main() {
	fmt.Printf("Outro programa em %s!\n", "Go")
}

/*
 - go help COMMAND
 - go version
 - go doc -http:6060 -> offline docs
 - go env
 - go doc cmd/vet
 	- go vet comandos.go
 - go build
 - ./comandos
 - go run comandos.go
 - go run *.go
 - go get -u github.com/go-sql-driver/mysql
	- ls $(go env GOPATH)/src/github.com
*/
