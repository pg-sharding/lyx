
//line lyx/lexer.rl:1
package lyx

import (
	"bufio"
	"fmt"
)


//line lyx/lexer.go:12
const lexer_start int = 2
const lexer_first_final int = 2
const lexer_error int = 0

const lexer_en_main int = 2


//line lyx/lexer.rl:14


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
    
//line lyx/lexer.go:37
	{
	 lex.cs = lexer_start
	 lex.ts = 0
	 lex.te = 0
	 lex.act = 0
	}

//line lyx/lexer.rl:30
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

    
//line lyx/lexer.go:79
	{
	if ( lex.p) == ( lex.pe) {
		goto _test_eof
	}
	switch  lex.cs {
	case 2:
		goto st_case_2
	case 0:
		goto st_case_0
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 1:
		goto st_case_1
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
	case 22:
		goto st_case_22
	case 23:
		goto st_case_23
	case 24:
		goto st_case_24
	case 25:
		goto st_case_25
	case 26:
		goto st_case_26
	case 27:
		goto st_case_27
	case 28:
		goto st_case_28
	case 29:
		goto st_case_29
	case 30:
		goto st_case_30
	case 31:
		goto st_case_31
	case 32:
		goto st_case_32
	case 33:
		goto st_case_33
	case 34:
		goto st_case_34
	case 35:
		goto st_case_35
	case 36:
		goto st_case_36
	case 37:
		goto st_case_37
	case 38:
		goto st_case_38
	case 39:
		goto st_case_39
	case 40:
		goto st_case_40
	}
	goto st_out
tr2:
//line lyx/lexer.rl:112
 lex.te = ( lex.p)+1
{ lval.str = string(lex.data[lex.ts:lex.te]); tok = int(SCONST); {( lex.p)++;  lex.cs = 2; goto _out }}
	goto st2
tr7:
//line lyx/lexer.rl:118
 lex.te = ( lex.p)+1
{ lval.str = string(lex.data[lex.ts:lex.te]); tok = int(TOPENBR); {( lex.p)++;  lex.cs = 2; goto _out }}
	goto st2
tr8:
//line lyx/lexer.rl:119
 lex.te = ( lex.p)+1
{ lval.str = string(lex.data[lex.ts:lex.te]); tok = int(TCLOSEBR); {( lex.p)++;  lex.cs = 2; goto _out }}
	goto st2
tr11:
//line lyx/lexer.rl:117
 lex.te = ( lex.p)+1
{ lval.str = string(lex.data[lex.ts:lex.te]); tok = int(TCOMMA); {( lex.p)++;  lex.cs = 2; goto _out }}
	goto st2
tr13:
//line lyx/lexer.rl:122
 lex.te = ( lex.p)+1
{ lval.str = string(lex.data[lex.ts:lex.te]); tok = TDOT; {( lex.p)++;  lex.cs = 2; goto _out }}
	goto st2
tr16:
//line lyx/lexer.rl:124
 lex.te = ( lex.p)+1
{ lval.str = string(lex.data[lex.ts:lex.te]); tok = int(TCOLON); {( lex.p)++;  lex.cs = 2; goto _out }}
	goto st2
tr17:
//line lyx/lexer.rl:123
 lex.te = ( lex.p)+1
{ lval.str = string(lex.data[lex.ts:lex.te]); tok = int(TSEMICOLON); {( lex.p)++;  lex.cs = 2; goto _out }}
	goto st2
tr30:
//line lyx/lexer.rl:120
 lex.te = ( lex.p)+1
{ lval.str = string(lex.data[lex.ts:lex.te]); tok = int(TSQCLOSEBR); {( lex.p)++;  lex.cs = 2; goto _out }}
	goto st2
tr31:
//line lyx/lexer.rl:121
 lex.te = ( lex.p)+1
{ lval.str = string(lex.data[lex.ts:lex.te]); tok = int(TSQCLOSEBR); {( lex.p)++;  lex.cs = 2; goto _out }}
	goto st2
tr33:
//line lyx/lexer.rl:96
 lex.te = ( lex.p)
( lex.p)--
{ /* do nothing */ }
	goto st2
tr34:
//line lyx/lexer.rl:139
 lex.te = ( lex.p)
( lex.p)--
{ 



                lval.str = string(lex.data[lex.ts:lex.te]); tok = int(OP);    
                {( lex.p)++;  lex.cs = 2; goto _out }
            }
	goto st2
tr36:
//line NONE:1
	switch  lex.act {
	case 3:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = SELECT; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 4:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = UPDATE; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 5:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = DELETE; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 6:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = TABLE; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 7:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = SET; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 8:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = FROM; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 9:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = WHERE; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 10:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = ORDER; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 11:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = BY; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 12:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = int(IDENT); {( lex.p)++;  lex.cs = 2; goto _out }}
	case 22:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = int(TPLUS); {( lex.p)++;  lex.cs = 2; goto _out }}
	case 23:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = int(TMINUS); {( lex.p)++;  lex.cs = 2; goto _out }}
	case 24:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = int(TMUL); {( lex.p)++;  lex.cs = 2; goto _out }}
	case 25:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = int(TMOD); {( lex.p)++;  lex.cs = 2; goto _out }}
	case 26:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = int(TPOW); {( lex.p)++;  lex.cs = 2; goto _out }}
	case 28:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = int(TGREATER); {( lex.p)++;  lex.cs = 2; goto _out }}
	case 29:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = TEQ; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 30:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = TNOT_EQUALS; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 31:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = TNOT_EQUALS; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 32:
	{( lex.p) = ( lex.te) - 1
 



                lval.str = string(lex.data[lex.ts:lex.te]); tok = int(OP);    
                {( lex.p)++;  lex.cs = 2; goto _out }
            }
	}
	
	goto st2
