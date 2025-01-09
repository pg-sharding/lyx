package lyx

import (
    "strconv"
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

        integer = digit+;
        ninteger = '-' integer;
        param = '$' integer;
        

#        space		=	[ \t\n\r\f];
        horiz_space	= [ \t\f];
        newline		=	[\n\r];
        non_newline	=	[^\n\r];
        dquote      =   ["];


        decimal	= ((digit*'.'digit+)|(digit+'.'digit*));
        real = (decimal)|('-'decimal);

        ident_start	=	[A-Za-z\200-\377_];
        ident_cont	=	[A-Za-z\200-\377_0-9$];

        identifier	=	ident_start ident_cont*;

        qidentifier	=	dquote (any - newline - dquote)* dquote ;


        sql_comment = '-''-' non_newline*;
        c_style_comment = '/''*' (any - '*''/')* '*''/';
        comment		= sql_comment | c_style_comment;


#       whitespace	=	({space}+|{comment});
        whitespace	=	space+;


        op_chars	=	( '~' | '!' | '@' | '#' | '^' | '&' | '|' | '`' | '?' | '+' | '-' | '*' | '\\' | '%' | '<' | '>' | '=' ) ;
        operator	=	op_chars+;


        sconst = '\'' (any-'\'')* '\'';
        
        main := |*
            whitespace => { /* do nothing */ };
            # integer const is string const 
            comment => {/* nothing */};
            
            # skip dollar, get only param number
            param =>  {
                i, _ := strconv.Atoi(string(lex.data[lex.ts+1:lex.te]))
                lval.int = int(i); tok = PARAM; fbreak;
            };

            integer =>  { lval.int, _ = strconv.Atoi(string(lex.data[lex.ts:lex.te])); tok = ICONST; fbreak;};
            ninteger => { lval.int, _ = strconv.Atoi(string(lex.data[lex.ts:lex.te])); tok = ICONST; fbreak;};

            real =>  { lval.str = string(lex.data[lex.ts:lex.te]); tok = SCONST; fbreak;};

            '::' =>  { lval.str = string(lex.data[lex.ts:lex.te]); tok = TYPECAST; fbreak;};

            /cast/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = CAST; fbreak;};
            
            /setof/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = SETOF; fbreak;};
            /int/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = INT_P; fbreak;};
            /integer/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = INTEGER; fbreak;};
            /smallint/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = SMALLINT; fbreak;};
            /bigint/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = BIGINT; fbreak;};
            /real/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = REAL; fbreak;};
            /float/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = FLOAT_P; fbreak;};
            /double/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = DOUBLE_P; fbreak;};
            /decimal/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = DECIMAL_P; fbreak;};
            /dec/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = DEC; fbreak;};
            /numeric/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = NUMERIC; fbreak;};
            /boolean/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = BOOLEAN_P; fbreak;};
            /bit/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = BIT; fbreak;};
            

            /interval/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = INTERVAL; fbreak;};
            /year/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = YEAR_P; fbreak;};
            /month/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = MONTH_P; fbreak;};
            /day/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = DAY_P; fbreak;};
            /hour/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = HOUR_P; fbreak;};
            /minute/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = MINUTE_P; fbreak;};
            /second/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = SECOND_P; fbreak;};


            /character/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = CHARACTER; fbreak;};
            /char/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = CHAR_P; fbreak;};
            /varchar/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = VARCHAR; fbreak;};
            /national/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = NATIONAL; fbreak;};
            /nchar/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = NCHAR; fbreak;};


            /without/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = WITHOUT; fbreak;};
            /time/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TIME; fbreak;};
            /zone/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ZONE; fbreak;};

            /if/i =>  { lval.str = string(lex.data[lex.ts:lex.te]); tok = IF_P; fbreak;};
            /OF/i =>  { lval.str = string(lex.data[lex.ts:lex.te]); tok = OF; fbreak;};

            /prepare/i =>  { lval.str = string(lex.data[lex.ts:lex.te]); tok = PREPARE; fbreak;};
            /deallocate/i =>  { lval.str = string(lex.data[lex.ts:lex.te]); tok = DEALLOCATE; fbreak;};
            /execute/i =>  { lval.str = string(lex.data[lex.ts:lex.te]); tok = EXECUTE; fbreak;};

            /select/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = SELECT; fbreak;};
            /insert/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = INSERT; fbreak;};
            /into/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = INTO; fbreak;};
            /values/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = VALUES; fbreak;};
            /update/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = UPDATE; fbreak;};
            /delete/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = DELETE; fbreak;};
            /create/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = CREATE; fbreak;};
            /truncate/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TRUNCATE; fbreak;};
            /table/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TABLE; fbreak;};
            /schema/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = SCHEMA; fbreak;};
            /database/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = DATABASE; fbreak;};
            /role/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ROLE; fbreak;};
            /primary/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = PRIMARY; fbreak;};
            /foreign/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = FOREIGN; fbreak;};
            /references/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = REFERENCES; fbreak;};
            /key/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = KEY; fbreak;};
            /set/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = SET; fbreak;};
            /reset/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = RESET; fbreak;};
            /show/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = SHOW; fbreak;};
            /from/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = FROM; fbreak;};
            /where/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = WHERE; fbreak;};
            /order/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ORDER; fbreak;};
            /group/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = GROUP; fbreak;};
            /by/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = BY; fbreak;};
            /having/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = HAVING; fbreak;};
            /as/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = AS; fbreak;};
            /and/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = AND; fbreak;};
            /or/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = OR; fbreak;};

            /explain/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = EXPLAIN; fbreak;};

            /returning/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = RETURNING; fbreak;};
            /default/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = DEFAULT; fbreak;};
            
            /copy/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = COPY; fbreak;};
            /to/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TO; fbreak;};
            /stdout/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = STDOUT; fbreak;};

            /limit/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = LIMIT; fbreak;};
            /distinct/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = DISTINCT; fbreak;};
            /like/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = LIKE; fbreak;};
            
            /is/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IS; fbreak;};
            /isnull/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ISNULL; fbreak;};
            /null/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = NULL_P; fbreak;};
            /nulls/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = NULLS_LA; fbreak;};
            /not/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = NOT; fbreak;};
            /notnull/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = NOTNULL; fbreak;};
            /lateral/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = LATERAL_P; fbreak;};
            /ordinality/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ORDINALITY; fbreak;};
            /with/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = WITH; fbreak;};
            /true/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TRUE_P; fbreak;};
            /false/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = FALSE_P; fbreak;};

            /FIRST/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = FIRST_P; fbreak;};
            /LAST/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = LAST_P; fbreak;};
            /ASC/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ASC; fbreak;};
            /DESC/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = DESC; fbreak;};

            /CHARACTERISTICS/i =>  { lval.str = string(lex.data[lex.ts:lex.te]); tok = CHARACTERISTICS; fbreak;};
            /SESSION/i =>  { lval.str = string(lex.data[lex.ts:lex.te]); tok = SESSION; fbreak;};
            /ISOLATION/i =>  { lval.str = string(lex.data[lex.ts:lex.te]); tok = ISOLATION; fbreak;};
            /LEVEL/i =>  { lval.str = string(lex.data[lex.ts:lex.te]); tok = LEVEL; fbreak;};

            /COMMITTED/i =>  { lval.str = string(lex.data[lex.ts:lex.te]); tok = COMMITTED; fbreak;};


            /array/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ARRAY; fbreak;};
            # explicit row for c_expr 
            /row/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ROW; fbreak;};
            /exists/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = EXISTS; fbreak;};


            /join/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = JOIN; fbreak;};
            /cross/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = CROSS; fbreak;};
            /left/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = LEFT; fbreak;};
            /right/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = RIGHT; fbreak;};
            /full/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = FULL; fbreak;};
            /outer/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = OUTER_P; fbreak;};
            /inner/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = INNER_P; fbreak;};
            /on/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ON; fbreak;};
            /in/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IN_P; fbreak;};
            /for/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = FOR; fbreak;};


            /locked/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = LOCKED; fbreak;};
            /skip/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = SKIP; fbreak;};

            /RECURSIVE/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = RECURSIVE; fbreak;};
            /SEARCH/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = SEARCH; fbreak;};
            /CYCLE/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = CYCLE; fbreak;};
            /BREADTH/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = BREADTH; fbreak;};
            /DEPTH/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = DEPTH; fbreak;};

            /SHARE/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = SHARE; fbreak;};

            /USING/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = USING; fbreak;};

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
            /ELSE/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ELSE; fbreak;};
            /END/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = END_P; fbreak;};
            /TRANSACTION/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TRANSACTION; fbreak;};
            /READ/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = READ; fbreak;};
            /ONLY/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ONLY; fbreak;};
            /WRITE/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = WRITE; fbreak;};
            /DEFERRABLE/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = DEFERRABLE; fbreak;};
            /ISOLATION/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ISOLATION; fbreak;};
            /LEVEL/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = LEVEL; fbreak;};
            /UNCOMMITTED/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = UNCOMMITTED; fbreak;};
            /COMMITTED/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = COMMITTED; fbreak;};
            /REPEATABLE/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = REPEATABLE; fbreak;};
            /SERIALIZABLE/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = SERIALIZABLE; fbreak;};

            /START/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = START; fbreak;};
            /ABORT/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ABORT_P; fbreak;};
            /END/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = END_P; fbreak;};
            /ROLLBACK/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ROLLBACK; fbreak;};
            /COMMIT/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = COMMIT; fbreak;};

            /CASE/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = CASE; fbreak;};
            /WHEN/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = WHEN; fbreak;};
            /THEN/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = THEN; fbreak;};
            /END/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = END_P; fbreak;};

            /CONFLICT/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = CONFLICT; fbreak;};
            /ON/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ON; fbreak;};
            /NOTHING/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = NOTHING; fbreak;};
            /DO/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = DO; fbreak;};

            /UNION/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = UNION; fbreak;};
            /EXCEPT/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = EXCEPT; fbreak;};
            /INTERSECT/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = INTERSECT; fbreak;};
            /ALL/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ALL; fbreak;};

            /PARTITION/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = PARTITION; fbreak;};

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