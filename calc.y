%{
package main

import (
	"math"
	"math/rand"
)

var reg = map[string]float64{
	"pi": math.Pi,
	"e":  math.E,
}

var last float64

%}

%union{
	fval float64
	sval string
}

%token <fval> NUMBER
%token <sval> WORD

%type <fval> exp

%right '='
%left '+' '-'
%left '*' '/' '%'
%left UMINUS
%right '^'

%%

line
	:			/* empty */
	| exp			{
				  last = $1
				  yylex.(*yyLex).result = $1
				}
	| error
	;

exp
	: NUMBER
	| WORD			{
				  v, ok := reg[$1]
				  if !ok {
					yylex.Error("undefined: " + $1)
				  }
				  $$ = v
				}
	| WORD '=' exp		{
				  reg[$1] = $3
				  $$ = $3
				}
	| '_'			{ $$ = last }
	| '?'			{ $$ = rand.Float64() }
	| exp '+' exp		{ $$ = $1 + $3 }
	| exp '-' exp		{ $$ = $1 - $3 }
	| exp '*' exp		{ $$ = $1 * $3 }
	| exp '/' exp		{ $$ = $1 / $3 }
	| exp '%' exp		{ $$ = math.Mod($1, $3) }
	| exp '^' exp		{ $$ = math.Pow($1, $3) }
	| '-' exp %prec UMINUS	{ $$ = -$2 }
	| '(' exp ')'		{ $$ = $2 }
	| '|' exp '|'		{ $$ = math.Abs($2) }
	;

%%

func Parse(input string) (float64, error) {
	l := lex(input)
	yyParse(l)
	return l.result, l.err
}
