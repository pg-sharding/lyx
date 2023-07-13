
//line lx/lexer.rl:1
package lx

import (
	"bufio"
	"fmt"
)


//line lx/lexer.go:12
const lexer_start int = 1
const lexer_first_final int = 1
const lexer_error int = 0

const lexer_en_main int = 1


//line lx/lexer.rl:14



type yySymType struct {
    token string
}

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
    
//line lx/lexer.go:42
	{
	 lex.cs = lexer_start
	 lex.ts = 0
	 lex.te = 0
	 lex.act = 0
	}

//line lx/lexer.rl:35
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

    
//line lx/lexer.go:83
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
	}
	goto st_out
tr6:
//line lx/lexer.rl:99
 lex.te = ( lex.p)
( lex.p)--
{ /* do nothing */ }
	goto st1
tr7:
//line lx/lexer.rl:102
 lex.te = ( lex.p)
( lex.p)--
{ lval.token = string(lex.data[lex.ts:lex.te]); tok = int(9); {( lex.p)++;  lex.cs = 1; goto _out } }
	goto st1
tr8:
//line lx/lexer.rl:100
 lex.te = ( lex.p)
( lex.p)--
{ lval.token = string(lex.data[lex.ts:lex.te]); tok = int(7); {( lex.p)++;  lex.cs = 1; goto _out }}
	goto st1
tr9:
//line lx/lexer.rl:101
 lex.te = ( lex.p)
( lex.p)--
{ lval.token = string(lex.data[lex.ts:lex.te]); tok = int(8); {( lex.p)++;  lex.cs = 1; goto _out }}
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

//line lx/lexer.go:140
		switch  lex.data[( lex.p)] {
		case 32:
			goto st2
		case 33:
			goto st3
		case 35:
			goto st3
		case 45:
			goto st3
		case 55:
			goto st4
		case 92:
			goto st3
		case 95:
			goto st5
		case 124:
			goto st3
		case 126:
			goto st3
		}
		switch {
		case  lex.data[( lex.p)] < 52:
			switch {
			case  lex.data[( lex.p)] < 37:
				if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 13 {
					goto st2
				}
			case  lex.data[( lex.p)] > 38:
				switch {
				case  lex.data[( lex.p)] > 43:
					if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 51 {
						goto st4
					}
				case  lex.data[( lex.p)] >= 42:
					goto st3
				}
			default:
				goto st3
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] < 65:
				if 60 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 64 {
					goto st3
				}
			case  lex.data[( lex.p)] > 90:
				switch {
				case  lex.data[( lex.p)] > 96:
					if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
						goto st5
					}
				case  lex.data[( lex.p)] >= 94:
					goto st3
				}
			default:
				goto st5
			}
		default:
			goto st6
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
		if  lex.data[( lex.p)] == 32 {
			goto st2
		}
		if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 13 {
			goto st2
		}
		goto tr6
	st3:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof3
		}
	st_case_3:
		switch  lex.data[( lex.p)] {
		case 33:
			goto st3
		case 35:
			goto st3
		case 45:
			goto st3
		case 92:
			goto st3
		case 94:
			goto st3
		case 96:
			goto st3
		case 124:
			goto st3
		case 126:
			goto st3
		}
		switch {
		case  lex.data[( lex.p)] < 42:
			if 37 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 38 {
				goto st3
			}
		case  lex.data[( lex.p)] > 43:
			if 60 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 64 {
				goto st3
			}
		default:
			goto st3
		}
		goto tr7
	st4:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof4
		}
	st_case_4:
		switch  lex.data[( lex.p)] {
		case 36:
			goto st5
		case 95:
			goto st5
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto st4
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st5
			}
		default:
			goto st5
		}
		goto tr8
	st5:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof5
		}
	st_case_5:
		switch  lex.data[( lex.p)] {
		case 36:
			goto st5
		case 95:
			goto st5
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto st5
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st5
			}
		default:
			goto st5
		}
		goto tr9
	st6:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof6
		}
	st_case_6:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st6
		}
		goto tr8
	st_out:
	_test_eof1:  lex.cs = 1; goto _test_eof
	_test_eof2:  lex.cs = 2; goto _test_eof
	_test_eof3:  lex.cs = 3; goto _test_eof
	_test_eof4:  lex.cs = 4; goto _test_eof
	_test_eof5:  lex.cs = 5; goto _test_eof
	_test_eof6:  lex.cs = 6; goto _test_eof

	_test_eof: {}
	if ( lex.p) == eof {
		switch  lex.cs {
		case 2:
			goto tr6
		case 3:
			goto tr7
		case 4:
			goto tr8
		case 5:
			goto tr9
		case 6:
			goto tr8
		}
	}

	_out: {}
	}

//line lx/lexer.rl:106


    return int(tok);
}