package lyx_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/pg-sharding/lyx/lyx"
)

const benchTuples = 250000

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

// nested-quote literals exercise the '' escaping added in 40a34fd
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

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, _, err := lyx.Parse(query); err != nil {
			b.Fatalf("parse failed: %v", err)
		}
	}
}

func BenchmarkInsertNestedQuotes(b *testing.B) {
	query := buildInsertWithNestedQuotes(benchTuples)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, _, err := lyx.Parse(query); err != nil {
			b.Fatalf("parse failed: %v", err)
		}
	}
}

const perfRuns = 5

func TestInsertManyTuplesPerf(t *testing.T) {
	for _, tt := range []struct {
		name     string
		tuples   int
		deadline time.Duration
	}{
		{name: "10k", tuples: 10000, deadline: 20 * time.Millisecond},
		{name: "250k", tuples: 250000, deadline: time.Second},
	} {
		for _, gen := range []struct {
			name  string
			build func(int) string
		}{
			{name: "plain", build: buildInsertWithManyTuples},
			{name: "nested-quotes", build: buildInsertWithNestedQuotes},
		} {
			t.Run(tt.name+"/"+gen.name, func(t *testing.T) {
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
			})
		}
	}
}
