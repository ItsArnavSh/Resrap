package main

import (
	"fmt"
)

func parser(filename string) {
	lines, err := ParseFile(filename)
	if err != nil {
		return
	}
	fmt.Println("There are ", len(lines))
	sc := NewScanner()
	for _, line := range lines {
		sc.ScanLine(line)
	}
}
