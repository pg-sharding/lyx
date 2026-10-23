package lyx

import (
	"errors"
	"fmt"
)

// Tokenizer is the struct used to generate SQL
// tokens for the parser.
type Tokenizer struct {
	s string

	ParseTree []Node
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
	ret := t.l.Lex(lval)
	if ret == 0 {
		// EOF has no location, mark it as unknown.
		lval.loc = ParseLocUnknown
	} else {
		// Like YYLLOC_DEFAULT in PostgreSQL, the location of a
		// nonterminal is the start position of its first token.
		lval.loc = t.l.ts
	}
	return ret
}

func (t *Tokenizer) LexT() int {
	return t.l.Lex(new(yySymType))
}

func (t *Tokenizer) Reset(sql string) {
	t.s = sql
	ResetLexer(t.l, []byte(sql))
}

func setParseTree(yylex interface{}, stmt []Node) {
	yylex.(*Tokenizer).ParseTree = stmt
}

func Parse(sql string) ([]Node, int, error) {
	tokenizer := NewStringTokenizer(sql)
	p := yyNewParser()
	if p.Parse(tokenizer) != 0 {
		return nil, tokenizer.l.ts, errors.New(tokenizer.LastError)
	}
	ast := tokenizer.ParseTree
	return ast, 0, nil
}

func ParseWithLexerParser(l LyxParser, t *Tokenizer, sql string) ([]Node, int, error) {
	t.Reset(sql)
	if l.Parse(t) != 0 {
		return nil, 0, errors.New(t.LastError + fmt.Sprintf(" at or near %d", t.l.ts))
	}
	ast := t.ParseTree
	return ast, 0, nil
}
