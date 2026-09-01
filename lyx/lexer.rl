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

            /absent/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /absolute/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /access/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /action/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /add/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /admin/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /after/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /aggregate/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /also/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /always/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /analyse/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /any/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /asensitive/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /assertion/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /assignment/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /atomic/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /attach/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /attribute/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /authorization/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /backward/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /before/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /both/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /cache/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /called/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /call/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /cascaded/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /catalog/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /chain/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /checkpoint/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /class/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /close/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /coalesce/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /collation/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /columns/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /column/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /comments/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /comment/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /compression/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /configuration/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /connection/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /constraints/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /content/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /continue/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /cost/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /cube/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /cursor/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /data/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /declare/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /defaults/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /deferred/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /definer/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /depends/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /detach/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /dictionary/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /disable/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /document/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /domain/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /each/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /enable/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /encrypted/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /event/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /excluding/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /exclusive/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /exclude/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /expression/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /external/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /family/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /finalize/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /following/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /force/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /format/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /forward/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /freeze/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /functions/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /function/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /generated/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /global/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /granted/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /greatest/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /grouping/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /groups/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /handler/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /hold/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /identity/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /immediate/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /immutable/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /implicit/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /import/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /including/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /include/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /increment/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /indent/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /indexes/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /inherits/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /inherit/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /initially/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /inline/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /inout/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /input/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /insensitive/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /instead/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /invoker/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /json/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /keys/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /label/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /language/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /large/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /leading/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /leakproof/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /least/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /listen/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /load/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /location/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /lock/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /logged/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /mapping/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /matched/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /match/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /materialized/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /maxvalue/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /merge/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /method/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /minvalue/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /mode/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /move/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /name/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /natural/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /new/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /next/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /nfkc/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /nfkd/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /nfc/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /nfd/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /none/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /normalized/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /notify/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /nowait/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /nullif/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /object/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /off/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /old/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /options/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /option/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /others/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /out/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /overlaps/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /overlay/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /overriding/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /owned/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /owner/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /parallel/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /parameter/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /parser/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /partial/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /passing/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /password/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /placing/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /policy/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /position/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /preceding/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /precision/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /prepared/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /preserve/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /prior/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /privileges/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /procedural/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /procedures/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /procedure/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /publication/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /range/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /reassign/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /recheck/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /referencing/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /refresh/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /ref/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /reindex/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /relative/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /release/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /rename/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /replace/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /replica/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /restart/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /returns/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /return/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /rollup/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /routines/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /routine/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /rule/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /scalar/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /schemas/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /scroll/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /security/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /sequence/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /server/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /sets/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /similar/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /simple/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /snapshot/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /some/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /sql/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /stable/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /standalone/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /statement/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /storage/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /stored/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /strict/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /strip/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /subscription/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /substring/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /support/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /symmetric/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /sysid/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /system/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /tablesample/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /tablespace/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /tables/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /template/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /text/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /ties/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /trailing/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /transform/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /treat/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /trigger/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /trim/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /trusted/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /types/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /uescape/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /unbounded/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /unencrypted/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /unknown/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /unlisten/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /unlogged/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /until/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /user/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /validate/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /validator/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /valid/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /value/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /variadic/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /varying/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /verbose/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /version/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /views/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /view/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /volatile/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /whitespace/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /window/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /within/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /work/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /wrapper/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /xmlattributes/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /xmlconcat/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /xmlelement/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /xmlexists/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /xmlforest/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /xmlnamespaces/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /xmlparse/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /xmlpi/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /xmlroot/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /xmlserialize/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /xmltable/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /xml/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};
            /yes/i => { lval.str = string(lex.data[lex.ts:lex.te]); tok = IDENT; fbreak;};

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
            '/' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TDIV; fbreak;};
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
