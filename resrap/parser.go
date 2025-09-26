package resrap

import "strings"

type lang struct {
	graph *syntaxGraph
	nodes int
}

func newLang() lang {
	return lang{}
}
func (l *lang) GetGraph() *syntaxGraph {
	return l.graph
}
func (l *lang) ParserFile(filename string) {
	lines, err := ParseFile(filename)
	if err != nil {
		return
	}
	sc := NewScanner()
	for _, line := range lines {
		sc.ScanLine(line)
	}
	l.nodes = len(sc.synG.nodeRef)
	l.graph = &sc.synG
}

func (l *lang) ParserString(data string) {
	lines := strings.Split(data, "\n")
	sc := NewScanner()
	for _, line := range lines {
		sc.ScanLine(line)
	}
	l.nodes = len(sc.synG.nodeRef)
	l.graph = &sc.synG
}
