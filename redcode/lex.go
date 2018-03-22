
//line redcode.rl:1
package redcode

import (
    "fmt"
    "strconv"
)


//line lex.go:12
const redcodeLexer_start int = 1
const redcodeLexer_first_final int = 1
const redcodeLexer_error int = 0

const redcodeLexer_en_main int = 1


//line redcode.rl:14


type lexer struct {
    data                             []byte
    p, pe, cs                        int
    ts, te, act                      int
    instructions                     []Instruction
    directives                       Directives
    emitted_eof, seen_end, force_eof bool
    err                              error
    filename                         string
    line                             int
}

func newLexer(data []byte, filename string) *lexer {
    lex := &lexer{ 
        data:       data,
        directives: make(Directives),
        pe:         len(data),
        filename:   filename,
        line:       1,
    }
    
//line lex.go:44
	{
	 lex.cs = redcodeLexer_start
	 lex.ts = 0
	 lex.te = 0
	 lex.act = 0
	}

//line redcode.rl:37
    return lex
}

func (lex *lexer) Lex(out *yySymType) int {
    eof := lex.pe
    tok := 0

    if lex.force_eof {
        if lex.emitted_eof {
            return tok
        } else {
            lex.emitted_eof = true
            return EOF
        }
    }

    
//line lex.go:70
	{
	if ( lex.p) == ( lex.pe) {
		goto _test_eof
	}
	switch  lex.cs {
	case 1:
		goto st_case_1
	case 0:
		goto st_case_0
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	}
	goto st_out
tr0:
//line redcode.rl:75
 lex.te = ( lex.p)+1

	goto st1
tr2:
//line redcode.rl:74
 lex.te = ( lex.p)+1
{ lex.line++; tok = NEWLINE; lex.force_eof = lex.seen_end; {( lex.p)++;  lex.cs = 1; goto _out } }
	goto st1
tr4:
//line redcode.rl:68
 lex.te = ( lex.p)+1
{ tok = ADDRIMMEDIATE; {( lex.p)++;  lex.cs = 1; goto _out } }
	goto st1
tr5:
//line redcode.rl:69
 lex.te = ( lex.p)+1
{ tok = ADDRDIRECT; {( lex.p)++;  lex.cs = 1; goto _out } }
	goto st1
tr6:
//line redcode.rl:67
 lex.te = ( lex.p)+1
{ tok = PLUS; {( lex.p)++;  lex.cs = 1; goto _out } }
	goto st1
tr7:
//line redcode.rl:65
 lex.te = ( lex.p)+1
{ tok = COMMA; {( lex.p)++;  lex.cs = 1; goto _out } }
	goto st1
tr8:
//line redcode.rl:66
 lex.te = ( lex.p)+1
{ tok = MINUS; {( lex.p)++;  lex.cs = 1; goto _out } }
	goto st1
tr11:
//line redcode.rl:71
 lex.te = ( lex.p)+1
{ tok = ADDRDECREMENT; {( lex.p)++;  lex.cs = 1; goto _out } }
	goto st1
tr12:
//line redcode.rl:70
 lex.te = ( lex.p)+1
{ tok = ADDRINDIRECT; {( lex.p)++;  lex.cs = 1; goto _out } }
	goto st1
tr21:
//line redcode.rl:74
 lex.te = ( lex.p)
( lex.p)--
{ lex.line++; tok = NEWLINE; lex.force_eof = lex.seen_end; {( lex.p)++;  lex.cs = 1; goto _out } }
	goto st1
tr22:
//line redcode.rl:72
 lex.te = ( lex.p)
( lex.p)--
{ out.number, _ = strconv.Atoi(string(lex.data[lex.ts:lex.te])); tok = NUMBER; {( lex.p)++;  lex.cs = 1; goto _out } }
	goto st1
tr23:
//line redcode.rl:76
 lex.te = ( lex.p)
( lex.p)--
{ tok = COMMENT; out.comment = string(lex.data[lex.ts+1:lex.te]); {( lex.p)++;  lex.cs = 1; goto _out } }
	goto st1
tr24:
//line redcode.rl:73
 lex.te = ( lex.p)
( lex.p)--
{ out.identifier = string(lex.data[lex.ts:lex.te]); tok = IDENTIFIER; {( lex.p)++;  lex.cs = 1; goto _out } }
	goto st1
tr26:
//line NONE:1
	switch  lex.act {
	case 1:
	{( lex.p) = ( lex.te) - 1
 tok = OPDAT; {( lex.p)++;  lex.cs = 1; goto _out } }
	case 2:
	{( lex.p) = ( lex.te) - 1
 tok = OPMOV; {( lex.p)++;  lex.cs = 1; goto _out } }
	case 3:
	{( lex.p) = ( lex.te) - 1
 tok = OPADD; {( lex.p)++;  lex.cs = 1; goto _out } }
	case 4:
	{( lex.p) = ( lex.te) - 1
 tok = OPSUB; {( lex.p)++;  lex.cs = 1; goto _out } }
	case 5:
	{( lex.p) = ( lex.te) - 1
 tok = OPJMP; {( lex.p)++;  lex.cs = 1; goto _out } }
	case 6:
	{( lex.p) = ( lex.te) - 1
 tok = OPJMZ; {( lex.p)++;  lex.cs = 1; goto _out } }
	case 7:
	{( lex.p) = ( lex.te) - 1
 tok = OPDJN; {( lex.p)++;  lex.cs = 1; goto _out } }
	case 8:
	{( lex.p) = ( lex.te) - 1
 tok = OPCMP; {( lex.p)++;  lex.cs = 1; goto _out } }
	case 9:
	{( lex.p) = ( lex.te) - 1
 tok = OPSPL; {( lex.p)++;  lex.cs = 1; goto _out } }
	case 10:
	{( lex.p) = ( lex.te) - 1
 tok = OPEND; lex.seen_end = true; {( lex.p)++;  lex.cs = 1; goto _out } }
	case 19:
	{( lex.p) = ( lex.te) - 1
 out.identifier = string(lex.data[lex.ts:lex.te]); tok = IDENTIFIER; {( lex.p)++;  lex.cs = 1; goto _out } }
	}
	
	goto st1
	st1:
//line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof1
		}
	st_case_1:
//line NONE:1
 lex.ts = ( lex.p)

//line lex.go:241
		switch  lex.data[( lex.p)] {
		case 9:
			goto tr0
		case 10:
			goto tr2
		case 13:
			goto st2
		case 32:
			goto tr0
		case 35:
			goto tr4
		case 36:
			goto tr5
		case 43:
			goto tr6
		case 44:
			goto tr7
		case 45:
			goto tr8
		case 59:
			goto st4
		case 60:
			goto tr11
		case 64:
			goto tr12
		case 65:
			goto st5
		case 67:
			goto st8
		case 68:
			goto st10
		case 69:
			goto st13
		case 74:
			goto st15
		case 77:
			goto st17
		case 83:
			goto st19
		case 97:
			goto st5
		case 99:
			goto st8
		case 100:
			goto st10
		case 101:
			goto st13
		case 106:
			goto st15
		case 109:
			goto st17
		case 115:
			goto st19
		}
		switch {
		case  lex.data[( lex.p)] < 66:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto st3
			}
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr14
			}
		default:
			goto tr14
		}
		goto st0
