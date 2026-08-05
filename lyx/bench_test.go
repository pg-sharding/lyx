package lyx_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/pg-sharding/lyx/lyx"
)

const benchTuples = 250000

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

// buildSelectNoLiterals builds a long SELECT with only integer comparisons:
//
//	SELECT col0, col1, ... FROM t WHERE col0 = 0 AND col1 = 1 AND ...
func buildSelectNoLiterals(cols int) string {
	var b strings.Builder
	b.WriteString("SELECT ")
	for i := 0; i < cols; i++ {
		if i > 0 {
			b.WriteString(", ")
		}
		fmt.Fprintf(&b, "col%d", i)
	}
	b.WriteString(" FROM t WHERE ")
	for i := 0; i < cols; i++ {
		if i > 0 {
			b.WriteString(" AND ")
		}
		fmt.Fprintf(&b, "col%d = %d", i, i)
	}
	return b.String()
}

// buildInsertNoLiterals builds a long INSERT with only integer values:
//
//	INSERT INTO t (col0, col1, ...) VALUES (0, 1, ...)
func buildInsertNoLiterals(cols int) string {
	var b strings.Builder
	b.WriteString("INSERT INTO t (")
	for i := 0; i < cols; i++ {
		if i > 0 {
			b.WriteString(", ")
		}
		fmt.Fprintf(&b, "col%d", i)
	}
	b.WriteString(") VALUES (")
	for i := 0; i < cols; i++ {
		if i > 0 {
			b.WriteString(", ")
		}
		fmt.Fprintf(&b, "%d", i)
	}
	b.WriteByte(')')
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
	query := buildInsertWithManyTuples(benchTuples)

	b.SetBytes(int64(len(query)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, _, err := lyx.Parse(query); err != nil {
			b.Fatalf("parse failed: %v", err)
		}
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

func BenchmarkInsertNestedQuotes(b *testing.B) {
	query := buildInsertWithNestedQuotes(benchTuples)

	b.SetBytes(int64(len(query)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, _, err := lyx.Parse(query); err != nil {
			b.Fatalf("parse failed: %v", err)
		}
	}
}

// BenchmarkLexerSelectNoLiterals measures the Ragel lexer throughput on a
// wide SELECT with only identifier and integer tokens — no string literals.
func BenchmarkLexerSelectNoLiterals(b *testing.B) {
	for _, cols := range []int{100, 1000, 10000} {
		query := buildSelectNoLiterals(cols)
		b.Run(fmt.Sprintf("%d-cols", cols), func(b *testing.B) {
			b.SetBytes(int64(len(query)))
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				lexAll(query)
			}
		})
	}
}

// BenchmarkLexerInsertNoLiterals measures the Ragel lexer throughput on a
// wide INSERT with only integer values — no string literals.
func BenchmarkLexerInsertNoLiterals(b *testing.B) {
	for _, cols := range []int{100, 1000, 10000} {
		query := buildInsertNoLiterals(cols)
		b.Run(fmt.Sprintf("%d-cols", cols), func(b *testing.B) {
			b.SetBytes(int64(len(query)))
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				lexAll(query)
			}
		})
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
