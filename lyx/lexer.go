
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
	case 1:
		goto st_case_1
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
	}
	goto st_out
tr2:
//line lyx/lexer.rl:108
 lex.te = ( lex.p)+1
{ lval.str = string(lex.data[lex.ts:lex.te]); tok = int(SCONST); {( lex.p)++;  lex.cs = 2; goto _out }}
	goto st2
tr5:
//line lyx/lexer.rl:120
 lex.te = ( lex.p)+1
{ lval.str = string(lex.data[lex.ts:lex.te]); tok = int(TOPENBR); {( lex.p)++;  lex.cs = 2; goto _out }}
	goto st2
tr6:
//line lyx/lexer.rl:121
 lex.te = ( lex.p)+1
{ lval.str = string(lex.data[lex.ts:lex.te]); tok = int(TCLOSEBR); {( lex.p)++;  lex.cs = 2; goto _out }}
	goto st2
tr7:
//line lyx/lexer.rl:119
 lex.te = ( lex.p)+1
{ lval.str = string(lex.data[lex.ts:lex.te]); tok = int(TCOMMA); {( lex.p)++;  lex.cs = 2; goto _out }}
	goto st2
tr15:
//line lyx/lexer.rl:122
 lex.te = ( lex.p)+1
{ lval.str = string(lex.data[lex.ts:lex.te]); tok = int(TSQCLOSEBR); {( lex.p)++;  lex.cs = 2; goto _out }}
	goto st2
tr16:
//line lyx/lexer.rl:123
 lex.te = ( lex.p)+1
{ lval.str = string(lex.data[lex.ts:lex.te]); tok = int(TSQCLOSEBR); {( lex.p)++;  lex.cs = 2; goto _out }}
	goto st2
tr17:
//line lyx/lexer.rl:96
 lex.te = ( lex.p)
( lex.p)--
{ /* do nothing */ }
	goto st2
tr18:
//line lyx/lexer.rl:109
 lex.te = ( lex.p)
( lex.p)--
{ 



                lval.str = string(lex.data[lex.ts:lex.te]); tok = int(OP);    
                {( lex.p)++;  lex.cs = 2; goto _out }
            }
	goto st2
tr19:
//line lyx/lexer.rl:98
 lex.te = ( lex.p)
( lex.p)--
{ lval.str = string(lex.data[lex.ts:lex.te]); tok = int(SCONST); {( lex.p)++;  lex.cs = 2; goto _out }}
	goto st2
tr20:
//line NONE:1
	switch  lex.act {
	case 3:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = int(SELECT); {( lex.p)++;  lex.cs = 2; goto _out }}
	case 4:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = int(UPDATE); {( lex.p)++;  lex.cs = 2; goto _out }}
	case 5:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = int(DELETE); {( lex.p)++;  lex.cs = 2; goto _out }}
	case 6:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = int(TABLE); {( lex.p)++;  lex.cs = 2; goto _out }}
	case 7:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = int(SET); {( lex.p)++;  lex.cs = 2; goto _out }}
	case 8:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = int(IDENT); {( lex.p)++;  lex.cs = 2; goto _out }}
	}
	
	goto st2
tr21:
//line lyx/lexer.rl:107
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

//line lyx/lexer.go:236
		switch  lex.data[( lex.p)] {
		case 32:
			goto st3
		case 33:
			goto st4
		case 35:
			goto st4
		case 39:
			goto st1
		case 40:
			goto tr5
		case 41:
			goto tr6
		case 44:
			goto tr7
		case 55:
			goto st5
		case 68:
			goto st8
		case 83:
			goto st13
		case 84:
			goto st18
		case 85:
			goto st22
		case 91:
			goto tr15
		case 92:
			goto st4
		case 93:
			goto tr16
		case 94:
			goto st4
		case 96:
			goto st4
		case 100:
			goto st8
		case 115:
			goto st13
		case 116:
			goto st18
		case 117:
			goto st22
		case 124:
			goto st4
		case 126:
			goto st4
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			switch {
			case  lex.data[( lex.p)] > 13:
				if 37 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 45 {
					goto st4
				}
			case  lex.data[( lex.p)] >= 9:
				goto st3
			}
		case  lex.data[( lex.p)] > 51:
			switch {
			case  lex.data[( lex.p)] < 60:
				if 52 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
					goto st7
				}
			case  lex.data[( lex.p)] > 64:
				if 65 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto tr10
				}
			default:
				goto st4
			}
		default:
			goto st5
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
		goto tr17
	st4:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof4
		}
	st_case_4:
		switch  lex.data[( lex.p)] {
		case 33:
			goto st4
		case 35:
			goto st4
		case 45:
			goto st4
		case 92:
			goto st4
		case 94:
			goto st4
		case 96:
			goto st4
		case 124:
			goto st4
		case 126:
			goto st4
		}
		switch {
		case  lex.data[( lex.p)] < 42:
			if 37 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 38 {
				goto st4
			}
		case  lex.data[( lex.p)] > 43:
			if 60 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 64 {
				goto st4
			}
		default:
			goto st4
		}
		goto tr18
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
	st5:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof5
		}
	st_case_5:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr10
		case 95:
			goto tr10
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto st5
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr19
tr10:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:107
 lex.act = 8;
	goto st6
tr26:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:102
 lex.act = 5;
	goto st6
tr29:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:104
 lex.act = 7;
	goto st6
tr32:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:100
 lex.act = 3;
	goto st6
