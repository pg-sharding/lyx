package lyx_test

import (
	"lyx/lyx"
	"testing"

	"github.com/stretchr/testify/assert"
)

// import (
// 	"lyx/lyx"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// /*

//  */

// func TestSelectComplex(t *testing.T) {
// 	assert := assert.New(t)

// 	type tcase struct {
// 		query string
// 		exp   lyx.Statement
// 		err   error
// 	}

// 	for _, tt := range []tcase{
// 		{
// 			query: `
// 			SELECT * FROM "exschema"."extable" WHERE id='83912839012903' AND utype='2' AND btype='sample' AND state = 0 AND is_something = true AND (keys @> '{reshke,denchick}' OR keys @> '{munakoiso,werelaxe,x4mmm}') AND c_id = 'trunk' ORDER BY entity_id asc;
// 			`,
// 			exp: &lyx.Select{
// 				FromClause: []FromClauseNode{&RangeVar{
// 					RelationName: "extable",
// 					SchemaName:   "exschema",
// 				},
// 				},
// 				Where: &AExprOp{
// 					Op: "AND",
// 					Left: &AExprOp{
// 						ColRef: ColumnRef{
// 							TableAlias: "",
// 							ColName:    "id",
// 						},
// 						Value: "83912839012903",
// 					},
// 					Right: &WhereClauseOp{
// 						Op: "AND",
// 						Left: &WhereClauseLeaf{
// 							ColRef: ColumnRef{
// 								TableAlias: "",
// 								ColName:    "utype",
// 							},
// 							Value: "2",
// 						},
// 						Right: &WhereClauseOp{
// 							Op: "AND",
// 							Left: &WhereClauseLeaf{
// 								ColRef: ColumnRef{
// 									TableAlias: "",
// 									ColName:    "btype",
// 								},
// 								Value: "sample",
// 							},
// 							Right: &WhereClauseOp{
// 								Op: "AND",
// 								Left: &WhereClauseLeaf{
// 									ColRef: ColumnRef{
// 										TableAlias: "",
// 										ColName:    "state",
// 									},
// 									Value: "0",
// 								},
// 								Right: &WhereClauseOp{
// 									Op: "AND",
// 									Left: &WhereClauseLeaf{
// 										ColRef: ColumnRef{
// 											TableAlias: "",
// 											ColName:    "is_something",
// 										},
// 										Value: "true",
// 									},
// 									Right: &WhereClauseOp{
// 										Op: "AND",
// 										Left: &WhereClauseOp{
// 											Left: &WhereClauseLeaf{
// 												ColRef: ColumnRef{
// 													ColName: "keys",
// 												},
// 												Value: "{reshke,denchick}",
// 											},
// 											Right: &WhereClauseLeaf{
// 												ColRef: ColumnRef{
// 													ColName: "keys",
// 												},
// 												Value: "{munakoiso,werelaxe,x4mmm}",
// 											},
// 											Op: "OR",
// 										},
// 										Right: &WhereClauseLeaf{
// 											ColRef: ColumnRef{
// 												TableAlias: "",
// 												ColName:    "c_id",
// 											},
// 											Value: "trunk",
// 										},
// 									},
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},

// 			err: nil,
// 		},

