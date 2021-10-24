SHELL = /bin/bash

tidy:
	go mod tidy
	go mod vendor

run:
	go build app/sales-api/main.go;./main