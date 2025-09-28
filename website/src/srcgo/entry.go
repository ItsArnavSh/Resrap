package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	println("Go WASM: Starting with Go 1.24.4...")

	// Your custom function
	js.Global().Set("generateText", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) < 4 {
			return "Error: Need 4 parameters: grammar, startpoint, seed, tokenlen"
		}

		grammar := args[0].String()
		startpoint := args[1].String()
		seed := args[2].Int()
		tokenlen := args[3].Int()

		println("Go WASM: generateText called with:")
		println("  grammar:", grammar)
		println("  startpoint:", startpoint)
		println("  seed:", seed)
		println("  tokenlen:", tokenlen)

		// Your logic here - for now just returning a formatted string
		result := fmt.Sprintf("Generated text with grammar='%s', startpoint='%s', seed=%d, tokenlen=%d",
			grammar, startpoint, seed, tokenlen)

		println("Go WASM: returning:", result)
		return result
	}))

	// Keep the existing test functions
	js.Global().Set("goTest", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return "Hello from Go 1.24.4 WASM!"
	}))

	js.Global().Set("goAdd", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) < 2 {
			return 0
		}
		return args[0].Int() + args[1].Int()
	}))

	println("Go WASM: All functions registered successfully")

	// Keep the program running
	<-make(chan bool)
}
