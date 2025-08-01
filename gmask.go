// Code generated by goyacc -o gmask.go gmask.y. DO NOT EDIT.

//line gmask.y:4

package main

import __yyfmt__ "fmt"

//line gmask.y:5

import (
	"bufio"
	"fmt"
	"github.com/fggp/gmasklib"
	"io"
	"log"
	"os"
)

//line gmask.y:18
type yySymType struct {
	yys int
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

const SCO = 57346
const NUMBER = 57347
const F = 57348
const PARAM = 57349
const IPL = 57350
const COS = 57351
const OFF = 57352
const PREC = 57353
const ITEM = 57354
const CYCLE = 57355
const SWING = 57356
const HEAP = 57357
const RANDOM = 57358
const CONST = 57359
const SEG = 57360
const RANGE = 57361
const RND = 57362
const UNI = 57363
const LIN = 57364
const RLIN = 57365
const TRI = 57366
const EXP = 57367
const REXP = 57368
const BEXP = 57369
const GAUSS = 57370
const CAUCHY = 57371
const BETA = 57372
const WEI = 57373
const OSC = 57374
const SIN = 57375
const SQUARE = 57376
const TRIANGLE = 57377
const SAWUP = 57378
const SAWDOWN = 57379
const POWUP = 57380
const POWDOWN = 57381
const MASK = 57382
const MAP = 57383
const QUANT = 57384
const ACCUM = 57385
const ON = 57386
const LIMIT = 57387
const MIRROR = 57388
const WRAP = 57389
const INIT = 57390

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"SCO",
	"NUMBER",
	"F",
	"PARAM",
	"IPL",
	"COS",
	"OFF",
	"PREC",
	"ITEM",
	"CYCLE",
	"SWING",
	"HEAP",
	"RANDOM",
	"CONST",
	"SEG",
	"RANGE",
	"RND",
	"UNI",
	"LIN",
	"RLIN",
	"TRI",
	"EXP",
	"REXP",
	"BEXP",
	"GAUSS",
	"CAUCHY",
	"BETA",
	"WEI",
	"OSC",
	"SIN",
	"SQUARE",
	"TRIANGLE",
	"SAWUP",
	"SAWDOWN",
	"POWUP",
	"POWDOWN",
	"MASK",
	"MAP",
	"QUANT",
	"ACCUM",
	"ON",
	"LIMIT",
	"MIRROR",
	"WRAP",
	"INIT",
	"'('",
	"')'",
	"'['",
	"']'",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line gmask.y:192

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
	input, err := os.ReadFile(os.Args[1])
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

//line yacctab:1
var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 163

var yyAct = [...]uint8{
	33, 106, 120, 118, 128, 122, 116, 114, 32, 90,
	102, 100, 69, 64, 98, 96, 108, 105, 42, 44,
	94, 92, 3, 5, 85, 83, 34, 123, 35, 72,
	74, 76, 124, 7, 75, 73, 125, 126, 71, 43,
	113, 112, 84, 86, 41, 111, 34, 34, 35, 35,
	34, 34, 35, 35, 34, 34, 35, 35, 34, 34,
	35, 35, 104, 110, 34, 34, 35, 35, 34, 34,
	35, 35, 93, 95, 97, 99, 101, 103, 34, 34,
	35, 35, 34, 34, 35, 35, 109, 91, 34, 57,
	35, 65, 66, 67, 68, 27, 89, 115, 117, 119,
	121, 16, 107, 88, 87, 108, 10, 12, 13, 19,
	82, 81, 127, 56, 58, 59, 60, 61, 62, 63,
	80, 20, 78, 70, 28, 36, 29, 25, 45, 46,
	47, 48, 49, 50, 51, 52, 53, 54, 55, 37,
	38, 39, 40, 31, 21, 8, 1, 30, 26, 24,
	23, 18, 17, 79, 77, 11, 15, 14, 9, 6,
	4, 22, 2,
}

var yyPact = [...]int16{
	18, 17, -1000, -1000, 26, 140, -1000, 89, 139, 84,
	138, -41, -23, 120, -1000, -1000, 126, 39, 34, 107,
	80, -1000, -1000, -28, -1000, 47, -36, 118, 33, 30,
	29, -1000, 117, -1000, 115, 106, 105, -1000, -1000, -1000,
	-1000, 20, 19, 99, 98, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 91, -39, -1000, -1000, -1000, 82,
	-1000, 16, 15, 10, 9, 6, 5, 12, -1000, 97,
	81, 58, -1000, -1000, -1000, -1000, -1000, 40, 36, -1000,
	35, -1000, -1000, -1000, -1000, -1000, 2, 1, -2, -3,
	-1000, -1000, -1000, -1000, -1000, -1000, -45, 22, 27, -1000,
	8, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -48, -1000,
}

var yyPgo = [...]uint8{
	0, 162, 161, 160, 159, 158, 0, 157, 156, 155,
	154, 153, 1, 152, 151, 150, 149, 148, 147, 146,
}

var yyR1 = [...]int8{
	0, 19, 19, 1, 3, 3, 4, 5, 5, 5,
	5, 5, 5, 5, 5, 5, 5, 5, 9, 9,
	9, 9, 10, 10, 6, 6, 11, 11, 12, 12,
	12, 12, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 7, 7, 7, 7, 7, 7, 7,
	14, 14, 14, 14, 14, 14, 14, 14, 8, 8,
	8, 8, 8, 8, 15, 15, 15, 15, 15, 16,
	16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
	16, 16, 16, 18, 18, 18, 17, 17, 17, 17,
	17, 2, 2,
}

var yyR2 = [...]int8{
	0, 1, 2, 1, 3, 2, 3, 2, 4, 2,
	3, 1, 1, 2, 2, 3, 5, 2, 2, 2,
	2, 2, 1, 2, 4, 5, 2, 3, 0, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 1, 2, 2, 3, 3, 3, 3,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	3, 3, 4, 4, 3, 3, 3, 3, 3, 2,
	2, 3, 3, 3, 3, 4, 4, 4, 4, 4,
	4, 4, 4, 2, 2, 2, 3, 3, 3, 3,
	3, 0, 2,
}

var yyChk = [...]int16{
	-1000, -19, -1, 4, -3, 6, -4, 7, 5, -5,
	17, -9, 18, 19, -7, -8, 12, -13, -14, 20,
	32, 5, -2, -15, -16, 43, -17, 11, 40, 42,
	-18, 5, 49, -6, 49, 51, 5, 13, 14, 15,
	16, 5, -6, 5, -6, 21, 22, 23, 24, 25,
	26, 27, 28, 29, 30, 31, 33, 9, 34, 35,
	36, 37, 38, 39, 41, 44, 45, 46, 47, 48,
	5, 5, -6, 5, -6, 5, -6, -10, 5, -11,
	5, 5, 5, 5, -6, 5, -6, 5, 5, 5,
	48, 5, 5, -6, 5, -6, 5, -6, 5, -6,
	5, -6, 5, -6, 50, 5, -12, 5, 8, 5,
	5, 5, 5, 5, 5, -6, 5, -6, 5, -6,
	5, -6, 50, 5, 5, 9, 10, -12, 52,
}

var yyDef = [...]int8{
	0, -2, 1, 3, 2, 0, 5, 0, 0, 91,
	0, 0, 0, 0, 11, 12, 0, 43, 0, 0,
	0, 4, 6, 13, 14, 0, 17, 0, 0, 0,
	0, 7, 0, 9, 0, 0, 0, 18, 19, 20,
	21, 44, 45, 58, 59, 32, 33, 34, 35, 36,
	37, 38, 39, 40, 41, 42, 50, 51, 52, 53,
	54, 55, 56, 57, 0, 15, 83, 84, 85, 0,
	92, 0, 0, 69, 70, 0, 0, 0, 22, 28,
	0, 0, 10, 46, 48, 47, 49, 60, 61, 68,
	0, 90, 64, 66, 65, 67, 71, 72, 73, 74,
	86, 88, 87, 89, 8, 23, 0, 0, 0, 26,
	28, 62, 63, 16, 75, 76, 77, 78, 79, 80,
	81, 82, 24, 27, 29, 30, 31, 0, 25,
}

var yyTok1 = [...]int8{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	49, 50, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 51, 3, 52,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48,
}

var yyTok3 = [...]int8{
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
	base := int(yyPact[state])
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && int(yyChk[int(yyAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || int(yyExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := int(yyExca[i])
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
		token = int(yyTok1[0])
		goto out
	}
	if char < len(yyTok1) {
		token = int(yyTok1[char])
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = int(yyTok2[char-yyPrivate])
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = int(yyTok3[i+0])
		if token == char {
			token = int(yyTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(yyTok2[1]) /* unknown char */
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
	yyn = int(yyPact[yystate])
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
	yyn = int(yyAct[yyn])
	if int(yyChk[yyn]) == yytoken { /* valid shift */
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
	yyn = int(yyDef[yystate])
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && int(yyExca[xi+1]) == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = int(yyExca[xi+0])
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = int(yyExca[xi+1])
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
				yyn = int(yyPact[yyS[yyp].yys]) + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = int(yyAct[yyn]) /* simulate a shift of "error" */
					if int(yyChk[yystate]) == yyErrCode {
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

	yyp -= int(yyR2[yyn])
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = int(yyR1[yyn])
	yyg := int(yyPgo[yyn])
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = int(yyAct[yyg])
	} else {
		yystate = int(yyAct[yyj])
		if int(yyChk[yystate]) != -yyn {
			yystate = int(yyAct[yyg])
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gmask.y:62
		{
			fmt.Fprintf(w, "%s\n", yyDollar[1].sco)
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:63
		{
			fieldNum++
			yyDollar[2].fie.EvalToScore(w, fieldNum)
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gmask.y:66
		{
			yyVAL.sco = yyDollar[1].sco
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gmask.y:69
		{
			yyVAL.fie = gmasklib.NewField(yyDollar[2].val, yyDollar[3].val)
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:70
		{
			yyDollar[1].fie.AddParam(yyDollar[2].par)
		}
	case 6:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gmask.y:73
		{
			yyVAL.par = gmasklib.NewParam(yyDollar[1].n, yyDollar[2].gen, yyDollar[3].n)
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:76
		{
			yyVAL.gen = gmasklib.ConstGen(yyDollar[2].val)
		}
	case 8:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gmask.y:77
		{
			yyVAL.gen = gmasklib.ItemGen(yyDollar[1].itm, yyDollar[3].lst.GetVal())
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:78
		{
			yyVAL.gen = yyDollar[2].gen
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gmask.y:79
		{
			yyVAL.gen = gmasklib.RangeGen(yyDollar[2].val, yyDollar[3].val)
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:82
		{
			yyVAL.gen = gmasklib.MaskGen(yyDollar[1].gen, yyDollar[2].sli...)
		}
	case 14:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:83
		{
			yyVAL.gen = gmasklib.QuantGen(yyDollar[1].gen, yyDollar[2].sli...)
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gmask.y:84
		{
			yyVAL.gen = gmasklib.AccumGen(yyDollar[1].gen, gmasklib.ON)
		}
	case 16:
		yyDollar = yyS[yypt-5 : yypt+1]
//line gmask.y:85
		{
			yyVAL.gen = gmasklib.AccumGen(yyDollar[1].gen, gmasklib.ON, yyDollar[5].val)
		}
	case 17:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:86
		{
			yyVAL.gen = gmasklib.AccumGen(yyDollar[1].gen, yyDollar[2].sli...)
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:89
		{
			yyVAL.itm = gmasklib.CYCLE
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:90
		{
			yyVAL.itm = gmasklib.SWING
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:91
		{
			yyVAL.itm = gmasklib.HEAP
		}
	case 21:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:92
		{
			yyVAL.itm = gmasklib.RANDOM
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gmask.y:95
		{
			yyVAL.lst = NewList(yyDollar[1].val)
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:96
		{
			yyDollar[1].lst.AddVal(yyDollar[2].val)
		}
	case 24:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gmask.y:99
		{
			yyVAL.gen = gmasklib.BpfGen(yyDollar[2].lst.GetVal(), yyDollar[3].ipl)
		}
	case 25:
		yyDollar = yyS[yypt-5 : yypt+1]
//line gmask.y:100
		{
			yyVAL.gen = gmasklib.BpfGen([]float64{yyDollar[2].val, yyDollar[3].val}, yyDollar[4].ipl)
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:103
		{
			yyVAL.lst = NewBpList(yyDollar[1].val, yyDollar[2].val)
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gmask.y:104
		{
			yyDollar[1].lst.AddBp(yyDollar[2].val, yyDollar[3].val)
		}
	case 28:
		yyDollar = yyS[yypt-0 : yypt+1]
//line gmask.y:107
		{
			yyVAL.ipl = gmasklib.NewInterpolation(0, gmasklib.IPLNUM)
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:108
		{
			yyVAL.ipl = gmasklib.NewInterpolation(yyDollar[2].val, gmasklib.IPLNUM)
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:109
		{
			yyVAL.ipl = gmasklib.NewInterpolation(0, gmasklib.IPLCOS)
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:110
		{
			yyVAL.ipl = gmasklib.NewInterpolation(0, gmasklib.IPLOFF)
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:113
		{
			yyVAL.rmd = gmasklib.UNI
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:114
		{
			yyVAL.rmd = gmasklib.LIN
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:115
		{
			yyVAL.rmd = gmasklib.RLIN
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:116
		{
			yyVAL.rmd = gmasklib.TRI
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:117
		{
			yyVAL.rmd = gmasklib.EXP
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:118
		{
			yyVAL.rmd = gmasklib.REXP
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:119
		{
			yyVAL.rmd = gmasklib.BEXP
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:120
		{
			yyVAL.rmd = gmasklib.GAUSS
		}
	case 40:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:121
		{
			yyVAL.rmd = gmasklib.CAUCHY
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:122
		{
			yyVAL.rmd = gmasklib.BETA
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:123
		{
			yyVAL.rmd = gmasklib.WEI
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gmask.y:126
		{
			yyVAL.gen = gmasklib.RndGen(yyDollar[1].rmd)
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:127
		{
			yyVAL.gen = gmasklib.RndGen(yyDollar[1].rmd, yyDollar[2].val)
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:128
		{
			yyVAL.gen = gmasklib.RndGen(yyDollar[1].rmd, yyDollar[2].gen)
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gmask.y:129
		{
			yyVAL.gen = gmasklib.RndGen(yyDollar[1].rmd, yyDollar[2].val, yyDollar[3].val)
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gmask.y:130
		{
			yyVAL.gen = gmasklib.RndGen(yyDollar[1].rmd, yyDollar[2].gen, yyDollar[3].val)
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gmask.y:131
		{
			yyVAL.gen = gmasklib.RndGen(yyDollar[1].rmd, yyDollar[2].val, yyDollar[3].gen)
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gmask.y:132
		{
			yyVAL.gen = gmasklib.RndGen(yyDollar[1].rmd, yyDollar[2].gen, yyDollar[3].gen)
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:135
		{
			yyVAL.omd = gmasklib.SIN
		}
	case 51:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:136
		{
			yyVAL.omd = gmasklib.COS
		}
	case 52:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:137
		{
			yyVAL.omd = gmasklib.SQUARE
		}
	case 53:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:138
		{
			yyVAL.omd = gmasklib.TRIANGLE
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:139
		{
			yyVAL.omd = gmasklib.SAWUP
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:140
		{
			yyVAL.omd = gmasklib.SAWDOWN
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:141
		{
			yyVAL.omd = gmasklib.POWUP
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:142
		{
			yyVAL.omd = gmasklib.POWDOWN
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:145
		{
			yyVAL.gen = gmasklib.OscGen(yyDollar[1].omd, yyDollar[2].val)
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:146
		{
			yyVAL.gen = gmasklib.OscGen(yyDollar[1].omd, yyDollar[2].gen)
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gmask.y:147
		{
			yyVAL.gen = gmasklib.OscGen(yyDollar[1].omd, yyDollar[2].val, yyDollar[3].val)
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gmask.y:148
		{
			yyVAL.gen = gmasklib.OscGen(yyDollar[1].omd, yyDollar[2].gen, yyDollar[3].val)
		}
	case 62:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gmask.y:149
		{
			yyVAL.gen = gmasklib.OscGen(yyDollar[1].omd, yyDollar[2].val, yyDollar[3].val, yyDollar[4].val)
		}
	case 63:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gmask.y:150
		{
			yyVAL.gen = gmasklib.OscGen(yyDollar[1].omd, yyDollar[2].gen, yyDollar[3].val, yyDollar[4].val)
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gmask.y:153
		{
			yyVAL.sli = NewInterfaceSlice(yyDollar[2].val, yyDollar[3].val)
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gmask.y:154
		{
			yyVAL.sli = NewInterfaceSlice(yyDollar[2].gen, yyDollar[3].val)
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gmask.y:155
		{
			yyVAL.sli = NewInterfaceSlice(yyDollar[2].val, yyDollar[3].gen)
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gmask.y:156
		{
			yyVAL.sli = NewInterfaceSlice(yyDollar[2].gen, yyDollar[3].gen)
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gmask.y:157
		{
			yyVAL.sli = append(yyDollar[1].sli, yyDollar[3].val)
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:160
		{
			yyVAL.sli = NewInterfaceSlice(yyDollar[2].val)
		}
	case 70:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:161
		{
			yyVAL.sli = NewInterfaceSlice(yyDollar[2].gen)
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gmask.y:162
		{
			yyVAL.sli = NewInterfaceSlice(yyDollar[2].val, yyDollar[3].val)
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gmask.y:163
		{
			yyVAL.sli = NewInterfaceSlice(yyDollar[2].val, yyDollar[3].gen)
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gmask.y:164
		{
			yyVAL.sli = NewInterfaceSlice(yyDollar[2].gen, yyDollar[3].val)
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gmask.y:165
		{
			yyVAL.sli = NewInterfaceSlice(yyDollar[2].gen, yyDollar[3].gen)
		}
	case 75:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gmask.y:166
		{
			yyVAL.sli = NewInterfaceSlice(yyDollar[2].val, yyDollar[3].val, yyDollar[4].val)
		}
	case 76:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gmask.y:167
		{
			yyVAL.sli = NewInterfaceSlice(yyDollar[2].val, yyDollar[3].val, yyDollar[4].gen)
		}
	case 77:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gmask.y:168
		{
			yyVAL.sli = NewInterfaceSlice(yyDollar[2].val, yyDollar[3].gen, yyDollar[4].val)
		}
	case 78:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gmask.y:169
		{
			yyVAL.sli = NewInterfaceSlice(yyDollar[2].val, yyDollar[3].gen, yyDollar[4].gen)
		}
	case 79:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gmask.y:170
		{
			yyVAL.sli = NewInterfaceSlice(yyDollar[2].gen, yyDollar[3].val, yyDollar[4].val)
		}
	case 80:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gmask.y:171
		{
			yyVAL.sli = NewInterfaceSlice(yyDollar[2].gen, yyDollar[3].val, yyDollar[4].gen)
		}
	case 81:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gmask.y:172
		{
			yyVAL.sli = NewInterfaceSlice(yyDollar[2].gen, yyDollar[3].gen, yyDollar[4].val)
		}
	case 82:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gmask.y:173
		{
			yyVAL.sli = NewInterfaceSlice(yyDollar[2].gen, yyDollar[3].gen, yyDollar[4].gen)
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:176
		{
			yyVAL.amd = gmasklib.LIMIT
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:177
		{
			yyVAL.amd = gmasklib.MIRROR
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:178
		{
			yyVAL.amd = gmasklib.WRAP
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gmask.y:181
		{
			yyVAL.sli = NewInterfaceSlice(yyDollar[1].amd, yyDollar[2].val, yyDollar[3].val)
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gmask.y:182
		{
			yyVAL.sli = NewInterfaceSlice(yyDollar[1].amd, yyDollar[2].gen, yyDollar[3].val)
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gmask.y:183
		{
			yyVAL.sli = NewInterfaceSlice(yyDollar[1].amd, yyDollar[2].val, yyDollar[3].gen)
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gmask.y:184
		{
			yyVAL.sli = NewInterfaceSlice(yyDollar[1].amd, yyDollar[2].gen, yyDollar[3].gen)
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gmask.y:185
		{
			yyVAL.sli = append(yyDollar[1].sli, yyDollar[3].val)
		}
	case 91:
		yyDollar = yyS[yypt-0 : yypt+1]
//line gmask.y:188
		{
			yyVAL.n = -1
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gmask.y:189
		{
			yyVAL.n = int(yyDollar[2].val)
		}
	}
	goto yystack /* stack new state and value */
}
