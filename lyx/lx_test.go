package lyx_test

import (
	"testing"

	"github.com/pg-sharding/lyx/lyx"

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
			exp: []int{lyx.SELECT, lyx.ICONST},
			err: nil,
		},
		{
			query: `
			14
			`,
			exp: []int{lyx.ICONST},
			err: nil,
		},
		{
			query: `
			$14
			`,
			exp: []int{lyx.PARAM},
			err: nil,
		},
		{
			query: `
			'$14'
			`,
			exp: []int{lyx.SCONST},
			err: nil,
		},
		{
			query: `
			select /*jjiewjow****/1
			`,
			exp: []int{lyx.SELECT, lyx.ICONST},
			err: nil,
		},
		{
			query: `
			update table x set i = 19;
			`,
			exp: []int{lyx.UPDATE, lyx.TABLE, lyx.IDENT, lyx.SET, lyx.IDENT, lyx.TEQ, lyx.ICONST, lyx.TSEMICOLON},
			err: nil,
		},
		{
			query: `
			select * from xx where i = 1 order by i;
			`,
			exp: []int{lyx.SELECT, lyx.TMUL, lyx.FROM, lyx.IDENT, lyx.WHERE, lyx.IDENT, lyx.TEQ, lyx.ICONST, lyx.ORDER, lyx.BY, lyx.IDENT, lyx.TSEMICOLON},
			err: nil,
		},
		{
			query: `
			select * from xx where i + j != 1;
			`,
			exp: []int{lyx.SELECT, lyx.TMUL, lyx.FROM, lyx.IDENT, lyx.WHERE, lyx.IDENT, lyx.TPLUS, lyx.IDENT, lyx.TNOT_EQUALS, lyx.ICONST, lyx.TSEMICOLON},
			err: nil,
		},
		{
			query: `
			SELECT kind, sum(len) AS
			 total FROM films GROUP BY kind;
			`,
			exp: []int{
				lyx.SELECT, lyx.IDENT, lyx.TCOMMA,
				lyx.IDENT, lyx.TOPENBR, lyx.IDENT, lyx.TCLOSEBR, lyx.AS, lyx.IDENT,
				lyx.FROM, lyx.IDENT, lyx.GROUP, lyx.BY, lyx.IDENT, lyx.TSEMICOLON},
			err: nil,
		},
		{
			query: `
			PREPARE P(TEXT) AS SELECT $1
			`,
			exp: []int{
				lyx.PREPARE, lyx.IDENT,
				lyx.TOPENBR, lyx.IDENT,
				lyx.TCLOSEBR, lyx.AS, lyx.SELECT, lyx.PARAM,
			},
			err: nil,
		},
		{
			query: `
			INSERT INTO xx VALUES
				(1,2,'gyuguy');
			`,
			exp: []int{
				lyx.INSERT, lyx.INTO, lyx.IDENT,
				lyx.VALUES,
				lyx.TOPENBR,
				lyx.ICONST, lyx.TCOMMA, lyx.ICONST, lyx.TCOMMA, lyx.SCONST,
				lyx.TCLOSEBR, lyx.TSEMICOLON},
			err: nil,
		},
		{
			query: `INSERT INTO films (code, title, did, date_prod, kind)
			VALUES ('T_601', 'Yojimbo', 106, '1961-06-16', 'Drama');`,
			exp: []int{
				lyx.INSERT, lyx.INTO, lyx.IDENT,

				lyx.TOPENBR,

				lyx.IDENT,

				lyx.TCOMMA,
				lyx.IDENT,

				lyx.TCOMMA,
				lyx.IDENT,

				lyx.TCOMMA,
				lyx.IDENT,

				lyx.TCOMMA,
				lyx.IDENT,

				lyx.TCLOSEBR,
				lyx.VALUES,
				lyx.TOPENBR,
				lyx.SCONST,

				lyx.TCOMMA,
				lyx.SCONST,

				lyx.TCOMMA,
				lyx.ICONST,

				lyx.TCOMMA,
				lyx.SCONST,

				lyx.TCOMMA,
				lyx.SCONST,

				lyx.TCLOSEBR, lyx.TSEMICOLON},
			err: nil,
		},
		{
			query: `
			UPDATE customer1
        	SET c_balance= -4755.000000
			`,
			exp: []int{
				lyx.UPDATE, lyx.IDENT, lyx.SET,
				lyx.IDENT, lyx.TEQ, lyx.SCONST,
			},
			err: nil,
		},
		{
			query: `
			SELECT 1 UNION SELECT 2
			`,
			exp: []int{
				lyx.SELECT, lyx.ICONST, lyx.UNION, lyx.SELECT, lyx.ICONST,
			},
			err: nil,
		},
		{
			query: `
			SELECT * FROM t UNION ALL SELECT * FROM a
			`,
			exp: []int{
				lyx.SELECT, lyx.TMUL, lyx.FROM, lyx.IDENT, lyx.UNION, lyx.ALL, lyx.SELECT, lyx.TMUL, lyx.FROM, lyx.IDENT,
			},
			err: nil,
		},
		{
			query: `select 1; select 2;`,
			exp: []int{
				lyx.SELECT,
				lyx.ICONST,
				lyx.TSEMICOLON,
				lyx.SELECT,
				lyx.ICONST,
				lyx.TSEMICOLON,
			},
		},
		{
			query: `create table "xxx" ( i int )`,
			exp: []int{
				lyx.CREATE,
				lyx.TABLE,
				lyx.IDENT,
				lyx.TOPENBR,
				lyx.IDENT,
				lyx.INT_P,
				lyx.TCLOSEBR,
			},
		},
		{
			query: `create table "xx-x" ( i int )`,
			exp: []int{
				lyx.CREATE,
				lyx.TABLE,
				lyx.IDENT,
				lyx.TOPENBR,
				lyx.IDENT,
				lyx.INT_P,
				lyx.TCLOSEBR,
			},
		},
		{
			query: `create table sh.'xx-x' ( i int )`,
			exp: []int{
				lyx.CREATE,
				lyx.TABLE,
				lyx.IDENT,
				lyx.TDOT,
				lyx.SCONST,
				lyx.TOPENBR,
				lyx.IDENT,
				lyx.INT_P,
				lyx.TCLOSEBR,
			},
		},
		{
			query: `DROP table sh.'xxx'`,
			exp: []int{
				lyx.DROP,
				lyx.TABLE,
				lyx.IDENT,
				lyx.TDOT,
				lyx.SCONST,
			},
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

		assert.Equal(tt.exp, res, tt.query)
	}
}
