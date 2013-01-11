# Makefile for pwmake

#include $(GOROOT)/src/Make.$(GOARCH)

TARGET=pwmake
TPKGS=pwdgen pwdgen/symbol pwdgen/rand pwdgen/columns
PKGS=main $(TPKGS)
GOPATH := $(shell pwd -L)
export GOPATH

all:
	go build -o $(TARGET) main

test:
	go test $(TPKGS)

fix:
	go fix $(TPKGS)

bench:
	go test -test.bench 'Benchmark.*' $(TPKGS)

clean:
	rm -f $(TARGET)

main:$(TARGETS)

# build an executable from a go source:
% : %.go
	go build $<