tr37:
//line lyx/lexer.rl:98
 lex.te = ( lex.p)
( lex.p)--
{ lval.str = string(lex.data[lex.ts:lex.te]); tok = SCONST; {( lex.p)++;  lex.cs = 2; goto _out }}
	goto st2
tr38:
//line lyx/lexer.rl:131
 lex.te = ( lex.p)
( lex.p)--
{ lval.str = string(lex.data[lex.ts:lex.te]); tok = int(TLESS); {( lex.p)++;  lex.cs = 2; goto _out }}
	goto st2
tr40:
//line lyx/lexer.rl:111
 lex.te = ( lex.p)
( lex.p)--
{ lval.str = string(lex.data[lex.ts:lex.te]); tok = int(IDENT); {( lex.p)++;  lex.cs = 2; goto _out }}
	goto st2
	st2:
//line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof2
		}
	st_case_2:
//line NONE:1
 lex.ts = ( lex.p)

//line lyx/lexer.go:333
		switch  lex.data[( lex.p)] {
		case 32:
			goto st3
		case 33:
			goto st4
		case 35:
			goto tr5
		case 37:
			goto tr6
		case 38:
			goto tr5
		case 39:
			goto st1
		case 40:
			goto tr7
		case 41:
			goto tr8
		case 42:
			goto tr9
		case 43:
			goto tr10
		case 44:
			goto tr11
		case 45:
			goto tr12
		case 46:
			goto tr13
		case 55:
			goto st6
		case 58:
			goto tr16
		case 59:
			goto tr17
		case 60:
			goto st9
		case 61:
			goto tr19
		case 62:
			goto tr20
		case 66:
			goto st10
		case 68:
			goto st11
		case 70:
			goto st16
		case 79:
			goto st19
		case 83:
			goto st23
		case 84:
			goto st28
		case 85:
			goto st32
		case 87:
			goto st37
		case 91:
			goto tr30
		case 92:
			goto tr5
		case 93:
			goto tr31
		case 94:
			goto tr32
		case 96:
			goto tr5
		case 98:
			goto st10
		case 100:
			goto st11
		case 102:
			goto st16
		case 111:
			goto st19
		case 115:
			goto st23
		case 116:
			goto st28
		case 117:
			goto st32
		case 119:
			goto st37
		case 124:
			goto tr5
		case 126:
			goto tr5
		}
		switch {
		case  lex.data[( lex.p)] < 52:
			switch {
			case  lex.data[( lex.p)] > 13:
				if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 51 {
					goto st6
				}
			case  lex.data[( lex.p)] >= 9:
				goto st3
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 64:
				if 65 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto tr21
				}
			case  lex.data[( lex.p)] >= 63:
				goto tr5
			}
		default:
			goto st8
		}
		goto st0
st_case_0:
	st0:
		 lex.cs = 0
		goto _out
	st3:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof3
		}
	st_case_3:
		if  lex.data[( lex.p)] == 32 {
			goto st3
		}
		if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 13 {
			goto st3
		}
		goto tr33
	st4:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof4
		}
	st_case_4:
		switch  lex.data[( lex.p)] {
		case 33:
			goto tr5
		case 35:
			goto tr5
		case 45:
			goto tr5
		case 61:
			goto tr35
		case 92:
			goto tr5
		case 94:
			goto tr5
		case 96:
			goto tr5
		case 124:
			goto tr5
		case 126:
			goto tr5
		}
		switch {
		case  lex.data[( lex.p)] < 42:
			if 37 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 38 {
				goto tr5
			}
		case  lex.data[( lex.p)] > 43:
			if 60 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 64 {
				goto tr5
			}
		default:
			goto tr5
		}
		goto tr34
