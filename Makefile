PROG := calc
GO := go
GOFILES := $(wildcard *.go)
YYFILES := $(wildcard *.y)
AUTOGEN := $(YYFILES:.y=.go)

$(PROG): $(GOFILES) $(AUTOGEN)
	$(GO) build

%.go: %.y
	$(GO) tool yacc -o $@ $<

clean:
	$(GO) clean
	$(RM) $(AUTOGEN) y.output

install:
	$(GO) install

.PHONY: clean
