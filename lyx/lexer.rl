package lyx

import (
	"bufio"
	"fmt"
)

%%{ 
    machine lexer;
    write data;
    access lex.;
    variable p lex.p;
    variable pe lex.pe;
}%%

type Lexer struct {
	data         []byte
	p, pe, cs    int
	ts, te, act  int

	result []string
}

func NewLexer(data []byte) *Lexer {
    lex := &Lexer{ 
        data: data,
        pe: len(data),
    }
    %% write init;
    return lex
}


func Run(scanner *bufio.Scanner) {
	for scanner.Scan() {
		l := scanner.Text()

		tt := NewLexer([]byte(l))
		ySym := new(yySymType)
		for {
			v := tt.Lex(ySym)
			if v == 0 {
				fmt.Println("end")
                break;
			} else {
				fmt.Printf("token type %d\n", v)
			}
		}
	}
}


func (l *Lexer) Error(msg string) {
	println(msg)
}


func (lex *Lexer) Lex(lval *yySymType) int {
    eof := lex.pe
    var tok int

    %%{
        # /* digit = [0-9] ; already defined */


#        xcstart		=	\/\*{op_chars}*;
#        xcstop		=	\*+\/;
#        xcinside	=	[^*/]+;

        ident_start	=	[A-Za-z\200-\377_];
        ident_cont	=	[A-Za-z\200-\377_0-9$];

        identifier	=	ident_start ident_cont*;


#        space		=	[ \t\n\r\f];
        horiz_space	= [ \t\f];
        newline		=	[\n\r];
#        non_newline	=	[^\n\r];

#        comment		=	("--"{non_newline}*);


#       whitespace	=	({space}+|{comment});
        whitespace	=	space+;


        op_chars	=	( '~' | '!' | '@' | '#' | '^' | '&' | '|' | '`' | '?' | '+' | '-' | '*' | '\\' | '%' | '<' | '>' | '=' ) ;
        operator	=	op_chars+;

        integer = digit+;

        sconst = '\'' ident_cont* '\'';
        
        main := |*
            whitespace => { /* do nothing */ };
            # integer const is string const 
            integer =>  { lval.str = string(lex.data[lex.ts:lex.te]); tok = SCONST; fbreak;};

            /select/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = SELECT; fbreak;};
            /update/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = UPDATE; fbreak;};
            /delete/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = DELETE; fbreak;};
            /table/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TABLE; fbreak;};
            /set/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = SET; fbreak;};
            /from/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = FROM; fbreak;};
            /where/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = WHERE; fbreak;};
            /order/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ORDER; fbreak;};
            /by/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = BY; fbreak;};
            /as/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = AS; fbreak;};
            /group/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = GROUP; fbreak;};


            identifier      => { lval.str = string(lex.data[lex.ts:lex.te]); tok = int(IDENT); fbreak;};
            sconst      => { lval.str = string(lex.data[lex.ts:lex.te]); tok = int(SCONST); fbreak;};


#           self		=	(',' | '(' | ')' | '[' | ']' | '.' | ';'| ':' | '+' | '-' | '*' | '\\' | '%' | '^' | '<' | '>' | '=');

            ',' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = int(TCOMMA); fbreak;};
            '(' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = int(TOPENBR); fbreak;};
            ')' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = int(TCLOSEBR); fbreak;};
            '[' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = int(TSQCLOSEBR); fbreak;};
            ']' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = int(TSQCLOSEBR); fbreak;};
            '.' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TDOT; fbreak;};
            ';' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = int(TSEMICOLON); fbreak;};
            ':' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = int(TCOLON); fbreak;};
            '+' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = int(TPLUS); fbreak;};
            '-' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = int(TMINUS); fbreak;};
            '*' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = int(TMUL); fbreak;};
           # TODO: support '\\' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = int(TMUL); fbreak;};
            '%' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = int(TMOD); fbreak;};
            '^' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = int(TPOW); fbreak;};
            '<' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = int(TLESS); fbreak;};
            '>' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = int(TGREATER); fbreak;};
            '=' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TEQ; fbreak;};

            '<>' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TNOT_EQUALS; fbreak;};
            '<=' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TLESS_EQUALS; fbreak;};
            '>=' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TGREATER_EQUALS; fbreak;};
            '!=' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TNOT_EQUALS; fbreak;};


            operator => { 



                lval.str = string(lex.data[lex.ts:lex.te]); tok = int(OP);    
                fbreak;
            };

#            self => {
#
#            }
        *|;

        write exec;
    }%%

    return int(tok);
}