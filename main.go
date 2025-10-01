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
	code := rs.GenerateRandom("C", "program", 10)
	fmt.Println(code)
	//Lets get a multithreaded API set up quick
	r := resrap.NewResrapMT(20, 1000) //20 worker pool and 1000 wait queue max size
	err = r.ParseGrammarFile("C", "example/C.g4")
	if err != nil {
		fmt.Println(err)
		return
	}
	//Receive from this
	r.StartResrap()
	defer r.ShutDownResrap()
	codeChan := r.GetCodeChannel()
	id := "12321"
	r.GenerateRandom(id, "C", "program", 10)
	res := <-codeChan
	fmt.Println(res.Code)
}
