%{

package main

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

%}

%union{
	dval Number
	vval Interval
	rval rune
	sval string
}

%token <rval> DREG VREG
%token <dval> NUMBER
%token <sval> STRING
%token EQ LBR RBR COM

%type <dval> dexp
%type <vval> vexp

%left SUM SUB
%left MUL DIV
%left UMINUS

%%

line
	: dexp			{ fmt.Println($1) }
	| vexp			{ fmt.Println($1) }
	| DREG EQ dexp		{ $3.Set($1) }
	| VREG EQ vexp		{ $3.Set($1) }
	| STRING		{ fmt.Println($1) }
	;

dexp
	: NUMBER
	| DREG			{ $$.Get($1) }
	| dexp SUM dexp		{ $$ = $1 + $3 }
	| dexp SUB dexp		{ $$ = $1 - $3 }
	| dexp MUL dexp		{ $$ = $1 * $3 }
	| dexp DIV dexp		{ $$ = $1 / $3 }
	| SUB dexp %prec UMINUS	{ $$ = -$2 }
	| LBR dexp RBR		{ $$ = $2 }
	| error			{ $$ = Number(math.NaN()) }
	;

vexp
	: LBR dexp COM dexp RBR	{
		if $2 > $4 {
			$$ = Interval{$4, $2}
		} else {
			$$ = Interval{$2, $4}
		}
	}
	| VREG			{ $$.Get($1) }
	| vexp SUM vexp		{ $$ = $3.vadd($1) }
	| dexp SUM vexp		{ $$ = $3.vadd(Interval{$1, $1}) }
	| vexp SUM dexp		{ $$ = Interval{$3, $3}.vadd($1) }
	| vexp SUB vexp		{ $$ = $3.vsub($1) }
	| dexp SUB vexp		{ $$ = $3.vsub(Interval{$1, $1}) }
	| vexp SUB dexp		{ $$ = Interval{$3, $3}.vsub($1) }
	| vexp MUL vexp		{ $$ = $3.vmul($1) }
	| dexp MUL vexp		{ $$ = $3.vmul(Interval{$1, $1}) }
	| vexp MUL dexp		{ $$ = Interval{$3, $3}.vmul($1) }
	| vexp DIV vexp		{
		if $3.dcheck() {
			yylex.Error("divisor interval contains 0")
		}
		$$ = $3.vdiv($1)
	}
	| dexp DIV vexp		{
		if $3.dcheck() {
			yylex.Error("divisor interval contains 0")
		}
		$$ = $3.vdiv(Interval{$1, $1})
	}
	| vexp DIV dexp		{
		if $3 == 0 {
			yylex.Error("divisor interval contains 0")
		}
		$$ = Interval{$3, $3}.vdiv($1)
	}
	| SUB vexp %prec UMINUS	{ $$ = $2.vneg() }
	| LBR vexp RBR		{ $$ = $2 }
	;

%%

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
		if c > v.hi { v.hi = c }
		if d < v.lo { v.lo = d }
	} else {
		if d > v.hi { v.hi = d }
		if c < v.lo { v.lo = c }
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
