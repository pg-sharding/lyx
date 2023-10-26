package lyx_test

import (
	"testing"

	"github.com/pg-sharding/lyx/lyx"

	"github.com/stretchr/testify/assert"
)

func TestSelectComplex(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: `
			SELECT * FROM "exschema"."extable" WHERE id='83912839012903' AND utype='2' AND btype='sample' AND state = 0 AND is_something = true AND (keys @> '{reshke,denchick}' OR keys @> '{munakoiso,werelaxe,x4mmm}') AND c_id = 'trunk' ORDER BY entity_id asc;
			`,
			exp: &lyx.Select{

				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{&lyx.RangeVar{
					RelationName: "extable",
					SchemaName:   "exschema",
				},
				},

				Where: &lyx.AExprOp{
					Left: &lyx.AExprOp{
						Left: &lyx.AExprOp{
							Left: &lyx.AExprOp{
								Left: &lyx.AExprOp{
									Left: &lyx.AExprOp{
										Left: &lyx.AExprOp{
											Left: &lyx.ColumnRef{
												ColName: "id",
											},
											Right: &lyx.AExprConst{
												Value: "83912839012903",
											},
											Op: "=",
										},
										Right: &lyx.AExprOp{
											Left: &lyx.ColumnRef{
												ColName: "utype",
											},
											Right: &lyx.AExprConst{
												Value: "2",
											},
											Op: "=",
										},
										Op: "AND",
									},
									Right: &lyx.AExprOp{
										Left: &lyx.ColumnRef{
											ColName: "btype",
										},
										Right: &lyx.AExprConst{
											Value: "sample",
										},
										Op: "=",
									},
									Op: "AND",
								},
								Right: &lyx.AExprOp{
									Left: &lyx.ColumnRef{
										ColName: "state",
									},
									Right: &lyx.AExprConst{
										Value: "0",
									},
									Op: "=",
								},
								Op: "AND",
							},
							Right: &lyx.AExprOp{
								Left: &lyx.ColumnRef{
									ColName: "is_something",
								},
								Right: &lyx.AExprConst{
									Value: "true",
								},
								Op: "=",
							},
							Op: "AND",
						},
						Right: &lyx.AExprOp{
							Left: &lyx.ColumnRef{
								ColName: "keys",
							},
							Right: &lyx.AExprOp{
								Left: &lyx.AExprConst{
									Value: "{reshke,denchick}",
								},
								Right: &lyx.AExprOp{
									Left: &lyx.ColumnRef{
										ColName: "keys",
									},
									Right: &lyx.AExprConst{
										Value: "{munakoiso,werelaxe,x4mmm}",
									},
									Op: "@>",
								},
								Op: "OR",
							},
							Op: "@>",
						},
						Op: "AND",
					},
					Right: &lyx.AExprOp{
						Left: &lyx.ColumnRef{
							ColName: "c_id",
						},
						Right: &lyx.AExprConst{
							Value: "trunk",
						},
						Op: "=",
					},
					Op: "AND",
				},
			},

			err: nil,
		},

		{
			query: `
	SELECT *
	FROM "exschema"."extable"
	WHERE
		id='83912839012903'
		AND utype='2'
		AND btype='sample'
		AND state = 0
		AND is_something = true
		AND (keys @> '{reshke,denchick}' OR keys @> '{munakoiso,werelaxe,x4mmm}')
		AND c_id = 'trunk'
	ORDER BY entity_id asc;
			`,
			exp: &lyx.Select{

				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{&lyx.RangeVar{
					RelationName: "extable",
					SchemaName:   "exschema",
				},
				},
				Where: &lyx.AExprOp{
					Left: &lyx.AExprOp{
						Left: &lyx.AExprOp{
							Left: &lyx.AExprOp{
								Left: &lyx.AExprOp{
									Left: &lyx.AExprOp{
										Left: &lyx.AExprOp{
											Left: &lyx.ColumnRef{
												ColName: "id",
											},
											Right: &lyx.AExprConst{
												Value: "83912839012903",
											},
											Op: "=",
										},
										Right: &lyx.AExprOp{
											Left: &lyx.ColumnRef{
												ColName: "utype",
											},
											Right: &lyx.AExprConst{
												Value: "2",
											},
											Op: "=",
										},
										Op: "AND",
									},
									Right: &lyx.AExprOp{
										Left: &lyx.ColumnRef{
											ColName: "btype",
										},
										Right: &lyx.AExprConst{
											Value: "sample",
										},
										Op: "=",
									},
									Op: "AND",
								},
								Right: &lyx.AExprOp{
									Left: &lyx.ColumnRef{
										ColName: "state",
									},
									Right: &lyx.AExprConst{
										Value: "0",
									},
									Op: "=",
								},
								Op: "AND",
							},
							Right: &lyx.AExprOp{
								Left: &lyx.ColumnRef{
									ColName: "is_something",
								},
								Right: &lyx.AExprConst{
									Value: "true",
								},
								Op: "=",
							},
							Op: "AND",
						},
						Right: &lyx.AExprOp{
							Left: &lyx.ColumnRef{
								ColName: "keys",
							},
							Right: &lyx.AExprOp{
								Left: &lyx.AExprConst{
									Value: "{reshke,denchick}",
								},
								Right: &lyx.AExprOp{
									Left: &lyx.ColumnRef{
										ColName: "keys",
									},
									Right: &lyx.AExprConst{
										Value: "{munakoiso,werelaxe,x4mmm}",
									},
									Op: "@>",
								},
								Op: "OR",
							},
							Op: "@>",
						},
						Op: "AND",
					},
					Right: &lyx.AExprOp{
						Left: &lyx.ColumnRef{
							ColName: "c_id",
						},
						Right: &lyx.AExprConst{
							Value: "trunk",
						},
						Op: "=",
					},
					Op: "AND",
				},
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp)
	}
}

