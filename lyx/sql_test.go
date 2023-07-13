package routerparser_test

import (
	"testing"

	routerparser "github.com/pg-sharding/spqr/yacc/router"
	"github.com/stretchr/testify/assert"
)

/*

 */

func TestSelectComplex(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   routerparser.Statement
		err   error
	}

	for _, tt := range []tcase{
		{
			query: `
			SELECT * FROM "exschema"."extable" WHERE id='83912839012903' AND utype='2' AND btype='sample' AND state = 0 AND is_something = true AND (keys @> '{reshke,denchick}' OR keys @> '{munakoiso,werelaxe,x4mmm}') AND c_id = 'trunk' ORDER BY entity_id asc;
			`,
			exp: &routerparser.Select{
				FromClause: []routerparser.FromClauseNode{&routerparser.RangeVar{
					RelationName: "extable",
					SchemaName:   "exschema",
				},
				},
				Where: &routerparser.AExprOp{
					Op: "AND",
					Left: &routerparser.AExprOp{
						ColRef: routerparser.ColumnRef{
							TableAlias: "",
							ColName:    "id",
						},
						Value: "83912839012903",
					},
					Right: &routerparser.WhereClauseOp{
						Op: "AND",
						Left: &routerparser.WhereClauseLeaf{
							ColRef: routerparser.ColumnRef{
								TableAlias: "",
								ColName:    "utype",
							},
							Value: "2",
						},
						Right: &routerparser.WhereClauseOp{
							Op: "AND",
							Left: &routerparser.WhereClauseLeaf{
								ColRef: routerparser.ColumnRef{
									TableAlias: "",
									ColName:    "btype",
								},
								Value: "sample",
							},
							Right: &routerparser.WhereClauseOp{
								Op: "AND",
								Left: &routerparser.WhereClauseLeaf{
									ColRef: routerparser.ColumnRef{
										TableAlias: "",
										ColName:    "state",
									},
									Value: "0",
								},
								Right: &routerparser.WhereClauseOp{
									Op: "AND",
									Left: &routerparser.WhereClauseLeaf{
										ColRef: routerparser.ColumnRef{
											TableAlias: "",
											ColName:    "is_something",
										},
										Value: "true",
									},
									Right: &routerparser.WhereClauseOp{
										Op: "AND",
										Left: &routerparser.WhereClauseOp{
											Left: &routerparser.WhereClauseLeaf{
												ColRef: routerparser.ColumnRef{
													ColName: "keys",
												},
												Value: "{reshke,denchick}",
											},
											Right: &routerparser.WhereClauseLeaf{
												ColRef: routerparser.ColumnRef{
													ColName: "keys",
												},
												Value: "{munakoiso,werelaxe,x4mmm}",
											},
											Op: "OR",
										},
										Right: &routerparser.WhereClauseLeaf{
											ColRef: routerparser.ColumnRef{
												TableAlias: "",
												ColName:    "c_id",
											},
											Value: "trunk",
										},
									},
								},
							},
						},
					},
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
			exp: &routerparser.Select{
				FromClause: []routerparser.FromClauseNode{&routerparser.RangeVar{
					RelationName: "extable",
					SchemaName:   "exschema",
				},
				},
				Where: &routerparser.WhereClauseOp{
					Op: "AND",
					Left: &routerparser.WhereClauseLeaf{
						ColRef: routerparser.ColumnRef{
							TableAlias: "",
							ColName:    "id",
						},
						Value: "83912839012903",
					},
					Right: &routerparser.WhereClauseOp{
						Op: "AND",
						Left: &routerparser.WhereClauseLeaf{
							ColRef: routerparser.ColumnRef{
								TableAlias: "",
								ColName:    "utype",
							},
							Value: "2",
						},
						Right: &routerparser.WhereClauseOp{
							Op: "AND",
							Left: &routerparser.WhereClauseLeaf{
								ColRef: routerparser.ColumnRef{
									TableAlias: "",
									ColName:    "btype",
								},
								Value: "sample",
							},
							Right: &routerparser.WhereClauseOp{
								Op: "AND",
								Left: &routerparser.WhereClauseLeaf{
									ColRef: routerparser.ColumnRef{
										TableAlias: "",
										ColName:    "state",
									},
									Value: "0",
								},
								Right: &routerparser.WhereClauseOp{
									Op: "AND",
									Left: &routerparser.WhereClauseLeaf{
										ColRef: routerparser.ColumnRef{
											TableAlias: "",
											ColName:    "is_something",
										},
										Value: "true",
									},
									Right: &routerparser.WhereClauseOp{
										Op: "AND",
										Left: &routerparser.WhereClauseOp{
											Left: &routerparser.WhereClauseLeaf{
												ColRef: routerparser.ColumnRef{
													ColName: "keys",
												},
												Value: "{reshke,denchick}",
											},
											Right: &routerparser.WhereClauseLeaf{
												ColRef: routerparser.ColumnRef{
													ColName: "keys",
												},
												Value: "{munakoiso,werelaxe,x4mmm}",
											},
											Op: "OR",
										},
										Right: &routerparser.WhereClauseLeaf{
											ColRef: routerparser.ColumnRef{
												TableAlias: "",
												ColName:    "c_id",
											},
											Value: "trunk",
										},
									},
								},
							},
						},
					},
				},
			},

			err: nil,
		},
	} {
		tmp, err := routerparser.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp)
	}
}

