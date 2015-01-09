GOFILES := $(wildcard *.go)
YYFILES := $(wildcard *.y)
AUTOGEN := $(YYFILES:.y=.go)
GO := go

build: $(GOFILES) $(AUTOGEN)
	$(GO) build

%.go: %.y
	$(GO) tool yacc -o $@ $<

clean:
	$(GO) clean
	$(RM) $(AUTOGEN) y.output

install: build
	$(GO) install

.PHONY: clean
