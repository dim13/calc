package main

import (
	"log"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

const eof = 0

type item struct {
	typ int
	val string
}

type yyLex struct {
	input string
	start int
	pos   int
	width int
	items chan item
}

func (y *yyLex) Error(s string) {
	log.Println(s)
}

func (y *yyLex) Lex(lval *yySymType) int {
	item := <-y.items
	switch item.typ {
	case NUMBER:
		n, err := strconv.ParseFloat(item.val, 64)
		if err != nil {
			log.Println(err)
		}
		lval.dval = Number(n)
	case VREG, DREG:
		lval.rval = rune(item.val[0])
	case STRING:
		lval.sval = item.val[1:len(item.val)-1]
	}
	return int(item.typ)
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
			y.emit(NUMBER)
		case unicode.IsUpper(c):
			y.emit(VREG)
		case unicode.IsLower(c):
			y.emit(DREG)
		case unicode.IsSpace(c):
			y.ignore()
		case c == '\'':
			if y.lexQuoted() {
				y.emit(STRING)
			} else {
				y.emit(int(c))
			}
		case c == eof:
			return
		default:
			y.emit(int(c))
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
}

func (y *yyLex) lexQuoted() bool {
	if n := strings.IndexRune(y.input[y.pos:], '\''); n >= 0 {
		y.pos += n
		y.next()
		return true
	}
	return false
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
	if strings.IndexRune(valid, y.next()) >= 0 {
		return true
	}
	y.backup()
	return false
}

func (y *yyLex) acceptRun(valid string) {
	for strings.IndexRune(valid, y.next()) >= 0 {
	}
	y.backup()
}
