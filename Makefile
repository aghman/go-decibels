GO=go
APP-NAME=go-decibels
.PHONY: run
run: 
	sudo go run main.go

.PHONY: build
build: clean
	-mkdir build
	go build -o ${PWD}/build/${APP-NAME}

.PHONY: clean
clean: 
	-rm -rf build

.PHONY: prebuild
prebuild: 
	go get ./...

