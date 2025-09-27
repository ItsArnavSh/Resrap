package main

import (
	"fmt"
	"time"

	"github.com/ItsArnavSh/Resrap/resrap"
)

func main() {
	// Config
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
