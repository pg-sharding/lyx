package lyx_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/pg-sharding/lyx/lyx"
)

// lexAll drives only the Ragel lexer (no yacc) until EOF and returns the
// number of tokens produced. It is the correct way to benchmark the lexer in
// isolation.
func lexAll(query string) int {
	tok := lyx.NewStringTokenizer(query)
	n := 0
	for tok.LexT() != 0 {
		n++
	}
	return n
}

func buildSelectTokNames(ln int) string {
	var b strings.Builder
	b.WriteString("SELECT ")
	toknames := lyx.TokNames()
	for i := 0; i < ln; i++ {
		b.WriteString(toknames[i%len(toknames)])
	}
	return b.String()
}

func buildSelectLiterals(ln int, lit_len int) string {

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var b strings.Builder
	b.Grow(ln*lit_len + 2*lit_len + ln /* spaces */)
	for i := 0; i < ln; i++ {
		b.WriteByte('\'')
		for j := 0; j < lit_len; j++ {
			b.WriteByte(charset[j%len(charset)])
		}
		b.WriteByte('\'')
		b.WriteByte(' ')
	}
	return b.String()
}

/*

    ',' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TCOMMA; fbreak;};
    '(' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TOPENBR; fbreak;};
    ')' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TCLOSEBR; fbreak;};
    '[' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TSQOPENBR; fbreak;};
    ']' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TSQCLOSEBR; fbreak;};
    '.' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TDOT; fbreak;};
    ';' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TSEMICOLON; fbreak;};
    ':' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TCOLON; fbreak;};
    '+' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TPLUS; fbreak;};
    '-' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TMINUS; fbreak;};
    '*' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TMUL; fbreak;};
   # TODO: support '\\' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = int(TMUL); fbreak;};
    '%' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TMOD; fbreak;};
    '^' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TPOW; fbreak;};
    '<' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TLESS; fbreak;};
    '>' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TGREATER; fbreak;};
    '=' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TEQ; fbreak;};

    '<>' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TNOT_EQUALS; fbreak;};
    '<=' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TLESS_EQUALS; fbreak;};
    '>=' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TGREATER_EQUALS; fbreak;};
    '!=' => { lval.str = string(lex.data[lex.ts:lex.te]); tok = TNOT_EQUALS; fbreak;};
*/

func buildOperators(ln int) string {

	var ops = []string{",", ")", "(", "[", "]", ".", "::", ";", ":", "+", "-", "*", "$", "%", "^", "<", ">", "=", "<>", "<=", ">=", "!="}
	var b strings.Builder

	for i := 0; i < ln; i++ {
		b.WriteString(ops[i%len(ops)])
		b.WriteByte(' ')
	}
	return b.String()
}

func buildInsertWithLongStrings(n, strLen int) string {
	val := strings.Repeat("x", strLen)
	var b strings.Builder
	b.Grow(n * (strLen + 8))
	b.WriteString("INSERT INTO t (i) VALUES ")
	for i := 0; i < n; i++ {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteByte('(')
		b.WriteByte('\'')
		b.WriteString(val)
		b.WriteByte('\'')
		b.WriteByte(')')
	}
	return b.String()
}

func buildInsertWithManyTuples(n int) string {
	var b strings.Builder
	b.WriteString("INSERT INTO t (i) VALUES ")
	for i := 0; i < n; i++ {
		if i > 0 {
			b.WriteString(", ")
		}
		fmt.Fprintf(&b, "('str%d')", i)
	}
	return b.String()
}

// nested-quote literals exercise the ” escaping added in 40a34fd
func buildInsertWithNestedQuotes(n int) string {
	var b strings.Builder
	b.WriteString("INSERT INTO t (i) VALUES ")
	for i := 0; i < n; i++ {
		if i > 0 {
			b.WriteString(", ")
		}
		fmt.Fprintf(&b, "('str''%d''val')", i)
	}
	return b.String()
}

func BenchmarkInsertManyTuples(b *testing.B) {
	for _, benchTuples := range []int{1e4, 1e5, 3e5} {
		query := buildInsertWithManyTuples(benchTuples)

		b.Run(fmt.Sprintf("%d-len", benchTuples), func(b *testing.B) {
			b.SetBytes(int64(len(query)))
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {

				if _, _, err := lyx.Parse(query); err != nil {
					b.Fatalf("parse failed: %v", err)
				}

			}
		})
	}
}

func BenchmarkLexInsertManyTuples(b *testing.B) {

	for _, benchTuples := range []int{1e4, 1e5, 3e5} {
		query := buildInsertWithManyTuples(benchTuples)

		b.Run(fmt.Sprintf("%d-len", benchTuples), func(b *testing.B) {
			b.SetBytes(int64(len(query)))
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				lexAll(query)
			}
		})
	}
}

func BenchmarkInsertNestedQuotes(b *testing.B) {
	for _, benchTuples := range []int{1e4, 1e5, 3e5} {
		query := buildInsertWithNestedQuotes(benchTuples)

		b.Run(fmt.Sprintf("%d-len", benchTuples), func(b *testing.B) {
			b.SetBytes(int64(len(query)))
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {

				if _, _, err := lyx.Parse(query); err != nil {
					b.Fatalf("parse failed: %v", err)
				}

			}
		})
	}
}

