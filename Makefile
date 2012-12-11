# Makefile for something

#include $(GOROOT)/src/Make.$(GOARCH)

TARGETS=welcome

all:main

clean:
	rm -f $(TARGETS)

main:$(TARGETS)

# build an executable from a go source:
% : %.go
	go build $<