tr5:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:139
 lex.act = 32;
	goto st5
tr6:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:129
 lex.act = 25;
	goto st5
tr9:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:127
 lex.act = 24;
	goto st5
tr10:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:125
 lex.act = 22;
	goto st5
tr12:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:126
 lex.act = 23;
	goto st5
tr19:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:133
 lex.act = 29;
	goto st5
tr20:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:132
 lex.act = 28;
	goto st5
tr32:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:130
 lex.act = 26;
	goto st5
tr35:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:136
 lex.act = 31;
	goto st5
tr39:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:135
 lex.act = 30;
	goto st5
	st5:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof5
		}
	st_case_5:
//line lyx/lexer.go:572
		switch  lex.data[( lex.p)] {
		case 33:
			goto tr5
		case 35:
			goto tr5
		case 45:
			goto tr5
		case 92:
			goto tr5
		case 94:
			goto tr5
		case 96:
			goto tr5
		case 124:
			goto tr5
		case 126:
			goto tr5
		}
		switch {
		case  lex.data[( lex.p)] < 42:
			if 37 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 38 {
				goto tr5
			}
		case  lex.data[( lex.p)] > 43:
			if 60 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 64 {
				goto tr5
			}
		default:
			goto tr5
		}
		goto tr36
	st1:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof1
		}
	st_case_1:
		switch  lex.data[( lex.p)] {
		case 36:
			goto st1
		case 39:
			goto tr2
		case 95:
			goto st1
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto st1
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto st0
	st6:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof6
		}
	st_case_6:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr21
		case 95:
			goto tr21
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto st6
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr21
			}
		default:
			goto tr21
		}
		goto tr37
tr21:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:111
 lex.act = 12;
	goto st7
tr41:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:108
 lex.act = 11;
	goto st7
tr46:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:102
 lex.act = 5;
	goto st7
tr49:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:105
 lex.act = 8;
	goto st7
tr53:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:107
 lex.act = 10;
	goto st7
tr56:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:104
 lex.act = 7;
	goto st7
tr59:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:100
 lex.act = 3;
	goto st7
tr63:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:103
 lex.act = 6;
	goto st7
tr68:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:101
 lex.act = 4;
	goto st7
tr72:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:106
 lex.act = 9;
	goto st7
	st7:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof7
		}
	st_case_7:
