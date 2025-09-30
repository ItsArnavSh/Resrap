package main

import (
	"fmt"

	"github.com/osdc/resrap/resrap"
)

func main() {
	r := resrap.NewResrap()
	r.ParseGrammarFile("C", "example/C.g4")
	code := r.GenerateRandom("C", "program", 100)
	fmt.Println(code)
}
