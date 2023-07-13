package lyx_test

import (
	"lyx/lyx"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   []int
		err   error
	}

	for _, tt := range []tcase{
		{
			query: `
			select 1
			`,
			exp: []int{lyx.SELECT, lyx.SCONST},
			err: nil,
		},
		{
			query: `
			update table x set i = 19;
			`,
			exp: []int{lyx.UPDATE, lyx.TABLE, lyx.IDENT, lyx.SET, lyx.IDENT, lyx.TEQ, lyx.SCONST, lyx.TSEMICOLON},
			err: nil,
		},
		{
			query: `
			select * from xx where i = 1 order by i;
			`,
			exp: []int{lyx.SELECT, lyx.TMUL, lyx.FROM, lyx.IDENT, lyx.WHERE, lyx.IDENT, lyx.TEQ, lyx.SCONST, lyx.ORDER, lyx.BY, lyx.IDENT, lyx.TSEMICOLON},
			err: nil,
		},
		{
			query: `
			select * from xx where i + j != 1;
			`,
			exp: []int{lyx.SELECT, lyx.TMUL, lyx.FROM, lyx.IDENT, lyx.WHERE, lyx.IDENT, lyx.TPLUS, lyx.IDENT, lyx.TNOT_EQUALS, lyx.SCONST, lyx.TSEMICOLON},
			err: nil,
		},
	} {
		t := lyx.NewStringTokenizer(tt.query)
		var res []int
		for {
			v := t.LexT()
			if v == 0 {
				break
			}
			res = append(res, v)
		}

		assert.Equal(tt.exp, res)
	}
}
