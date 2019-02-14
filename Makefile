.PHONY : build test cover

build:
	go build

test:
	go test -count=1

cover:
	go test -coverprofile=coverage.txt -covermode=atomic