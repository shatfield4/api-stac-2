# set default shell
SHELL = bash -e -o pipefail

# variables
VERSION                  ?= $(shell cat ./VERSION)

now=$(shell date +"%Y%m%d%H%M%S")

default: run

.PHONY:	install
install:
	go mod tidy
	go mod vendor

.PHONY:	lint
lint:
	golangci-lint run 

.PHONY:	build
build:
	mkdir -p bin
	go build -race -o bin/api-grandtheftbacon \
	    cmd/main.go

.PHONY:	test
test:
	go test -race -v -p 1 ./...
	
.PHONY:	run
run:	build
	./bin/api-grandtheftbacon

.PHONY: up
up:
	docker-compose down
	docker-compose up -d --build

.PHONY: bacon-contract-build
bacon-contract-build:
	rm -rf infrastructures/smartcontracts/Bacon.go
	abigen --abi=infrastructures/smartcontracts/build/Bacon.abi --pkg=bacon --out=infrastructures/smartcontracts/bacon/Bacon.go --alias _contractURI=contractURI1

.PHONY: thefryingpan-contract-build
thefryingpan-contract-build:
	rm -rf infrastructures/smartcontracts/TheFryingPan.go
	abigen --abi=infrastructures/smartcontracts/build/TheFryingPan.abi --pkg=thefryingpan --out=infrastructures/smartcontracts/thefryingpan/TheFryingPan.go
	