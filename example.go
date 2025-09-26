package main

import (
	"fmt"
	"m/resrap"
)

func examplecode() {
	graphs := resrap.NewResrap()
	graphs.ParseGrammarFile("C", "example/C.g4")
	random_content := graphs.GenerateRandom("C", "program")
	fmt.Println(random_content)
	seeded_content := graphs.GenerateWithSeeded("C", "program", 20)
	fmt.Println(seeded_content)
}