// 		{
// 			query: `
// 	SELECT *
// 	FROM "exschema"."extable"
// 	WHERE
// 		id='83912839012903'
// 		AND utype='2'
// 		AND btype='sample'
// 		AND state = 0
// 		AND is_something = true
// 		AND (keys @> '{reshke,denchick}' OR keys @> '{munakoiso,werelaxe,x4mmm}')
// 		AND c_id = 'trunk'
// 	ORDER BY entity_id asc;
// 			`,
// 			exp: &Select{
// 				FromClause: []FromClauseNode{&RangeVar{
// 					RelationName: "extable",
// 					SchemaName:   "exschema",
// 				},
// 				},
// 				Where: &WhereClauseOp{
// 					Op: "AND",
// 					Left: &WhereClauseLeaf{
// 						ColRef: ColumnRef{
// 							TableAlias: "",
// 							ColName:    "id",
// 						},
// 						Value: "83912839012903",
// 					},
// 					Right: &WhereClauseOp{
// 						Op: "AND",
// 						Left: &WhereClauseLeaf{
// 							ColRef: ColumnRef{
// 								TableAlias: "",
// 								ColName:    "utype",
// 							},
// 							Value: "2",
// 						},
// 						Right: &WhereClauseOp{
// 							Op: "AND",
// 							Left: &WhereClauseLeaf{
// 								ColRef: ColumnRef{
// 									TableAlias: "",
// 									ColName:    "btype",
// 								},
// 								Value: "sample",
// 							},
// 							Right: &WhereClauseOp{
// 								Op: "AND",
// 								Left: &WhereClauseLeaf{
// 									ColRef: ColumnRef{
// 										TableAlias: "",
// 										ColName:    "state",
// 									},
// 									Value: "0",
// 								},
// 								Right: &WhereClauseOp{
// 									Op: "AND",
// 									Left: &WhereClauseLeaf{
// 										ColRef: ColumnRef{
// 											TableAlias: "",
// 											ColName:    "is_something",
// 										},
// 										Value: "true",
// 									},
// 									Right: &WhereClauseOp{
// 										Op: "AND",
// 										Left: &WhereClauseOp{
// 											Left: &WhereClauseLeaf{
// 												ColRef: ColumnRef{
// 													ColName: "keys",
// 												},
// 												Value: "{reshke,denchick}",
// 											},
// 											Right: &WhereClauseLeaf{
// 												ColRef: ColumnRef{
// 													ColName: "keys",
// 												},
// 												Value: "{munakoiso,werelaxe,x4mmm}",
// 											},
// 											Op: "OR",
// 										},
// 										Right: &WhereClauseLeaf{
// 											ColRef: ColumnRef{
// 												TableAlias: "",
// 												ColName:    "c_id",
// 											},
// 											Value: "trunk",
// 										},
// 									},
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},

// 			err: nil,
// 		},
// 	} {
// 		tmp, err := Parse(tt.query)

// 		assert.NoError(err, "query %s", tt.query)

// 		assert.Equal(tt.exp, tmp)
// 	}
// }

// func TestSelect(t *testing.T) {
// 	assert := assert.New(t)

// 	type tcase struct {
// 		query string
// 		exp   Statement
// 		err   error
// 	}

// 	for _, tt := range []tcase{
// 		{
// 			query: "select 42",
// 			exp: &Select{
// 				Where: &WhereClauseEmpty{},
// 			},
// 			err: nil,
// 		},
// 		{
// 			query: "select * from xx where i = 1 ",
// 			exp: &Select{
// 				FromClause: []FromClauseNode{&RangeVar{
// 					RelationName: "xx",
// 				},
// 				},
// 				Where: &WhereClauseLeaf{
// 					ColRef: ColumnRef{
// 						ColName:    "i",
// 						TableAlias: "",
// 					},
// 					Value: "1",
// 				},
// 			},
// 			err: nil,
// 		},

// 		{
// 			query: "select * from xx where i = 1 order by i ",
// 			exp: &Select{
// 				FromClause: []FromClauseNode{&RangeVar{
// 					RelationName: "xx",
// 				},
// 				},
// 				Where: &WhereClauseLeaf{
// 					ColRef: ColumnRef{
// 						ColName:    "i",
// 						TableAlias: "",
// 					},
// 					Value: "1",
// 				},
// 			},
// 			err: nil,
// 		},

// 		{
// 			query: "select * from xx where i = 1 limit 7 ",
// 			exp: &Select{
// 				FromClause: []FromClauseNode{&RangeVar{
// 					RelationName: "xx",
// 				},
// 				},
// 				Where: &WhereClauseLeaf{
// 					ColRef: ColumnRef{
// 						ColName:    "i",
// 						TableAlias: "",
// 					},
// 					Value: "1",
// 				},
// 			},
// 			err: nil,
// 		},

// 		{
// 			query: "select * from xx where i = 1 group by i ",
// 			exp: &Select{
// 				FromClause: []FromClauseNode{&RangeVar{
// 					RelationName: "xx",
// 				},
// 				},
// 				Where: &WhereClauseLeaf{
// 					ColRef: ColumnRef{
// 						ColName:    "i",
// 						TableAlias: "",
// 					},
// 					Value: "1",
// 				},
// 			},
// 			err: nil,
// 		},

