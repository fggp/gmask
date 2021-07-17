// Gmask parser.

%{

package main

import (
	"bufio"
	"fmt"
	"github.com/fggp/gmasklib"
	"io"
	"io/ioutil"
	"log"
	"os"
)

%}

%union {
  val float64
  n   int
  sco string
  fie *gmasklib.Field
  par gmasklib.Param
  gen gmasklib.Generator
  itm gmasklib.ItemMode
  lst *List
  ipl *gmasklib.Interpolation
  rmd gmasklib.RndMode
  omd gmasklib.OscMode
  amd gmasklib.AccumMode
  sli []interface{}
}

%token SCO
%token NUMBER
%token F PARAM
%token IPL COS OFF
%token PREC
%token ITEM CYCLE SWING HEAP RANDOM
%token CONST SEG RANGE
%token RND UNI LIN RLIN TRI EXP REXP BEXP GAUSS CAUCHY BETA WEI
%token OSC SIN SQUARE TRIANGLE SAWUP SAWDOWN POWUP POWDOWN
%token MASK MAP QUANT
%token ACCUM ON LIMIT MIRROR WRAP INIT

%type <sco> SCO, csco
%type <val> NUMBER
%type <n>   PARAM prec
%type <fie> field
%type <par> param
%type <gen> generator bpf rndgen oscgen
%type <itm> item
%type <lst> list bplist
%type <ipl> ipl
%type <rmd> rnd
%type <omd> osc
%type <sli> mask quant accum
%type <amd> accummode

%% // Grammar rules and actions follow.

input: csco     { fmt.Fprintf(w, "%s\n", $1) }
| input field	{ fieldNum++; $2.EvalToScore(w, fieldNum) }
;

csco: SCO       { $$ = $1 }
;

field: F NUMBER NUMBER    { $$ = gmasklib.NewField($2, $3) }
| field param             { $1.AddParam($2) }
;

param: PARAM generator prec   { $$ = gmasklib.NewParam($1, $2, $3) }
;

generator: CONST NUMBER         { $$ = gmasklib.ConstGen($2) }
| item '(' list ')'             { $$ = gmasklib.ItemGen($1, $3.GetVal()) }
| SEG bpf                       { $$ = $2 }
| RANGE NUMBER NUMBER           { $$ = gmasklib.RangeGen($2, $3) }
| rndgen
| oscgen
| generator mask                  { $$ = gmasklib.MaskGen($1, $2...) }
| generator quant                 { $$ = gmasklib.QuantGen($1, $2...) }
| generator ACCUM ON              { $$ = gmasklib.AccumGen($1, gmasklib.ON) }
| generator ACCUM ON INIT NUMBER  { $$ = gmasklib.AccumGen($1, gmasklib.ON, $5) }
| generator accum                 { $$ = gmasklib.AccumGen($1, $2...) }
;

item: ITEM CYCLE          { $$ = gmasklib.CYCLE }
| ITEM SWING              { $$ = gmasklib.SWING }
| ITEM HEAP               { $$ = gmasklib.HEAP }
| ITEM RANDOM             { $$ = gmasklib.RANDOM }
;

list: NUMBER              { $$ = NewList($1) } 
| list NUMBER             { $1.AddVal($2) }
;

bpf: '(' bplist ipl ')'      { $$ = gmasklib.BpfGen($2.GetVal(), $3) }
| '[' NUMBER NUMBER ipl ']'  { $$ = gmasklib.BpfGen([]float64{$2, $3}, $4) }
;

bplist: NUMBER NUMBER     { $$ = NewBpList($1, $2) } 
| bplist NUMBER NUMBER    { $1.AddBp($2, $3) }
;

ipl: /* empty */          { $$ = gmasklib.NewInterpolation(0, false, false) }
| IPL NUMBER              { $$ = gmasklib.NewInterpolation($2, false, false) }
| IPL COS                 { $$ = gmasklib.NewInterpolation(0, true, false) }
| IPL OFF                 { $$ = gmasklib.NewInterpolation(0, false, true) }
;

rnd: RND UNI    { $$ = gmasklib.UNI }
| RND LIN       { $$ = gmasklib.LIN }
| RND RLIN      { $$ = gmasklib.RLIN }
| RND TRI       { $$ = gmasklib.TRI }
| RND EXP       { $$ = gmasklib.EXP }
| RND REXP      { $$ = gmasklib.REXP }
| RND BEXP      { $$ = gmasklib.BEXP }
| RND GAUSS     { $$ = gmasklib.GAUSS }
| RND CAUCHY    { $$ = gmasklib.CAUCHY }
| RND BETA      { $$ = gmasklib.BETA }
| RND WEI       { $$ = gmasklib.WEI }
;

rndgen: rnd            { $$ = gmasklib.RndGen($1) }
| rnd NUMBER           { $$ = gmasklib.RndGen($1, $2) }
| rnd bpf              { $$ = gmasklib.RndGen($1, $2) }
| rnd NUMBER NUMBER    { $$ = gmasklib.RndGen($1, $2, $3) }
| rnd bpf NUMBER       { $$ = gmasklib.RndGen($1, $2, $3) }
| rnd NUMBER bpf       { $$ = gmasklib.RndGen($1, $2, $3) }
| rnd bpf bpf          { $$ = gmasklib.RndGen($1, $2, $3) }
;