func TestSelectDistinct(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "select count(distinct id) from a.b where other_id ~ 1337;",
			exp: &lyx.Select{
				TargetList: []lyx.Node{nil},
				FromClause: []lyx.FromClauseNode{
					&lyx.RangeVar{
						SchemaName:   "a",
						RelationName: "b",
					},
				},
				Where: &lyx.AExprOp{
					Left: &lyx.ColumnRef{
						ColName: "other_id",
					},
					Right: &lyx.AExprConst{
						Value: "1337",
					},
					Op: "~",
				},
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, tt.query)

		assert.Equal(tt.exp, tmp, tt.query)
	}
}

func TestSelect(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "select 42",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprConst{Value: "42"}},
				Where:      &lyx.AExprEmpty{},
			},
			err: nil,
		},
		{
			query: "select * from xx where i = 1 ",
			exp: &lyx.Select{

				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{&lyx.RangeVar{
					RelationName: "xx",
				},
				},
				Where: &lyx.AExprOp{
					Left: &lyx.ColumnRef{
						ColName:    "i",
						TableAlias: "",
					},
					Right: &lyx.AExprConst{
						Value: "1",
					},
					Op: "=",
				},
			},
			err: nil,
		},

		{
			query: "select * from xx where i = 1 order by i ",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{&lyx.RangeVar{
					RelationName: "xx",
				},
				},
				Where: &lyx.AExprOp{
					Left: &lyx.ColumnRef{
						ColName:    "i",
						TableAlias: "",
					},
					Right: &lyx.AExprConst{
						Value: "1",
					},
					Op: "=",
				},
			},
			err: nil,
		},

		{
			query: "select * from xx where i = 1 limit 7 ",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{&lyx.RangeVar{
					RelationName: "xx",
				},
				},
				Where: &lyx.AExprOp{
					Left: &lyx.ColumnRef{
						ColName:    "i",
						TableAlias: "",
					},
					Right: &lyx.AExprConst{
						Value: "1",
					},
					Op: "=",
				},
			},
			err: nil,
		},

		{
			query: "select * from xx where i = 1 group by i ",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{&lyx.RangeVar{
					RelationName: "xx",
				},
				},
				Where: &lyx.AExprOp{
					Left: &lyx.ColumnRef{
						ColName:    "i",
						TableAlias: "",
					},
					Right: &lyx.AExprConst{
						Value: "1",
					},
					Op: "=",
				},
			},
			err: nil,
		},

		{
			query: "select * from xx where i = 1 group by i having sum(i)",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{&lyx.RangeVar{
					RelationName: "xx",
				},
				},
				Where: &lyx.AExprOp{
					Left: &lyx.ColumnRef{
						ColName:    "i",
						TableAlias: "",
					},
					Right: &lyx.AExprConst{
						Value: "1",
					},
					Op: "=",
				},
			},
			err: nil,
		},

		{
			query: "select * from xx where i = 1 order by i limit 7 ",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{
					&lyx.RangeVar{
						RelationName: "xx",
					},
				},
				Where: &lyx.AExprOp{
					Left: &lyx.ColumnRef{
						ColName:    "i",
						TableAlias: "",
					},
					Right: &lyx.AExprConst{
						Value: "1",
					},
					Op: "=",
				},
			},
			err: nil,
		},
		{
			query: "select * from xx where i = 1 AND j = 2 ",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{&lyx.RangeVar{
					RelationName: "xx",
				},
				},
				Where: &lyx.AExprOp{
					Op: "AND",
					Left: &lyx.AExprOp{
						Left: &lyx.ColumnRef{
							ColName:    "i",
							TableAlias: "",
						},
						Right: &lyx.AExprConst{
							Value: "1",
						},
						Op: "=",
					},
					Right: &lyx.AExprOp{
						Left: &lyx.ColumnRef{
							ColName:    "j",
							TableAlias: "",
						},
						Right: &lyx.AExprConst{
							Value: "2",
						},
						Op: "=",
					},
				},
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, tt.query)

		assert.Equal(tt.exp, tmp, tt.query)
	}
}

func TestSelectMultipleRelations(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "select * from xx, xx2 ",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{&lyx.RangeVar{
					RelationName: "xx",
				},
					&lyx.RangeVar{
						RelationName: "xx2",
					},
				},
				Where: &lyx.AExprEmpty{},
			},
			err: nil,
		},

		{
			query: "select * from xx, xx2 b where b.i = 1",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{&lyx.RangeVar{
					RelationName: "xx",
				},
					&lyx.RangeVar{
						RelationName: "xx2",
						Alias:        "b",
					},
				},
				Where: &lyx.AExprOp{
					Left: &lyx.ColumnRef{
						ColName:    "i",
						TableAlias: "b",
					},
					Right: &lyx.AExprConst{Value: "1"},
					Op:    "=",
				},
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err)

		assert.Equal(tt.exp, tmp)
	}
}

func TestSelectBetween(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "SELECT * FROM xxmixed WHERE id BETWEEN 21 AND 22 ORDER BY id;",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{&lyx.RangeVar{
					RelationName: "xxmixed",
				},
				},
				Where: &lyx.AExprOp{
					Left: &lyx.ColumnRef{
						ColName: "id",
					},
					Right: &lyx.AExprList{
						List: []lyx.Node{
							&lyx.AExprConst{
								Value: "21",
							},
							&lyx.AExprConst{
								Value: "22",
							},
						},
					},
					Op: "BETWEEN",
				},
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp, "query %s", tt.query)
	}
}

