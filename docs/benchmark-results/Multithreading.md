

## Performance Boosts Using the Multithreaded Version (`ResrapMT`)

### Test Apparatus

* **CPU:** Intel i7-12700H (20 logical cores)
* **Test Program:** Generating 10,000 jobs of 100 tokens each

```go
func main() {
    // Configuration
    numJobs := 10000
    tokensPerJob := 100

    // --- SINGLE-THREADED RESRAP ---
    fmt.Println("Starting single-threaded test...")
    resrapSync := resrap.NewResrap()
    resrapSync.ParseGrammarFile("C", "example/C.g4")

    start := time.Now()
    for i := 0; i < numJobs; i++ {
        _ = resrapSync.GenerateRandom("C", "program", tokensPerJob)
    }
    elapsedSync := time.Since(start)
    fmt.Printf("Single-threaded: %v for %d jobs of %d tokens\n", elapsedSync, numJobs, tokensPerJob)

    // --- MULTI-THREADED RESRAPMT ---
    fmt.Println("Starting multi-threaded test...")
    poolSize := 20
    waitQueueSize := 10000
    resrapMT := resrap.NewResrapMT(poolSize, waitQueueSize)
    resrapMT.ParseGrammarFile("C", "example/C.g4")
    resrapMT.StartResrap()

    start = time.Now()
    for i := 0; i < numJobs; i++ {
        jobID := fmt.Sprintf("%d", i)
        resrapMT.GenerateRandom(jobID, "C", "program", tokensPerJob)
    }

    // Collect all results
    for i := 0; i < numJobs; i++ {
        <-resrapMT.CodeChannel
    }
    elapsedMT := time.Since(start)
    fmt.Printf("Multi-threaded: %v for %d jobs of %d tokens\n", elapsedMT, numJobs, tokensPerJob)
}
```

### Results

| Version                     | Time Taken | Jobs   | Tokens per Job |
| --------------------------- | ---------- | ------ | -------------- |
| Single-threaded (`Resrap`)  | 769.6 ms   | 10,000 | 100            |
| Multi-threaded (`ResrapMT`) | 153.9 ms   | 10,000 | 100            |

### Analysis

* **Speedup:** ~5× faster with `ResrapMT` for concurrent small jobs.
* **Reason:** Worker pool efficiently utilizes multiple cores, and grammar graphs are immutable, allowing lock-free concurrent generation.
* **Implication:** For **multi-user or bursty workloads**, `ResrapMT` significantly outperforms the single-threaded version, reducing latency and improving throughput.
