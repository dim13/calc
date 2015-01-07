package main

//go:generate go tool yacc -o calc.go calc.y

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
		//line, err := in.ReadBytes('\n')
		line, err := in.ReadString('\n')
		if err == io.EOF {
			return
		}
		//yyParse(&yyLex{input: line})
		yyParse(lex(line))
	}
}
