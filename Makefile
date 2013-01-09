# Makefile for something

#include $(GOROOT)/src/Make.$(GOARCH)

TARGETS=mpw
TPKGS=pwdgen pwdgen/symbol pwdgen/rand
PKGS=main $(TPKGS)
GOPATH := $(shell pwd -L)
export GOPATH

all:
	go build -o mpw main

test:
	go test $(TPKGS)

bench:
	go test -test.bench 'Benchmark.*' $(TPKGS)

clean:
	rm -f $(TARGETS)

main:$(TARGETS)

# build an executable from a go source:
% : %.go
	go build $<

