# lyx

lyx is a SQL parser. It tokenizes, parses SQL queries and creates Query structures. Inspired by [src/backend/parser](https://github.com/postgres/postgres/tree/master/src/backend/parser).


```
ast.go     contains some useful structures
gram.go    generated from gram.y by goyacc
gram.y     parse the tokens and produce a "raw" parse tree
lexer.go   generated file from lexer.rl by regel
lexer.rl   lexer is lexer
lx_test.go test lexer
util.go    util commands
yx_test.go test SQL queries parse
```

See also [SPQR](github.com/pg-sharding/spqr).
