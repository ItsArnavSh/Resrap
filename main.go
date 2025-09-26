package main

import (
	"github.com/ItsArnavSh/Resrap/resrap"
)

func main() {
	resrap := resrap.NewResrap()
	resrap.ParseGrammarFile("code", "example/C.g4")
	_ = resrap.GenerateRandom("code", "program", 400)
	print("Done")
}
