PROG := calc

$(PROG): $(wildcard *.go) $(PROG).go
	go build

%.go: %.y
	go tool yacc -o $@ $<

clean:
	$(RM) $(PROG) $(PROG).go y.output

.PHONY: clean