func TestSelect(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   routerparser.Statement
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "select 42",
			exp: &routerparser.Select{
				Where: &routerparser.WhereClauseEmpty{},
			},
			err: nil,
		},
		{
			query: "select * from xx where i = 1 ",
			exp: &routerparser.Select{
				FromClause: []routerparser.FromClauseNode{&routerparser.RangeVar{
					RelationName: "xx",
				},
				},
				Where: &routerparser.WhereClauseLeaf{
					ColRef: routerparser.ColumnRef{
						ColName:    "i",
						TableAlias: "",
					},
					Value: "1",
				},
			},
			err: nil,
		},

		{
			query: "select * from xx where i = 1 order by i ",
			exp: &routerparser.Select{
				FromClause: []routerparser.FromClauseNode{&routerparser.RangeVar{
					RelationName: "xx",
				},
				},
				Where: &routerparser.WhereClauseLeaf{
					ColRef: routerparser.ColumnRef{
						ColName:    "i",
						TableAlias: "",
					},
					Value: "1",
				},
			},
			err: nil,
		},

		{
			query: "select * from xx where i = 1 limit 7 ",
			exp: &routerparser.Select{
				FromClause: []routerparser.FromClauseNode{&routerparser.RangeVar{
					RelationName: "xx",
				},
				},
				Where: &routerparser.WhereClauseLeaf{
					ColRef: routerparser.ColumnRef{
						ColName:    "i",
						TableAlias: "",
					},
					Value: "1",
				},
			},
			err: nil,
		},

		{
			query: "select * from xx where i = 1 group by i ",
			exp: &routerparser.Select{
				FromClause: []routerparser.FromClauseNode{&routerparser.RangeVar{
					RelationName: "xx",
				},
				},
				Where: &routerparser.WhereClauseLeaf{
					ColRef: routerparser.ColumnRef{
						ColName:    "i",
						TableAlias: "",
					},
					Value: "1",
				},
			},
			err: nil,
		},

		{
			query: "select * from xx where i = 1 group by i having sum(i)",
			exp: &routerparser.Select{
				FromClause: []routerparser.FromClauseNode{&routerparser.RangeVar{
					RelationName: "xx",
				},
				},
				Where: &routerparser.WhereClauseLeaf{
					ColRef: routerparser.ColumnRef{
						ColName:    "i",
						TableAlias: "",
					},
					Value: "1",
				},
			},
			err: nil,
		},

		{
			query: "select * from xx where i = 1 order by i limit 7 ",
			exp: &routerparser.Select{
				FromClause: []routerparser.FromClauseNode{&routerparser.RangeVar{
					RelationName: "xx",
				},
				},
				Where: &routerparser.WhereClauseLeaf{
					ColRef: routerparser.ColumnRef{
						ColName:    "i",
						TableAlias: "",
					},
					Value: "1",
				},
			},
			err: nil,
		},
		{
			query: "select * from xx where i = 1 AND j = 2 ",
			exp: &routerparser.Select{
				FromClause: []routerparser.FromClauseNode{&routerparser.RangeVar{
					RelationName: "xx",
				},
				},
				Where: &routerparser.WhereClauseOp{
					Op: "AND",
					Left: &routerparser.WhereClauseLeaf{
						ColRef: routerparser.ColumnRef{
							ColName:    "i",
							TableAlias: "",
						},
						Value: "1",
					},
					Right: &routerparser.WhereClauseLeaf{
						ColRef: routerparser.ColumnRef{
							ColName:    "j",
							TableAlias: "",
						},
						Value: "2",
					},
				},
			},
			err: nil,
		},
	} {
		tmp, err := routerparser.Parse(tt.query)

		assert.NoError(err)

		assert.Equal(tt.exp, tmp)
	}
}