func TestSelectAlias(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{

		{
			query: "SELECT kind, sum(len) AS total FROM films GROUP BY kind;",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.ColumnRef{ColName: "kind"}, nil},
				FromClause: []lyx.FromClauseNode{&lyx.RangeVar{
					RelationName: "films",
				},
				},
				Where: &lyx.AExprEmpty{},
			},
			err: nil,
		},

		{
			query: "SELECT kind, sum(len) AS total FROM films AS f GROUP BY kind;",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.ColumnRef{ColName: "kind"}, nil},
				FromClause: []lyx.FromClauseNode{&lyx.RangeVar{
					RelationName: "films",
					Alias:        "f",
				},
				},
				Where: &lyx.AExprEmpty{},
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp)
	}
}

func TestCreate(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "create table xx ( i int )",
			exp: &lyx.CreateTable{
				TableName: "xx",
				TableElts: []lyx.TableElt{
					{
						ColName: "i",
						ColType: "int",
					},
				},
			},
			err: nil,
		},
		/* same test, no spaces between tokens */
		{
			query: "create table xx(i int)",
			exp: &lyx.CreateTable{
				TableName: "xx",
				TableElts: []lyx.TableElt{
					{
						ColName: "i",
						ColType: "int",
					},
				},
			},
			err: nil,
		},

		/* same test, no spaces between tokens */
		{
			query: "create table tt(i int);",
			exp: &lyx.CreateTable{
				TableName: "tt",
				TableElts: []lyx.TableElt{
					{
						ColName: "i",
						ColType: "int",
					},
				},
			},
			err: nil,
		},

		{
			query: "CREATE TABLE sshjt1(i int, j int);",
			exp: &lyx.CreateTable{
				TableName: "sshjt1",
				TableElts: []lyx.TableElt{
					{
						ColName: "i",
						ColType: "int",
					},
					{
						ColName: "j",
						ColType: "int",
					},
				},
			},
			err: nil,
		},

		{
			query: "CREATE TABLE orders(id INT PRIMARY KEY);",
			exp: &lyx.CreateTable{
				TableName: "orders",
				TableElts: []lyx.TableElt{
					{
						ColName: "id",
						ColType: "INT",
					},
				},
			},
			err: nil,
		},

		{
			query: "CREATE TABLE delivery(id INT PRIMARY KEY, order_id INT, FOREIGN KEY(order_id) REFERENCES orders(id));",
			exp: &lyx.CreateTable{
				TableName: "delivery",
				TableElts: []lyx.TableElt{
					{
						ColName: "id",
						ColType: "INT",
					},
					{
						ColName: "order_id",
						ColType: "INT",
					},
				},
			},
			err: nil,
		},

		{
			query: "create database reg",
			exp:   &lyx.CreateDatabase{},
			err:   nil,
		},

		{
			query: "create role reg",
			exp:   &lyx.CreateRole{},
			err:   nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp)
	}
}

