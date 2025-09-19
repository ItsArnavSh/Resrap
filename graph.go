package main

type SyntaxNode struct {
	next []*SyntaxNode
	name string
}

func (s *SyntaxNode) AddEdge(node *SyntaxNode) {
	s.next = append(s.next, node)
}

type SyntaxGraph struct {
	nodeRef   map[string]*SyntaxNode
	pointedTo map[string]bool
}

func (s *SyntaxGraph) GetNode(name string) *SyntaxNode {
	if s.nodeRef[name] != nil {
		return s.nodeRef[name]
	}
	return &SyntaxNode{nil, name}
}
