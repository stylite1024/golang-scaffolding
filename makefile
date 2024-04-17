APP := $(shell cat go.mod | grep "^module " | sed -e "s/module //g")
GOARCH := $(shell go env GOARCH)

define LDFLAGS
endef

GOBUILD=CGO_ENABLED=0 go build -a -ldflags ${LDFLAGS}

# 使用内置目标名.PHONY声明这些“伪目标”名是“伪目标”，而不是与“伪目标”同名的文件
.PHONY: help all build windows linux darwin

default:help

all:build windows linux darwin ls

build:
	@${GOBUILD} -o ./bin/${APP}
windows:
	@GOOS=windows ${GOBUILD} -o ./bin/${APP}-windows-${GOARCH}.exe
linux:
	@GOOS=linux ${GOBUILD} -o ./bin/${APP}-linux-${GOARCH}
darwin:
	@GOOS=darwin ${GOBUILD} -o ./bin/${APP}-darwin-${GOARCH}
ls:
	@ls ./bin

clean:
	@echo "Cleaning up all the generated files"
	@if [ -d bin ] ; then rm -rf ./bin ; fi

help:
	@echo "usage: make <option>"
	@echo "options and effects:"
	@echo "    help   	: Show help"
	@echo "    all    	: Build multiple binary of this project"
	@echo "    build  	: Build the binary of this project for current platform"
	@echo "    windows	: Build the windows binary of this project"
	@echo "    linux  	: Build the linux binary of this project"
	@echo "    darwin 	: Build the darwin binary of this project"
	@echo "    clean  	: Cleaning up all the generated files"
