package main

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type Scanner struct {
	synG SyntaxGraph
}

func (s *Scanner) ScanLine(line string) {
	// Remove newlines and trim spaces
	line = strings.TrimSpace(line)

	parts := strings.SplitN(line, ":", 2)
	if len(parts) < 2 {
		return // no colon, nothing to scan
	}

	// Take the part before colon and split by whitespace
	fields := strings.Fields(parts[0])
	if len(fields) == 0 {
		return
	}

	heading := fields[len(fields)-1] // FIRST word before :

	content := parts[1]
	parentNode := s.synG.GetNode(heading)
	s.contentHandler(parentNode, content)

	fmt.Println(heading)
}
func (s *Scanner) contentHandler(parentNode *SyntaxNode, content string) {
	var buffer strings.Builder
	for i := 0; i < len(content); i++ {
		switch content[i] {
		case '|':
			node := s.synG.GetNode(strings.TrimSpace(buffer.String()))
			buffer.Reset()
			parentNode.AddEdge(node)
		case '+':
			node := s.synG.GetNode(strings.TrimSpace(buffer.String()))
			buffer.Reset()
			parentNode.AddEdge(node)
			node.AddEdge(s.synG.GetNode(":~exit~:"))
			node.AddEdge(node)
		case '*':

			node := s.synG.GetNode(strings.TrimSpace(buffer.String()))
			buffer.Reset()
			parentNode.AddEdge(s.synG.GetNode(":~exit~:"))
			node.AddEdge(s.synG.GetNode(":~exit~:"))
			node.AddEdge(node)
		case '(':
			//find last index of )
			closeIndex := strings.LastIndex(content, ")")
			//String between i and closeindex
			substr := content[i+1 : closeIndex]
			hollowNode := s.synG.GetNode(uuid.NewString()) //To replace the bracket

			s.contentHandler(hollowNode, substr)
			i = closeIndex + 1
		case '\'':
			//Find the next ' and save the name of the entry here
		case ' ':
			//Also find a way to append normal stuff
		default:
			buffer.WriteByte(content[i])
		}

	}
}
func NewScanner() Scanner {
	return Scanner{
		synG: SyntaxGraph{},
	}
}