func TestSelectMultipleRelations(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   routerparser.Statement
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "select * from xx, xx2 ",
			exp: &routerparser.Select{
				FromClause: []routerparser.FromClauseNode{&routerparser.RangeVar{
					RelationName: "xx",
				},
					&routerparser.RangeVar{
						RelationName: "xx2",
					},
				},
				Where: &routerparser.WhereClauseEmpty{},
			},
			err: nil,
		},

		{
			query: "select * from xx, xx2 b where b.i = 1",
			exp: &routerparser.Select{
				FromClause: []routerparser.FromClauseNode{&routerparser.RangeVar{
					RelationName: "xx",
				},
					&routerparser.RangeVar{
						RelationName: "xx2",
						Alias:        "b",
					},
				},
				Where: &routerparser.WhereClauseLeaf{
					ColRef: routerparser.ColumnRef{
						ColName:    "i",
						TableAlias: "b",
					},
					Value: "1",
				},
			},
			err: nil,
		},
	} {
		tmp, err := routerparser.Parse(tt.query)

		assert.NoError(err)

		assert.Equal(tt.exp, tmp)
	}
}

func TestSelectAlias(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   routerparser.Statement
		err   error
	}

	for _, tt := range []tcase{

		{
			query: "SELECT kind, sum(len) AS total FROM films GROUP BY kind;",
			exp: &routerparser.Select{
				FromClause: []routerparser.FromClauseNode{&routerparser.RangeVar{
					RelationName: "films",
				},
				},
				Where: &routerparser.WhereClauseEmpty{},
			},
			err: nil,
		},

		{
			query: "SELECT kind, sum(len) AS total FROM films AS f GROUP BY kind;",
			exp: &routerparser.Select{
				FromClause: []routerparser.FromClauseNode{&routerparser.RangeVar{
					RelationName: "films",
					Alias:        "f",
				},
				},
				Where: &routerparser.WhereClauseEmpty{},
			},
			err: nil,
		},
	} {
		tmp, err := routerparser.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp)
	}
}

