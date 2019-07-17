package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	eof = iota
	number
	word
	char
)

type item struct {
	typ int
	val string
}

type yyLex struct {
	input  string
	start  int
	pos    int
	width  int
	items  chan item
	result float64
	err    error
}

func (y *yyLex) Error(s string) {
	y.err = errors.New(s)
}

func (y *yyLex) Lex(lval *yySymType) int {
	switch item := <-y.items; item.typ {
	case number:
		n, err := strconv.ParseFloat(item.val, 64)
		if err != nil {
			y.Error(err.Error())
		}
		lval.fval = n
		return NUMBER
	case word:
		lval.sval = item.val
		return WORD
	case char:
		return int(item.val[0])
	default:
		return eof
	}
}

func lex(input string) *yyLex {
	l := &yyLex{
		input: input,
		items: make(chan item),
	}
	go l.run()
	return l
}

func (y *yyLex) run() {
	defer close(y.items)
	for {
		switch c := y.next(); {
		case unicode.IsDigit(c):
			y.lexNumber()
		case unicode.IsLetter(c):
			y.lexWord()
		case unicode.IsSpace(c):
			y.ignore()
		case c == eof:
			return
		default:
			y.emit(char)
		}
	}
}

func (y *yyLex) lexNumber() {
	y.acceptDigits()
	if y.acceptRune('.') {
		y.acceptDigits()
	}
	if y.acceptRune('e', 'E') {
		y.acceptRune('-')
		y.acceptDigits()
	}
	y.emit(number)
}

func (y *yyLex) lexWord() {
	y.acceptLetters()
	y.emit(word)
}

func (y *yyLex) emit(t int) {
	y.items <- item{
		typ: t,
		val: y.input[y.start:y.pos],
	}
	y.start = y.pos
}

func (y *yyLex) next() (r rune) {
	if y.pos >= len(y.input) {
		y.width = 0
		return eof
	}
	r, y.width = utf8.DecodeRuneInString(y.input[y.pos:])
	y.pos += y.width
	return r
}

func (y *yyLex) ignore() {
	y.start = y.pos
}

func (y *yyLex) backup() {
	y.pos -= y.width
}

func (y *yyLex) peek() rune {
	defer y.backup()
	return y.next()
}

func (y *yyLex) acceptDigits() {
	defer y.backup()
	for unicode.IsDigit(y.next()) {
	}
}

func (y *yyLex) acceptLetters() {
	defer y.backup()
	for unicode.IsLetter(y.next()) {
	}
}

func (y *yyLex) acceptRune(valid ...rune) bool {
	for _, r := range valid {
		if y.next() == r {
			return true
		}
		y.backup()
	}
	return false
}

func (y *yyLex) accept(valid string) bool {
	if strings.ContainsRune(valid, y.next()) {
		return true
	}
	y.backup()
	return false
}

func (y *yyLex) acceptRun(valid string) {
	for strings.ContainsRune(valid, y.next()) {
	}
	y.backup()
}
