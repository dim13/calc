GOFILES := $(wildcard *.go)
YYFILES := $(wildcard *.y)
AUTOGEN := $(YYFILES:.y=.go)
GO := go

build: $(GOFILES) $(AUTOGEN)
	$(GO) $@

run: $(GOFILES) $(AUTOGEN)
	$(GO) $@ $^

generate:
	$(GO) $@

clean:
	$(GO) $@
	$(RM) $(AUTOGEN) y.output

install: build
	$(GO) $@

%.go: %.y
	$(GO) tool yacc -o $@ $<

.PHONY: clean
