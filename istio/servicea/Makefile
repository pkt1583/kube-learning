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

acr-build:
	az acr build --registry myacrmyrg --image servicea .

all: deps build acr-build