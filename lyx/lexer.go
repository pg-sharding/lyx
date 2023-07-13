
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
	case 41:
		goto st_case_41
	case 42:
		goto st_case_42
	case 43:
		goto st_case_43
	case 44:
		goto st_case_44
	case 45:
		goto st_case_45
	case 46:
		goto st_case_46
	case 47:
		goto st_case_47
	case 48:
		goto st_case_48
	case 49:
		goto st_case_49
	case 50:
		goto st_case_50
	case 51:
		goto st_case_51
	case 52:
		goto st_case_52
	case 53:
		goto st_case_53
	case 54:
		goto st_case_54
	case 55:
		goto st_case_55
	case 56:
		goto st_case_56
	case 57:
		goto st_case_57
	case 58:
		goto st_case_58
	case 59:
		goto st_case_59
	case 60:
		goto st_case_60
	case 61:
		goto st_case_61
	case 62:
		goto st_case_62
	case 63:
		goto st_case_63
	case 64:
		goto st_case_64
	case 65:
		goto st_case_65
	case 66:
		goto st_case_66
	case 67:
		goto st_case_67
	case 68:
		goto st_case_68
	case 69:
		goto st_case_69
	case 70:
		goto st_case_70
	case 71:
		goto st_case_71
	case 72:
		goto st_case_72
	case 73:
		goto st_case_73
	case 74:
		goto st_case_74
	case 75:
		goto st_case_75
	case 76:
		goto st_case_76
	case 77:
		goto st_case_77
	case 78:
		goto st_case_78
	case 79:
		goto st_case_79
	case 80:
		goto st_case_80
	case 81:
		goto st_case_81
	case 82:
		goto st_case_82
	case 83:
		goto st_case_83
	case 84:
		goto st_case_84
	case 85:
		goto st_case_85
	case 86:
		goto st_case_86
	case 87:
		goto st_case_87
	case 88:
		goto st_case_88
	case 89:
		goto st_case_89
	case 90:
		goto st_case_90
	case 91:
		goto st_case_91
	case 92:
		goto st_case_92
	case 93:
		goto st_case_93
	case 94:
		goto st_case_94
	case 95:
		goto st_case_95
	case 96:
		goto st_case_96
	case 97:
		goto st_case_97
	case 98:
		goto st_case_98
	case 99:
		goto st_case_99
	case 100:
		goto st_case_100
	case 101:
		goto st_case_101
	case 102:
		goto st_case_102
	case 103:
		goto st_case_103
	case 104:
		goto st_case_104
	case 105:
		goto st_case_105
	case 106:
		goto st_case_106
	case 107:
		goto st_case_107
	case 108:
		goto st_case_108
	case 109:
		goto st_case_109
	case 110:
		goto st_case_110
	case 111:
		goto st_case_111
	case 112:
		goto st_case_112
	case 113:
		goto st_case_113
	case 114:
		goto st_case_114
	case 115:
		goto st_case_115
	case 116:
		goto st_case_116
	case 117:
		goto st_case_117
	case 118:
		goto st_case_118
	case 119:
		goto st_case_119
	case 120:
		goto st_case_120
	case 121:
		goto st_case_121
	case 122:
		goto st_case_122
	case 123:
		goto st_case_123
	case 124:
		goto st_case_124
	case 125:
		goto st_case_125
	case 126:
		goto st_case_126
	case 127:
		goto st_case_127
	case 128:
		goto st_case_128
	case 129:
		goto st_case_129
	case 130:
		goto st_case_130
	case 131:
		goto st_case_131
	case 132:
		goto st_case_132
	case 133:
		goto st_case_133
	case 134:
		goto st_case_134
	case 135:
		goto st_case_135
	case 136:
		goto st_case_136
	case 137:
		goto st_case_137
	}
	goto st_out
tr1:
//line lyx/lexer.rl:148
 lex.te = ( lex.p)+1
{ lval.str = string(lex.data[lex.ts + 1:lex.te - 1]); tok = int(SCONST); {( lex.p)++;  lex.cs = 2; goto _out }}
	goto st2
tr7:
//line lyx/lexer.rl:154
 lex.te = ( lex.p)+1
{ lval.str = string(lex.data[lex.ts:lex.te]); tok = TOPENBR; {( lex.p)++;  lex.cs = 2; goto _out }}
	goto st2
tr8:
//line lyx/lexer.rl:155
 lex.te = ( lex.p)+1
{ lval.str = string(lex.data[lex.ts:lex.te]); tok = TCLOSEBR; {( lex.p)++;  lex.cs = 2; goto _out }}
	goto st2
tr11:
//line lyx/lexer.rl:153
 lex.te = ( lex.p)+1
{ lval.str = string(lex.data[lex.ts:lex.te]); tok = TCOMMA; {( lex.p)++;  lex.cs = 2; goto _out }}
	goto st2
tr13:
//line lyx/lexer.rl:158
 lex.te = ( lex.p)+1
{ lval.str = string(lex.data[lex.ts:lex.te]); tok = TDOT; {( lex.p)++;  lex.cs = 2; goto _out }}
	goto st2
tr16:
//line lyx/lexer.rl:160
 lex.te = ( lex.p)+1
{ lval.str = string(lex.data[lex.ts:lex.te]); tok = TCOLON; {( lex.p)++;  lex.cs = 2; goto _out }}
	goto st2
tr17:
//line lyx/lexer.rl:159
 lex.te = ( lex.p)+1
{ lval.str = string(lex.data[lex.ts:lex.te]); tok = TSEMICOLON; {( lex.p)++;  lex.cs = 2; goto _out }}
	goto st2
