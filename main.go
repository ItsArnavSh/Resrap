package main

import (
	"fmt"

	"github.com/osdc/resrap/resrap"
)

func main() {
	r := resrap.NewResrap()
	err := r.ParseGrammarFile("C", "example/C.g4")
	if err != nil {
		fmt.Println(err)
		return
	}
	code := r.GenerateRandom("C", "program", 100)
	fmt.Println(code)

}
