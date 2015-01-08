package main

import (
	"log"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	eof = iota
	digit
	lower
	upper
	quoted
	char
)

var charmap = map[rune]int{
	'+': SUM,
	'-': SUB,
	'*': MUL,
	'/': DIV,
	'=': EQ,
	'(': LBR,
	')': RBR,
	',': COM,
}

var typemap = map[int]int{
	digit:  NUMBER,
	upper:  VREG,
	lower:  DREG,
	quoted: STRING,
	char:   eof,
}

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

func (y *yyLex) Lex(lval *yySymType) (ret int) {
	item := <-y.items
	ret = typemap[item.typ]

	switch item.typ {
	case digit:
		n, err := strconv.ParseFloat(item.val, 64)
		if err != nil {
			log.Println(err)
		}
		lval.dval = Number(n)
	case upper:
		lval.rval = rune(item.val[0])
	case lower:
		lval.rval = rune(item.val[0])
	case quoted:
		lval.sval = item.val[1 : len(item.val)-1]
	case char:
		c := rune(item.val[0])
		if ch, ok := charmap[c]; ok {
			ret = ch
		} else {
			ret = int(c)
		}
	}
	return ret
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
		case unicode.IsUpper(c):
			y.emit(upper)
		case unicode.IsLower(c):
			y.emit(lower)
		case unicode.IsSpace(c):
			y.ignore()
		case c == eof:
			return
		case c == '\'':
			y.lexQuoted()
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
	y.emit(digit)
}

func (y *yyLex) lexQuoted() {
	if n := strings.IndexRune(y.input[y.pos:], '\''); n >= 0 {
		y.pos += n
		y.next()
		y.emit(quoted)
	} else {
		y.emit(char)
	}
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