func TestCreate(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   routerparser.Statement
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "create table xx ( i int )",
			exp: &routerparser.CreateTable{
				TableName: "xx",
				TableElts: []routerparser.TableElt{
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
			exp: &routerparser.CreateTable{
				TableName: "xx",
				TableElts: []routerparser.TableElt{
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
			exp: &routerparser.CreateTable{
				TableName: "tt",
				TableElts: []routerparser.TableElt{
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
			exp: &routerparser.CreateTable{
				TableName: "sshjt1",
				TableElts: []routerparser.TableElt{
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
			exp: &routerparser.CreateTable{
				TableName: "orders",
				TableElts: []routerparser.TableElt{
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
			exp: &routerparser.CreateTable{
				TableName: "delivery",
				TableElts: []routerparser.TableElt{
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
			exp:   &routerparser.CreateDatabase{},
			err:   nil,
		},

		{
			query: "create role reg",
			exp:   &routerparser.CreateRole{},
			err:   nil,
		},
	} {
		tmp, err := routerparser.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp)
	}
}

func TestInsert(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   routerparser.Statement
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "insert into xx (id) values(1)",
			exp: &routerparser.Insert{
				TableRef: &routerparser.RangeVar{
					RelationName: "xx",
					SchemaName:   "",
					Alias:        "",
				},
				Columns: []string{"id"},
				Values:  []string{"1"},
			},
			err: nil,
		},

		{
			query: "Insert into xx (i) select * from yy where i = 8",
			exp: &routerparser.Insert{
				TableRef: &routerparser.RangeVar{
					RelationName: "xx",
					SchemaName:   "",
					Alias:        "",
				},
				Columns: nil,
				Values:  nil,
				SubSelect: &routerparser.Select{
					FromClause: []routerparser.FromClauseNode{
						&routerparser.RangeVar{
							RelationName: "yy",
						},
					},
					Where: &routerparser.WhereClauseLeaf{
						ColRef: routerparser.ColumnRef{
							ColName: "i",
						},
						Value: "8",
					},
				},
			},
			err: nil,
		},

		{
			query: `INSERT INTO films (code, title, did, date_prod, kind)
			VALUES ('T_601', 'Yojimbo', 106, '1961-06-16', 'Drama');`,
			exp: &routerparser.Insert{
				TableRef: &routerparser.RangeVar{
					RelationName: "films",
					SchemaName:   "",
					Alias:        "",
				},
				Columns: []string{"kind", "date_prod", "did", "title", "code"},
				Values:  []string{"Drama", "1961-06-16", "106", "Yojimbo", "T_601"},
			},
			err: nil,
		},

		///

		{
			query: `INSERT INTO films SELECT * FROM tmp_films WHERE date_prod < '2004-05-07';`,
			exp: &routerparser.Insert{
				TableRef: &routerparser.RangeVar{
					RelationName: "films",
					SchemaName:   "",
					Alias:        "",
				},
				SubSelect: &routerparser.Select{
					FromClause: []routerparser.FromClauseNode{&routerparser.RangeVar{
						RelationName: "tmp_films",
					},
					},
					Where: &routerparser.WhereClauseLeaf{
						ColRef: routerparser.ColumnRef{
							ColName:    "date_prod",
							TableAlias: "",
						},
						Value: "2004-05-07",
					},
				},
			},
			err: nil,
		},

		{
			query: "insert into xx (id,id2) values(1,2)",
			exp: &routerparser.Insert{
				TableRef: &routerparser.RangeVar{
					RelationName: "xx",
					SchemaName:   "",
					Alias:        "",
				},
				Columns: []string{"id2", "id"},
				Values:  []string{"2", "1"},
			},
			err: nil,
		},
		{
			query: "INSERT INTO xxtt1 (j, i, w_id) VALUES(2121221, -211212, 21);",
			exp: &routerparser.Insert{
				TableRef: &routerparser.RangeVar{
					RelationName: "xxtt1",
					SchemaName:   "",
					Alias:        "",
				},
				Columns: []string{"w_id", "i", "j"},
				Values:  []string{"21", "2121221", "2121221"},
			},
			err: nil,
		},

		/* only first tuple from values clause parsed  */
		{
			query: "insert into xx (id,id2) values(1,2), (2,3), ( 4, 5)",
			exp: &routerparser.Insert{
				TableRef: &routerparser.RangeVar{
					RelationName: "xx",
					SchemaName:   "",
					Alias:        "",
				},
				Columns: []string{"id2", "id"},
				Values:  []string{"2", "1"},
			},
			err: nil,
		},

		/* insert from select */
		{
			query: "insert into xx select * from xx2 where id2 = 7",
			exp: &routerparser.Insert{
				TableRef: &routerparser.RangeVar{
					RelationName: "xx",
					SchemaName:   "",
					Alias:        "",
				},
				SubSelect: &routerparser.Select{
					FromClause: []routerparser.FromClauseNode{&routerparser.RangeVar{
						RelationName: "xx2",
					},
					},
					Where: &routerparser.WhereClauseLeaf{
						ColRef: routerparser.ColumnRef{
							TableAlias: "",
							ColName:    "id2",
						},
						Value: "7",
					},
				},
			},
			err: nil,
		},
	} {
		tmp, err := routerparser.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp)
	}
}

func TestInsertComplex(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   routerparser.Statement
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
			exp: &routerparser.Insert{
				TableRef: &routerparser.RangeVar{
					RelationName: "extable",
					SchemaName:   "exschema",
					Alias:        "",
				},
				Values: []string{
					"0",
					"0",
					"bfuiqwefbIUGEIUWhgui..012-2-03849012381==-=-~~~?!@$#@#%%^&*^*(*)../././",
					"31337",
					"{9223372036854775806}",
					"{280,fb8,909,e6,ffc}",
					"true",
					"trunk",
					"0",
					"*()*()Q*D()beiwe",
					"thing",
					"some",
					"2",
					"9223372036854775807",
					"NULL",
					"NULL",
					"111111",
					"1970-01-01 12:00:00.5",
				},
				Columns: []string{
					"state",
					"value_type",
					"key_value",
					"key_id",
					"tags",
					"keys",
					"isnotok",
					"isok",
					"version",
					"id",
					"type",
					"kek",
					"lol",
					"bee",
					"fly",
					"dog",
					"fox", "cat",
				}, SubSelect: nil},
			err: nil,
		},
	} {
		tmp, err := routerparser.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp)
	}
}

func TestUpdate(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   routerparser.Statement
		err   error
	}

	for _, tt := range []tcase{
		/* simple update */
		{
			query: "UPDATE films SET kind = 'Dramatic' WHERE kind = 'Drama';",
			exp: &routerparser.Update{
				TableRef: &routerparser.RangeVar{
					RelationName: "films",
					SchemaName:   "",
					Alias:        "",
				},
				Where: &routerparser.WhereClauseLeaf{
					ColRef: routerparser.ColumnRef{
						TableAlias: "",
						ColName:    "kind",
					},
					Value: "Drama",
				},
			},
			err: nil,
		},

		/* another simple update */
		{
			query: `UPDATE weather SET temp_lo = temp_lo+1, temp_hi = temp_lo+15, prcp = DEFAULT
			WHERE city = 'San Francisco' AND date = '2003-07-03';`,
			exp: &routerparser.Update{
				TableRef: &routerparser.RangeVar{
					RelationName: "weather",
					SchemaName:   "",
					Alias:        "",
				},
				Where: &routerparser.WhereClauseOp{
					Op: "AND",
					Left: &routerparser.WhereClauseLeaf{
						ColRef: routerparser.ColumnRef{
							TableAlias: "",
							ColName:    "city",
						},
						Value: "San Francisco",
					},
					Right: &routerparser.WhereClauseLeaf{
						ColRef: routerparser.ColumnRef{
							TableAlias: "",
							ColName:    "date",
						},
						Value: "2003-07-03",
					},
				},
			},
			err: nil,
		},

		/* another simple update  with returning clause */
		{
			query: `UPDATE weather SET temp_lo = temp_lo+1, temp_hi = temp_lo+15, prcp = DEFAULT
			WHERE city = 'San Francisco' AND date = '2003-07-03'
			RETURNING temp_lo, temp_hi, prcp;`,
			exp: &routerparser.Update{
				TableRef: &routerparser.RangeVar{
					RelationName: "weather",
					SchemaName:   "",
					Alias:        "",
				},
				Where: &routerparser.WhereClauseOp{
					Op: "AND",
					Left: &routerparser.WhereClauseLeaf{
						ColRef: routerparser.ColumnRef{
							TableAlias: "",
							ColName:    "city",
						},
						Value: "San Francisco",
					},
					Right: &routerparser.WhereClauseLeaf{
						ColRef: routerparser.ColumnRef{
							TableAlias: "",
							ColName:    "date",
						},
						Value: "2003-07-03",
					},
				},
			},
			err: nil,
		},
	} {
		tmp, err := routerparser.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp)
	}
}

func TestDelete(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   routerparser.Statement
		err   error
	}

	for _, tt := range []tcase{
		/* delete without where */
		{
			query: "DELETE FROM films;",
			exp: &routerparser.Delete{
				TableRef: &routerparser.RangeVar{
					RelationName: "films",
					SchemaName:   "",
					Alias:        "",
				},
				Where: &routerparser.WhereClauseEmpty{},
			},
			err: nil,
		},

		/* simple delete */
		{
			query: "DELETE FROM films WHERE kind <> 'Musical';",
			exp: &routerparser.Delete{
				TableRef: &routerparser.RangeVar{
					RelationName: "films",
					SchemaName:   "",
					Alias:        "",
				},
				Where: &routerparser.WhereClauseLeaf{
					ColRef: routerparser.ColumnRef{
						TableAlias: "",
						ColName:    "kind",
					},
					Value: "Musical",
				},
			},
			err: nil,
		},

		/* simple delete with returning */
		{
			query: "DELETE FROM tasks WHERE status = 'DONE' RETURNING *;",
			exp: &routerparser.Delete{
				TableRef: &routerparser.RangeVar{
					RelationName: "tasks",
					SchemaName:   "",
					Alias:        "",
				},
				Where: &routerparser.WhereClauseLeaf{
					ColRef: routerparser.ColumnRef{
						TableAlias: "",
						ColName:    "status",
					},
					Value: "DONE",
				},
			},
			err: nil,
		},
	} {
		tmp, err := routerparser.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp)
	}
}

