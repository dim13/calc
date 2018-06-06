package main

//go:generate -command yacc goyacc
//go:generate yacc -o calc.go calc.y

import (
	"bufio"
	"fmt"
	"os"
)

const promt = "\t"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(promt)
	for scanner.Scan() {
		line := scanner.Text()
		result, ok := Parse(line)
		if ok {
			fmt.Printf("%v\n\n", result)
		}
		fmt.Print(promt)
	}
}