tr40:
//line lyx/lexer.rl:156
 lex.te = ( lex.p)+1
{ lval.str = string(lex.data[lex.ts:lex.te]); tok = TSQCLOSEBR; {( lex.p)++;  lex.cs = 2; goto _out }}
	goto st2
tr41:
//line lyx/lexer.rl:157
 lex.te = ( lex.p)+1
{ lval.str = string(lex.data[lex.ts:lex.te]); tok = TSQCLOSEBR; {( lex.p)++;  lex.cs = 2; goto _out }}
	goto st2
tr43:
//line lyx/lexer.rl:96
 lex.te = ( lex.p)
( lex.p)--
{ /* do nothing */ }
	goto st2
tr44:
//line lyx/lexer.rl:177
 lex.te = ( lex.p)
( lex.p)--
{ 



                lval.str = string(lex.data[lex.ts:lex.te]); tok = int(OP);    
                {( lex.p)++;  lex.cs = 2; goto _out }
            }
	goto st2
tr46:
//line NONE:1
	switch  lex.act {
	case 3:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = SELECT; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 4:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = INSERT; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 5:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = INTO; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 6:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = VALUES; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 7:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = UPDATE; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 8:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = DELETE; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 9:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = CREATE; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 10:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = TABLE; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 11:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = DATABASE; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 12:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = ROLE; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 13:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = PRIMARY; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 14:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = FOREIGN; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 15:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = REFERENCES; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 16:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = KEY; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 17:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = SET; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 18:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = FROM; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 19:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = WHERE; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 20:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = ORDER; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 21:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = GROUP; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 22:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = BY; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 23:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = AS; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 24:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = AND; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 26:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = RETURNING; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 27:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = DEFAULT; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 28:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = COPY; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 29:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = TO; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 30:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = STDOUT; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 31:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = LIMIT; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 32:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = JOIN; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 33:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = CROSS; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 34:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = FULL; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 35:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = OUTER_P; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 36:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = INNER_P; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 37:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = ON; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 39:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = VACUUM; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 40:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = CLUSTER; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 41:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = ANALYZE; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 42:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = int(IDENT); {( lex.p)++;  lex.cs = 2; goto _out }}
	case 52:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = TPLUS; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 53:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = TMINUS; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 54:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = TMUL; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 55:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = TMOD; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 56:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = TPOW; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 59:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = TEQ; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 60:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = TNOT_EQUALS; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 61:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = TLESS_EQUALS; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 62:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = TGREATER_EQUALS; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 63:
	{( lex.p) = ( lex.te) - 1
 lval.str = string(lex.data[lex.ts:lex.te]); tok = TNOT_EQUALS; {( lex.p)++;  lex.cs = 2; goto _out }}
	case 64:
	{( lex.p) = ( lex.te) - 1
 



                lval.str = string(lex.data[lex.ts:lex.te]); tok = int(OP);    
                {( lex.p)++;  lex.cs = 2; goto _out }
            }
	}
	
	goto st2
tr47:
//line lyx/lexer.rl:98
 lex.te = ( lex.p)
( lex.p)--
{ lval.str = string(lex.data[lex.ts:lex.te]); tok = SCONST; {( lex.p)++;  lex.cs = 2; goto _out }}
	goto st2
tr48:
//line lyx/lexer.rl:167
 lex.te = ( lex.p)
( lex.p)--
{ lval.str = string(lex.data[lex.ts:lex.te]); tok = TLESS; {( lex.p)++;  lex.cs = 2; goto _out }}
	goto st2
tr51:
//line lyx/lexer.rl:168
 lex.te = ( lex.p)
( lex.p)--
{ lval.str = string(lex.data[lex.ts:lex.te]); tok = TGREATER; {( lex.p)++;  lex.cs = 2; goto _out }}
	goto st2
tr53:
//line lyx/lexer.rl:147
 lex.te = ( lex.p)
( lex.p)--
{ lval.str = string(lex.data[lex.ts:lex.te]); tok = int(IDENT); {( lex.p)++;  lex.cs = 2; goto _out }}
	goto st2
tr101:
//line lyx/lexer.rl:140
 lex.te = ( lex.p)
( lex.p)--
{ lval.str = string(lex.data[lex.ts:lex.te]); tok = FOR; {( lex.p)++;  lex.cs = 2; goto _out }}
	goto st2
tr136:
//line lyx/lexer.rl:122
 lex.te = ( lex.p)
( lex.p)--
{ lval.str = string(lex.data[lex.ts:lex.te]); tok = OR; {( lex.p)++;  lex.cs = 2; goto _out }}
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

//line lyx/lexer.go:632
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
			goto st10
		case 65:
			goto st11
		case 66:
			goto st17
		case 67:
			goto st18
		case 68:
			goto st32
		case 70:
			goto st47
		case 71:
			goto st57
		case 73:
			goto st61
		case 74:
			goto st69
		case 75:
			goto st72
		case 76:
			goto st74
		case 79:
			goto st78
		case 80:
			goto st85
		case 82:
			goto st91
		case 83:
			goto st108
		case 84:
			goto st117
		case 85:
			goto st121
		case 86:
			goto st126
		case 87:
			goto st134
		case 91:
			goto tr40
		case 92:
			goto tr5
		case 93:
			goto tr41
		case 94:
			goto tr42
		case 96:
			goto tr5
		case 97:
			goto st11
		case 98:
			goto st17
		case 99:
			goto st18
		case 100:
			goto st32
		case 102:
			goto st47
		case 103:
			goto st57
		case 105:
			goto st61
		case 106:
			goto st69
		case 107:
			goto st72
		case 108:
			goto st74
		case 111:
			goto st78
		case 112:
			goto st85
		case 114:
			goto st91
		case 115:
			goto st108
		case 116:
			goto st117
		case 117:
			goto st121
		case 118:
			goto st126
		case 119:
			goto st134
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
				if 69 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto tr25
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
		goto tr43
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
			goto tr45
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
		goto tr44