func TestInsert(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "insert into xx (id) values(1)",
			exp: &lyx.Insert{
				TableRef: &lyx.RangeVar{
					RelationName: "xx",
					SchemaName:   "",
					Alias:        "",
				},
				Columns: []string{"id"},
				Values:  []lyx.Node{&lyx.AExprConst{Value: "1"}},
			},
			err: nil,
		},

		{
			query: "INSERT INTO xx (w_id) SELECT 20;",
			exp: &lyx.Insert{
				TableRef: &lyx.RangeVar{
					RelationName: "xx",
					SchemaName:   "",
					Alias:        "",
				},
				Columns: []string{"w_id"},
				SubSelect: &lyx.Select{
					TargetList: []lyx.Node{&lyx.AExprConst{Value: "20"}},
					Where:      &lyx.AExprEmpty{},
				},
			},
			err: nil,
		},

		{
			query: "Insert into xx (i) select * from yy where i = 8",
			exp: &lyx.Insert{
				TableRef: &lyx.RangeVar{
					RelationName: "xx",
					SchemaName:   "",
					Alias:        "",
				},
				Columns: []string{"i"},
				Values:  nil,
				SubSelect: &lyx.Select{
					TargetList: []lyx.Node{&lyx.AExprEmpty{}},
					FromClause: []lyx.FromClauseNode{
						&lyx.RangeVar{
							RelationName: "yy",
						},
					},
					Where: &lyx.AExprOp{
						Left: &lyx.ColumnRef{
							ColName: "i",
						},
						Right: &lyx.AExprConst{Value: "8"},
						Op:    "=",
					},
				},
			},
			err: nil,
		},

		{
			query: `INSERT INTO films (code, title, did, date_prod, kind)
			VALUES ('T_601', 'Yojimbo', 106, '1961-06-16', 'Drama');`,
			exp: &lyx.Insert{
				TableRef: &lyx.RangeVar{
					RelationName: "films",
					SchemaName:   "",
					Alias:        "",
				},
				Columns: []string{"code", "title", "did", "date_prod", "kind"},
				Values: []lyx.Node{
					&lyx.AExprConst{Value: "T_601"},
					&lyx.AExprConst{Value: "Yojimbo"},
					&lyx.AExprConst{Value: "106"},
					&lyx.AExprConst{Value: "1961-06-16"},
					&lyx.AExprConst{Value: "Drama"},
				},
			},
			err: nil,
		},

		///

		{
			query: `INSERT INTO films SELECT * FROM tmp_films WHERE date_prod < '2004-05-07';`,
			exp: &lyx.Insert{
				TableRef: &lyx.RangeVar{
					RelationName: "films",
					SchemaName:   "",
					Alias:        "",
				},
				SubSelect: &lyx.Select{
					TargetList: []lyx.Node{&lyx.AExprEmpty{}},
					FromClause: []lyx.FromClauseNode{&lyx.RangeVar{
						RelationName: "tmp_films",
					},
					},
					Where: &lyx.AExprOp{
						Left: &lyx.ColumnRef{
							ColName:    "date_prod",
							TableAlias: "",
						},
						Right: &lyx.AExprConst{
							Value: "2004-05-07",
						},
						Op: "<",
					},
				},
			},
			err: nil,
		},

		{
			query: "insert into xx (id,id2) values(1,2)",
			exp: &lyx.Insert{
				TableRef: &lyx.RangeVar{
					RelationName: "xx",
					SchemaName:   "",
					Alias:        "",
				},
				Columns: []string{
					"id",
					"id2",
				},
				Values: []lyx.Node{
					&lyx.AExprConst{Value: "1"},
					&lyx.AExprConst{Value: "2"},
				},
			},
			err: nil,
		},
		{
			query: "INSERT INTO xxtt1 (j, i, w_id) VALUES(2121221, -211212, 21);",
			exp: &lyx.Insert{
				TableRef: &lyx.RangeVar{
					RelationName: "xxtt1",
					SchemaName:   "",
					Alias:        "",
				},
				Columns: []string{"j", "i", "w_id"},
				Values: []lyx.Node{
					&lyx.AExprConst{Value: "2121221"},
					&lyx.AExprConst{Value: "211212"},
					&lyx.AExprConst{Value: "21"},
				},
			},
			err: nil,
		},

		/* only first tuple from values clause parsed  */
		{
			query: "insert into xx (id,id2) values(1,2), (2,3), ( 4, 5)",
			exp: &lyx.Insert{
				TableRef: &lyx.RangeVar{
					RelationName: "xx",
					SchemaName:   "",
					Alias:        "",
				},
				Columns: []string{
					"id",
					"id2",
				},
				Values: []lyx.Node{
					&lyx.AExprConst{Value: "1"},
					&lyx.AExprConst{Value: "2"},
				},
			},
			err: nil,
		},

		/* insert from select */
		{
			query: "insert into xx select * from xx2 where id2 = 7",
			exp: &lyx.Insert{
				TableRef: &lyx.RangeVar{
					RelationName: "xx",
					SchemaName:   "",
					Alias:        "",
				},
				SubSelect: &lyx.Select{
					TargetList: []lyx.Node{&lyx.AExprEmpty{}},
					FromClause: []lyx.FromClauseNode{&lyx.RangeVar{
						RelationName: "xx2",
					},
					},
					Where: &lyx.AExprOp{
						Left: &lyx.ColumnRef{
							TableAlias: "",
							ColName:    "id2",
						},
						Right: &lyx.AExprConst{
							Value: "7",
						},
						Op: "=",
					},
				},
			},
			err: nil,
		},

		/* Explicit row */
		{
			query: `
			INSERT INTO lol.kek
(
 id,
 info
 )
VALUES(2822, ROW('none', '0'));
`,
			exp: &lyx.Insert{
				TableRef: &lyx.RangeVar{
					RelationName: "kek",
					SchemaName:   "lol",
					Alias:        "",
				},
				Values: []lyx.Node{
					&lyx.AExprConst{
						Value: "2822",
					},
					nil,
				},
				Columns:   []string{"id", "info"},
				SubSelect: nil,
			},
			err: nil,
		},
		{
			query: `
			INSERT INTO lol.kek
(
 id,
 info,
 bytes
 )
VALUES(2822, ROW('none', '0'), ''::bytea);
`,
			exp: &lyx.Insert{
				TableRef: &lyx.RangeVar{
					RelationName: "kek",
					SchemaName:   "lol",
					Alias:        "",
				}, Values: []lyx.Node{
					&lyx.AExprConst{
						Value: "2822",
					},
					nil,
					&lyx.AExprConst{
						Value: "",
					},
				}, Columns: []string{"id", "info", "bytes"},
				SubSelect: nil},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp, "query %s", tt.query)
	}
}

func TestInsertComplex(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: `

INSERT
	INTO "exschema"."extable"
("cat","fox","dog","fly","bee","lol","kek","type","id","version","isok","isnotok","keys","tags","key_id",
	"key_value","value_type","state")
VALUES ('1970-01-01 12:00:00.5',111111,NULL,NULL,'9223372036854775807',
	'2','some','thing','*()*()Q*D()beiwe','0','trunk',true,'{280,fb8,909,e6,ffc}',
	'{9223372036854775806}','31337','bfuiqwefbIUGEIUWhgui..012-2-03849012381==-=-~~~?!@$#@#%%^&*^*(*)../././','0',0)

			`,
			exp: &lyx.Insert{
				TableRef: &lyx.RangeVar{
					RelationName: "extable",
					SchemaName:   "exschema",
					Alias:        "",
				},
				Values: []lyx.Node{
					&lyx.AExprConst{Value: "1970-01-01 12:00:00.5"},
					&lyx.AExprConst{Value: "111111"},
					&lyx.AExprConst{Value: "NULL"},
					&lyx.AExprConst{Value: "NULL"},
					&lyx.AExprConst{Value: "9223372036854775807"},
					&lyx.AExprConst{Value: "2"},
					&lyx.AExprConst{Value: "some"},
					&lyx.AExprConst{Value: "thing"},
					&lyx.AExprConst{Value: "*()*()Q*D()beiwe"},
					&lyx.AExprConst{Value: "0"},
					&lyx.AExprConst{Value: "trunk"},
					&lyx.AExprConst{Value: "true"},
					&lyx.AExprConst{Value: "{280,fb8,909,e6,ffc}"},
					&lyx.AExprConst{Value: "{9223372036854775806}"},
					&lyx.AExprConst{Value: "31337"},
					&lyx.AExprConst{Value: "bfuiqwefbIUGEIUWhgui..012-2-03849012381==-=-~~~?!@$#@#%%^&*^*(*)../././"},
					&lyx.AExprConst{Value: "0"},
					&lyx.AExprConst{Value: "0"},
				},
				Columns: []string{
					"cat",
					"fox",
					"dog",
					"fly",
					"bee",
					"lol",
					"kek",
					"type",
					"id",
					"version",
					"isok",
					"isnotok",
					"keys",
					"tags",
					"key_id",
					"key_value",
					"value_type",
					"state",
				}, SubSelect: nil},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp)
	}
}

