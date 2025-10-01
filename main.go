package main

import (
	"fmt"

	"github.com/osdc/resrap/resrap"
)

func main() {
	r := resrap.NewResrap()
	err := r.ParseGrammarFile("C", "example/Infinity.g4")
	if err != nil {
		fmt.Println(err)
	}

	code := r.GenerateRandom("C", "program", 1000)

	fmt.Println(code)
}
