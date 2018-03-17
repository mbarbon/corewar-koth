//line redcode.y:2
package redcode

import __yyfmt__ "fmt"

//line redcode.y:2
//line redcode.y:5
type yySymType struct {
	yys        int
	number     int
	identifier string
	lines      []Instruction
	line       Instruction
	operand    Operand
	opcode     Opcode
	expression *Expression
}

const NUMBER = 57346
const IDENTIFIER = 57347
const COMMA = 57348
const NEWLINE = 57349
const PLUS = 57350
const MINUS = 57351
const EOF = 57352
const OPDAT = 57353
const OPMOV = 57354
const OPADD = 57355
const OPSUB = 57356
const OPJMP = 57357
const OPJMZ = 57358
const OPDJZ = 57359
const OPCMP = 57360
const OPSPL = 57361
const OPEND = 57362
const ADDRIMMEDIATE = 57363
const ADDRDIRECT = 57364
const ADDRINDIRECT = 57365

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"NUMBER",
	"IDENTIFIER",
	"COMMA",
	"NEWLINE",
	"PLUS",
	"MINUS",
	"EOF",
	"OPDAT",
	"OPMOV",
	"OPADD",
	"OPSUB",
	"OPJMP",
	"OPJMZ",
	"OPDJZ",
	"OPCMP",
	"OPSPL",
	"OPEND",
	"ADDRIMMEDIATE",
	"ADDRDIRECT",
	"ADDRINDIRECT",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 59

var yyAct = [...]int{

	23, 33, 8, 22, 4, 40, 39, 5, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 4, 29,
	28, 5, 2, 30, 31, 36, 37, 38, 29, 28,
	21, 27, 30, 31, 7, 42, 43, 41, 34, 35,
	9, 20, 6, 32, 1, 25, 24, 26, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 3,
}
var yyPact = [...]int{

	-1000, -3, -1000, -1000, -1000, -1000, 37, 11, -1000, 24,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	11, -1000, -5, 30, 15, 15, 15, -1000, -1000, -1000,
	2, 1, -1000, 24, 15, 15, 30, 30, 30, -1000,
	-1000, -1000, 30, 30,
}
var yyPgo = [...]int{

	0, 58, 34, 44, 42, 40, 3, 0, 31, 22,
}
var yyR1 = [...]int{

	0, 3, 3, 3, 1, 1, 9, 9, 4, 2,
	2, 5, 5, 5, 5, 5, 5, 5, 5, 5,
	5, 6, 6, 6, 6, 7, 7, 7, 7, 8,
	8, 8,
}
var yyR2 = [...]int{

	0, 0, 2, 2, 3, 2, 1, 1, 1, 2,
	4, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 2, 2, 2, 1, 1, 3, 3, 1,
	2, 2,
}
var yyChk = [...]int{

	-1000, -3, -9, -1, 7, 10, -4, -2, 5, -5,
	11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	-2, -9, -6, -7, 22, 21, 23, -8, 5, 4,
	8, 9, -9, 6, 8, 9, -7, -7, -7, 4,
	4, -6, -7, -7,
}
var yyDef = [...]int{

	1, -2, 2, 3, 6, 7, 0, 0, 8, 0,
	11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	0, 5, 9, 21, 0, 0, 0, 25, 26, 29,
	0, 0, 4, 0, 0, 0, 22, 23, 24, 30,
	31, 10, 27, 28,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23,
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
		yyDollar = yyS[yypt-0 : yypt+1]
		//line redcode.y:32
		{
			yyVAL.lines = yylex.(*lexer).instructions
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line redcode.y:34
		{
			yyVAL.lines = yylex.(*lexer).instructions
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line redcode.y:36
		{
			yyVAL.lines = append(yyDollar[1].lines, yyDollar[2].line)
			yylex.(*lexer).instructions = yyVAL.lines
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line redcode.y:40
		{
			yyVAL.line = yyDollar[2].line
			yyVAL.line.Label = yyDollar[1].identifier
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line redcode.y:42
		{
			yyVAL.line = yyDollar[1].line
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line redcode.y:50
		{
			yyVAL.line = Instruction{Opcode: yyDollar[1].opcode, A: yyDollar[2].operand}
		}
	case 10:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line redcode.y:52
		{
			yyVAL.line = Instruction{Opcode: yyDollar[1].opcode, A: yyDollar[2].operand, B: yyDollar[4].operand}
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line redcode.y:55
		{
			yyVAL.opcode = OpDat
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line redcode.y:56
		{
			yyVAL.opcode = OpMov
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line redcode.y:57
		{
			yyVAL.opcode = OpAdd
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line redcode.y:58
		{
			yyVAL.opcode = OpSub
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line redcode.y:59
		{
			yyVAL.opcode = OpJmp
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line redcode.y:60
		{
			yyVAL.opcode = OpJmz
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line redcode.y:61
		{
			yyVAL.opcode = OpDjz
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line redcode.y:62
		{
			yyVAL.opcode = OpCmp
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line redcode.y:63
		{
			yyVAL.opcode = OpSpl
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line redcode.y:64
		{
			yyVAL.opcode = OpEnd
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line redcode.y:68
		{
			yyVAL.operand = Operand{Mode: Relative, Expression: yyDollar[1].expression}
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line redcode.y:70
		{
			yyVAL.operand = Operand{Mode: Relative, Expression: yyDollar[2].expression}
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line redcode.y:72
		{
			yyVAL.operand = Operand{Mode: Immediate, Expression: yyDollar[2].expression}
		}
	case 24:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line redcode.y:74
		{
			yyVAL.operand = Operand{Mode: Indirect, Expression: yyDollar[2].expression}
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line redcode.y:78
		{
			yyVAL.expression = &Expression{Operation: Number, Number: yyDollar[1].number}
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line redcode.y:80
		{
			yyVAL.expression = &Expression{Operation: Label, Label: yyDollar[1].identifier}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line redcode.y:82
		{
			yyVAL.expression = &Expression{Operation: Add, Left: yyDollar[1].expression, Right: yyDollar[3].expression}
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line redcode.y:84
		{
			yyVAL.expression = &Expression{Operation: Sub, Left: yyDollar[1].expression, Right: yyDollar[3].expression}
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line redcode.y:88
		{
			yyVAL.number = yyDollar[1].number
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line redcode.y:90
		{
			yyVAL.number = yyDollar[2].number
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line redcode.y:92
		{
			yyVAL.number = -yyDollar[2].number
		}
	}
	goto yystack /* stack new state and value */
}