func TestUpdate(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		/* simple update */
		{
			query: "UPDATE films SET kind = 'Dramatic' WHERE kind = 'Drama';",
			exp: &lyx.Update{
				TableRef: &lyx.RangeVar{
					RelationName: "films",
					SchemaName:   "",
					Alias:        "",
				},
				Where: &lyx.AExprOp{
					Left: &lyx.ColumnRef{
						TableAlias: "",
						ColName:    "kind",
					},
					Right: &lyx.AExprConst{Value: "Drama"},
					Op:    "=",
				},
			},
			err: nil,
		},

		/* another simple update */
		{
			query: `UPDATE weather SET temp_lo = temp_lo+1, temp_hi = temp_lo+15, prcp = DEFAULT
			WHERE city = 'San Francisco' AND date = '2003-07-03';`,
			exp: &lyx.Update{
				TableRef: &lyx.RangeVar{
					RelationName: "weather",
					SchemaName:   "",
					Alias:        "",
				},
				Where: &lyx.AExprOp{
					Left: &lyx.AExprOp{
						Left: &lyx.ColumnRef{
							TableAlias: "",
							ColName:    "city",
						},
						Right: &lyx.AExprConst{
							Value: "San Francisco",
						},
						Op: "=",
					},
					Right: &lyx.AExprOp{
						Left: &lyx.ColumnRef{
							TableAlias: "",
							ColName:    "date",
						},
						Right: &lyx.AExprConst{
							Value: "2003-07-03",
						},
						Op: "=",
					},
					Op: "AND",
				},
			},
			err: nil,
		},

		/* another simple update  with returning clause */
		{
			query: `UPDATE weather SET temp_lo = temp_lo+1, temp_hi = temp_lo+15, prcp = DEFAULT
			WHERE city = 'San Francisco' AND date = '2003-07-03'
			RETURNING temp_lo, temp_hi, prcp;`,
			exp: &lyx.Update{
				TableRef: &lyx.RangeVar{
					RelationName: "weather",
					SchemaName:   "",
					Alias:        "",
				},
				Where: &lyx.AExprOp{
					Left: &lyx.AExprOp{
						Left: &lyx.ColumnRef{
							TableAlias: "",
							ColName:    "city",
						},
						Right: &lyx.AExprConst{
							Value: "San Francisco",
						},
						Op: "=",
					},
					Right: &lyx.AExprOp{
						Left: &lyx.ColumnRef{
							TableAlias: "",
							ColName:    "date",
						},
						Right: &lyx.AExprConst{
							Value: "2003-07-03",
						},
						Op: "=",
					},
					Op: "AND",
				},
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp)
	}
}

func TestDelete(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		/* delete without where */
		{
			query: "DELETE FROM films;",
			exp: &lyx.Delete{
				TableRef: &lyx.RangeVar{
					RelationName: "films",
					SchemaName:   "",
					Alias:        "",
				},
				Where: &lyx.AExprEmpty{},
			},
			err: nil,
		},

		/* simple delete */
		{
			query: "DELETE FROM films WHERE kind <> 'Musical';",
			exp: &lyx.Delete{
				TableRef: &lyx.RangeVar{
					RelationName: "films",
					SchemaName:   "",
					Alias:        "",
				},
				Where: &lyx.AExprOp{
					Left: &lyx.ColumnRef{
						TableAlias: "",
						ColName:    "kind",
					},
					Right: &lyx.AExprConst{
						Value: "Musical",
					},
					Op: "<>",
				},
			},
			err: nil,
		},

		/* simple delete with returning */
		{
			query: "DELETE FROM tasks WHERE status = 'DONE' RETURNING *;",
			exp: &lyx.Delete{
				TableRef: &lyx.RangeVar{
					RelationName: "tasks",
					SchemaName:   "",
					Alias:        "",
				},
				Where: &lyx.AExprOp{
					Left: &lyx.ColumnRef{
						TableAlias: "",
						ColName:    "status",
					},
					Right: &lyx.AExprConst{
						Value: "DONE",
					},
					Op: "=",
				},
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp)
	}
}

func TestCopy(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		/* copy with where */
		{
			query: "COPY (SELECT * FROM country WHERE country_name = 'R') TO  STDOUT",
			exp: &lyx.Copy{
				IsFrom: false,
				SubStmt: &lyx.Select{
					TargetList: []lyx.Node{
						&lyx.AExprEmpty{},
					},
					FromClause: []lyx.FromClauseNode{&lyx.RangeVar{
						RelationName: "country",
					}},
					Where: &lyx.AExprOp{
						Left: &lyx.ColumnRef{
							ColName: "country_name",
						},
						Right: &lyx.AExprConst{
							Value: "R",
						},
						Op: "=",
					},
				},
			},
			err: nil,
		},
		/* copy with where */
		{
			query: "COPY copy_test FROM STDIN WHERE id <= 30;",
			exp: &lyx.Copy{
				TableRef: &lyx.RangeVar{
					RelationName: "copy_test",
				},
				Where: &lyx.AExprOp{
					Left: &lyx.ColumnRef{
						ColName: "id",
					},
					Right: &lyx.AExprConst{
						Value: "30",
					},
					Op: "<=",
				},
				IsFrom: true,
			},
			err: nil,
		},
		{
			query: `copy pgbench_accounts from stdin with (freeze on)`,
			exp: &lyx.Copy{
				TableRef: &lyx.RangeVar{
					RelationName: "pgbench_accounts",
				},
				Where:  &lyx.AExprEmpty{},
				IsFrom: true,
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp)
	}
}

