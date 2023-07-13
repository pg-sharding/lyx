package lyx

import (
	"errors"
	"fmt"
)

// Tokenizer is the struct used to generate SQL
// tokens for the parser.
type Tokenizer struct {
	s   string
	pos int

	ParseTree Statement
	LastError string
	l         *Lexer
}

func (t *Tokenizer) Error(s string) {
	t.LastError = s
}

func NewStringTokenizer(sql string) *Tokenizer {
	return &Tokenizer{
		s: sql,
		l: NewLexer([]byte(sql)),
	}
}

func (t *Tokenizer) Lex(lval *yySymType) int {
	return t.l.Lex(lval)
}

func (t *Tokenizer) LexT() int {
	return t.l.Lex(new(yySymType))
}

func setParseTree(yylex interface{}, stmt Statement) {
	yylex.(*Tokenizer).ParseTree = stmt
}

func Parse(sql string) (Statement, error) {
	tokenizer := NewStringTokenizer(sql)
	if yyParse(tokenizer) != 0 {
		return nil, errors.New(tokenizer.LastError + fmt.Sprintf(" on pos %d", tokenizer.l.ts))
	}
	ast := tokenizer.ParseTree
	return ast, nil
}
