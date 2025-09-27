# ResrapMT — Multithreaded Grammar-Based Code Generation

`ResrapMT` is the **multithreaded version** of `Resrap`, designed to efficiently generate code or text from grammars in parallel. It is ideal for handling **many simultaneous requests**, making it perfect for server environments or batch processing.

> For benchmarks and performance comparisons, see [benchmark-results/Multithreading.md](benchmark-results/Multithreading.md).

---

## Overview

`ResrapMT` allows:

- Parsing and storing multiple grammars (like `Resrap`)
- Generating content concurrently using a **worker pool**
- Handling large numbers of requests efficiently
- Non-blocking job submission with a **blocking result channel**

> **Note:** You are responsible for creating unique IDs for each request. The result channel returns a `CodeGenRes` struct with both the generated code and the request ID, allowing you to map results back to requests.

---

## Public Structs

### `ResrapMT`

```go
type ResrapMT struct {
    languageGraph map[string]lang
    poolsize      int // Number of threads in the pool
    waitqueuesize int
    pendingjobs   chan codeGenReq
    CodeChannel   chan CodeGenRes
}
````

* `poolsize` — number of worker goroutines processing jobs.
* `waitqueuesize` — buffer size for the pending job queue.
* `pendingjobs` — internal buffered channel for job requests.
* `CodeChannel` — public channel returning completed jobs with IDs.

---

## Public Methods

### `NewResrapMT(poolsize, waitqueuesize int) *ResrapMT`

Creates a new `ResrapMT` instance.

```go
resrapMT := resrap.NewResrapMT(10, 100)
```

* `poolsize` — number of worker threads.
* `waitqueuesize` — size of the pending job queue.

---

### `ParseGrammar(name, grammar string)`

Parses a grammar from a string and stores it under the given name.

```go
resrapMT.ParseGrammar("C", "grammar content here")
```

* `name` — unique grammar identifier (e.g., `"C"`).
* `grammar` — grammar definition as a string.
* The grammar is **normalized internally** after parsing.

---

### `ParseGrammarFile(name, location string)`

Parses a grammar from a file and stores it under the given name.

```go
resrapMT.ParseGrammarFile("C", "example/C.g4")
```

* `name` — unique grammar identifier.
* `location` — path to the grammar file.
* The grammar is **normalized internally** after parsing.

---

### `GenerateRandom(id, name, starting_node string, tokens int)`

Submits a **non-deterministic generation job** to the worker pool.

```go
resrapMT.GenerateRandom("job-100", "C", "program", 100)
```

* `id` — unique request identifier (used to map results).
* `name` — grammar identifier.
* `starting_node` — starting symbol in the grammar.
* `tokens` — number of tokens to generate.

> The result will be available on `resrapMT.CodeChannel` as a `CodeGenRes` struct containing the `Id` and generated `Code`. You are responsible for reading the channel and handling results.

---

### `GenerateWithSeeded(id, name, starting_node string, seed uint64, tokens int)`

Submits a **deterministic generation job** using a numeric seed.

```go
resrapMT.GenerateWithSeeded("job-101", "C", "program", 12345, 100)
```

* `id` — unique request identifier.
* `name` — grammar identifier.
* `starting_node` — starting symbol in the grammar.
* `seed` — numeric seed for deterministic generation.
* `tokens` — number of tokens to generate.

> Results appear on `CodeChannel` like `GenerateRandom`.

---

### `StartResrap()`

Starts the worker pool. Should be called **once** after parsing grammars and before submitting jobs.

```go
resrapMT.StartResrap()
```

* Spawns `poolsize` goroutines that process jobs from `pendingjobs`.
* Workers continuously listen for jobs and write results to `CodeChannel`.

---

## Example Usage

```go
package main

import (
    "fmt"
    "yourmodule/resrap"
)

func main() {
    // Create multi-threaded Resrap
    rmt := resrap.NewResrapMT(10, 100)
    rmt.ParseGrammarFile("C", "example/C.g4")
    rmt.StartResrap()

    // Submit jobs with unique IDs
    for i := 0; i < 10; i++ {
        jobID := fmt.Sprintf("job-%d", i)
        rmt.GenerateRandom(jobID, "C", "program", 100)
    }

    // Collect results
    for i := 0; i < 10; i++ {
        res := <-rmt.CodeChannel
        fmt.Printf("Job %s generated code: %s\n", res.Id, res.Code[:50])
    }
}
```

**Notes:**

* `CodeChannel` is a **blocking, unbounded channel**.
* You must **handle the channel yourself**, mapping results back to your request IDs.
* `ResrapMT` does **not automatically manage IDs** — they are your responsibility.

---

## Why Multithreaded?

`ResrapMT` is designed to handle **many concurrent jobs efficiently**.

* Ideal for **multi-user applications** or **batch processing** of grammar-based code generation.
* Fully utilizes multiple CPU cores while keeping grammar graphs **immutable and lock-free**.
* Delivers **significant speedup** compared to single-threaded `Resrap` for workloads with many small or medium-sized jobs.

> See [benchmark-results/Multithreading.md](benchmark-results/Multithreading.md) for detailed performance benchmarks.

---