func TestFuncApplication(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}
	for _, tt := range []tcase{
		{
			query: "INSERT INTO xxtt1 (j, w_id) SELECT a, 20 from unnest(ARRAY[10]) a;",
			exp: &lyx.Insert{
				TableRef: &lyx.RangeVar{RelationName: "xxtt1"},
				Columns:  []string{"j", "w_id"},
				SubSelect: &lyx.Select{
					TargetList: []lyx.Node{
						&lyx.ColumnRef{
							ColName: "a",
						},
						&lyx.AExprConst{
							Value: "20",
						},
					},
					FromClause: []lyx.FromClauseNode{nil},
					Where:      &lyx.AExprEmpty{},
				},
			},
			err: nil,
		},

		{
			query: ` UPDATE xxtt1 
				set i=a.i, j=a.j 
			from unnest(ARRAY[(1,10)]) as a(i int, j int) 
			where w_id=20 and xxtt1.j=a.j;`,
			exp: &lyx.Update{
				TableRef: &lyx.RangeVar{RelationName: "xxtt1"},
				Where: &lyx.AExprOp{
					Left: &lyx.AExprOp{
						Left: &lyx.ColumnRef{
							ColName: "w_id",
						},
						Right: &lyx.AExprConst{
							Value: "20",
						},
						Op: "=",
					},
					Right: &lyx.AExprOp{
						Left: &lyx.ColumnRef{
							TableAlias: "xxtt1",
							ColName:    "j",
						},
						Right: &lyx.ColumnRef{
							TableAlias: "a",
							ColName:    "j",
						},
						Op: "=",
					},
					Op: "and",
				},
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp, "query %s", tt.query)
	}
}

func TestSelectTargetLists(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "SELECT pg_is_in_recovery(), id FROM tsa_test WHERE id = 22;",
			exp: &lyx.Select{

				TargetList: []lyx.Node{nil, &lyx.ColumnRef{ColName: "id"}},
				FromClause: []lyx.FromClauseNode{
					&lyx.RangeVar{
						RelationName: "tsa_test",
					},
				},
				Where: &lyx.AExprOp{
					Left: &lyx.ColumnRef{
						ColName: "id",
					},
					Right: &lyx.AExprConst{
						Value: "22",
					},
					Op: "=",
				},
			},
			err: nil,
		},
		{
			query: "select count(*) from pgbench_branches;",
			exp: &lyx.Select{
				TargetList: []lyx.Node{nil},
				FromClause: []lyx.FromClauseNode{
					&lyx.RangeVar{
						RelationName: "pgbench_branches",
					},
				},
				Where: &lyx.AExprEmpty{},
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp, "query %s", tt.query)
	}
}

func TestWithComments(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "SELECT * /* o lo lo l ol **** *!*&*(&!89**/ FROM  tsa_test WHERE id = 22;",
			exp: &lyx.Select{

				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{
					&lyx.RangeVar{
						RelationName: "tsa_test",
					},
				},
				Where: &lyx.AExprOp{
					Left: &lyx.ColumnRef{
						ColName: "id",
					},
					Right: &lyx.AExprConst{
						Value: "22",
					},
					Op: "=",
				},
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp, "query %s", tt.query)
	}
}

func TestJoins(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "SELECT * FROM delivery JOIN orders ON order_id = id;",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{
					&lyx.JoinExpr{
						Larg: &lyx.RangeVar{
							RelationName: "delivery",
						},
						Rarg: &lyx.RangeVar{
							RelationName: "orders",
						},
					},
				},
				Where: &lyx.AExprEmpty{},
			},
			err: nil,
		},

		{
			query: "SELECT * FROM delivery JOIN orders ON order_id = id;",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{
					&lyx.JoinExpr{
						Larg: &lyx.RangeVar{
							RelationName: "delivery",
						},
						Rarg: &lyx.RangeVar{
							RelationName: "orders",
						},
					},
				},
				Where: &lyx.AExprEmpty{},
			},
			err: nil,
		},

		{
			query: "SELECT * FROM sshjt1 a join sshjt1 b ON TRUE WHERE a.i = 12 AND b.j = a.j;",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{
					&lyx.JoinExpr{
						Larg: &lyx.RangeVar{
							RelationName: "sshjt1",
							Alias:        "a",
						},
						Rarg: &lyx.RangeVar{
							RelationName: "sshjt1",
							Alias:        "b",
						},
					},
				},
				Where: &lyx.AExprOp{
					Left: &lyx.AExprOp{
						Left: &lyx.ColumnRef{
							ColName:    "i",
							TableAlias: "a",
						},
						Right: &lyx.AExprConst{
							Value: "12",
						},
						Op: "=",
					},
					Right: &lyx.AExprOp{
						Left: &lyx.ColumnRef{
							ColName:    "j",
							TableAlias: "b",
						},
						Right: &lyx.ColumnRef{
							ColName:    "j",
							TableAlias: "a",
						},
						Op: "=",
					},
					Op: "AND",
				},
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp, "query %s", tt.query)
	}
}

