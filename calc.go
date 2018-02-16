//line calc.y:2
package main

import __yyfmt__ "fmt"

//line calc.y:3
import (
	"fmt"
	"math"
)

type Number float64

type Interval struct {
	lo Number
	hi Number
}

var nreg = make(map[rune]Number)
var ireg = make(map[rune]Interval)

//line calc.y:22
type yySymType struct {
	yys  int
	dval Number
	vval Interval
	rval rune
	sval string
}

const DREG = 57346
const VREG = 57347
const NUMBER = 57348
const STRING = 57349
const EQ = 57350
const LBR = 57351
const RBR = 57352
const COM = 57353
const SUM = 57354
const SUB = 57355
const MUL = 57356
const DIV = 57357
const UMINUS = 57358

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"DREG",
	"VREG",
	"NUMBER",
	"STRING",
	"EQ",
	"LBR",
	"RBR",
	"COM",
	"SUM",
	"SUB",
	"MUL",
	"DIV",
	"UMINUS",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line calc.y:103

func (i Interval) String() string {
	return fmt.Sprint("(", i.lo, ",", i.hi, ")")
}

func hilo(a, b, c, d Number) (v Interval) {
	if a > b {
		v.hi = a
		v.lo = b
	} else {
		v.hi = b
		v.lo = a
	}
	if c > d {
		if c > v.hi {
			v.hi = c
		}
		if d < v.lo {
			v.lo = d
		}
	} else {
		if d > v.hi {
			v.hi = d
		}
		if c < v.lo {
			v.lo = c
		}
	}
	return
}

func (v Interval) vmul(a Interval) Interval {
	return hilo(a.lo*v.hi, a.lo*v.lo, a.hi*v.hi, a.hi*v.lo)
}

func (v Interval) dcheck() bool {
	return v.hi >= 0 && v.lo <= 0
}

func (v Interval) vdiv(a Interval) Interval {
	return hilo(a.lo/v.hi, a.lo/v.lo, a.hi/v.hi, a.hi/v.lo)
}

func (v Interval) vadd(a Interval) Interval {
	return Interval{a.lo + v.lo, a.hi + v.hi}
}

func (v Interval) vsub(a Interval) Interval {
	return Interval{a.lo - v.lo, a.hi - v.hi}
}

func (v Interval) vneg() Interval {
	return Interval{lo: -v.hi, hi: -v.lo}
}

func (n *Number) Set(key rune) { nreg[key] = *n }
func (n *Number) Get(key rune) { *n = nreg[key] }

func (n *Interval) Set(key rune) { ireg[key] = *n }
func (n *Interval) Get(key rune) { *n = ireg[key] }

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 110

