
test: 
	go test ./lyx/...

lextest: 
	go test ./lyx/lx_test.go

bench:
	go test -run '^$$' -bench 'BenchmarkInsert' -benchmem ./lyx/...

yaccgen:
	goyacc -o lyx/gram.go -p yy lyx/gram.y

gen:
	ragel -Z -G2 -o lyx/lexer.go lyx/lexer.rl

build: gen yaccgen