func TestMisc(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "vacuum xx;",
			exp:   &lyx.Vacuum{},
			err:   nil,
		},
		{
			query: "drop table xx;",
			exp:   &lyx.Drop{},
			err:   nil,
		},
		{
			query: "analyze xx;",
			exp:   &lyx.Analyze{},
			err:   nil,
		},
		{
			query: "alter table xx add column i int;",
			exp:   &lyx.Alter{},
			err:   nil,
		},
		{
			query: "cluster xx;",
			exp:   &lyx.Cluster{},
			err:   nil,
		},
		{
			query: ";",
			exp:   nil,
			err:   nil,
		},
		{
			query: "-- comment",
			exp:   nil,
			err:   nil,
		},
		{
			query: "/* comment */",
			exp:   nil,
			err:   nil,
		},

		{
			query: "select 11 /* comment */",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprConst{Value: "11"}},
				Where:      &lyx.AExprEmpty{},
			},
			err: nil,
		},

		{
			query: "SELECT  * FROM delivery  JOIN  /* comment */ orders ON order_id = id;",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{
					&lyx.JoinExpr{
						Larg: &lyx.RangeVar{
							RelationName: "delivery",
						},
						Rarg: &lyx.RangeVar{
							RelationName: "orders",
						},
					},
				},
				Where: &lyx.AExprEmpty{},
			},
			err: nil,
		},

		{
			query: "SELECT  * FROM delivery  JOIN  /* comment */ orders ON order_id = id;-- ololol",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{
					&lyx.JoinExpr{
						Larg: &lyx.RangeVar{
							RelationName: "delivery",
						},
						Rarg: &lyx.RangeVar{
							RelationName: "orders",
						},
					},
				},
				Where: &lyx.AExprEmpty{},
			},
			err: nil,
		},
		{
			query: "SELECT /* comment */  * FROM delivery  JOIN  orders ON order_id = id;",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{
					&lyx.JoinExpr{
						Larg: &lyx.RangeVar{
							RelationName: "delivery",
						},
						Rarg: &lyx.RangeVar{
							RelationName: "orders",
						},
					},
				},
				Where: &lyx.AExprEmpty{},
			},
			err: nil,
		},
		{
			query: "begin;",
			exp: &lyx.TransactionStmt{
				Kind: lyx.TRANS_STMT_BEGIN,
			},
			err: nil,
		},
		{
			query: "begin transaction read only;",
			exp: &lyx.TransactionStmt{
				Kind: lyx.TRANS_STMT_BEGIN,
				Options: []lyx.TransactionModeItem{
					lyx.TransactionReadOnly,
				},
			},
			err: nil,
		},
		{
			query: "rollback;",
			exp: &lyx.TransactionStmt{
				Kind: lyx.TRANS_STMT_ROLLBACK,
			},
			err: nil,
		},
		{
			query: "commit;",
			exp: &lyx.TransactionStmt{
				Kind: lyx.TRANS_STMT_COMMIT,
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp)
	}
}

func TestErrors(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "SELECT * FROM sshjt1 a join sshjt1 b WHERE a.i = 12 ON TRUE;",
			exp:   nil,
			err:   nil,
		},
		{
			query: "SELECT * FROM x WHERE ixxxd = 1 iuwehiuhweui;",
			exp:   nil,
			err:   nil,
		},
	} {
		_, err := lyx.Parse(tt.query)

		assert.Error(err, "query %s", tt.query)
	}
}

func TestOperators(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "SELECT 1 ~~~~~~~~ 'dewoijoi'",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprOp{
					Left:  &lyx.AExprConst{Value: "1"},
					Right: &lyx.AExprConst{Value: "dewoijoi"},
					Op:    "~~~~~~~~",
				}},
				Where: &lyx.AExprEmpty{},
			},
			err: nil,
		},
		{
			query: "SELECT 1 ~~~^%^%^^~~~~~ 'dewoijoi'",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprOp{
					Left:  &lyx.AExprConst{Value: "1"},
					Right: &lyx.AExprConst{Value: "dewoijoi"},
					Op:    "~~~^%^%^^~~~~~",
				}},
				Where: &lyx.AExprEmpty{},
			},
			err: nil,
		},
		{
			query: "SELECT * FROM xxtt1 a WHERE a.w_id = 21 and j + i != 0;'",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{
					&lyx.RangeVar{
						RelationName: "xxtt1",
						Alias:        "a",
					},
				},
				Where: &lyx.AExprOp{
					Left: &lyx.AExprOp{
						Left: &lyx.ColumnRef{
							ColName:    "w_id",
							TableAlias: "a",
						},
						Right: &lyx.AExprConst{
							Value: "21",
						},
						Op: "=",
					},
					Right: &lyx.AExprOp{
						Left: &lyx.AExprOp{
							Left: &lyx.ColumnRef{
								ColName: "j",
							},
							Right: &lyx.ColumnRef{
								ColName: "i",
							},
							Op: "+",
						},
						Right: &lyx.AExprConst{
							Value: "0",
						},
						Op: "!=",
					},
					Op: "and",
				},
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp)
	}
}

func TestParams(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "PREPARE P(TEXT) AS SELECT $1",
			exp: &lyx.PrepareStmt{
				Name: "P",
				Statement: &lyx.Select{
					TargetList: []lyx.Node{
						&lyx.ParamRef{
							Number: 1,
						},
					},
					Where: &lyx.AExprEmpty{},
				},
			},
			err: nil,
		},
		{
			query: "SELECT * FROM x WHERE i = $1",
			exp: &lyx.Select{
				TargetList: []lyx.Node{&lyx.AExprEmpty{}},
				FromClause: []lyx.FromClauseNode{
					&lyx.RangeVar{RelationName: "x"},
				},
				Where: &lyx.AExprOp{
					Op:    "=",
					Left:  &lyx.ColumnRef{ColName: "i"},
					Right: &lyx.ParamRef{Number: 1},
				},
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp)
	}
}

func TestCreateTableWithPrimaryKey(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
	}

	for _, tt := range []tcase{
		{
			query: `
			create table IF NOT EXISTS item (i_id int not null, i_im_id int, i_name varchar(24), i_price decimal(5,2), i_data varchar(50),PRIMARY KEY(i_id) )`,
			exp: &lyx.Drop{},
		},
	} {
		_, err := lyx.Parse(tt.query)

		assert.NoError(err, tt.query)
	}
}

func TestCreateTableWithCompositePrimaryKey(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
	}

	for _, tt := range []tcase{
		{
			query: `
			create table IF NOT EXISTS customer (c_id int not null, c_d_id int not null, c_w_id int not null, c_first varchar(16), c_middle char(2), c_last varchar(16), c_street_1 varchar(20), c_street_2 varchar(20), c_city varchar(20), c_state char(2), c_zip char(9), c_phone char(16), c_credit char(2), c_credit_lim bigint, c_discount decimal(4,2), c_balance decimal(12,2), c_ytd_payment decimal(12,2), c_payment_cnt int, c_delivery_cnt int, c_data text, PRIMARY KEY(c_w_id, c_d_id, c_id) )`,
			exp: &lyx.Drop{},
		},
	} {
		_, err := lyx.Parse(tt.query)

		assert.NoError(err, tt.query)
	}
}