//line lyx/lexer.go:729
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr21
		case 95:
			goto tr21
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr21
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr21
			}
		default:
			goto tr21
		}
		goto tr36
	st8:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof8
		}
	st_case_8:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st8
		}
		goto tr37
	st9:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof9
		}
	st_case_9:
		switch  lex.data[( lex.p)] {
		case 33:
			goto tr5
		case 35:
			goto tr5
		case 45:
			goto tr5
		case 62:
			goto tr39
		case 92:
			goto tr5
		case 94:
			goto tr5
		case 96:
			goto tr5
		case 124:
			goto tr5
		case 126:
			goto tr5
		}
		switch {
		case  lex.data[( lex.p)] < 42:
			if 37 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 38 {
				goto tr5
			}
		case  lex.data[( lex.p)] > 43:
			if 60 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 64 {
				goto tr5
			}
		default:
			goto tr5
		}
		goto tr38
	st10:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof10
		}
	st_case_10:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr21
		case 89:
			goto tr41
		case 95:
			goto tr21
		case 121:
			goto tr41
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr21
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr21
			}
		default:
			goto tr21
		}
		goto tr40
	st11:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof11
		}
	st_case_11:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr21
		case 69:
			goto st12
		case 95:
			goto tr21
		case 101:
			goto st12
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr21
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr21
			}
		default:
			goto tr21
		}
		goto tr40
	st12:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof12
		}
	st_case_12:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr21
		case 76:
			goto st13
		case 95:
			goto tr21
		case 108:
			goto st13
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr21
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr21
			}
		default:
			goto tr21
		}
		goto tr40
	st13:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof13
		}
	st_case_13:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr21
		case 69:
			goto st14
		case 95:
			goto tr21
		case 101:
			goto st14
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr21
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr21
			}
		default:
			goto tr21
		}
		goto tr40
	st14:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof14
		}
	st_case_14:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr21
		case 84:
			goto st15
		case 95:
			goto tr21
		case 116:
			goto st15
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr21
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr21
			}
		default:
			goto tr21
		}
		goto tr40
	st15:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof15
		}
	st_case_15:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr21
		case 69:
			goto tr46
		case 95:
			goto tr21
		case 101:
			goto tr46
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr21
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr21
			}
		default:
			goto tr21
		}
		goto tr40
	st16:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof16
		}
	st_case_16:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr21
		case 82:
			goto st17
		case 95:
			goto tr21
		case 114:
			goto st17
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr21
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr21
			}
		default:
			goto tr21
		}
		goto tr40
	st17:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof17
		}
	st_case_17:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr21
		case 79:
			goto st18
		case 95:
			goto tr21
		case 111:
			goto st18
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr21
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr21
			}
		default:
			goto tr21
		}
		goto tr40
	st18:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof18
		}
	st_case_18:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr21
		case 77:
			goto tr49
		case 95:
			goto tr21
		case 109:
			goto tr49
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr21
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr21
			}
		default:
			goto tr21
		}
		goto tr40
	st19:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof19
		}
	st_case_19:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr21
		case 82:
			goto st20
		case 95:
			goto tr21
		case 114:
			goto st20
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr21
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr21
			}
		default:
			goto tr21
		}
		goto tr40
	st20:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof20
		}
	st_case_20:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr21
		case 68:
			goto st21
		case 95:
			goto tr21
		case 100:
			goto st21
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr21
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr21
			}
		default:
			goto tr21
		}
		goto tr40
	st21:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof21
		}
	st_case_21:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr21
		case 69:
			goto st22
		case 95:
			goto tr21
		case 101:
			goto st22
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr21
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr21
			}
		default:
			goto tr21
		}
		goto tr40
	st22:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof22
		}
	st_case_22:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr21
		case 82:
			goto tr53
		case 95:
			goto tr21
		case 114:
			goto tr53
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr21
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr21
			}
		default:
			goto tr21
		}
		goto tr40
	st23:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof23
		}
	st_case_23:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr21
		case 69:
			goto st24
		case 95:
			goto tr21
		case 101:
			goto st24
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr21
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr21
			}
		default:
			goto tr21
		}
		goto tr40
	st24:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof24
		}
	st_case_24:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr21
		case 76:
			goto st25
		case 84:
			goto tr56
		case 95:
			goto tr21
		case 108:
			goto st25
		case 116:
			goto tr56
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr21
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr21
			}
		default:
			goto tr21
		}
		goto tr40
	st25:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof25
		}
	st_case_25:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr21
		case 69:
			goto st26
		case 95:
			goto tr21
		case 101:
			goto st26
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr21
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr21
			}
		default:
			goto tr21
		}
		goto tr40
	st26:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof26
		}
	st_case_26:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr21
		case 67:
			goto st27
		case 95:
			goto tr21
		case 99:
			goto st27
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr21
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr21
			}
		default:
			goto tr21
		}
		goto tr40
	st27:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof27
		}
	st_case_27:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr21
		case 84:
			goto tr59
		case 95:
			goto tr21
		case 116:
			goto tr59
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr21
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr21
			}
		default:
			goto tr21
		}
		goto tr40
	st28:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof28
		}
	st_case_28:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr21
		case 65:
			goto st29
		case 95:
			goto tr21
		case 97:
			goto st29
		}
		switch {
		case  lex.data[( lex.p)] < 66:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr21
			}
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr21
			}
		default:
			goto tr21
		}
		goto tr40
	st29:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof29
		}
	st_case_29:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr21
		case 66:
			goto st30
		case 95:
			goto tr21
		case 98:
			goto st30
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr21
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr21
			}
		default:
			goto tr21
		}
		goto tr40
	st30:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof30
		}
	st_case_30:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr21
		case 76:
			goto st31
		case 95:
			goto tr21
		case 108:
			goto st31
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr21
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr21
			}
		default:
			goto tr21
		}
		goto tr40
	st31:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof31
		}
	st_case_31:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr21
		case 69:
			goto tr63
		case 95:
			goto tr21
		case 101:
			goto tr63
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr21
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr21
			}
		default:
			goto tr21
		}
		goto tr40
	st32:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof32
		}
	st_case_32:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr21
		case 80:
			goto st33
		case 95:
			goto tr21
		case 112:
			goto st33
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr21
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr21
			}
		default:
			goto tr21
		}
		goto tr40
	st33:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof33
		}
	st_case_33:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr21
		case 68:
			goto st34
		case 95:
			goto tr21
		case 100:
			goto st34
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr21
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr21
			}
		default:
			goto tr21
		}
		goto tr40
	st34:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof34
		}
	st_case_34:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr21
		case 65:
			goto st35
		case 95:
			goto tr21
		case 97:
			goto st35
		}
		switch {
		case  lex.data[( lex.p)] < 66:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr21
			}
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr21
			}
		default:
			goto tr21
		}
		goto tr40
	st35:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof35
		}
	st_case_35:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr21
		case 84:
			goto st36
		case 95:
			goto tr21
		case 116:
			goto st36
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr21
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr21
			}
		default:
			goto tr21
		}
		goto tr40
	st36:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof36
		}
	st_case_36:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr21
		case 69:
			goto tr68
		case 95:
			goto tr21
		case 101:
			goto tr68
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr21
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr21
			}
		default:
			goto tr21
		}
		goto tr40
	st37:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof37
		}
	st_case_37:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr21
		case 72:
			goto st38
		case 95:
			goto tr21
		case 104:
			goto st38
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr21
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr21
			}
		default:
			goto tr21
		}
		goto tr40
	st38:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof38
		}
	st_case_38:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr21
		case 69:
			goto st39
		case 95:
			goto tr21
		case 101:
			goto st39
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr21
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr21
			}
		default:
			goto tr21
		}
		goto tr40
	st39:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof39
		}
	st_case_39:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr21
		case 82:
			goto st40
		case 95:
			goto tr21
		case 114:
			goto st40
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr21
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr21
			}
		default:
			goto tr21
		}
		goto tr40
	st40:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof40
		}
	st_case_40:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr21
		case 69:
			goto tr72
		case 95:
			goto tr21
		case 101:
			goto tr72
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr21
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr21
			}
		default:
			goto tr21
		}
		goto tr40
	st_out:
	_test_eof2:  lex.cs = 2; goto _test_eof
	_test_eof3:  lex.cs = 3; goto _test_eof
	_test_eof4:  lex.cs = 4; goto _test_eof
	_test_eof5:  lex.cs = 5; goto _test_eof
	_test_eof1:  lex.cs = 1; goto _test_eof
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
	_test_eof22:  lex.cs = 22; goto _test_eof
	_test_eof23:  lex.cs = 23; goto _test_eof
	_test_eof24:  lex.cs = 24; goto _test_eof
	_test_eof25:  lex.cs = 25; goto _test_eof
	_test_eof26:  lex.cs = 26; goto _test_eof
	_test_eof27:  lex.cs = 27; goto _test_eof
	_test_eof28:  lex.cs = 28; goto _test_eof
	_test_eof29:  lex.cs = 29; goto _test_eof
	_test_eof30:  lex.cs = 30; goto _test_eof
	_test_eof31:  lex.cs = 31; goto _test_eof
	_test_eof32:  lex.cs = 32; goto _test_eof
	_test_eof33:  lex.cs = 33; goto _test_eof
	_test_eof34:  lex.cs = 34; goto _test_eof
	_test_eof35:  lex.cs = 35; goto _test_eof
	_test_eof36:  lex.cs = 36; goto _test_eof
	_test_eof37:  lex.cs = 37; goto _test_eof
	_test_eof38:  lex.cs = 38; goto _test_eof
	_test_eof39:  lex.cs = 39; goto _test_eof
	_test_eof40:  lex.cs = 40; goto _test_eof

	_test_eof: {}
	if ( lex.p) == eof {
		switch  lex.cs {
		case 3:
			goto tr33
		case 4:
			goto tr34
		case 5:
			goto tr36
		case 6:
			goto tr37
		case 7:
			goto tr36
		case 8:
			goto tr37
		case 9:
			goto tr38
		case 10:
			goto tr40
		case 11:
			goto tr40
		case 12:
			goto tr40
		case 13:
			goto tr40
		case 14:
			goto tr40
		case 15:
			goto tr40
		case 16:
			goto tr40
		case 17:
			goto tr40
		case 18:
			goto tr40
		case 19:
			goto tr40
		case 20:
			goto tr40
		case 21:
			goto tr40
		case 22:
			goto tr40
		case 23:
			goto tr40
		case 24:
			goto tr40
		case 25:
			goto tr40
		case 26:
			goto tr40
		case 27:
			goto tr40
		case 28:
			goto tr40
		case 29:
			goto tr40
		case 30:
			goto tr40
		case 31:
			goto tr40
		case 32:
			goto tr40
		case 33:
			goto tr40
		case 34:
			goto tr40
		case 35:
			goto tr40
		case 36:
			goto tr40
		case 37:
			goto tr40
		case 38:
			goto tr40
		case 39:
			goto tr40
		case 40:
			goto tr40
		}
	}

	_out: {}
	}

//line lyx/lexer.rl:153


    return int(tok);
}