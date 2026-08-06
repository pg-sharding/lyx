
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

generate:
	docker build -f docker/generator/Dockerfile -t lyx-generator .
	docker run --name lyx-generator-1 lyx-generator
	docker cp lyx-generator-1:/lyx/lyx/gram.go lyx/gram.go
	docker cp lyx-generator-1:/lyx/lyx/lexer.go lyx/lexer.go
	docker container rm lyx-generator-1