st_case_0:
	st0:
		 lex.cs = 0
		goto _out
	st2:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof2
		}
	st_case_2:
		if  lex.data[( lex.p)] == 10 {
			goto tr2
		}
		goto tr21
	st3:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof3
		}
	st_case_3:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st3
		}
		goto tr22
	st4:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof4
		}
	st_case_4:
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr23
		case 13:
			goto tr23
		}
		goto st4
	st5:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof5
		}
	st_case_5:
		switch  lex.data[( lex.p)] {
		case 68:
			goto st7
		case 100:
			goto st7
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr14
			}
		case  lex.data[( lex.p)] >= 65:
			goto tr14
		}
		goto tr24
tr14:
//line NONE:1
 lex.te = ( lex.p)+1

//line redcode.rl:73
 lex.act = 19;
	goto st6
tr27:
//line NONE:1
 lex.te = ( lex.p)+1

//line redcode.rl:57
 lex.act = 3;
	goto st6
tr29:
//line NONE:1
 lex.te = ( lex.p)+1

//line redcode.rl:62
 lex.act = 8;
	goto st6
tr32:
//line NONE:1
 lex.te = ( lex.p)+1

//line redcode.rl:55
 lex.act = 1;
	goto st6
tr33:
//line NONE:1
 lex.te = ( lex.p)+1

