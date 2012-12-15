# Makefile for something

#include $(GOROOT)/src/Make.$(GOARCH)

TARGETS=mpw

GOPATH := $(shell pwd -L)
export GOPATH

all:
	go build -o mpw main

test:
	go test main pwdgen

clean:
	rm -f $(TARGETS)

main:$(TARGETS)

# build an executable from a go source:
% : %.go
	go build $<

