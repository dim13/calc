package main

//go:generate -command yacc goyacc
//go:generate yacc -o calc.go calc.y

import (
	"bufio"
	"io"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	yyDebug = 1

	for {
		os.Stdout.WriteString("\t")
		line, err := in.ReadString('\n')
		if err == io.EOF {
			return
		}
		yyParse(lex(line))
	}
}