tr36:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:103
 lex.act = 6;
	goto st6
tr41:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:101
 lex.act = 4;
	goto st6
	st6:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof6
		}
	st_case_6:
//line lyx/lexer.go:461
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr10
		case 95:
			goto tr10
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr10
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr20
	st7:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof7
		}
	st_case_7:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st7
		}
		goto tr19
	st8:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof8
		}
	st_case_8:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr10
		case 69:
			goto st9
		case 95:
			goto tr10
		case 101:
			goto st9
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr10
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr21
	st9:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof9
		}
	st_case_9:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr10
		case 76:
			goto st10
		case 95:
			goto tr10
		case 108:
			goto st10
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr10
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr21
	st10:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof10
		}
	st_case_10:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr10
		case 69:
			goto st11
		case 95:
			goto tr10
		case 101:
			goto st11
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr10
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr21
	st11:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof11
		}
	st_case_11:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr10
		case 84:
			goto st12
		case 95:
			goto tr10
		case 116:
			goto st12
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr10
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr21
	st12:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof12
		}
	st_case_12:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr10
		case 69:
			goto tr26
		case 95:
			goto tr10
		case 101:
			goto tr26
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr10
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr21
	st13:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof13
		}
	st_case_13:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr10
		case 69:
			goto st14
		case 95:
			goto tr10
		case 101:
			goto st14
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr10
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr21
	st14:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof14
		}
	st_case_14:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr10
		case 76:
			goto st15
		case 84:
			goto tr29
		case 95:
			goto tr10
		case 108:
			goto st15
		case 116:
			goto tr29
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr10
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr21
	st15:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof15
		}
	st_case_15:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr10
		case 69:
			goto st16
		case 95:
			goto tr10
		case 101:
			goto st16
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr10
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr21
	st16:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof16
		}
	st_case_16:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr10
		case 67:
			goto st17
		case 95:
			goto tr10
		case 99:
			goto st17
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr10
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr21
	st17:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof17
		}
	st_case_17:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr10
		case 84:
			goto tr32
		case 95:
			goto tr10
		case 116:
			goto tr32
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr10
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr21
	st18:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof18
		}
	st_case_18:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr10
		case 65:
			goto st19
		case 95:
			goto tr10
		case 97:
			goto st19
		}
		switch {
		case  lex.data[( lex.p)] < 66:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr10
			}
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr21
	st19:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof19
		}
	st_case_19:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr10
		case 66:
			goto st20
		case 95:
			goto tr10
		case 98:
			goto st20
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr10
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr21
	st20:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof20
		}
	st_case_20:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr10
		case 76:
			goto st21
		case 95:
			goto tr10
		case 108:
			goto st21
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr10
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr21
	st21:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof21
		}
	st_case_21:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr10
		case 69:
			goto tr36
		case 95:
			goto tr10
		case 101:
			goto tr36
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr10
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr21
	st22:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof22
		}
	st_case_22:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr10
		case 80:
			goto st23
		case 95:
			goto tr10
		case 112:
			goto st23
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr10
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr21
	st23:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof23
		}
	st_case_23:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr10
		case 68:
			goto st24
		case 95:
			goto tr10
		case 100:
			goto st24
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr10
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr21
	st24:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof24
		}
	st_case_24:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr10
		case 65:
			goto st25
		case 95:
			goto tr10
		case 97:
			goto st25
		}
		switch {
		case  lex.data[( lex.p)] < 66:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr10
			}
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr21
	st25:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof25
		}
	st_case_25:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr10
		case 84:
			goto st26
		case 95:
			goto tr10
		case 116:
			goto st26
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr10
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr21
	st26:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof26
		}
	st_case_26:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr10
		case 69:
			goto tr41
		case 95:
			goto tr10
		case 101:
			goto tr41
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr10
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr10
			}
		default:
			goto tr10
		}
		goto tr21
	st_out:
	_test_eof2:  lex.cs = 2; goto _test_eof
	_test_eof3:  lex.cs = 3; goto _test_eof
	_test_eof4:  lex.cs = 4; goto _test_eof
	_test_eof1:  lex.cs = 1; goto _test_eof
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
	_test_eof22:  lex.cs = 22; goto _test_eof
	_test_eof23:  lex.cs = 23; goto _test_eof
	_test_eof24:  lex.cs = 24; goto _test_eof
	_test_eof25:  lex.cs = 25; goto _test_eof
	_test_eof26:  lex.cs = 26; goto _test_eof

	_test_eof: {}
	if ( lex.p) == eof {
		switch  lex.cs {
		case 3:
			goto tr17
		case 4:
			goto tr18
		case 5:
			goto tr19
		case 6:
			goto tr20
		case 7:
			goto tr19
		case 8:
			goto tr21
		case 9:
			goto tr21
		case 10:
			goto tr21
		case 11:
			goto tr21
		case 12:
			goto tr21
		case 13:
			goto tr21
		case 14:
			goto tr21
		case 15:
			goto tr21
		case 16:
			goto tr21
		case 17:
			goto tr21
		case 18:
			goto tr21
		case 19:
			goto tr21
		case 20:
			goto tr21
		case 21:
			goto tr21
		case 22:
			goto tr21
		case 23:
			goto tr21
		case 24:
			goto tr21
		case 25:
			goto tr21
		case 26:
			goto tr21
		}
	}

	_out: {}
	}

//line lyx/lexer.rl:131


    return int(tok);
}