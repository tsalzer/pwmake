# Makefile for something

#include $(GOROOT)/src/Make.$(GOARCH)

TARGETS=welcome.bin

GOPATH := $(shell pwd -L)
export GOPATH

all:
	go build -o welcome.bin welcome


clean:
	rm -f $(TARGETS)

main:$(TARGETS)

# build an executable from a go source:
% : %.go
	go build $<

