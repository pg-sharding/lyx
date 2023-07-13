gen:
	ragel -Z -G2 -o lx/lexer.go lx/lexer.rl
build: gen
	go build  -o main ./main.go
