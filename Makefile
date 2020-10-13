PROJECTNAME=$(shell basename "$(PWD)")

GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin

all: build

install:
	go mod download

start:
	go build -o $(GOBIN)/app ./main.go || exit
	$(GOBIN)/app