func TestCopy(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   routerparser.Statement
		err   error
	}

	for _, tt := range []tcase{
		/* copy with where */
		{
			query: "COPY xx TO STDOUT WHERE id = 1",
			exp: &routerparser.Copy{
				TableRef: &routerparser.RangeVar{
					RelationName: "xx",
				},
				Where: &routerparser.WhereClauseLeaf{
					ColRef: routerparser.ColumnRef{
						ColName: "id",
					},
					Value: "1",
				},
			},
			err: nil,
		},
	} {
		tmp, err := routerparser.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp)
	}
}

func TestSelectTargetLists(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   routerparser.Statement
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "SELECT pg_is_in_recovery(), id FROM tsa_test WHERE id = 22;",
			exp: &routerparser.Select{
				FromClause: []routerparser.FromClauseNode{
					&routerparser.RangeVar{
						RelationName: "tsa_test",
					},
				},
				Where: &routerparser.WhereClauseLeaf{
					ColRef: routerparser.ColumnRef{
						ColName: "id",
					},
					Value: "22",
				},
			},
			err: nil,
		},
	} {
		tmp, err := routerparser.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp)
	}
}

func TestJoins(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   routerparser.Statement
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "SELECT * FROM delivery JOIN orders ON order_id = id;",
			exp: &routerparser.Select{
				FromClause: []routerparser.FromClauseNode{
					&routerparser.JoinExpr{
						Larg: &routerparser.RangeVar{
							RelationName: "delivery",
						},
						Rarg: &routerparser.RangeVar{
							RelationName: "orders",
						},
					},
				},
				Where: &routerparser.WhereClauseEmpty{},
			},
			err: nil,
		},

		{
			query: "SELECT * FROM sshjt1 a join sshjt1 b ON TRUE WHERE a.i = 12 AND b.j = a.j;",
			exp: &routerparser.Select{
				FromClause: []routerparser.FromClauseNode{
					&routerparser.JoinExpr{
						Larg: &routerparser.RangeVar{
							RelationName: "sshjt1",
							Alias:        "a",
						},
						Rarg: &routerparser.RangeVar{
							RelationName: "sshjt1",
							Alias:        "b",
						},
					},
				},
				Where: &routerparser.WhereClauseOp{
					Left: &routerparser.WhereClauseLeaf{
						ColRef: routerparser.ColumnRef{
							ColName:    "i",
							TableAlias: "a",
						},
						Value: "12",
					},
					Right: &routerparser.WhereClauseEmpty{},
					Op:    "AND",
				},
			},
			err: nil,
		},
	} {
		tmp, err := routerparser.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp)
	}
}