// 		{
// 			query: "select * from xx where i = 1 group by i having sum(i)",
// 			exp: &Select{
// 				FromClause: []FromClauseNode{&RangeVar{
// 					RelationName: "xx",
// 				},
// 				},
// 				Where: &WhereClauseLeaf{
// 					ColRef: ColumnRef{
// 						ColName:    "i",
// 						TableAlias: "",
// 					},
// 					Value: "1",
// 				},
// 			},
// 			err: nil,
// 		},

// 		{
// 			query: "select * from xx where i = 1 order by i limit 7 ",
// 			exp: &Select{
// 				FromClause: []FromClauseNode{&RangeVar{
// 					RelationName: "xx",
// 				},
// 				},
// 				Where: &WhereClauseLeaf{
// 					ColRef: ColumnRef{
// 						ColName:    "i",
// 						TableAlias: "",
// 					},
// 					Value: "1",
// 				},
// 			},
// 			err: nil,
// 		},
// 		{
// 			query: "select * from xx where i = 1 AND j = 2 ",
// 			exp: &Select{
// 				FromClause: []FromClauseNode{&RangeVar{
// 					RelationName: "xx",
// 				},
// 				},
// 				Where: &WhereClauseOp{
// 					Op: "AND",
// 					Left: &WhereClauseLeaf{
// 						ColRef: ColumnRef{
// 							ColName:    "i",
// 							TableAlias: "",
// 						},
// 						Value: "1",
// 					},
// 					Right: &WhereClauseLeaf{
// 						ColRef: ColumnRef{
// 							ColName:    "j",
// 							TableAlias: "",
// 						},
// 						Value: "2",
// 					},
// 				},
// 			},
// 			err: nil,
// 		},
// 	} {
// 		tmp, err := Parse(tt.query)

// 		assert.NoError(err)

// 		assert.Equal(tt.exp, tmp)
// 	}
// }