func TestDrop(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
	}

	for _, tt := range []tcase{
		{
			query: `
			drop table if exists pgbench_accounts, pgbench_branches, pgbench_history, pgbench_tellers`,
			exp: &lyx.Drop{},
		},
	} {
		_, err := lyx.Parse(tt.query)

		assert.NoError(err, tt.query)
	}
}

func TestCreateFail(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
	}

	for _, tt := range []tcase{
		{
			query: "create table SELECT ( i int )",
		},
		{
			query: "create table ASYMMETRIC ( i int )",
		},
		{
			query: "create table INTEGER ( i int )",
		},
		{
			query: "create table xx ( i SELECT )",
		},
		{
			query: "create table xx ( i WHERE )",
		},
		{
			query: "create table xx ( i AS )",
		}} {
		_, err := lyx.Parse(tt.query)
		assert.Error(err, tt.query)
	}
}

func TestTruncateTable(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}
	for _, tt := range []tcase{
		{
			query: `
			truncate table pgbench_accounts, pgbench_branches, pgbench_history, pgbench_tellers
			`,

			exp: &lyx.Truncate{},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err)

		assert.Equal(tt.exp, tmp)
	}
}

func TestCreateTableWith(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: `
			create table 
				pgbench_tellers(tid int not null,bid int,tbalance int,filler char(84)) with (fillfactor=100)
			`,
			exp: &lyx.CreateTable{
				TableName: "pgbench_tellers",
				TableElts: []lyx.TableElt{
					{
						ColName: "tid",
						ColType: "int",
					},
					{
						ColName: "bid",
						ColType: "int",
					},
					{
						ColName: "tbalance",
						ColType: "int",
					},
					{
						ColName: "filler",
						ColType: "char",
					},
				},
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err)

		assert.Equal(tt.exp, tmp)
	}
}

func TestCreateSuccess(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "create table xx ( ADMIN int, ATOMIC int, CLASS int, LIKE int )",
			exp: &lyx.CreateTable{
				TableName: "xx",
				TableElts: []lyx.TableElt{
					{
						ColName: "ADMIN",
						ColType: "int",
					},
					{
						ColName: "ATOMIC",
						ColType: "int",
					},
					{
						ColName: "CLASS",
						ColType: "int",
					},
					{
						ColName: "LIKE",
						ColType: "int",
					},
				},
			},
			err: nil,
		},
		{
			query: "create table xx ( i0 BINARY, i1 CURRENT_SCHEMA, i2 IS, i3 JOIN, i4 NOTNULL, i5 TABLESAMPLE)",
			exp: &lyx.CreateTable{
				TableName: "xx",
				TableElts: []lyx.TableElt{
					{
						ColName: "i0",
						ColType: "BINARY",
					},
					{
						ColName: "i1",
						ColType: "CURRENT_SCHEMA",
					},
					{
						ColName: "i2",
						ColType: "IS",
					},
					{
						ColName: "i3",
						ColType: "JOIN",
					},
					{
						ColName: "i4",
						ColType: "NOTNULL",
					},
					{
						ColName: "i5",
						ColType: "TABLESAMPLE",
					},
				},
			},
			err: nil,
		},
		{
			query: "create table JSON_ARRAYAGG ( i int )",
			exp: &lyx.CreateTable{
				TableName: "JSON_ARRAYAGG",
				TableElts: []lyx.TableElt{
					{
						ColName: "i",
						ColType: "int",
					},
				},
			},
			err: nil,
		},
		{
			query: "create table XMLTABLE ( i int )",
			exp: &lyx.CreateTable{
				TableName: "XMLTABLE",
				TableElts: []lyx.TableElt{
					{
						ColName: "i",
						ColType: "int",
					},
				},
			},
			err: nil,
		}} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, tt.query)

		assert.Equal(tt.exp, tmp, tt.query)
	}
}

func TestMiscQ(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: `
SELECT n.nspname as "Schema",
  c.relname as "Name",
  CASE c.relkind WHEN 'r' THEN 'table' WHEN 'v' THEN 'view' WHEN 'm' THEN 'materialized view' WHEN 'i' THEN 'index' WHEN 'S' THEN 'sequence' WHEN 't' THEN 'TOAST table' WHEN 'f' THEN 'foreign table' WHEN 'p' THEN 'partitioned table' WHEN 'I' THEN 'partitioned index' END as "Type",
  pg_catalog.pg_get_userbyid(c.relowner) as "Owner"
FROM pg_catalog.pg_class c
     LEFT JOIN pg_catalog.pg_namespace n ON n.oid = c.relnamespace
     LEFT JOIN pg_catalog.pg_am am ON am.oid = c.relam
WHERE c.relkind IN ('r','p','v','m','S','f','')
      AND n.nspname <> 'pg_catalog'
      AND n.nspname !~ '^pg_toast'
      AND n.nspname <> 'information_schema'
  AND pg_catalog.pg_table_is_visible(c.oid)
ORDER BY 1,2;
`,
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, tt.query)

		assert.Equal(tt.exp, tmp, tt.query)
	}
}

func TestExplain(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Node
		err   error
	}

	for _, tt := range []tcase{
		{
			query: `
			EXPLAIN select 1
`,
			exp: &lyx.Explain{
				Stmt: &lyx.Select{
					Where: &lyx.AExprEmpty{},
					TargetList: []lyx.Node{
						&lyx.AExprConst{
							Value: "1",
						},
					},
				},
			},
			err: nil,
		},
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, tt.query)

		assert.Equal(tt.exp, tmp, tt.query)
	}
}