//line redcode.rl:61
 lex.act = 7;
	goto st6
tr35:
//line NONE:1
 lex.te = ( lex.p)+1

//line redcode.rl:64
 lex.act = 10;
	goto st6
tr37:
//line NONE:1
 lex.te = ( lex.p)+1

//line redcode.rl:59
 lex.act = 5;
	goto st6
tr38:
//line NONE:1
 lex.te = ( lex.p)+1

//line redcode.rl:60
 lex.act = 6;
	goto st6
tr40:
//line NONE:1
 lex.te = ( lex.p)+1

//line redcode.rl:56
 lex.act = 2;
	goto st6
tr43:
//line NONE:1
 lex.te = ( lex.p)+1

//line redcode.rl:63
 lex.act = 9;
	goto st6
tr44:
//line NONE:1
 lex.te = ( lex.p)+1

//line redcode.rl:58
 lex.act = 4;
	goto st6
	st6:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof6
		}
	st_case_6:
//line lex.go:445
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr14
			}
		case  lex.data[( lex.p)] >= 65:
			goto tr14
		}
		goto tr26
	st7:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof7
		}
	st_case_7:
		switch  lex.data[( lex.p)] {
		case 68:
			goto tr27
		case 100:
			goto tr27
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr14
			}
		case  lex.data[( lex.p)] >= 65:
			goto tr14
		}
		goto tr24
	st8:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof8
		}
	st_case_8:
		switch  lex.data[( lex.p)] {
		case 77:
			goto st9
		case 109:
			goto st9
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr14
			}
		case  lex.data[( lex.p)] >= 65:
			goto tr14
		}
		goto tr24
	st9:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof9
		}
	st_case_9:
		switch  lex.data[( lex.p)] {
		case 80:
			goto tr29
		case 112:
			goto tr29
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr14
			}
		case  lex.data[( lex.p)] >= 65:
			goto tr14
		}
		goto tr24
	st10:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof10
		}
	st_case_10:
		switch  lex.data[( lex.p)] {
		case 65:
			goto st11
		case 74:
			goto st12
		case 97:
			goto st11
		case 106:
			goto st12
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr14
			}
		case  lex.data[( lex.p)] >= 66:
			goto tr14
		}
		goto tr24
	st11:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof11
		}
	st_case_11:
		switch  lex.data[( lex.p)] {
		case 84:
			goto tr32
		case 116:
			goto tr32
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr14
			}
		case  lex.data[( lex.p)] >= 65:
			goto tr14
		}
		goto tr24
	st12:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof12
		}
	st_case_12:
		switch  lex.data[( lex.p)] {
		case 78:
			goto tr33
		case 110:
			goto tr33
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr14
			}
		case  lex.data[( lex.p)] >= 65:
			goto tr14
		}
		goto tr24
	st13:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof13
		}
	st_case_13:
		switch  lex.data[( lex.p)] {
		case 78:
			goto st14
		case 110:
			goto st14
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr14
			}
		case  lex.data[( lex.p)] >= 65:
			goto tr14
		}
		goto tr24
	st14:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof14
		}
	st_case_14:
		switch  lex.data[( lex.p)] {
		case 68:
			goto tr35
		case 100:
			goto tr35
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr14
			}
		case  lex.data[( lex.p)] >= 65:
			goto tr14
		}
		goto tr24
	st15:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof15
		}
	st_case_15:
		switch  lex.data[( lex.p)] {
		case 77:
			goto st16
		case 109:
			goto st16
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr14
			}
		case  lex.data[( lex.p)] >= 65:
			goto tr14
		}
		goto tr24
	st16:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof16
		}
	st_case_16:
		switch  lex.data[( lex.p)] {
		case 80:
			goto tr37
		case 90:
			goto tr38
		case 112:
			goto tr37
		case 122:
			goto tr38
		}
		switch {
		case  lex.data[( lex.p)] > 89:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 121 {
				goto tr14
			}
		case  lex.data[( lex.p)] >= 65:
			goto tr14
		}
		goto tr24
	st17:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof17
		}
	st_case_17:
		switch  lex.data[( lex.p)] {
		case 79:
			goto st18
		case 111:
			goto st18
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr14
			}
		case  lex.data[( lex.p)] >= 65:
			goto tr14
		}
		goto tr24
	st18:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof18
		}
	st_case_18:
		switch  lex.data[( lex.p)] {
		case 86:
			goto tr40
		case 118:
			goto tr40
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr14
			}
		case  lex.data[( lex.p)] >= 65:
			goto tr14
		}
		goto tr24
	st19:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof19
		}
	st_case_19:
		switch  lex.data[( lex.p)] {
		case 80:
			goto st20
		case 85:
			goto st21
		case 112:
			goto st20
		case 117:
			goto st21
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr14
			}
		case  lex.data[( lex.p)] >= 65:
			goto tr14
		}
		goto tr24
	st20:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof20
		}
	st_case_20:
		switch  lex.data[( lex.p)] {
		case 76:
			goto tr43
		case 108:
			goto tr43
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr14
			}
		case  lex.data[( lex.p)] >= 65:
			goto tr14
		}
		goto tr24
	st21:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof21
		}
	st_case_21:
		switch  lex.data[( lex.p)] {
		case 66:
			goto tr44
		case 98:
			goto tr44
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr14
			}
		case  lex.data[( lex.p)] >= 65:
			goto tr14
		}
		goto tr24
	st_out:
	_test_eof1:  lex.cs = 1; goto _test_eof
	_test_eof2:  lex.cs = 2; goto _test_eof
	_test_eof3:  lex.cs = 3; goto _test_eof
	_test_eof4:  lex.cs = 4; goto _test_eof
	_test_eof5:  lex.cs = 5; goto _test_eof
	_test_eof6:  lex.cs = 6; goto _test_eof
	_test_eof7:  lex.cs = 7; goto _test_eof
	_test_eof8:  lex.cs = 8; goto _test_eof
	_test_eof9:  lex.cs = 9; goto _test_eof
	_test_eof10:  lex.cs = 10; goto _test_eof
	_test_eof11:  lex.cs = 11; goto _test_eof
	_test_eof12:  lex.cs = 12; goto _test_eof
	_test_eof13:  lex.cs = 13; goto _test_eof
	_test_eof14:  lex.cs = 14; goto _test_eof
	_test_eof15:  lex.cs = 15; goto _test_eof
	_test_eof16:  lex.cs = 16; goto _test_eof
	_test_eof17:  lex.cs = 17; goto _test_eof
	_test_eof18:  lex.cs = 18; goto _test_eof
	_test_eof19:  lex.cs = 19; goto _test_eof
	_test_eof20:  lex.cs = 20; goto _test_eof
	_test_eof21:  lex.cs = 21; goto _test_eof

	_test_eof: {}
	if ( lex.p) == eof {
		switch  lex.cs {
		case 2:
			goto tr21
		case 3:
			goto tr22
		case 4:
			goto tr23
		case 5:
			goto tr24
		case 6:
			goto tr26
		case 7:
			goto tr24
		case 8:
			goto tr24
		case 9:
			goto tr24
		case 10:
			goto tr24
		case 11:
			goto tr24
		case 12:
			goto tr24
		case 13:
			goto tr24
		case 14:
			goto tr24
		case 15:
			goto tr24
		case 16:
			goto tr24
		case 17:
			goto tr24
		case 18:
			goto tr24
		case 19:
			goto tr24
		case 20:
			goto tr24
		case 21:
			goto tr24
		}
	}

	_out: {}
	}

//line redcode.rl:80


    if tok == 0 && lex.p == lex.pe && !lex.emitted_eof {
        lex.emitted_eof = true
        return EOF
    }

    return tok;
}

func (lex *lexer) Error(e string) {
    lex.err = fmt.Errorf("%s in %s at line %d", e, lex.filename, lex.line)
}