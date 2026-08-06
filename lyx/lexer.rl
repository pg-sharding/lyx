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
    
    variable cs lex.cs;
    variable top lex.top;
    variable stack lex.stack;
}%%

type Lexer struct {
	data         []byte
	p, pe, cs    int
	ts, te, act  int

    top          int
    stack        []int

	result []string
}

func NewLexer(data []byte) *Lexer {
    lex := &Lexer{ 
        data: data,
        pe: len(data),
        stack: make([]int, 128),
    }
    %% write init;
    return lex
}

func ResetLexer(lex *Lexer, data []byte) {
    lex.pe = len(data)
    lex.data = data
    lex.stack = make([]int, 128)

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
        # https://www.colm.net/files/ragel/ragel-guide-6.6.pdf
        c_style_comment = '/''*' (any* - (any* '*/' any* )) '*''/';
        comment		= sql_comment | c_style_comment;


#       whitespace	=	({space}+|{comment});
        whitespace	=	space+;


        op_chars	=	( '~' | '!' | '@' | '#' | '^' | '&' | '|' | '`' | '?' | '+' | '-' | '*' | '\\' | '%' | '<' | '>' | '=' ) ;
        operator	=	op_chars;
        #https://github.com/postgres/postgres/blob/master/src/backend/parser/scan.l
        # ...For SQL compatibility, '+' and '-' cannot be the last char of a multi-char operator...
        #... The idea is to lex '=-' as two operators,..
        operator_multi	= op_chars+ (op_chars - ( '+' | '-' ));


        singleQuoteString := |*
            any => {
                tok = SCONST;

                if  lex.data[( lex.p)] == '\'' {

                    if lex.p + 1 < lex.pe  && lex.data[( lex.p + 1)] == '\'' {
                        lex.p++;
                    } else {
                        /* XXX: fix this mess */
                        {
                            ( lex.top)--; 
                            ( lex.cs) = ( lex.stack)[( lex.top)];
                            ( lex.p)++; 
                            lval.str = lval.strB.String();
                            goto _out
                        }
                    }
                }

                lval.strB.WriteByte(lex.data[lex.p])
            };
        *|;
        
        main := |*
            whitespace => { /* do nothing */ };
            # integer const is string const 
            comment => {/* nothing */};

            [']  => { 
                lval.strB.Reset();
                fcall singleQuoteString;
            };
            
            # skip dollar, get only param number
            param =>  {
                inp := string(lex.data[lex.ts+1:lex.te])
                if v, err := strconv.ParseInt(inp, 10, 64); err != nil {
                    lval.int = 0; tok = INVALID_ICONST; fbreak;
                } else {
                    lval.int = v; tok = PARAM; fbreak;
                }
            };

            integer =>  { 
                inp := string(lex.data[lex.ts:lex.te])
                if v, err := strconv.ParseInt(inp, 10, 64);  err != nil {
                    /* tooo big */
                    lval.str = inp; tok = SCONST; fbreak;
                } else {
                    lval.int = v; tok = ICONST; fbreak;
                }
            };


            real =>  { lval.str = string(lex.data[lex.ts:lex.te]); tok = SCONST; fbreak;};

            '::' =>  { lval.str = string(lex.data[lex.ts:lex.te]); tok = TYPECAST; fbreak;};

            /cast/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = CAST; fbreak;};
            /at/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = AT; fbreak;};
            
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
            /timestamp/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TIMESTAMP; fbreak;};
            /type/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TYPE_P; fbreak;};
            /enum/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ENUM_P; fbreak;};
            /zone/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ZONE; fbreak;};

            /if/i =>  { lval.str = string(lex.data[lex.ts:lex.te]); tok = IF_P; fbreak;};
            /of/i =>  { lval.str = string(lex.data[lex.ts:lex.te]); tok = OF; fbreak;};

            /prepare/i =>  { lval.str = string(lex.data[lex.ts:lex.te]); tok = PREPARE; fbreak;};
            /operator/i =>  { lval.str = string(lex.data[lex.ts:lex.te]); tok = OPERATOR; fbreak;};
            /collate/i =>  { lval.str = string(lex.data[lex.ts:lex.te]); tok = COLLATE; fbreak;};
            /deallocate/i =>  { lval.str = string(lex.data[lex.ts:lex.te]); tok = DEALLOCATE; fbreak;};
            /execute/i =>  { lval.str = string(lex.data[lex.ts:lex.te]); tok = EXECUTE; fbreak;};

            /select/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = SELECT; fbreak;};
            /insert/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = INSERT; fbreak;};
            /into/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = INTO; fbreak;};
            /values/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = VALUES; fbreak;};
            /update/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = UPDATE; fbreak;};
            /delete/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = DELETE; fbreak;};
            /create/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = CREATE; fbreak;};
            /TEMPORARY/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TEMPORARY; fbreak;};
            /TEMP/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TEMP; fbreak;};
            /OIDS/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = OIDS; fbreak;};
            /truncate/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TRUNCATE; fbreak;};
            /table/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TABLE; fbreak;};
            /extension/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = EXTENSION; fbreak;};
            /schema/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = SCHEMA; fbreak;};
            /database/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = DATABASE; fbreak;};
            /role/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ROLE; fbreak;};
            /primary/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = PRIMARY; fbreak;};
            /unique/i  => { lval.str = string(lex.data[lex.ts:lex.te]); tok = UNIQUE; fbreak;};
            /CONSTRAINT/i  => { lval.str = string(lex.data[lex.ts:lex.te]); tok = CONSTRAINT; fbreak;};
            /CONCURRENTLY/i  => { lval.str = string(lex.data[lex.ts:lex.te]); tok = CONCURRENTLY; fbreak;};
            /foreign/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = FOREIGN; fbreak;};
            /check/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = CHECK; fbreak;};
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

            /GRANT/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = GRANT; fbreak;};
            /REVOKE/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = REVOKE; fbreak;};

            /explain/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = EXPLAIN; fbreak;};

            /returning/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = RETURNING; fbreak;};
            /default/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = DEFAULT; fbreak;};
            
            /copy/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = COPY; fbreak;};
            /extract/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = EXTRACT; fbreak;};
            /to/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TO; fbreak;};
            /stdout/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = STDOUT; fbreak;};

            /limit/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = LIMIT; fbreak;};
            /offset/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = OFFSET; fbreak;};
            /distinct/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = DISTINCT; fbreak;};
            /like/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = LIKE; fbreak;};
            /ilike/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ILIKE; fbreak;};
            
            /is/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IS; fbreak;};
            /isnull/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ISNULL; fbreak;};
            /null/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = NULL_P; fbreak;};
            /nulls/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = NULLS_LA; fbreak;};
            /not/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = NOT; fbreak;};
            /normalize/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = NORMALIZE; fbreak;};
            /notnull/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = NOTNULL; fbreak;};
            /lateral/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = LATERAL_P; fbreak;};
            /ordinality/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ORDINALITY; fbreak;};
            /with/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = WITH; fbreak;};
            /true/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TRUE_P; fbreak;};
            /false/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = FALSE_P; fbreak;};

            /FETCH/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = FETCH; fbreak;};
            /FIRST/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = FIRST_P; fbreak;};
            /FILTER/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = FILTER; fbreak;};
            /LAST/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = LAST_P; fbreak;};
            /ASC/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ASC; fbreak;};
            /DESC/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = DESC; fbreak;};

            /SESSION/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = SESSION; fbreak;};
            /LOCAL/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = LOCAL; fbreak;};
            /CURRENT_USER/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = CURRENT_USER; fbreak;};

            /CHARACTERISTICS/i =>  { lval.str = string(lex.data[lex.ts:lex.te]); tok = CHARACTERISTICS; fbreak;};
            /SESSION/i =>  { lval.str = string(lex.data[lex.ts:lex.te]); tok = SESSION; fbreak;};
            /ISOLATION/i =>  { lval.str = string(lex.data[lex.ts:lex.te]); tok = ISOLATION; fbreak;};
            /LEVEL/i =>  { lval.str = string(lex.data[lex.ts:lex.te]); tok = LEVEL; fbreak;};

            /COMMITTED/i =>  { lval.str = string(lex.data[lex.ts:lex.te]); tok = COMMITTED; fbreak;};

            /CASCADE/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = CASCADE; fbreak;};
            /RESTRICT/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = RESTRICT; fbreak;};

            /array/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ARRAY; fbreak;};
            # explicit row for c_expr 
            /row/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ROW; fbreak;};
            /rows/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ROWS; fbreak;};
            /exists/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = EXISTS; fbreak;};


            /discard/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = DISCARD; fbreak;};
            /plans/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = PLANS; fbreak;};
            /sequences/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = SEQUENCES; fbreak;};
            /names/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = NAMES; fbreak;};

            /join/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = JOIN; fbreak;};
            /cross/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = CROSS; fbreak;};
            /left/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = LEFT; fbreak;};
            /right/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = RIGHT; fbreak;};
            /full/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = FULL; fbreak;};
            /outer/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = OUTER_P; fbreak;};
            /inner/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = INNER_P; fbreak;};
            /on/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ON; fbreak;};
            /no/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = NO; fbreak;};
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
            /SAVEPOINT/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = SAVEPOINT; fbreak;};
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
            /OVER/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = OVER; fbreak;};

            /CONVERSION/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = CONVERSION_P; fbreak;};
            /STATISTICS/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = STATISTICS; fbreak;};

            /absent/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ABSENT; fbreak;};
            /absolute/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ABSOLUTE_P; fbreak;};
            /access/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ACCESS; fbreak;};
            /action/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ACTION; fbreak;};
            /add/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ADD_P; fbreak;};
            /admin/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ADMIN; fbreak;};
            /after/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = AFTER; fbreak;};
            /aggregate/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = AGGREGATE; fbreak;};
            /also/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ALSO; fbreak;};
            /always/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ALWAYS; fbreak;};
            /analyse/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ANALYSE; fbreak;};
            /any/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT /* FIXME  */; fbreak;};
            /asensitive/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ASENSITIVE; fbreak;};
            /assertion/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ASSERTION; fbreak;};
            /assignment/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ASSIGNMENT; fbreak;};
            /atomic/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ATOMIC; fbreak;};
            /attach/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ATTACH; fbreak;};
            /attribute/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ATTRIBUTE; fbreak;};
            /authorization/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = AUTHORIZATION; fbreak;};
            /backward/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = BACKWARD; fbreak;};
            /before/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = BEFORE; fbreak;};
            /both/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = BOTH; fbreak;};
            /cache/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = CACHE; fbreak;};
            /called/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = CALLED; fbreak;};
            /call/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = CALL; fbreak;};
            /cascaded/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = CASCADED; fbreak;};
            /catalog/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = CATALOG_P; fbreak;};
            /chain/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = CHAIN; fbreak;};
            /checkpoint/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = CHECKPOINT; fbreak;};
            /class/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = CLASS; fbreak;};
            /close/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = CLOSE; fbreak;};
            /coalesce/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = COALESCE; fbreak;};
            /collation/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = COLLATION; fbreak;};
            /columns/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = COLUMNS; fbreak;};
            /column/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = COLUMN; fbreak;};
            /comments/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = COMMENTS; fbreak;};
            /comment/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = COMMENT; fbreak;};
            /compression/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = COMPRESSION; fbreak;};
            /configuration/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = CONFIGURATION; fbreak;};
            /connection/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = CONNECTION; fbreak;};
            /constraints/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = CONSTRAINTS; fbreak;};
            /content/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = CONTENT_P; fbreak;};
            /continue/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = CONTINUE_P; fbreak;};
            /cost/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = COST; fbreak;};
            /cube/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = CUBE; fbreak;};
            /cursor/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = CURSOR; fbreak;};
            /data/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = DATA_P; fbreak;};
            /declare/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = DECLARE; fbreak;};
            /defaults/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = DEFAULTS; fbreak;};
            /deferred/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = DEFERRED; fbreak;};
            /definer/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = DEFINER; fbreak;};
            /depends/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = DEPENDS; fbreak;};
            /detach/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = DETACH; fbreak;};
            /dictionary/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = DICTIONARY; fbreak;};
            /disable/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = DISABLE_P; fbreak;};
            /document/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = DOCUMENT_P; fbreak;};
            /domain/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = DOMAIN_P; fbreak;};
            /each/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = EACH; fbreak;};
            /enable/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ENABLE_P; fbreak;};
            /encrypted/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ENCRYPTED; fbreak;};
            /event/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = EVENT; fbreak;};
            /excluding/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = EXCLUDING; fbreak;};
            /exclusive/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = EXCLUSIVE; fbreak;};
            /exclude/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = EXCLUDE; fbreak;};
            /expression/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = EXPRESSION; fbreak;};
            /external/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = EXTERNAL; fbreak;};
            /family/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = FAMILY; fbreak;};
            /finalize/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = FINALIZE; fbreak;};
            /following/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = FOLLOWING; fbreak;};
            /force/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = FORCE; fbreak;};
            /format/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = FORMAT; fbreak;};
            /forward/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = FORWARD; fbreak;};
            /freeze/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = FREEZE; fbreak;};
            /functions/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = FUNCTIONS; fbreak;};
            /function/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = FUNCTION; fbreak;};
            /generated/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = GENERATED; fbreak;};
            /global/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = GLOBAL; fbreak;};
            /granted/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = GRANTED; fbreak;};
            /greatest/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = GREATEST; fbreak;};
            /grouping/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = GROUPING; fbreak;};
            /groups/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = GROUPS; fbreak;};
            /handler/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = HANDLER; fbreak;};
            /hold/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = HOLD; fbreak;};
            /identity/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENTITY_P; fbreak;};
            /immediate/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IMMEDIATE; fbreak;};
            /immutable/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IMMUTABLE; fbreak;};
            /implicit/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IMPLICIT_P; fbreak;};
            /import/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IMPORT_P; fbreak;};
            /including/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = INCLUDING; fbreak;};
            /include/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = INCLUDE; fbreak;};
            /increment/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = INCREMENT; fbreak;};
            /indent/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = INDENT; fbreak;};
            /indexes/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = INDEXES; fbreak;};
            /inherits/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = INHERITS; fbreak;};
            /inherit/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = INHERIT; fbreak;};
            /initially/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = INITIALLY; fbreak;};
            /inline/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = INLINE_P; fbreak;};
            /inout/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = INOUT; fbreak;};
            /input/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = INPUT_P; fbreak;};
            /insensitive/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = INSENSITIVE; fbreak;};
            /instead/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = INSTEAD; fbreak;};
            /invoker/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = INVOKER; fbreak;};
            /json/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = JSON; fbreak;};
            /keys/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = KEYS; fbreak;};
            /label/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = LABEL; fbreak;};
            /language/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = LANGUAGE; fbreak;};
            /large/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = LARGE_P; fbreak;};
            /leading/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = LEADING; fbreak;};
            /leakproof/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = LEAKPROOF; fbreak;};
            /least/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = LEAST; fbreak;};
            /listen/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = LISTEN; fbreak;};
            /load/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = LOAD; fbreak;};
            /location/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = LOCATION; fbreak;};
            /lock/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = LOCK_P; fbreak;};
            /logged/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = LOGGED; fbreak;};
            /mapping/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = MAPPING; fbreak;};
            /matched/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = MATCHED; fbreak;};
            /match/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = MATCH; fbreak;};
            /materialized/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = MATERIALIZED; fbreak;};
            /maxvalue/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = MAXVALUE; fbreak;};
            /merge/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = MERGE; fbreak;};
            /method/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = METHOD; fbreak;};
            /minvalue/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = MINVALUE; fbreak;};
            /mode/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = MODE; fbreak;};
            /move/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = MOVE; fbreak;};
            /name/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = NAME_P; fbreak;};
            /natural/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = NATURAL; fbreak;};
            /new/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = NEW; fbreak;};
            /next/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = NEXT; fbreak;};
            /nfkc/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = NFKC; fbreak;};
            /nfkd/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = NFKD; fbreak;};
            /nfc/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = NFC; fbreak;};
            /nfd/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = NFD; fbreak;};
            /none/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = NONE; fbreak;};
            /normalized/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = NORMALIZED; fbreak;};
            /notify/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = NOTIFY; fbreak;};
            /nowait/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = NOWAIT; fbreak;};
            /nullif/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = NULLIF; fbreak;};
            /object/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = OBJECT_P; fbreak;};
            /off/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = OFF; fbreak;};
            /old/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = OLD; fbreak;};
            /options/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = OPTIONS; fbreak;};
            /option/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = OPTION; fbreak;};
            /others/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = OTHERS; fbreak;};
            /out/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = OUT_P; fbreak;};
            /overlaps/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = OVERLAPS; fbreak;};
            /overlay/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = OVERLAY; fbreak;};
            /overriding/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = OVERRIDING; fbreak;};
            /owned/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = OWNED; fbreak;};
            /owner/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = OWNER; fbreak;};
            /parallel/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = PARALLEL; fbreak;};
            /parameter/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = PARAMETER; fbreak;};
            /parser/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = PARSER; fbreak;};
            /partial/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = PARTIAL; fbreak;};
            /passing/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = PASSING; fbreak;};
            /password/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = PASSWORD; fbreak;};
            /placing/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = PLACING; fbreak;};
            /policy/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = POLICY; fbreak;};
            /position/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = POSITION; fbreak;};
            /preceding/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = PRECEDING; fbreak;};
            /precision/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = PRECISION; fbreak;};
            /prepared/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = PREPARED; fbreak;};
            /preserve/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = PRESERVE; fbreak;};
            /prior/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = PRIOR; fbreak;};
            /privileges/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = PRIVILEGES; fbreak;};
            /procedural/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = PROCEDURAL; fbreak;};
            /procedures/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = PROCEDURES; fbreak;};
            /procedure/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = PROCEDURE; fbreak;};
            /publication/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = PUBLICATION; fbreak;};
            /range/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = RANGE; fbreak;};
            /reassign/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = REASSIGN; fbreak;};
            /recheck/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = RECHECK; fbreak;};
            /referencing/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = REFERENCING; fbreak;};
            /refresh/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = REFRESH; fbreak;};
            /ref/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = REF_P; fbreak;};
            /reindex/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = REINDEX; fbreak;};
            /relative/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = RELATIVE_P; fbreak;};
            /release/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = RELEASE; fbreak;};
            /rename/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = RENAME; fbreak;};
            /replace/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = REPLACE; fbreak;};
            /replica/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = REPLICA; fbreak;};
            /restart/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = RESTART; fbreak;};
            /returns/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = RETURNS; fbreak;};
            /return/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = RETURN; fbreak;};
            /rollup/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ROLLUP; fbreak;};
            /routines/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ROUTINES; fbreak;};
            /routine/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = ROUTINE; fbreak;};
            /rule/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = RULE; fbreak;};
            /scalar/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = SCALAR; fbreak;};
            /schemas/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = SCHEMAS; fbreak;};
            /scroll/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = SCROLL; fbreak;};
            /security/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = SECURITY; fbreak;};
            /sequence/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = SEQUENCE; fbreak;};
            /server/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = SERVER; fbreak;};
            /sets/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = SETS; fbreak;};
            /similar/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = SIMILAR; fbreak;};
            /simple/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = SIMPLE; fbreak;};
            /snapshot/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = SNAPSHOT; fbreak;};
            /some/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = SOME; fbreak;};
            /sql/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = SQL_P; fbreak;};
            /stable/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = STABLE; fbreak;};
            /standalone/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = STANDALONE_P; fbreak;};
            /statement/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = STATEMENT; fbreak;};
            /storage/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = STORAGE; fbreak;};
            /stored/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = STORED; fbreak;};
            /strict/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = STRICT_P; fbreak;};
            /strip/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = STRIP_P; fbreak;};
            /subscription/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = SUBSCRIPTION; fbreak;};
            /substring/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = SUBSTRING; fbreak;};
            /support/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = SUPPORT; fbreak;};
            /symmetric/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = SYMMETRIC; fbreak;};
            /sysid/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = SYSID; fbreak;};
            /system/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = SYSTEM_P; fbreak;};
            /tablesample/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TABLESAMPLE; fbreak;};
            /tablespace/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TABLESPACE; fbreak;};
            /tables/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TABLES; fbreak;};
            /template/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TEMPLATE; fbreak;};
            /text/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TEXT_P; fbreak;};
            /ties/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TIES; fbreak;};
            /trailing/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TRAILING; fbreak;};
            /transform/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TRANSFORM; fbreak;};
            /treat/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TREAT; fbreak;};
            /trigger/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TRIGGER; fbreak;};
            /trim/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TRIM; fbreak;};
            /trusted/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TRUSTED; fbreak;};
            /types/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TYPES_P; fbreak;};
            /uescape/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = UESCAPE; fbreak;};
            /unbounded/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = UNBOUNDED; fbreak;};
            /unencrypted/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = UNENCRYPTED; fbreak;};
            /unknown/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = UNKNOWN; fbreak;};
            /unlisten/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = UNLISTEN; fbreak;};
            /unlogged/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = UNLOGGED; fbreak;};
            /until/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = UNTIL; fbreak;};
            /user/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = USER; fbreak;};
            /validate/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = VALIDATE; fbreak;};
            /validator/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = VALIDATOR; fbreak;};
            /valid/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = VALID; fbreak;};
            /value/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = VALUE_P; fbreak;};
            /variadic/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = VARIADIC; fbreak;};
            /varying/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = VARYING; fbreak;};
            /verbose/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = VERBOSE; fbreak;};
            /version/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = VERSION_P; fbreak;};
            /views/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = VIEWS; fbreak;};
            /view/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = VIEW; fbreak;};
            /volatile/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = VOLATILE; fbreak;};
            /whitespace/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = WHITESPACE_P; fbreak;};
            /window/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = WINDOW; fbreak;};
            /within/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = WITHIN; fbreak;};
            /work/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = WORK; fbreak;};
            /wrapper/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = WRAPPER; fbreak;};
            /xmlattributes/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = XMLATTRIBUTES; fbreak;};
            /xmlconcat/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = XMLCONCAT; fbreak;};
            /xmlelement/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = XMLELEMENT; fbreak;};
            /xmlexists/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = XMLEXISTS; fbreak;};
            /xmlforest/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = XMLFOREST; fbreak;};
            /xmlnamespaces/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = XMLNAMESPACES; fbreak;};
            /xmlparse/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = XMLPARSE; fbreak;};
            /xmlpi/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = XMLPI; fbreak;};
            /xmlroot/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = XMLROOT; fbreak;};
            /xmlserialize/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = XMLSERIALIZE; fbreak;};
            /xmltable/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = XMLTABLE; fbreak;};
            /xml/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = XML_P; fbreak;};
            /yes/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = YES_P; fbreak;};

            qidentifier      => { lval.str = string(lex.data[lex.ts + 1:lex.te - 1]); tok = IDENT; fbreak;};
            identifier      => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
 
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
            operator_multi => {
                lval.str = string(lex.data[lex.ts:lex.te]); tok = int(OP);    
                fbreak;
            };

        *|;

        write exec;
    }%%

    return int(tok);
}
