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

		assert.NoError(err)

		assert.Equal(tt.exp, tmp)
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
				TableRef: &lyx.RangeVar{
					RelationName: "xx",
				},
				Where: &lyx.AExprOp{
					Left: &lyx.ColumnRef{
						ColName: "id",
					},
					Right: &lyx.AExprConst{
						Value: "1",
					},
					Op: "=",
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
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp)
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
