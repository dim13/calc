%{

package main

import "math"

var reg = map[string]float64{
	"pi": math.Pi,
	"e":  math.E,
}

const last = "_"

%}

%union{
	fval float64
	sval string
}

%token <fval> NUMBER
%token <sval> WORD

%type <fval> exp

%left '+' '-'
%left '*' '/' '%'
%left '^'
%left UMINUS

%%

line
	:			/* empty */
	| exp			{
				  reg[last] = $1
				  yylex.(*yyLex).result = $1
				  yylex.(*yyLex).ok = true
				}
	| WORD '=' exp		{ reg[$1] = $3 }
	;

exp
	: NUMBER
	| WORD			{ $$ = reg[$1] }
	| '_'			{ $$ = reg[last] }
	| exp '+' exp		{ $$ = $1 + $3 }
	| exp '-' exp		{ $$ = $1 - $3 }
	| exp '*' exp		{ $$ = $1 * $3 }
	| exp '/' exp		{ $$ = $1 / $3 }
	| exp '%' exp		{ $$ = math.Mod($1, $3) }
	| exp '^' exp		{ $$ = math.Pow($1, $3) }
	| '-' exp %prec UMINUS	{ $$ = -$2 }
	| '(' exp ')'		{ $$ = $2 }
	| '|' exp '|'		{ $$ = math.Abs($2) }
	| error			{ $$ = math.NaN() }
	;

%%

func Parse(input string) (float64, bool) {
	l := lex(input)
	yyParse(l)
	return l.result, l.ok
}