osc: OSC SIN      { $$ = gmasklib.SIN }
| OSC COS         { $$ = gmasklib.COS }
| OSC SQUARE      { $$ = gmasklib.SQUARE }
| OSC TRIANGLE    { $$ = gmasklib.TRIANGLE }
| OSC SAWUP       { $$ = gmasklib.SAWUP }
| OSC SAWDOWN     { $$ = gmasklib.SAWDOWN }
| OSC POWUP       { $$ = gmasklib.POWUP }
| OSC POWDOWN     { $$ = gmasklib.POWDOWN }
;

oscgen: osc NUMBER          { $$ = gmasklib.OscGen($1, $2) }
| osc bpf                   { $$ = gmasklib.OscGen($1, $2) }
| osc NUMBER NUMBER         { $$ = gmasklib.OscGen($1, $2, $3) }
| osc bpf NUMBER            { $$ = gmasklib.OscGen($1, $2, $3) }
| osc NUMBER NUMBER NUMBER  { $$ = gmasklib.OscGen($1, $2, $3, $4) }
| osc bpf NUMBER NUMBER     { $$ = gmasklib.OscGen($1, $2, $3, $4) }
;

mask: MASK NUMBER NUMBER    { $$ = NewInterfaceSlice($2, $3) }
| MASK bpf NUMBER           { $$ = NewInterfaceSlice($2, $3) }
| MASK NUMBER bpf           { $$ = NewInterfaceSlice($2, $3) }
| MASK bpf bpf              { $$ = NewInterfaceSlice($2, $3) }
| mask MAP NUMBER           { $$ = append($1, $3) }
;

quant: QUANT NUMBER            { $$ = NewInterfaceSlice($2) }
| QUANT bpf                    { $$ = NewInterfaceSlice($2) }
| QUANT NUMBER NUMBER          { $$ = NewInterfaceSlice($2, $3) }
| QUANT NUMBER bpf             { $$ = NewInterfaceSlice($2, $3) }
| QUANT bpf NUMBER             { $$ = NewInterfaceSlice($2, $3) }
| QUANT bpf bpf                { $$ = NewInterfaceSlice($2, $3) }
| QUANT NUMBER NUMBER NUMBER   { $$ = NewInterfaceSlice($2, $3, $4) }
| QUANT NUMBER NUMBER bpf      { $$ = NewInterfaceSlice($2, $3, $4) }
| QUANT NUMBER bpf NUMBER      { $$ = NewInterfaceSlice($2, $3, $4) }
| QUANT NUMBER bpf bpf         { $$ = NewInterfaceSlice($2, $3, $4) }
| QUANT bpf NUMBER NUMBER      { $$ = NewInterfaceSlice($2, $3, $4) }
| QUANT bpf NUMBER bpf         { $$ = NewInterfaceSlice($2, $3, $4) }
| QUANT bpf bpf NUMBER         { $$ = NewInterfaceSlice($2, $3, $4) }
| QUANT bpf bpf bpf            { $$ = NewInterfaceSlice($2, $3, $4) }
;

accummode: ACCUM LIMIT    { $$ = gmasklib.LIMIT }
| ACCUM MIRROR            { $$ = gmasklib.MIRROR }
| ACCUM WRAP              { $$ = gmasklib.WRAP }
;

accum: accummode NUMBER NUMBER  { $$ = NewInterfaceSlice($1, $2, $3) }
| accummode bpf NUMBER          { $$ = NewInterfaceSlice($1, $2, $3) }
| accummode NUMBER bpf          { $$ = NewInterfaceSlice($1, $2, $3) }
| accummode bpf bpf             { $$ = NewInterfaceSlice($1, $2, $3) }
| accum INIT NUMBER             { $$ = append($1, $3) }
;

prec: /* empty */         { $$ = -1 }
| PREC NUMBER             { $$ = int($2) }
;

%%
	
const SLICEINC = 10

type List struct {
	pos int
	val []float64
}

func NewList(x float64) *List {
	v := make([]float64, SLICEINC)
	v[0] = x
	l := List{pos: 1, val: v}
	return &l
}

func NewBpList(t, x float64) *List {
	v := make([]float64, SLICEINC)
	v[0], v[1] = t, x
	l := List{pos: 2, val: v}
	return &l
}

func (l *List) AddVal(x float64) {
	if l.pos == len(l.val) {
		l.val = append(l.val, make([]float64, SLICEINC)...)
	}
	l.val[l.pos] = x
	l.pos++
}

func (l *List) AddBp(t, x float64) {
	if l.pos == len(l.val) {
		l.val = append(l.val, make([]float64, SLICEINC)...)
	}
	l.val[l.pos], l.val[l.pos+1] = t, x
	l.pos += 2
}

func (l *List) GetVal() []float64 {
	return l.val[:l.pos]
}

func NewInterfaceSlice(params ...interface{}) []interface{} {
	return params
}

var w io.Writer
var fieldNum int

func main() {
	input, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) <= 2 {
		w = os.Stdout
	} else {
		fo, err := os.Create(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		defer fo.Close()
		w = bufio.NewWriter(fo)
	}
	yyParse(NewLexer(input))
}
