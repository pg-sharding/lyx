
test: 
	go test ./lyx/...
yaccgen:
	goyacc -o lyx/gram.go -p yy lyx/gram.y
gen:
	ragel -Z -G2 -o lyx/lexer.go lyx/lexer.rl
sync_reserved_words:
	go run cmd/reserved_words/main.go

build: gen yaccgen
