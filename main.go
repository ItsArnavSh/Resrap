package main

import (
	"fmt"

	"github.com/osdc/resrap/resrap"
)

func main() {
	//Resrap with Single threaded
	rs := resrap.NewResrap()
	err := rs.ParseGrammarFile("C", "example/C.g4")
	if err != nil {
		fmt.Println(err)
		return
	}
	code := rs.GenerateWithSeeded("C", "program", 100, 100)
	fmt.Println(code)

}
