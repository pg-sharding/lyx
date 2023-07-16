package lyx

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

func ResetLexer(lex *Lexer, data []byte) {
    lex.pe = len(data)
    lex.data = data
    %% write init;
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

        qidentifier	=	'"' ident_start ident_cont* '"' ;


#        space		=	[ \t\n\r\f];
        horiz_space	= [ \t\f];
        newline		=	[\n\r];
#        non_newline	=	[^\n\r];

#        comment		=	("--"{non_newline}*);


        comment		= '/''*' (any - '*''/')* '*''/';


#       whitespace	=	({space}+|{comment});
        whitespace	=	space+;


        op_chars	=	( '~' | '!' | '@' | '#' | '^' | '&' | '|' | '`' | '?' | '+' | '-' | '*' | '\\' | '%' | '<' | '>' | '=' ) ;
        operator	=	op_chars+;

        integer = digit+;

        sconst = '\'' (any-'\'')* '\'';
        
        main := |*
            whitespace => { /* do nothing */ };
            # integer const is string const 
            comment => {/* nothing */};
            integer =>  { lval.str = string(lex.data[lex.ts:lex.te]); tok = SCONST; fbreak;};

            /select/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = SELECT; fbreak;};
            /insert/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = INSERT; fbreak;};
            /into/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = INTO; fbreak;};
            /values/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = VALUES; fbreak;};
            /update/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = UPDATE; fbreak;};
            /delete/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = DELETE; fbreak;};
            /create/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = CREATE; fbreak;};
            /table/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TABLE; fbreak;};
            /database/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = DATABASE; fbreak;};
            /role/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ROLE; fbreak;};
            /primary/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = PRIMARY; fbreak;};
            /foreign/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = FOREIGN; fbreak;};
            /references/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = REFERENCES; fbreak;};
            /key/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = KEY; fbreak;};
            /set/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = SET; fbreak;};
            /from/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = FROM; fbreak;};
            /where/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = WHERE; fbreak;};
            /order/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ORDER; fbreak;};
            /group/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = GROUP; fbreak;};
            /by/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = BY; fbreak;};
            /as/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = AS; fbreak;};
            /and/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = AND; fbreak;};
            /or/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = OR; fbreak;};

            /returning/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = RETURNING; fbreak;};
            /default/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = DEFAULT; fbreak;};
            
            /copy/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = COPY; fbreak;};
            /to/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TO; fbreak;};
            /stdout/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = STDOUT; fbreak;};

            /limit/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = LIMIT; fbreak;};
            
            /is/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IS; fbreak;};
            /isnull/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ISNULL; fbreak;};
            /null/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = NULL_P; fbreak;};
            /nulls/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = NULLS_LA; fbreak;};
            /not/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = NOT; fbreak;};
            /notnull/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = NOTNULL; fbreak;};
            /lateral/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = LATERAL_P; fbreak;};
            /ordinality/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ORDINALITY; fbreak;};
            /with/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = WITH_LA; fbreak;};
            /true/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TRUE_P; fbreak;};
            /false/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = FALSE_P; fbreak;};

            /FIRST/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = FIRST_P; fbreak;};
            /LAST/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = LAST_P; fbreak;};
            /ASC/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ASC; fbreak;};
            /DESC/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = DESC; fbreak;};


            /array/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ARRAY; fbreak;};


            /join/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = JOIN; fbreak;};
            /cross/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = CROSS; fbreak;};
            /full/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = FULL; fbreak;};
            /outer/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = OUTER_P; fbreak;};
            /inner/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = INNER_P; fbreak;};
            /on/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ON; fbreak;};
            /for/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = FOR; fbreak;};

            /vacuum/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = VACUUM; fbreak;};
            /cluster/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = CLUSTER; fbreak;};
            /analyze/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ANALYZE; fbreak;};

            /alter/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ALTER; fbreak;};

            /index/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = INDEX; fbreak;};

            /binary/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = BINARY; fbreak;};
            /DELIMITERS/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = DELIMITERS; fbreak;};
            /DELIMITER/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = DELIMITER; fbreak;};
            /CSV/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = CSV; fbreak;};
            /HEADER/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = HEADER_P; fbreak;};
            /QUOTE/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = QUOTE; fbreak;};
            /ESCAPE/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ESCAPE; fbreak;};
            /ENCODING/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ENCODING; fbreak;};
            /PROGRAM/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = PROGRAM; fbreak;};
            /stdin/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = STDIN; fbreak;};

            /ASYMMETRIC/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ASYMMETRIC; fbreak;};
            /BETWEEN/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = BETWEEN; fbreak;};
            /DROP/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = DROP; fbreak;};


            /BEGIN/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = BEGIN; fbreak;};
            /ROLLBACK/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ROLLBACK; fbreak;};
            /COMMIT/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = COMMIT; fbreak;};

            qidentifier      => { lval.str = string(lex.data[lex.ts + 1:lex.te - 1]); tok = IDENT; fbreak;};
            identifier      => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            sconst      => { lval.str = string(lex.data[lex.ts + 1:lex.te - 1]); tok = SCONST; fbreak;};


#           self		=	(',' | '(' | ')' | '[' | ']' | '.' | ';'| ':' | '+' | '-' | '*' | '\\' | '%' | '^' | '<' | '>' | '=');

            ',' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TCOMMA; fbreak;};
            '(' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TOPENBR; fbreak;};
            ')' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TCLOSEBR; fbreak;};
            '[' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TSQOPENBR; fbreak;};
            ']' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TSQCLOSEBR; fbreak;};
            '.' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TDOT; fbreak;};
            ';' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TSEMICOLON; fbreak;};
            ':' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TCOLON; fbreak;};
            '+' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TPLUS; fbreak;};
            '-' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TMINUS; fbreak;};
            '*' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TMUL; fbreak;};
           # TODO: support '\\' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = int(TMUL); fbreak;};
            '%' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TMOD; fbreak;};
            '^' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TPOW; fbreak;};
            '<' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TLESS; fbreak;};
            '>' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TGREATER; fbreak;};
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