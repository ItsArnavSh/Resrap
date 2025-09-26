package main

import (
	"fmt"

	"github.com/ItsArnavSh/Resrap/resrap"
)

func main() {
	resrap := resrap.NewResrap()
	resrap.ParseGrammarFile("code", "example/Infinity.g4")
	code := resrap.GenerateRandom("code", "program", 400)
	fmt.Println(code)
}
