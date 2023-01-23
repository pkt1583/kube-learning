#Makefile for go build

# Set the build variables
BINARY=bin/goapp
VERSION=1.0.0
BUILD=`date +%FT%T%z`

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"
GOOS=linux
GOARCH=amd64
all: deps build

# deps: Get the dependencies
deps:
	go get -d -v

# build: Build the binary file
build:
	GOARCH=amd64 GOOS=linux go build ${LDFLAGS} -o ${BINARY} main.go
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY}.exe main.go

# run: Build and run the binary file
run: build
	./${BINARY}

# clean: Clean up the directory
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

#build docker image
docker:
	docker build --platform=linux/amd64 -t myrgmyacr.azurecr.io/netpol-serviceb .

#push docker image
push:
	docker push myrgmyacr.azurecr.io/netpol-serviceb

delete-pod:
	kubectl delete pod -l=app=netpol-serviceb -n netpol-demo

all: deps build docker push delete-pod