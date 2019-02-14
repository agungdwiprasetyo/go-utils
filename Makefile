.PHONY : build test

build:
	go build

test:
	go test -count=1