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
	fmt.Printf("There are %d nodes \n", len(sc.synG.nodeRef))

	sc.synG.Clean()
	sc.synG.ExportDOTFile("test.dot")
	//sc.synG.TraverseInteractive("statement")
	sc.synG.RandomWalker("a", 100)

}