func TestSelectMultipleRelations(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   lyx.Statement
		err   error
	}

	for _, tt := range []tcase{
		// {
		// 	query: "select * from xx, xx2 ",
		// 	exp: &lyx.Select{
		// 		FromClause: []lyx.FromClauseNode{&lyx.RangeVar{
		// 			RelationName: "xx",
		// 		},
		// 			&lyx.RangeVar{
		// 				RelationName: "xx2",
		// 			},
		// 		},
		// 		Where: &lyx.AExprEmpty{},
		// 	},
		// 	err: nil,
		// },

		{
			query: "select * from xx, xx2 b where b.i = 1",
			exp: &lyx.Select{
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
		exp   lyx.Statement
		err   error
	}

	for _, tt := range []tcase{

		{
			query: "SELECT kind, sum(len) AS total FROM films GROUP BY kind;",
			exp: &lyx.Select{
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
		exp   lyx.Statement
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
		exp   lyx.Statement
		err   error
	}

	for _, tt := range []tcase{
		// {
		// 	query: "insert into xx (id) values(1)",
		// 	exp: &lyx.Insert{
		// 		TableRef: &lyx.RangeVar{
		// 			RelationName: "xx",
		// 			SchemaName:   "",
		// 			Alias:        "",
		// 		},
		// 		Columns: []string{"id"},
		// 		Values:  []lyx.Statement{&lyx.AExprConst{Value: "1"}},
		// 	},
		// 	err: nil,
		// },

		// {
		// 	query: "Insert into xx (i) select * from yy where i = 8",
		// 	exp: &lyx.Insert{
		// 		TableRef: &lyx.RangeVar{
		// 			RelationName: "xx",
		// 			SchemaName:   "",
		// 			Alias:        "",
		// 		},
		// 		Columns: nil,
		// 		Values:  nil,
		// 		SubSelect: &lyx.Select{
		// 			FromClause: []lyx.FromClauseNode{
		// 				&lyx.RangeVar{
		// 					RelationName: "yy",
		// 				},
		// 			},
		// 			Where: &lyx.AExprOp{
		// 				Left: &lyx.ColumnRef{
		// 					ColName: "i",
		// 				},
		// 				Right: &lyx.AExprConst{Value: "8"},
		// 				Op:    "=",
		// 			},
		// 		},
		// 	},
		// 	err: nil,
		// },

		{
			query: `INSERT INTO films (code, title, did, date_prod, kind)
			VALUES ('T_601', 'Yojimbo', 106, '1961-06-16', 'Drama');`,
			exp: &lyx.Insert{
				TableRef: &lyx.RangeVar{
					RelationName: "films",
					SchemaName:   "",
					Alias:        "",
				},
				Columns: []string{"kind", "date_prod", "did", "title", "code"},

				Values: []lyx.Statement{&lyx.AExprConst{Value: "Drama"}, &lyx.AExprConst{Value: "1961-06-16"}, &lyx.AExprConst{Value: "106"}, &lyx.AExprConst{Value: "Yojimbo"}, &lyx.AExprConst{Value: "T_601"}},
			},
			err: nil,
		},

		///

		// {
		// 	query: `INSERT INTO films SELECT * FROM tmp_films WHERE date_prod < '2004-05-07';`,
		// 	exp: &lyx.Insert{
		// 		TableRef: &lyx.RangeVar{
		// 			RelationName: "films",
		// 			SchemaName:   "",
		// 			Alias:        "",
		// 		},
		// 		SubSelect: &lyx.Select{
		// 			FromClause: []lyx.FromClauseNode{&lyx.RangeVar{
		// 				RelationName: "tmp_films",
		// 			},
		// 			},
		// 			Where: &lyx.AExprOp{
		// 				Left: &lyx.ColumnRef{
		// 					ColName:    "date_prod",
		// 					TableAlias: "",
		// 				},
		// 				Right: &lyx.AExprConst{
		// 					Value: "2004-05-07",
		// 				},
		// 				Op: "=",
		// 			},
		// 		},
		// 	},
		// 	err: nil,
		// },

		// {
		// 	query: "insert into xx (id,id2) values(1,2)",
		// 	exp: &lyx.Insert{
		// 		TableRef: &lyx.RangeVar{
		// 			RelationName: "xx",
		// 			SchemaName:   "",
		// 			Alias:        "",
		// 		},
		// 		Columns: []string{"id2", "id"},
		// 		Values:  []lyx.Statement{&lyx.AExprConst{Value: "2"}, &lyx.AExprConst{Value: "1"}},
		// 	},
		// 	err: nil,
		// },
		// {
		// 	query: "INSERT INTO xxtt1 (j, i, w_id) VALUES(2121221, -211212, 21);",
		// 	exp: &lyx.Insert{
		// 		TableRef: &lyx.RangeVar{
		// 			RelationName: "xxtt1",
		// 			SchemaName:   "",
		// 			Alias:        "",
		// 		},
		// 		Columns: []string{"w_id", "i", "j"},
		// 		Values:  []lyx.Statement{&lyx.AExprConst{Value: "21"}, &lyx.AExprConst{Value: "211212"}, &lyx.AExprConst{Value: "2121221"}},
		// 	},
		// 	err: nil,
		// },

		// /* only first tuple from values clause parsed  */
		// {
		// 	query: "insert into xx (id,id2) values(1,2), (2,3), ( 4, 5)",
		// 	exp: &lyx.Insert{
		// 		TableRef: &lyx.RangeVar{
		// 			RelationName: "xx",
		// 			SchemaName:   "",
		// 			Alias:        "",
		// 		},
		// 		Columns: []string{"id2", "id"},
		// 		Values:  []lyx.Statement{&lyx.AExprConst{Value: "2"}, &lyx.AExprConst{Value: "1"}},
		// 	},
		// 	err: nil,
		// },

		// /* insert from select */
		// {
		// 	query: "insert into xx select * from xx2 where id2 = 7",
		// 	exp: &lyx.Insert{
		// 		TableRef: &lyx.RangeVar{
		// 			RelationName: "xx",
		// 			SchemaName:   "",
		// 			Alias:        "",
		// 		},
		// 		SubSelect: &lyx.Select{
		// 			FromClause: []lyx.FromClauseNode{&lyx.RangeVar{
		// 				RelationName: "xx2",
		// 			},
		// 			},
		// 			Where: &lyx.AExprOp{
		// 				Left: &lyx.ColumnRef{
		// 					TableAlias: "",
		// 					ColName:    "id2",
		// 				},
		// 				Right: &lyx.AExprConst{
		// 					Value: "7",
		// 				},
		// 				Op: "=",
		// 			},
		// 		},
		// 	},
		// 	err: nil,
		// },
	} {
		tmp, err := lyx.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp)
	}
}

// func TestInsertComplex(t *testing.T) {
// 	assert := assert.New(t)

// 	type tcase struct {
// 		query string
// 		exp   Statement
// 		err   error
// 	}

// 	for _, tt := range []tcase{
// 		{
// 			query: `

// INSERT
// 	INTO "exschema"."extable"
// ("cat","fox","dog","fly","bee","lol","kek","type","id","version","isok","isnotok","keys","tags","key_id",
// 	"key_value","value_type","state")
// VALUES ('1970-01-01 12:00:00.5',111111,NULL,NULL,'9223372036854775807',
// 	'2','some','thing','*()*()Q*D()beiwe','0','trunk',true,'{280,fb8,909,e6,ffc}',
// 	'{9223372036854775806}','31337','bfuiqwefbIUGEIUWhgui..012-2-03849012381==-=-~~~?!@$#@#%%^&*^*(*)../././','0',0)

// 			`,
// 			exp: &Insert{
// 				TableRef: &RangeVar{
// 					RelationName: "extable",
// 					SchemaName:   "exschema",
// 					Alias:        "",
// 				},
// 				Values: []string{
// 					"0",
// 					"0",
// 					"bfuiqwefbIUGEIUWhgui..012-2-03849012381==-=-~~~?!@$#@#%%^&*^*(*)../././",
// 					"31337",
// 					"{9223372036854775806}",
// 					"{280,fb8,909,e6,ffc}",
// 					"true",
// 					"trunk",
// 					"0",
// 					"*()*()Q*D()beiwe",
// 					"thing",
// 					"some",
// 					"2",
// 					"9223372036854775807",
// 					"NULL",
// 					"NULL",
// 					"111111",
// 					"1970-01-01 12:00:00.5",
// 				},
// 				Columns: []string{
// 					"state",
// 					"value_type",
// 					"key_value",
// 					"key_id",
// 					"tags",
// 					"keys",
// 					"isnotok",
// 					"isok",
// 					"version",
// 					"id",
// 					"type",
// 					"kek",
// 					"lol",
// 					"bee",
// 					"fly",
// 					"dog",
// 					"fox", "cat",
// 				}, SubSelect: nil},
// 			err: nil,
// 		},
// 	} {
// 		tmp, err := Parse(tt.query)

// 		assert.NoError(err, "query %s", tt.query)

// 		assert.Equal(tt.exp, tmp)
// 	}
// }

// func TestUpdate(t *testing.T) {
// 	assert := assert.New(t)

// 	type tcase struct {
// 		query string
// 		exp   Statement
// 		err   error
// 	}

// 	for _, tt := range []tcase{
// 		/* simple update */
// 		{
// 			query: "UPDATE films SET kind = 'Dramatic' WHERE kind = 'Drama';",
// 			exp: &Update{
// 				TableRef: &RangeVar{
// 					RelationName: "films",
// 					SchemaName:   "",
// 					Alias:        "",
// 				},
// 				Where: &WhereClauseLeaf{
// 					ColRef: ColumnRef{
// 						TableAlias: "",
// 						ColName:    "kind",
// 					},
// 					Value: "Drama",
// 				},
// 			},
// 			err: nil,
// 		},

// 		/* another simple update */
// 		{
// 			query: `UPDATE weather SET temp_lo = temp_lo+1, temp_hi = temp_lo+15, prcp = DEFAULT
// 			WHERE city = 'San Francisco' AND date = '2003-07-03';`,
// 			exp: &Update{
// 				TableRef: &RangeVar{
// 					RelationName: "weather",
// 					SchemaName:   "",
// 					Alias:        "",
// 				},
// 				Where: &WhereClauseOp{
// 					Op: "AND",
// 					Left: &WhereClauseLeaf{
// 						ColRef: ColumnRef{
// 							TableAlias: "",
// 							ColName:    "city",
// 						},
// 						Value: "San Francisco",
// 					},
// 					Right: &WhereClauseLeaf{
// 						ColRef: ColumnRef{
// 							TableAlias: "",
// 							ColName:    "date",
// 						},
// 						Value: "2003-07-03",
// 					},
// 				},
// 			},
// 			err: nil,
// 		},

// 		/* another simple update  with returning clause */
// 		{
// 			query: `UPDATE weather SET temp_lo = temp_lo+1, temp_hi = temp_lo+15, prcp = DEFAULT
// 			WHERE city = 'San Francisco' AND date = '2003-07-03'
// 			RETURNING temp_lo, temp_hi, prcp;`,
// 			exp: &Update{
// 				TableRef: &RangeVar{
// 					RelationName: "weather",
// 					SchemaName:   "",
// 					Alias:        "",
// 				},
// 				Where: &WhereClauseOp{
// 					Op: "AND",
// 					Left: &WhereClauseLeaf{
// 						ColRef: ColumnRef{
// 							TableAlias: "",
// 							ColName:    "city",
// 						},
// 						Value: "San Francisco",
// 					},
// 					Right: &WhereClauseLeaf{
// 						ColRef: ColumnRef{
// 							TableAlias: "",
// 							ColName:    "date",
// 						},
// 						Value: "2003-07-03",
// 					},
// 				},
// 			},
// 			err: nil,
// 		},
// 	} {
// 		tmp, err := Parse(tt.query)

// 		assert.NoError(err, "query %s", tt.query)

// 		assert.Equal(tt.exp, tmp)
// 	}
// }

// func TestDelete(t *testing.T) {
// 	assert := assert.New(t)

// 	type tcase struct {
// 		query string
// 		exp   Statement
// 		err   error
// 	}

// 	for _, tt := range []tcase{
// 		/* delete without where */
// 		{
// 			query: "DELETE FROM films;",
// 			exp: &Delete{
// 				TableRef: &RangeVar{
// 					RelationName: "films",
// 					SchemaName:   "",
// 					Alias:        "",
// 				},
// 				Where: &WhereClauseEmpty{},
// 			},
// 			err: nil,
// 		},

// 		/* simple delete */
// 		{
// 			query: "DELETE FROM films WHERE kind <> 'Musical';",
// 			exp: &Delete{
// 				TableRef: &RangeVar{
// 					RelationName: "films",
// 					SchemaName:   "",
// 					Alias:        "",
// 				},
// 				Where: &WhereClauseLeaf{
// 					ColRef: ColumnRef{
// 						TableAlias: "",
// 						ColName:    "kind",
// 					},
// 					Value: "Musical",
// 				},
// 			},
// 			err: nil,
// 		},

// 		/* simple delete with returning */
// 		{
// 			query: "DELETE FROM tasks WHERE status = 'DONE' RETURNING *;",
// 			exp: &Delete{
// 				TableRef: &RangeVar{
// 					RelationName: "tasks",
// 					SchemaName:   "",
// 					Alias:        "",
// 				},
// 				Where: &WhereClauseLeaf{
// 					ColRef: ColumnRef{
// 						TableAlias: "",
// 						ColName:    "status",
// 					},
// 					Value: "DONE",
// 				},
// 			},
// 			err: nil,
// 		},
// 	} {
// 		tmp, err := Parse(tt.query)

// 		assert.NoError(err, "query %s", tt.query)

// 		assert.Equal(tt.exp, tmp)
// 	}
// }

// func TestCopy(t *testing.T) {
// 	assert := assert.New(t)

// 	type tcase struct {
// 		query string
// 		exp   Statement
// 		err   error
// 	}

// 	for _, tt := range []tcase{
// 		/* copy with where */
// 		{
// 			query: "COPY xx TO STDOUT WHERE id = 1",
// 			exp: &Copy{
// 				TableRef: &RangeVar{
// 					RelationName: "xx",
// 				},
// 				Where: &WhereClauseLeaf{
// 					ColRef: ColumnRef{
// 						ColName: "id",
// 					},
// 					Value: "1",
// 				},
// 			},
// 			err: nil,
// 		},
// 	} {
// 		tmp, err := Parse(tt.query)

// 		assert.NoError(err, "query %s", tt.query)

// 		assert.Equal(tt.exp, tmp)
// 	}
// }

// func TestSelectTargetLists(t *testing.T) {
// 	assert := assert.New(t)

// 	type tcase struct {
// 		query string
// 		exp   Statement
// 		err   error
// 	}

// 	for _, tt := range []tcase{
// 		{
// 			query: "SELECT pg_is_in_recovery(), id FROM tsa_test WHERE id = 22;",
// 			exp: &Select{
// 				FromClause: []FromClauseNode{
// 					&RangeVar{
// 						RelationName: "tsa_test",
// 					},
// 				},
// 				Where: &WhereClauseLeaf{
// 					ColRef: ColumnRef{
// 						ColName: "id",
// 					},
// 					Value: "22",
// 				},
// 			},
// 			err: nil,
// 		},
// 	} {
// 		tmp, err := Parse(tt.query)

// 		assert.NoError(err, "query %s", tt.query)

// 		assert.Equal(tt.exp, tmp)
// 	}
// }

// func TestJoins(t *testing.T) {
// 	assert := assert.New(t)

// 	type tcase struct {
// 		query string
// 		exp   Statement
// 		err   error
// 	}

// 	for _, tt := range []tcase{
// 		{
// 			query: "SELECT * FROM delivery JOIN orders ON order_id = id;",
// 			exp: &Select{
// 				FromClause: []FromClauseNode{
// 					&JoinExpr{
// 						Larg: &RangeVar{
// 							RelationName: "delivery",
// 						},
// 						Rarg: &RangeVar{
// 							RelationName: "orders",
// 						},
// 					},
// 				},
// 				Where: &WhereClauseEmpty{},
// 			},
// 			err: nil,
// 		},

// 		{
// 			query: "SELECT * FROM sshjt1 a join sshjt1 b ON TRUE WHERE a.i = 12 AND b.j = a.j;",
// 			exp: &Select{
// 				FromClause: []FromClauseNode{
// 					&JoinExpr{
// 						Larg: &RangeVar{
// 							RelationName: "sshjt1",
// 							Alias:        "a",
// 						},
// 						Rarg: &RangeVar{
// 							RelationName: "sshjt1",
// 							Alias:        "b",
// 						},
// 					},
// 				},
// 				Where: &WhereClauseOp{
// 					Left: &WhereClauseLeaf{
// 						ColRef: ColumnRef{
// 							ColName:    "i",
// 							TableAlias: "a",
// 						},
// 						Value: "12",
// 					},
// 					Right: &WhereClauseEmpty{},
// 					Op:    "AND",
// 				},
// 			},
// 			err: nil,
// 		},
// 	} {
// 		tmp, err := Parse(tt.query)

// 		assert.NoError(err, "query %s", tt.query)

// 		assert.Equal(tt.exp, tmp)
// 	}
// }

// func TestMisc(t *testing.T) {
// 	assert := assert.New(t)

// 	type tcase struct {
// 		query string
// 		exp   Statement
// 		err   error
// 	}

// 	for _, tt := range []tcase{
// 		{
// 			query: "vacuum xx;",
// 			exp:   &Vacuum{},
// 			err:   nil,
// 		},
// 		{
// 			query: "analyze xx;",
// 			exp:   &Analyze{},
// 			err:   nil,
// 		},
// 		{
// 			query: "cluster xx;",
// 			exp:   &Cluster{},
// 			err:   nil,
// 		},
// 		{
// 			query: ";",
// 			exp:   nil,
// 			err:   nil,
// 		},
// 	} {
// 		tmp, err := Parse(tt.query)

// 		assert.NoError(err, "query %s", tt.query)

// 		assert.Equal(tt.exp, tmp)
// 	}
// }

// func TestOperators(t *testing.T) {
// 	assert := assert.New(t)

// 	type tcase struct {
// 		query string
// 		exp   Statement
// 		err   error
// 	}

// 	for _, tt := range []tcase{
// 		{
// 			query: "SELECT 1 ~~~~~~~~ 'dewoijoi'",
// 			exp: &Select{
// 				Where: &AExprEmpty{},
// 			},
// 			err: nil,
// 		},
// 		{
// 			query: "SELECT 1 ~~~^%^%^^~~~~~ 'dewoijoi'",
// 			exp: &Select{
// 				Where: &AExprEmpty{},
// 			},
// 			err: nil,
// 		},
// 		{
// 			query: "SELECT * FROM xxtt1 a WHERE a.w_id = 21 and j + i != 0;'",
// 			exp: &Select{
// 				Where: &WhereClauseEmpty{},
// 			},
// 			err: nil,
// 		},
// 	} {
// 		tmp, err := Parse(tt.query)

// 		assert.NoError(err, "query %s", tt.query)

// 		assert.Equal(tt.exp, tmp)
// 	}
// }