var yyAct = [...]int{

	2, 48, 49, 11, 12, 13, 14, 53, 54, 21,
	25, 20, 27, 29, 31, 33, 36, 38, 40, 42,
	43, 47, 3, 62, 19, 51, 52, 53, 54, 13,
	14, 22, 26, 1, 28, 30, 32, 34, 35, 37,
	39, 41, 0, 46, 0, 55, 56, 17, 18, 0,
	57, 0, 58, 59, 60, 61, 48, 0, 51, 52,
	53, 54, 50, 0, 15, 16, 17, 18, 10, 0,
	4, 5, 7, 6, 0, 9, 0, 0, 0, 8,
	11, 12, 13, 14, 15, 16, 17, 18, 51, 52,
	53, 54, 10, 0, 23, 24, 7, 0, 10, 9,
	23, 0, 7, 8, 0, 45, 0, 0, 0, 44,
}
var yyPact = [...]int{

	66, -1000, 68, 72, 16, 3, -1000, -1000, 90, 90,
	-1000, 90, 90, 90, 90, 90, 90, 90, 90, 96,
	90, -1000, -1000, -1000, -1000, -9, 52, 15, 33, 15,
	33, -1000, -1000, -1000, -1000, 33, 15, 33, 15, -1000,
	-1000, -1000, -1000, 76, 96, 96, 72, 68, -1000, 96,
	-1000, 96, 96, 96, 96, -1000, 46, 13, -7, -7,
	-1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 0, 22, 33,
}
var yyR1 = [...]int{

	0, 3, 3, 3, 3, 3, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2,
}
var yyR2 = [...]int{

	0, 1, 1, 3, 3, 1, 1, 1, 3, 3,
	3, 3, 2, 3, 1, 5, 1, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 2,
	3,
}
var yyChk = [...]int{

	-1000, -3, -1, -2, 4, 5, 7, 6, 13, 9,
	2, 12, 13, 14, 15, 12, 13, 14, 15, 8,
	8, -1, -2, 4, 5, -1, -2, -1, -2, -1,
	-2, -1, -2, -1, -2, -2, -1, -2, -1, -2,
	-1, -2, -1, -1, 13, 9, -2, -1, 10, 11,
	10, 12, 13, 14, 15, -1, -1, -1, -1, -1,
	-1, -1, 10,
}
var yyDef = [...]int{

	0, -2, 1, 2, 7, 16, 5, 6, 0, 0,
	14, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 12, 29, 7, 16, 0, 0, 8, 18, 9,
	21, 10, 24, 11, 27, 17, 19, 20, 22, 23,
	25, 26, 28, 3, 0, 0, 4, 0, 13, 0,
	30, 0, 0, 0, 0, 12, 0, 0, 8, 9,
	10, 11, 15,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line calc.y:44
		{
			fmt.Println(yyDollar[1].dval)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line calc.y:45
		{
			fmt.Println(yyDollar[1].vval)
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line calc.y:46
		{
			yyDollar[3].dval.Set(yyDollar[1].rval)
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line calc.y:47
		{
			yyDollar[3].vval.Set(yyDollar[1].rval)
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line calc.y:48
		{
			fmt.Println(yyDollar[1].sval)
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line calc.y:53
		{
			yyVAL.dval.Get(yyDollar[1].rval)
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line calc.y:54
		{
			yyVAL.dval = yyDollar[1].dval + yyDollar[3].dval
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line calc.y:55
		{
			yyVAL.dval = yyDollar[1].dval - yyDollar[3].dval
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line calc.y:56
		{
			yyVAL.dval = yyDollar[1].dval * yyDollar[3].dval
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line calc.y:57
		{
			yyVAL.dval = yyDollar[1].dval / yyDollar[3].dval
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line calc.y:58
		{
			yyVAL.dval = -yyDollar[2].dval
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line calc.y:59
		{
			yyVAL.dval = yyDollar[2].dval
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line calc.y:60
		{
			yyVAL.dval = Number(math.NaN())
		}
	case 15:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line calc.y:64
		{
			if yyDollar[2].dval > yyDollar[4].dval {
				yyVAL.vval = Interval{yyDollar[4].dval, yyDollar[2].dval}
			} else {
				yyVAL.vval = Interval{yyDollar[2].dval, yyDollar[4].dval}
			}
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line calc.y:71
		{
			yyVAL.vval.Get(yyDollar[1].rval)
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line calc.y:72
		{
			yyVAL.vval = yyDollar[3].vval.vadd(yyDollar[1].vval)
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line calc.y:73
		{
			yyVAL.vval = yyDollar[3].vval.vadd(Interval{yyDollar[1].dval, yyDollar[1].dval})
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line calc.y:74
		{
			yyVAL.vval = Interval{yyDollar[3].dval, yyDollar[3].dval}.vadd(yyDollar[1].vval)
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line calc.y:75
		{
			yyVAL.vval = yyDollar[3].vval.vsub(yyDollar[1].vval)
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line calc.y:76
		{
			yyVAL.vval = yyDollar[3].vval.vsub(Interval{yyDollar[1].dval, yyDollar[1].dval})
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line calc.y:77
		{
			yyVAL.vval = Interval{yyDollar[3].dval, yyDollar[3].dval}.vsub(yyDollar[1].vval)
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line calc.y:78
		{
			yyVAL.vval = yyDollar[3].vval.vmul(yyDollar[1].vval)
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line calc.y:79
		{
			yyVAL.vval = yyDollar[3].vval.vmul(Interval{yyDollar[1].dval, yyDollar[1].dval})
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line calc.y:80
		{
			yyVAL.vval = Interval{yyDollar[3].dval, yyDollar[3].dval}.vmul(yyDollar[1].vval)
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line calc.y:81
		{
			if yyDollar[3].vval.dcheck() {
				yylex.Error("divisor interval contains 0")
			}
			yyVAL.vval = yyDollar[3].vval.vdiv(yyDollar[1].vval)
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line calc.y:87
		{
			if yyDollar[3].vval.dcheck() {
				yylex.Error("divisor interval contains 0")
			}
			yyVAL.vval = yyDollar[3].vval.vdiv(Interval{yyDollar[1].dval, yyDollar[1].dval})
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line calc.y:93
		{
			if yyDollar[3].dval == 0 {
				yylex.Error("divisor interval contains 0")
			}
			yyVAL.vval = Interval{yyDollar[3].dval, yyDollar[3].dval}.vdiv(yyDollar[1].vval)
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line calc.y:99
		{
			yyVAL.vval = yyDollar[2].vval.vneg()
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line calc.y:100
		{
			yyVAL.vval = yyDollar[2].vval
		}
	}
	goto yystack /* stack new state and value */
}
