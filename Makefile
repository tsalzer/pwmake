# Makefile for something

#include $(GOROOT)/src/Make.$(GOARCH)

TARGETS=mpw
PKGS=main pwdgen pwdgen/symbol pwdgen/rand

GOPATH := $(shell pwd -L)
export GOPATH

all:
	go build -o mpw main

test:
	go test $(PKGS)

clean:
	rm -f $(TARGETS)

main:$(TARGETS)

# build an executable from a go source:
% : %.go
	go build $<

