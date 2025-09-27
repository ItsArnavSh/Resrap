package main

import (
	"fmt"
	"time"

	"github.com/ItsArnavSh/Resrap/resrap"
)

func main() {
	// Create a new single-threaded Resrap
	resrapSync := resrap.NewResrap()

	// Parse grammar file (single-threaded)
	fmt.Println("Parsing grammar...")
	startParse := time.Now()
	resrapSync.ParseGrammarFile("C", "example/C.g4")
	parseElapsed := time.Since(startParse)
	fmt.Printf("Grammar parsed in %v\n", parseElapsed)

	// Generate a huge sequence (4 million tokens)
	numTokens := 4000000
	fmt.Printf("Generating %d tokens (single-threaded)...\n", numTokens)
	startGen := time.Now()
	code := resrapSync.GenerateRandom("C", "program", numTokens)
	genElapsed := time.Since(startGen)

	fmt.Printf("Generated %d tokens in %v\n", numTokens, genElapsed)
	fmt.Printf("First 500 chars of output: \n%s\n", code[:500])
}
