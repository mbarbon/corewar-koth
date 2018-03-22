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
	comment    string
}

const NUMBER = 57346
const IDENTIFIER = 57347
const COMMENT = 57348
const COMMA = 57349
const NEWLINE = 57350
const PLUS = 57351
const MINUS = 57352
const EOF = 57353
const OPDAT = 57354
const OPMOV = 57355
const OPADD = 57356
const OPSUB = 57357
const OPJMP = 57358
const OPJMZ = 57359
const OPDJN = 57360
const OPCMP = 57361
const OPSPL = 57362
const OPEND = 57363
const ADDRIMMEDIATE = 57364
const ADDRDIRECT = 57365
const ADDRINDIRECT = 57366
const ADDRDECREMENT = 57367

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"NUMBER",
	"IDENTIFIER",
	"COMMENT",
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
	"OPDJN",
	"OPCMP",
	"OPSPL",
	"OPEND",
	"ADDRIMMEDIATE",
	"ADDRDIRECT",
	"ADDRINDIRECT",
	"ADDRDECREMENT",
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

const yyLast = 71

var yyAct = [...]int{

	25, 2, 37, 24, 9, 4, 23, 5, 45, 22,
	6, 11, 12, 13, 14, 15, 16, 17, 18, 19,
	20, 44, 5, 30, 36, 6, 8, 40, 41, 42,
	43, 35, 38, 39, 21, 10, 7, 46, 1, 48,
	49, 47, 32, 31, 3, 0, 0, 33, 34, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 0,
	27, 26, 28, 29, 32, 31, 0, 0, 0, 33,
	34,
}
var yyPact = [...]int{

	-1000, -1, -1000, -1000, -1000, -1000, -1000, 37, 0, -1000,
	38, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 0, 14, -1000, -5, 23, 60, 60, 60, 60,
	-1000, -1000, -1000, 17, 4, 14, -1000, 38, 60, 60,
	23, 23, 23, 23, -1000, -1000, -1000, -1000, 23, 23,
}
var yyPgo = [...]int{

	0, 44, 26, 38, 36, 35, 3, 0, 23, 1,
	9,
}
var yyR1 = [...]int{

	0, 3, 3, 3, 3, 1, 1, 10, 10, 9,
	9, 4, 2, 2, 5, 5, 5, 5, 5, 5,
	5, 5, 5, 5, 6, 6, 6, 6, 6, 7,
	7, 7, 7, 8, 8, 8,
}
var yyR2 = [...]int{

	0, 0, 2, 2, 2, 4, 3, 0, 1, 1,
	1, 1, 2, 4, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 2, 2, 2, 1,
	1, 3, 3, 1, 2, 2,
}
var yyChk = [...]int{

	-1000, -3, -9, -1, 6, 8, 11, -4, -2, 5,
	-5, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	21, -2, -10, 6, -6, -7, 23, 22, 24, 25,
	-8, 5, 4, 9, 10, -10, -9, 7, 9, 10,
	-7, -7, -7, -7, 4, 4, -9, -6, -7, -7,
}
var yyDef = [...]int{

	1, -2, 2, 3, 4, 9, 10, 0, 7, 11,
	0, 14, 15, 16, 17, 18, 19, 20, 21, 22,
	23, 7, 0, 8, 12, 24, 0, 0, 0, 0,
	29, 30, 33, 0, 0, 0, 6, 0, 0, 0,
	25, 26, 27, 28, 34, 35, 5, 13, 31, 32,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25,
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
		//line redcode.y:34
		{
			yyVAL.lines = yylex.(*lexer).instructions
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line redcode.y:36
		{
			yyVAL.lines = yylex.(*lexer).instructions
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line redcode.y:38
		{
			yyVAL.lines = append(yyDollar[1].lines, yyDollar[2].line)
			yylex.(*lexer).instructions = yyVAL.lines
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line redcode.y:40
		{
			yyVAL.lines = yylex.(*lexer).instructions
			parseDirective(yylex.(*lexer), yyDollar[2].comment)
		}
	case 5:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line redcode.y:44
		{
			yyVAL.line = yyDollar[2].line
			yyVAL.line.Label = yyDollar[1].identifier
		}
	case 6:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line redcode.y:46
		{
			yyVAL.line = yyDollar[1].line
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line redcode.y:56
		{
			yyVAL.line = Instruction{Opcode: yyDollar[1].opcode, A: yyDollar[2].operand}
		}
	case 13:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line redcode.y:58
		{
			yyVAL.line = Instruction{Opcode: yyDollar[1].opcode, A: yyDollar[2].operand, B: yyDollar[4].operand}
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line redcode.y:61
		{
			yyVAL.opcode = OpDat
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line redcode.y:62
		{
			yyVAL.opcode = OpMov
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line redcode.y:63
		{
			yyVAL.opcode = OpAdd
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line redcode.y:64
		{
			yyVAL.opcode = OpSub
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line redcode.y:65
		{
			yyVAL.opcode = OpJmp
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line redcode.y:66
		{
			yyVAL.opcode = OpJmz
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line redcode.y:67
		{
			yyVAL.opcode = OpDjn
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line redcode.y:68
		{
			yyVAL.opcode = OpCmp
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line redcode.y:69
		{
			yyVAL.opcode = OpSpl
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line redcode.y:70
		{
			yyVAL.opcode = OpEnd
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line redcode.y:74
		{
			yyVAL.operand = Operand{Mode: Relative, Expression: yyDollar[1].expression}
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line redcode.y:76
		{
			yyVAL.operand = Operand{Mode: Relative, Expression: yyDollar[2].expression}
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line redcode.y:78
		{
			yyVAL.operand = Operand{Mode: Immediate, Expression: yyDollar[2].expression}
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line redcode.y:80
		{
			yyVAL.operand = Operand{Mode: Indirect, Expression: yyDollar[2].expression}
		}
	case 28:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line redcode.y:82
		{
			yyVAL.operand = Operand{Mode: DecrementIndirect, Expression: yyDollar[2].expression}
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line redcode.y:86
		{
			yyVAL.expression = &Expression{Operation: Number, Number: yyDollar[1].number}
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line redcode.y:88
		{
			yyVAL.expression = &Expression{Operation: Label, Label: yyDollar[1].identifier}
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line redcode.y:90
		{
			yyVAL.expression = &Expression{Operation: Add, Left: yyDollar[1].expression, Right: yyDollar[3].expression}
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line redcode.y:92
		{
			yyVAL.expression = &Expression{Operation: Sub, Left: yyDollar[1].expression, Right: yyDollar[3].expression}
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line redcode.y:96
		{
			yyVAL.number = yyDollar[1].number
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line redcode.y:98
		{
			yyVAL.number = yyDollar[2].number
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line redcode.y:100
		{
			yyVAL.number = -yyDollar[2].number
		}
	}
	goto yystack /* stack new state and value */
}
