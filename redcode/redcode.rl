package redcode

import (
    "fmt"
    "strconv"
)

%%{ 
    machine redcodeLexer;
    write data;
    access lex.;
    variable p lex.p;
    variable pe lex.pe;
}%%

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
    %% write init;
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

    %%{ 
        main := |*
            /dat/i => { tok = OPDAT; fbreak; };
            /mov/i => { tok = OPMOV; fbreak; };
            /add/i => { tok = OPADD; fbreak; };
            /sub/i => { tok = OPSUB; fbreak; };
            /jmp/i => { tok = OPJMP; fbreak; };
            /jmz/i => { tok = OPJMZ; fbreak; };
            /djn/i => { tok = OPDJN; fbreak; };
            /cmp/i => { tok = OPCMP; fbreak; };
            /spl/i => { tok = OPSPL; fbreak; };
            /end/i => { tok = OPEND; lex.seen_end = true; fbreak; };
            ',' { tok = COMMA; fbreak; };
            '-' { tok = MINUS; fbreak; };
            '+' { tok = PLUS; fbreak; };
            '#' { tok = ADDRIMMEDIATE; fbreak; };
            '$' { tok = ADDRDIRECT; fbreak; };
            '@' { tok = ADDRINDIRECT; fbreak; };
            '<' { tok = ADDRDECREMENT; fbreak; };
            digit+ => { out.number, _ = strconv.Atoi(string(lex.data[lex.ts:lex.te])); tok = NUMBER; fbreak; };
            [a-zA-Z]+ => { out.identifier = string(lex.data[lex.ts:lex.te]); tok = IDENTIFIER; fbreak; };
            '\r\n'|'\r'|'\n' => { lex.line++; tok = NEWLINE; lex.force_eof = lex.seen_end; fbreak; };
            /[ \t]/;
            ';' [^\r\n]* => { tok = COMMENT; out.comment = string(lex.data[lex.ts+1:lex.te]); fbreak; };
        *|;

         write exec;
    }%%

    if tok == 0 && lex.p == lex.pe && !lex.emitted_eof {
        lex.emitted_eof = true
        return EOF
    }

    return tok;
}

func (lex *lexer) Error(e string) {
    lex.err = fmt.Errorf("%s in %s at line %d", e, lex.filename, lex.line)
}