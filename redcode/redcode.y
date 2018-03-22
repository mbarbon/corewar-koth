%{
package redcode
%}

%union {
    number     int
    identifier string
    lines      []Instruction
    line       Instruction
    operand    Operand
    opcode     Opcode
    expression *Expression
    comment    string
}

%type <line> line instruction
%type <lines> lines
%type <identifier> label
%type <opcode> opcode
%type <operand> operand
%type <expression> expression
%type <number> number

%token <number> NUMBER
%token <identifier> IDENTIFIER
%token <comment> COMMENT
%token COMMA NEWLINE PLUS MINUS EOF
%token OPDAT OPMOV OPADD OPSUB OPJMP OPJMZ OPDJN OPCMP OPSPL OPEND
%token ADDRIMMEDIATE ADDRDIRECT ADDRINDIRECT ADDRDECREMENT

%%

lines: /* empty */
         { $$ = yylex.(*lexer).instructions }
     | lines line_end
         { $$ = yylex.(*lexer).instructions }
     | lines line     
         { $$ = append($1, $2); yylex.(*lexer).instructions = $$ }
     | lines COMMENT
         { $$ = yylex.(*lexer).instructions; parseDirective(yylex.(*lexer), $2) }
     ;

line: label instruction maybe_comment line_end
        { $$ = $2; $$.Label = $1 }
    | instruction maybe_comment line_end
        { $$ = $1 }
    ;

maybe_comment: /* empty */ | COMMENT;

line_end: NEWLINE | EOF;

label: IDENTIFIER;

instruction: opcode operand
               { $$ = Instruction{ Opcode: $1, A: $2 } }
           | opcode operand COMMA operand
               { $$ = Instruction{ Opcode: $1, A: $2, B: $4 } }
           ;

opcode: OPDAT { $$ = OpDat }
      | OPMOV { $$ = OpMov }
      | OPADD { $$ = OpAdd }
      | OPSUB { $$ = OpSub }
      | OPJMP { $$ = OpJmp }
      | OPJMZ { $$ = OpJmz }
      | OPDJN { $$ = OpDjn }
      | OPCMP { $$ = OpCmp }
      | OPSPL { $$ = OpSpl }
      | OPEND { $$ = OpEnd }
      ;

operand: expression
           { $$ = Operand{ Mode: Relative, Expression: $1 } }
       | ADDRDIRECT expression
           { $$ = Operand{ Mode: Relative, Expression: $2 } }
       | ADDRIMMEDIATE expression
           { $$ = Operand{ Mode: Immediate, Expression: $2 } }
       | ADDRINDIRECT expression
           { $$ = Operand{ Mode: Indirect, Expression: $2 } }
       | ADDRDECREMENT expression
           { $$ = Operand{ Mode: DecrementIndirect, Expression: $2 } }
       ;

expression: number
              { $$ = &Expression{ Operation: Number, Number: $1 } }
          | IDENTIFIER
              { $$ = &Expression{ Operation: Label, Label: $1 } }
          | expression PLUS expression
              { $$ = &Expression{ Operation: Add, Left: $1, Right: $3 } }
          | expression MINUS expression
              { $$ = &Expression{ Operation: Sub, Left: $1, Right: $3 } }
          ;

number: NUMBER
        { $$ = $1 }
      | PLUS NUMBER
        { $$ = $2 }
      | MINUS NUMBER
        { $$ = -$2 }
      ;