func BenchmarkInsertLongStrings(b *testing.B) {
	for _, tt := range []struct {
		name   string
		tuples int
		strLen int
	}{
		{name: "1k-str/1k-tuples", tuples: 1000, strLen: 1024},
		{name: "64k-str/100-tuples", tuples: 100, strLen: 64 * 1024},
		{name: "640k-str/2-tuples", tuples: 2, strLen: 640 * 1024},
	} {
		b.Run(tt.name, func(b *testing.B) {
			query := buildInsertWithLongStrings(tt.tuples, tt.strLen)

			b.SetBytes(int64(len(query)))
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				if _, _, err := lyx.Parse(query); err != nil {
					b.Fatalf("parse failed: %v", err)
				}
			}
		})
	}
}

func BenchmarkLexInsertLongStrings(b *testing.B) {
	for _, tt := range []struct {
		name   string
		tuples int
		strLen int
	}{
		{name: "1k-str/1k-tuples", tuples: 1000, strLen: 1024},
		{name: "64k-str/100-tuples", tuples: 100, strLen: 64 * 1024},
		{name: "640k-str/2-tuples", tuples: 2, strLen: 640 * 1024},
	} {
		b.Run(tt.name, func(b *testing.B) {
			query := buildInsertWithLongStrings(tt.tuples, tt.strLen)

			b.Run(fmt.Sprintf("%d-tups-%d-len", tt.tuples, tt.strLen), func(b *testing.B) {
				b.SetBytes(int64(len(query)))
				b.ReportAllocs()
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					lexAll(query)
				}
			})
		})
	}
}

func BenchmarkLexerSelectNoLiterals(b *testing.B) {
	for _, ln := range []int{1e5, 1e6, 1e7} {
		query := buildSelectTokNames(ln)
		b.Run(fmt.Sprintf("%d-len", ln), func(b *testing.B) {
			b.SetBytes(int64(len(query)))
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				lexAll(query)
			}
		})
	}
}

func BenchmarkLexerSelectLiterals(b *testing.B) {
	for _, ln := range []int{1e4, 1e5, 1e6} {
		for _, lit_len := range []int{2, 10, 1e2} {
			query := buildSelectLiterals(ln, lit_len)
			b.Run(fmt.Sprintf("%d-len-%d-lit-len", ln, lit_len), func(b *testing.B) {
				b.SetBytes(int64(len(query)))
				b.ReportAllocs()
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					lexAll(query)
				}
			})
		}
	}
}

func BenchmarkLexerOperators(b *testing.B) {
	for _, ln := range []int{1e4, 1e5, 1e6} {
		query := buildOperators(ln)
		b.Run(fmt.Sprintf("ops-%d-len", ln), func(b *testing.B) {
			b.SetBytes(int64(len(query)))
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				lexAll(query)
			}
		})
	}
}

func BenchmarkLexerSelectLongLiterals(b *testing.B) {
	for _, ln := range []int{1e1, 1e2, 1e3} {
		for _, lit_len := range []int{1e3, 1e4, 1e5} {
			query := buildSelectLiterals(ln, lit_len)
			b.Run(fmt.Sprintf("long-%d-len-%d-lit-len", ln, lit_len), func(b *testing.B) {
				b.SetBytes(int64(len(query)))
				b.ReportAllocs()
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					lexAll(query)
				}
			})
		}
	}
}

const perfRuns = 5

type testCase struct {
	name     string
	tuples   int
	deadline time.Duration
}

type testGen struct {
	name  string
	build func(int) string
}

var runner = func(t *testing.T, gen testGen, tt testCase) {
	query := gen.build(tt.tuples)

	// warm up to avoid measuring one-time init
	if _, _, err := lyx.Parse(query); err != nil {
		t.Fatalf("parse failed: %v", err)
	}

	var total time.Duration
	for i := 0; i < perfRuns; i++ {
		start := time.Now()
		if _, _, err := lyx.Parse(query); err != nil {
			t.Fatalf("parse failed: %v", err)
		}
		total += time.Since(start)
	}
	avg := total / perfRuns

	if avg >= tt.deadline {
		t.Fatalf("parsing %d tuples: avg over %d runs = %s, want < %s",
			tt.tuples, perfRuns, avg, tt.deadline)
	}
	t.Logf("parsed %d tuples: avg %s over %d runs (limit %s)",
		tt.tuples, avg, perfRuns, tt.deadline)
}

func TestInsertManyTuplesPerf(t *testing.T) {
	for _, tt := range []testCase{
		{name: "10k", tuples: 10000, deadline: 20 * time.Millisecond},
		{name: "250k", tuples: 250000, deadline: time.Second},
	} {
		for _, gen := range []testGen{
			{name: "plain", build: buildInsertWithManyTuples},
			{name: "nested-quotes", build: buildInsertWithNestedQuotes},
		} {
			t.Run(tt.name+"/"+gen.name, func(t *testing.T) {
				runner(t, gen, tt)
			})
		}
	}
}

func TestInsertManyTuplesLongStringsPerf(t *testing.T) {
	for _, tt := range []testCase{
		{name: "10k", tuples: 10000, deadline: 200 * time.Millisecond},
		{name: "250k", tuples: 250000, deadline: 5 * time.Second},
	} {
		for _, gen := range []testGen{
			{name: "long-strings", build: func(n int) string { return buildInsertWithLongStrings(n, 1024) }},
		} {
			t.Run(tt.name+"/"+gen.name, func(t *testing.T) {
				runner(t, gen, tt)
			})
		}
	}
}
