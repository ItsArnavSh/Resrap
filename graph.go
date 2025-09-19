package main

type SyntaxNode struct {
	down []*SyntaxNode
	next []*SyntaxNode
	name string
}

func (s *SyntaxNode) AddEdgeNext(node *SyntaxNode) {
	s.next = append(s.next, node)
}
func (s *SyntaxNode) AddEdgeDown(node *SyntaxNode) {
	s.down = append(s.down, node)
}

type SyntaxGraph struct {
	nodeRef   map[string]*SyntaxNode
	pointedTo map[string]bool
}

func (s *SyntaxGraph) GetNode(name string) *SyntaxNode {
	if s.nodeRef[name] != nil {
		return s.nodeRef[name]
	}
	return &SyntaxNode{nil, nil, name}
}