tr5:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:177
 lex.act = 64;
	goto st5
tr6:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:165
 lex.act = 55;
	goto st5
tr9:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:163
 lex.act = 54;
	goto st5
tr10:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:161
 lex.act = 52;
	goto st5
tr12:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:162
 lex.act = 53;
	goto st5
tr19:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:169
 lex.act = 59;
	goto st5
tr42:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:166
 lex.act = 56;
	goto st5
tr45:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:174
 lex.act = 63;
	goto st5
tr49:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:172
 lex.act = 61;
	goto st5
tr50:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:171
 lex.act = 60;
	goto st5
tr52:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:173
 lex.act = 62;
	goto st5
	st5:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof5
		}
	st_case_5:
//line lyx/lexer.go:918
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
		goto tr46
	st1:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof1
		}
	st_case_1:
		if  lex.data[( lex.p)] == 39 {
			goto tr1
		}
		goto st1
	st6:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof6
		}
	st_case_6:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 95:
			goto tr25
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto st6
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr47
tr25:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:147
 lex.act = 42;
	goto st7
tr55:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:120
 lex.act = 23;
	goto st7
tr57:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:121
 lex.act = 24;
	goto st7
tr61:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:144
 lex.act = 41;
	goto st7
tr62:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:119
 lex.act = 22;
	goto st7
tr70:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:143
 lex.act = 40;
	goto st7
tr72:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:127
 lex.act = 28;
	goto st7
tr77:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:106
 lex.act = 9;
	goto st7
tr79:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:135
 lex.act = 33;
	goto st7
tr87:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:108
 lex.act = 11;
	goto st7
tr93:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:125
 lex.act = 27;
	goto st7
tr96:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:105
 lex.act = 8;
	goto st7
tr105:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:111
 lex.act = 14;
	goto st7
tr107:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:115
 lex.act = 18;
	goto st7
tr109:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:136
 lex.act = 34;
	goto st7
tr113:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:118
 lex.act = 21;
	goto st7
tr119:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:138
 lex.act = 36;
	goto st7
tr122:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:101
 lex.act = 4;
	goto st7
tr123:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:102
 lex.act = 5;
	goto st7
tr126:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:134
 lex.act = 32;
	goto st7
tr128:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:113
 lex.act = 16;
	goto st7
tr132:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:131
 lex.act = 31;
	goto st7
tr133:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:139
 lex.act = 37;
	goto st7
tr139:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:117
 lex.act = 20;
	goto st7
tr142:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:137
 lex.act = 35;
	goto st7
tr148:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:110
 lex.act = 13;
	goto st7
tr159:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:112
 lex.act = 15;
	goto st7
tr165:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:124
 lex.act = 26;
	goto st7
tr167:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:109
 lex.act = 12;
	goto st7
tr171:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:114
 lex.act = 17;
	goto st7
tr174:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:100
 lex.act = 3;
	goto st7
tr178:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:129
 lex.act = 30;
	goto st7
tr180:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:128
 lex.act = 29;
	goto st7
tr183:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:107
 lex.act = 10;
	goto st7
tr188:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:104
 lex.act = 7;
	goto st7
tr194:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:142
 lex.act = 39;
	goto st7
tr197:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:103
 lex.act = 6;
	goto st7
tr201:
//line NONE:1
 lex.te = ( lex.p)+1

//line lyx/lexer.rl:116
 lex.act = 19;
	goto st7
	st7:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof7
		}
	st_case_7:
//line lyx/lexer.go:1254
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 95:
			goto tr25
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr46
	st8:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof8
		}
	st_case_8:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st8
		}
		goto tr47
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
		case 61:
			goto tr49
		case 62:
			goto tr50
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
		goto tr48
	st10:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof10
		}
	st_case_10:
		switch  lex.data[( lex.p)] {
		case 33:
			goto tr5
		case 35:
			goto tr5
		case 45:
			goto tr5
		case 61:
			goto tr52
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
		goto tr51
	st11:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof11
		}
	st_case_11:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 78:
			goto st12
		case 83:
			goto tr55
		case 95:
			goto tr25
		case 110:
			goto st12
		case 115:
			goto tr55
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st12:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof12
		}
	st_case_12:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 65:
			goto st13
		case 68:
			goto tr57
		case 95:
			goto tr25
		case 97:
			goto st13
		case 100:
			goto tr57
		}
		switch {
		case  lex.data[( lex.p)] < 66:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st13:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof13
		}
	st_case_13:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 76:
			goto st14
		case 95:
			goto tr25
		case 108:
			goto st14
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st14:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof14
		}
	st_case_14:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 89:
			goto st15
		case 95:
			goto tr25
		case 121:
			goto st15
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st15:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof15
		}
	st_case_15:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 90:
			goto st16
		case 95:
			goto tr25
		case 122:
			goto st16
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 89:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 121 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st16:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof16
		}
	st_case_16:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 69:
			goto tr61
		case 95:
			goto tr25
		case 101:
			goto tr61
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st17:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof17
		}
	st_case_17:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 89:
			goto tr62
		case 95:
			goto tr25
		case 121:
			goto tr62
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st18:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof18
		}
	st_case_18:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 76:
			goto st19
		case 79:
			goto st24
		case 82:
			goto st26
		case 95:
			goto tr25
		case 108:
			goto st19
		case 111:
			goto st24
		case 114:
			goto st26
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st19:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof19
		}
	st_case_19:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 85:
			goto st20
		case 95:
			goto tr25
		case 117:
			goto st20
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st20:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof20
		}
	st_case_20:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 83:
			goto st21
		case 95:
			goto tr25
		case 115:
			goto st21
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st21:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof21
		}
	st_case_21:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 84:
			goto st22
		case 95:
			goto tr25
		case 116:
			goto st22
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st22:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof22
		}
	st_case_22:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 69:
			goto st23
		case 95:
			goto tr25
		case 101:
			goto st23
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st23:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof23
		}
	st_case_23:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 82:
			goto tr70
		case 95:
			goto tr25
		case 114:
			goto tr70
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st24:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof24
		}
	st_case_24:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 80:
			goto st25
		case 95:
			goto tr25
		case 112:
			goto st25
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st25:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof25
		}
	st_case_25:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 89:
			goto tr72
		case 95:
			goto tr25
		case 121:
			goto tr72
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st26:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof26
		}
	st_case_26:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 69:
			goto st27
		case 79:
			goto st30
		case 95:
			goto tr25
		case 101:
			goto st27
		case 111:
			goto st30
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st27:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof27
		}
	st_case_27:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 65:
			goto st28
		case 95:
			goto tr25
		case 97:
			goto st28
		}
		switch {
		case  lex.data[( lex.p)] < 66:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st28:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof28
		}
	st_case_28:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 84:
			goto st29
		case 95:
			goto tr25
		case 116:
			goto st29
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st29:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof29
		}
	st_case_29:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 69:
			goto tr77
		case 95:
			goto tr25
		case 101:
			goto tr77
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st30:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof30
		}
	st_case_30:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 83:
			goto st31
		case 95:
			goto tr25
		case 115:
			goto st31
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st31:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof31
		}
	st_case_31:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 83:
			goto tr79
		case 95:
			goto tr25
		case 115:
			goto tr79
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st32:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof32
		}
	st_case_32:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 65:
			goto st33
		case 69:
			goto st39
		case 95:
			goto tr25
		case 97:
			goto st33
		case 101:
			goto st39
		}
		switch {
		case  lex.data[( lex.p)] < 66:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st33:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof33
		}
	st_case_33:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 84:
			goto st34
		case 95:
			goto tr25
		case 116:
			goto st34
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st34:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof34
		}
	st_case_34:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 65:
			goto st35
		case 95:
			goto tr25
		case 97:
			goto st35
		}
		switch {
		case  lex.data[( lex.p)] < 66:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st35:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof35
		}
	st_case_35:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 66:
			goto st36
		case 95:
			goto tr25
		case 98:
			goto st36
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st36:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof36
		}
	st_case_36:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 65:
			goto st37
		case 95:
			goto tr25
		case 97:
			goto st37
		}
		switch {
		case  lex.data[( lex.p)] < 66:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st37:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof37
		}
	st_case_37:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 83:
			goto st38
		case 95:
			goto tr25
		case 115:
			goto st38
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st38:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof38
		}
	st_case_38:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 69:
			goto tr87
		case 95:
			goto tr25
		case 101:
			goto tr87
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st39:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof39
		}
	st_case_39:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 70:
			goto st40
		case 76:
			goto st44
		case 95:
			goto tr25
		case 102:
			goto st40
		case 108:
			goto st44
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st40:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof40
		}
	st_case_40:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 65:
			goto st41
		case 95:
			goto tr25
		case 97:
			goto st41
		}
		switch {
		case  lex.data[( lex.p)] < 66:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st41:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof41
		}
	st_case_41:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 85:
			goto st42
		case 95:
			goto tr25
		case 117:
			goto st42
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st42:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof42
		}
	st_case_42:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 76:
			goto st43
		case 95:
			goto tr25
		case 108:
			goto st43
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st43:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof43
		}
	st_case_43:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 84:
			goto tr93
		case 95:
			goto tr25
		case 116:
			goto tr93
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st44:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof44
		}
	st_case_44:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 69:
			goto st45
		case 95:
			goto tr25
		case 101:
			goto st45
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st45:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof45
		}
	st_case_45:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 84:
			goto st46
		case 95:
			goto tr25
		case 116:
			goto st46
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st46:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof46
		}
	st_case_46:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 69:
			goto tr96
		case 95:
			goto tr25
		case 101:
			goto tr96
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st47:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof47
		}
	st_case_47:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 79:
			goto st48
		case 82:
			goto st53
		case 85:
			goto st55
		case 95:
			goto tr25
		case 111:
			goto st48
		case 114:
			goto st53
		case 117:
			goto st55
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st48:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof48
		}
	st_case_48:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 82:
			goto st49
		case 95:
			goto tr25
		case 114:
			goto st49
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st49:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof49
		}
	st_case_49:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 69:
			goto st50
		case 95:
			goto tr25
		case 101:
			goto st50
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr101
	st50:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof50
		}
	st_case_50:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 73:
			goto st51
		case 95:
			goto tr25
		case 105:
			goto st51
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st51:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof51
		}
	st_case_51:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 71:
			goto st52
		case 95:
			goto tr25
		case 103:
			goto st52
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st52:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof52
		}
	st_case_52:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 78:
			goto tr105
		case 95:
			goto tr25
		case 110:
			goto tr105
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st53:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof53
		}
	st_case_53:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 79:
			goto st54
		case 95:
			goto tr25
		case 111:
			goto st54
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st54:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof54
		}
	st_case_54:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 77:
			goto tr107
		case 95:
			goto tr25
		case 109:
			goto tr107
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st55:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof55
		}
	st_case_55:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 76:
			goto st56
		case 95:
			goto tr25
		case 108:
			goto st56
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st56:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof56
		}
	st_case_56:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 76:
			goto tr109
		case 95:
			goto tr25
		case 108:
			goto tr109
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st57:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof57
		}
	st_case_57:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 82:
			goto st58
		case 95:
			goto tr25
		case 114:
			goto st58
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st58:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof58
		}
	st_case_58:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 79:
			goto st59
		case 95:
			goto tr25
		case 111:
			goto st59
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st59:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof59
		}
	st_case_59:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 85:
			goto st60
		case 95:
			goto tr25
		case 117:
			goto st60
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st60:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof60
		}
	st_case_60:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 80:
			goto tr113
		case 95:
			goto tr25
		case 112:
			goto tr113
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st61:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof61
		}
	st_case_61:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 78:
			goto st62
		case 95:
			goto tr25
		case 110:
			goto st62
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st62:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof62
		}
	st_case_62:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 78:
			goto st63
		case 83:
			goto st65
		case 84:
			goto st68
		case 95:
			goto tr25
		case 110:
			goto st63
		case 115:
			goto st65
		case 116:
			goto st68
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st63:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof63
		}
	st_case_63:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 69:
			goto st64
		case 95:
			goto tr25
		case 101:
			goto st64
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st64:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof64
		}
	st_case_64:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 82:
			goto tr119
		case 95:
			goto tr25
		case 114:
			goto tr119
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st65:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof65
		}
	st_case_65:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 69:
			goto st66
		case 95:
			goto tr25
		case 101:
			goto st66
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st66:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof66
		}
	st_case_66:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 82:
			goto st67
		case 95:
			goto tr25
		case 114:
			goto st67
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st67:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof67
		}
	st_case_67:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 84:
			goto tr122
		case 95:
			goto tr25
		case 116:
			goto tr122
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st68:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof68
		}
	st_case_68:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 79:
			goto tr123
		case 95:
			goto tr25
		case 111:
			goto tr123
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st69:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof69
		}
	st_case_69:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 79:
			goto st70
		case 95:
			goto tr25
		case 111:
			goto st70
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st70:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof70
		}
	st_case_70:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 73:
			goto st71
		case 95:
			goto tr25
		case 105:
			goto st71
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st71:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof71
		}
	st_case_71:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 78:
			goto tr126
		case 95:
			goto tr25
		case 110:
			goto tr126
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st72:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof72
		}
	st_case_72:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 69:
			goto st73
		case 95:
			goto tr25
		case 101:
			goto st73
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st73:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof73
		}
	st_case_73:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 89:
			goto tr128
		case 95:
			goto tr25
		case 121:
			goto tr128
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st74:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof74
		}
	st_case_74:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 73:
			goto st75
		case 95:
			goto tr25
		case 105:
			goto st75
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st75:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof75
		}
	st_case_75:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 77:
			goto st76
		case 95:
			goto tr25
		case 109:
			goto st76
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st76:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof76
		}
	st_case_76:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 73:
			goto st77
		case 95:
			goto tr25
		case 105:
			goto st77
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st77:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof77
		}
	st_case_77:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 84:
			goto tr132
		case 95:
			goto tr25
		case 116:
			goto tr132
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st78:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof78
		}
	st_case_78:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 78:
			goto tr133
		case 82:
			goto st79
		case 85:
			goto st82
		case 95:
			goto tr25
		case 110:
			goto tr133
		case 114:
			goto st79
		case 117:
			goto st82
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st79:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof79
		}
	st_case_79:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 68:
			goto st80
		case 95:
			goto tr25
		case 100:
			goto st80
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr136
	st80:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof80
		}
	st_case_80:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 69:
			goto st81
		case 95:
			goto tr25
		case 101:
			goto st81
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st81:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof81
		}
	st_case_81:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 82:
			goto tr139
		case 95:
			goto tr25
		case 114:
			goto tr139
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st82:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof82
		}
	st_case_82:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 84:
			goto st83
		case 95:
			goto tr25
		case 116:
			goto st83
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st83:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof83
		}
	st_case_83:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 69:
			goto st84
		case 95:
			goto tr25
		case 101:
			goto st84
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st84:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof84
		}
	st_case_84:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 82:
			goto tr142
		case 95:
			goto tr25
		case 114:
			goto tr142
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st85:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof85
		}
	st_case_85:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 82:
			goto st86
		case 95:
			goto tr25
		case 114:
			goto st86
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st86:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof86
		}
	st_case_86:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 73:
			goto st87
		case 95:
			goto tr25
		case 105:
			goto st87
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st87:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof87
		}
	st_case_87:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 77:
			goto st88
		case 95:
			goto tr25
		case 109:
			goto st88
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st88:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof88
		}
	st_case_88:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 65:
			goto st89
		case 95:
			goto tr25
		case 97:
			goto st89
		}
		switch {
		case  lex.data[( lex.p)] < 66:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st89:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof89
		}
	st_case_89:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 82:
			goto st90
		case 95:
			goto tr25
		case 114:
			goto st90
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st90:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof90
		}
	st_case_90:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 89:
			goto tr148
		case 95:
			goto tr25
		case 121:
			goto tr148
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st91:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof91
		}
	st_case_91:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 69:
			goto st92
		case 79:
			goto st106
		case 95:
			goto tr25
		case 101:
			goto st92
		case 111:
			goto st106
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st92:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof92
		}
	st_case_92:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 70:
			goto st93
		case 84:
			goto st100
		case 95:
			goto tr25
		case 102:
			goto st93
		case 116:
			goto st100
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st93:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof93
		}
	st_case_93:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 69:
			goto st94
		case 95:
			goto tr25
		case 101:
			goto st94
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st94:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof94
		}
	st_case_94:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 82:
			goto st95
		case 95:
			goto tr25
		case 114:
			goto st95
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st95:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof95
		}
	st_case_95:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 69:
			goto st96
		case 95:
			goto tr25
		case 101:
			goto st96
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st96:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof96
		}
	st_case_96:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 78:
			goto st97
		case 95:
			goto tr25
		case 110:
			goto st97
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st97:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof97
		}
	st_case_97:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 67:
			goto st98
		case 95:
			goto tr25
		case 99:
			goto st98
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st98:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof98
		}
	st_case_98:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 69:
			goto st99
		case 95:
			goto tr25
		case 101:
			goto st99
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st99:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof99
		}
	st_case_99:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 83:
			goto tr159
		case 95:
			goto tr25
		case 115:
			goto tr159
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st100:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof100
		}
	st_case_100:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 85:
			goto st101
		case 95:
			goto tr25
		case 117:
			goto st101
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st101:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof101
		}
	st_case_101:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 82:
			goto st102
		case 95:
			goto tr25
		case 114:
			goto st102
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st102:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof102
		}
	st_case_102:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 78:
			goto st103
		case 95:
			goto tr25
		case 110:
			goto st103
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st103:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof103
		}
	st_case_103:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 73:
			goto st104
		case 95:
			goto tr25
		case 105:
			goto st104
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st104:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof104
		}
	st_case_104:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 78:
			goto st105
		case 95:
			goto tr25
		case 110:
			goto st105
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st105:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof105
		}
	st_case_105:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 71:
			goto tr165
		case 95:
			goto tr25
		case 103:
			goto tr165
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st106:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof106
		}
	st_case_106:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 76:
			goto st107
		case 95:
			goto tr25
		case 108:
			goto st107
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st107:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof107
		}
	st_case_107:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 69:
			goto tr167
		case 95:
			goto tr25
		case 101:
			goto tr167
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st108:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof108
		}
	st_case_108:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 69:
			goto st109
		case 84:
			goto st113
		case 95:
			goto tr25
		case 101:
			goto st109
		case 116:
			goto st113
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st109:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof109
		}
	st_case_109:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 76:
			goto st110
		case 84:
			goto tr171
		case 95:
			goto tr25
		case 108:
			goto st110
		case 116:
			goto tr171
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st110:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof110
		}
	st_case_110:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 69:
			goto st111
		case 95:
			goto tr25
		case 101:
			goto st111
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st111:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof111
		}
	st_case_111:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 67:
			goto st112
		case 95:
			goto tr25
		case 99:
			goto st112
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st112:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof112
		}
	st_case_112:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 84:
			goto tr174
		case 95:
			goto tr25
		case 116:
			goto tr174
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st113:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof113
		}
	st_case_113:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 68:
			goto st114
		case 95:
			goto tr25
		case 100:
			goto st114
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st114:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof114
		}
	st_case_114:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 79:
			goto st115
		case 95:
			goto tr25
		case 111:
			goto st115
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st115:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof115
		}
	st_case_115:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 85:
			goto st116
		case 95:
			goto tr25
		case 117:
			goto st116
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st116:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof116
		}
	st_case_116:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 84:
			goto tr178
		case 95:
			goto tr25
		case 116:
			goto tr178
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st117:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof117
		}
	st_case_117:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 65:
			goto st118
		case 79:
			goto tr180
		case 95:
			goto tr25
		case 97:
			goto st118
		case 111:
			goto tr180
		}
		switch {
		case  lex.data[( lex.p)] < 66:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st118:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof118
		}
	st_case_118:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 66:
			goto st119
		case 95:
			goto tr25
		case 98:
			goto st119
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st119:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof119
		}
	st_case_119:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 76:
			goto st120
		case 95:
			goto tr25
		case 108:
			goto st120
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st120:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof120
		}
	st_case_120:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 69:
			goto tr183
		case 95:
			goto tr25
		case 101:
			goto tr183
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st121:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof121
		}
	st_case_121:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 80:
			goto st122
		case 95:
			goto tr25
		case 112:
			goto st122
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st122:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof122
		}
	st_case_122:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 68:
			goto st123
		case 95:
			goto tr25
		case 100:
			goto st123
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st123:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof123
		}
	st_case_123:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 65:
			goto st124
		case 95:
			goto tr25
		case 97:
			goto st124
		}
		switch {
		case  lex.data[( lex.p)] < 66:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st124:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof124
		}
	st_case_124:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 84:
			goto st125
		case 95:
			goto tr25
		case 116:
			goto st125
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st125:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof125
		}
	st_case_125:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 69:
			goto tr188
		case 95:
			goto tr25
		case 101:
			goto tr188
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st126:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof126
		}
	st_case_126:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 65:
			goto st127
		case 95:
			goto tr25
		case 97:
			goto st127
		}
		switch {
		case  lex.data[( lex.p)] < 66:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 98 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st127:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof127
		}
	st_case_127:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 67:
			goto st128
		case 76:
			goto st131
		case 95:
			goto tr25
		case 99:
			goto st128
		case 108:
			goto st131
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st128:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof128
		}
	st_case_128:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 85:
			goto st129
		case 95:
			goto tr25
		case 117:
			goto st129
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st129:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof129
		}
	st_case_129:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 85:
			goto st130
		case 95:
			goto tr25
		case 117:
			goto st130
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st130:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof130
		}
	st_case_130:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 77:
			goto tr194
		case 95:
			goto tr25
		case 109:
			goto tr194
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st131:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof131
		}
	st_case_131:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 85:
			goto st132
		case 95:
			goto tr25
		case 117:
			goto st132
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st132:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof132
		}
	st_case_132:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 69:
			goto st133
		case 95:
			goto tr25
		case 101:
			goto st133
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st133:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof133
		}
	st_case_133:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 83:
			goto tr197
		case 95:
			goto tr25
		case 115:
			goto tr197
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st134:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof134
		}
	st_case_134:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 72:
			goto st135
		case 95:
			goto tr25
		case 104:
			goto st135
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st135:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof135
		}
	st_case_135:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 69:
			goto st136
		case 95:
			goto tr25
		case 101:
			goto st136
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st136:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof136
		}
	st_case_136:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 82:
			goto st137
		case 95:
			goto tr25
		case 114:
			goto st137
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
	st137:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof137
		}
	st_case_137:
		switch  lex.data[( lex.p)] {
		case 36:
			goto tr25
		case 69:
			goto tr201
		case 95:
			goto tr25
		case 101:
			goto tr201
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr25
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr53
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
	_test_eof41:  lex.cs = 41; goto _test_eof
	_test_eof42:  lex.cs = 42; goto _test_eof
	_test_eof43:  lex.cs = 43; goto _test_eof
	_test_eof44:  lex.cs = 44; goto _test_eof
	_test_eof45:  lex.cs = 45; goto _test_eof
	_test_eof46:  lex.cs = 46; goto _test_eof
	_test_eof47:  lex.cs = 47; goto _test_eof
	_test_eof48:  lex.cs = 48; goto _test_eof
	_test_eof49:  lex.cs = 49; goto _test_eof
	_test_eof50:  lex.cs = 50; goto _test_eof
	_test_eof51:  lex.cs = 51; goto _test_eof
	_test_eof52:  lex.cs = 52; goto _test_eof
	_test_eof53:  lex.cs = 53; goto _test_eof
	_test_eof54:  lex.cs = 54; goto _test_eof
	_test_eof55:  lex.cs = 55; goto _test_eof
	_test_eof56:  lex.cs = 56; goto _test_eof
	_test_eof57:  lex.cs = 57; goto _test_eof
	_test_eof58:  lex.cs = 58; goto _test_eof
	_test_eof59:  lex.cs = 59; goto _test_eof
	_test_eof60:  lex.cs = 60; goto _test_eof
	_test_eof61:  lex.cs = 61; goto _test_eof
	_test_eof62:  lex.cs = 62; goto _test_eof
	_test_eof63:  lex.cs = 63; goto _test_eof
	_test_eof64:  lex.cs = 64; goto _test_eof
	_test_eof65:  lex.cs = 65; goto _test_eof
	_test_eof66:  lex.cs = 66; goto _test_eof
	_test_eof67:  lex.cs = 67; goto _test_eof
	_test_eof68:  lex.cs = 68; goto _test_eof
	_test_eof69:  lex.cs = 69; goto _test_eof
	_test_eof70:  lex.cs = 70; goto _test_eof
	_test_eof71:  lex.cs = 71; goto _test_eof
	_test_eof72:  lex.cs = 72; goto _test_eof
	_test_eof73:  lex.cs = 73; goto _test_eof
	_test_eof74:  lex.cs = 74; goto _test_eof
	_test_eof75:  lex.cs = 75; goto _test_eof
	_test_eof76:  lex.cs = 76; goto _test_eof
	_test_eof77:  lex.cs = 77; goto _test_eof
	_test_eof78:  lex.cs = 78; goto _test_eof
	_test_eof79:  lex.cs = 79; goto _test_eof
	_test_eof80:  lex.cs = 80; goto _test_eof
	_test_eof81:  lex.cs = 81; goto _test_eof
	_test_eof82:  lex.cs = 82; goto _test_eof
	_test_eof83:  lex.cs = 83; goto _test_eof
	_test_eof84:  lex.cs = 84; goto _test_eof
	_test_eof85:  lex.cs = 85; goto _test_eof
	_test_eof86:  lex.cs = 86; goto _test_eof
	_test_eof87:  lex.cs = 87; goto _test_eof
	_test_eof88:  lex.cs = 88; goto _test_eof
	_test_eof89:  lex.cs = 89; goto _test_eof
	_test_eof90:  lex.cs = 90; goto _test_eof
	_test_eof91:  lex.cs = 91; goto _test_eof
	_test_eof92:  lex.cs = 92; goto _test_eof
	_test_eof93:  lex.cs = 93; goto _test_eof
	_test_eof94:  lex.cs = 94; goto _test_eof
	_test_eof95:  lex.cs = 95; goto _test_eof
	_test_eof96:  lex.cs = 96; goto _test_eof
	_test_eof97:  lex.cs = 97; goto _test_eof
	_test_eof98:  lex.cs = 98; goto _test_eof
	_test_eof99:  lex.cs = 99; goto _test_eof
	_test_eof100:  lex.cs = 100; goto _test_eof
	_test_eof101:  lex.cs = 101; goto _test_eof
	_test_eof102:  lex.cs = 102; goto _test_eof
	_test_eof103:  lex.cs = 103; goto _test_eof
	_test_eof104:  lex.cs = 104; goto _test_eof
	_test_eof105:  lex.cs = 105; goto _test_eof
	_test_eof106:  lex.cs = 106; goto _test_eof
	_test_eof107:  lex.cs = 107; goto _test_eof
	_test_eof108:  lex.cs = 108; goto _test_eof
	_test_eof109:  lex.cs = 109; goto _test_eof
	_test_eof110:  lex.cs = 110; goto _test_eof
	_test_eof111:  lex.cs = 111; goto _test_eof
	_test_eof112:  lex.cs = 112; goto _test_eof
	_test_eof113:  lex.cs = 113; goto _test_eof
	_test_eof114:  lex.cs = 114; goto _test_eof
	_test_eof115:  lex.cs = 115; goto _test_eof
	_test_eof116:  lex.cs = 116; goto _test_eof
	_test_eof117:  lex.cs = 117; goto _test_eof
	_test_eof118:  lex.cs = 118; goto _test_eof
	_test_eof119:  lex.cs = 119; goto _test_eof
	_test_eof120:  lex.cs = 120; goto _test_eof
	_test_eof121:  lex.cs = 121; goto _test_eof
	_test_eof122:  lex.cs = 122; goto _test_eof
	_test_eof123:  lex.cs = 123; goto _test_eof
	_test_eof124:  lex.cs = 124; goto _test_eof
	_test_eof125:  lex.cs = 125; goto _test_eof
	_test_eof126:  lex.cs = 126; goto _test_eof
	_test_eof127:  lex.cs = 127; goto _test_eof
	_test_eof128:  lex.cs = 128; goto _test_eof
	_test_eof129:  lex.cs = 129; goto _test_eof
	_test_eof130:  lex.cs = 130; goto _test_eof
	_test_eof131:  lex.cs = 131; goto _test_eof
	_test_eof132:  lex.cs = 132; goto _test_eof
	_test_eof133:  lex.cs = 133; goto _test_eof
	_test_eof134:  lex.cs = 134; goto _test_eof
	_test_eof135:  lex.cs = 135; goto _test_eof
	_test_eof136:  lex.cs = 136; goto _test_eof
	_test_eof137:  lex.cs = 137; goto _test_eof

	_test_eof: {}
	if ( lex.p) == eof {
		switch  lex.cs {
		case 3:
			goto tr43
		case 4:
			goto tr44
		case 5:
			goto tr46
		case 6:
			goto tr47
		case 7:
			goto tr46
		case 8:
			goto tr47
		case 9:
			goto tr48
		case 10:
			goto tr51
		case 11:
			goto tr53
		case 12:
			goto tr53
		case 13:
			goto tr53
		case 14:
			goto tr53
		case 15:
			goto tr53
		case 16:
			goto tr53
		case 17:
			goto tr53
		case 18:
			goto tr53
		case 19:
			goto tr53
		case 20:
			goto tr53
		case 21:
			goto tr53
		case 22:
			goto tr53
		case 23:
			goto tr53
		case 24:
			goto tr53
		case 25:
			goto tr53
		case 26:
			goto tr53
		case 27:
			goto tr53
		case 28:
			goto tr53
		case 29:
			goto tr53
		case 30:
			goto tr53
		case 31:
			goto tr53
		case 32:
			goto tr53
		case 33:
			goto tr53
		case 34:
			goto tr53
		case 35:
			goto tr53
		case 36:
			goto tr53
		case 37:
			goto tr53
		case 38:
			goto tr53
		case 39:
			goto tr53
		case 40:
			goto tr53
		case 41:
			goto tr53
		case 42:
			goto tr53
		case 43:
			goto tr53
		case 44:
			goto tr53
		case 45:
			goto tr53
		case 46:
			goto tr53
		case 47:
			goto tr53
		case 48:
			goto tr53
		case 49:
			goto tr101
		case 50:
			goto tr53
		case 51:
			goto tr53
		case 52:
			goto tr53
		case 53:
			goto tr53
		case 54:
			goto tr53
		case 55:
			goto tr53
		case 56:
			goto tr53
		case 57:
			goto tr53
		case 58:
			goto tr53
		case 59:
			goto tr53
		case 60:
			goto tr53
		case 61:
			goto tr53
		case 62:
			goto tr53
		case 63:
			goto tr53
		case 64:
			goto tr53
		case 65:
			goto tr53
		case 66:
			goto tr53
		case 67:
			goto tr53
		case 68:
			goto tr53
		case 69:
			goto tr53
		case 70:
			goto tr53
		case 71:
			goto tr53
		case 72:
			goto tr53
		case 73:
			goto tr53
		case 74:
			goto tr53
		case 75:
			goto tr53
		case 76:
			goto tr53
		case 77:
			goto tr53
		case 78:
			goto tr53
		case 79:
			goto tr136
		case 80:
			goto tr53
		case 81:
			goto tr53
		case 82:
			goto tr53
		case 83:
			goto tr53
		case 84:
			goto tr53
		case 85:
			goto tr53
		case 86:
			goto tr53
		case 87:
			goto tr53
		case 88:
			goto tr53
		case 89:
			goto tr53
		case 90:
			goto tr53
		case 91:
			goto tr53
		case 92:
			goto tr53
		case 93:
			goto tr53
		case 94:
			goto tr53
		case 95:
			goto tr53
		case 96:
			goto tr53
		case 97:
			goto tr53
		case 98:
			goto tr53
		case 99:
			goto tr53
		case 100:
			goto tr53
		case 101:
			goto tr53
		case 102:
			goto tr53
		case 103:
			goto tr53
		case 104:
			goto tr53
		case 105:
			goto tr53
		case 106:
			goto tr53
		case 107:
			goto tr53
		case 108:
			goto tr53
		case 109:
			goto tr53
		case 110:
			goto tr53
		case 111:
			goto tr53
		case 112:
			goto tr53
		case 113:
			goto tr53
		case 114:
			goto tr53
		case 115:
			goto tr53
		case 116:
			goto tr53
		case 117:
			goto tr53
		case 118:
			goto tr53
		case 119:
			goto tr53
		case 120:
			goto tr53
		case 121:
			goto tr53
		case 122:
			goto tr53
		case 123:
			goto tr53
		case 124:
			goto tr53
		case 125:
			goto tr53
		case 126:
			goto tr53
		case 127:
			goto tr53
		case 128:
			goto tr53
		case 129:
			goto tr53
		case 130:
			goto tr53
		case 131:
			goto tr53
		case 132:
			goto tr53
		case 133:
			goto tr53
		case 134:
			goto tr53
		case 135:
			goto tr53
		case 136:
			goto tr53
		case 137:
			goto tr53
		}
	}

	_out: {}
	}

//line lyx/lexer.rl:191


    return int(tok);
}