func TestMisc(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   routerparser.Statement
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "vacuum xx;",
			exp:   &routerparser.Vacuum{},
			err:   nil,
		},
		{
			query: "analyze xx;",
			exp:   &routerparser.Analyze{},
			err:   nil,
		},
		{
			query: "cluster xx;",
			exp:   &routerparser.Cluster{},
			err:   nil,
		},
		{
			query: ";",
			exp:   nil,
			err:   nil,
		},
	} {
		tmp, err := routerparser.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp)
	}
}

func TestOperators(t *testing.T) {
	assert := assert.New(t)

	type tcase struct {
		query string
		exp   routerparser.Statement
		err   error
	}

	for _, tt := range []tcase{
		{
			query: "SELECT 1 ~~~~~~~~ 'dewoijoi'",
			exp: &routerparser.Select{
				Where: &routerparser.AExprEmpty{},
			},
			err: nil,
		},
		{
			query: "SELECT 1 ~~~^%^%^^~~~~~ 'dewoijoi'",
			exp: &routerparser.Select{
				Where: &routerparser.AExprEmpty{},
			},
			err: nil,
		},
		{
			query: "SELECT * FROM xxtt1 a WHERE a.w_id = 21 and j + i != 0;'",
			exp: &routerparser.Select{
				Where: &routerparser.WhereClauseEmpty{},
			},
			err: nil,
		},
	} {
		tmp, err := routerparser.Parse(tt.query)

		assert.NoError(err, "query %s", tt.query)

		assert.Equal(tt.exp, tmp)
	}
}
