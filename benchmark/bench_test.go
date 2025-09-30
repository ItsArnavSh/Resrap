package benchmark_test

import (
	"testing"

	"github.com/osdc/resrap/resrap"
)

func TestGenerateRandomCCode(t *testing.T) {
	r := resrap.NewResrap()

	// Parse the grammar file for C
	r.ParseGrammarFile("C", "../example/C.g4")

	// Generate random C code
	code := r.GenerateWithSeeded("C", "program", 10, 1000000)

	// Optionally log the code for inspection
	t.Logf("Generated C code:\n%s", code)
}
func BenchmarkGenerateRandomCCode(b *testing.B) {
	r := resrap.NewResrap()
	r.ParseGrammarFile("C", "../example/C.g4")

	b.ReportAllocs() // <-- Track memory allocations

	for b.Loop() {
		r.ParseGrammarFile("C", "../example/C.g4")
	}
}
