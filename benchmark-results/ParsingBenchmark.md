

# Single-Threaded Benchmark — `Resrap`

### Apparatus

* **Grammar Used:** `C.g4` (from `example/`)
* **System:** Intel i7-12700H, 20 logical cores
* **Test Goal:** Measure time to parse grammar and generate a very long sequence of tokens in a **single-threaded** context.

---

### Benchmark Code

```go
package main

import (
	"fmt"
	"time"

	"yourmodule/resrap" // replace with the correct import path
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
```

---

### Benchmark Results

| Step                        | Time Taken |
| --------------------------- | ---------- |
| Grammar Parsing             | 332.489 µs |
| Generating 4,000,000 tokens | 2.663 s    |

---

### Analysis

* Parsing the grammar is **extremely fast** (<1 ms) because `C.g4` is parsed once and stored in an optimized in-memory structure.
* Generating **4 million tokens** in a single-threaded run takes **~2.66 seconds**, demonstrating the efficiency of `Resrap` even without multithreading.
* This benchmark serves as a baseline for **single-threaded performance**, which can later be compared against `ResrapMT` in multithreaded scenarios.
