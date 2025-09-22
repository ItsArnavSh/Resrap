package main

import (
	"fmt"
	"strings"

	"github.com/golang-collections/collections/stack"
)

type SyntaxNode struct {
	next []*SyntaxNode
	name string
}

func (s *SyntaxNode) AddEdgeNext(g *SyntaxGraph, node *SyntaxNode) {
	s.next = append(s.next, node)
	g.nodeRef[node.name] = node

}

type SyntaxGraph struct {
	nodeRef map[string]*SyntaxNode
	prng    PRNG
}

func (s *SyntaxGraph) GetNode(name string) *SyntaxNode {
	if s.nodeRef[name] != nil {
		return s.nodeRef[name]
	}
	newNode := &SyntaxNode{nil, name}
	s.nodeRef[name] = newNode
	return newNode
}

func NewSyntaxGraph() SyntaxGraph {
	return SyntaxGraph{
		nodeRef: make(map[string]*SyntaxNode),
	}
}

// RandomWalker does the walk and returns the generated string
func (s *SyntaxGraph) RandomWalker(prng *PRNG, start string, no int32) string {
	var result strings.Builder
	jumpStack := stack.New()
	startingNode := s.nodeRef[start]
	if startingNode == nil {
		return ""
	}
	current := startingNode
	visited := int32(0)
	for visited < no && current != nil {
		// Process logic only if name starts with ' or [
		if len(current.name) > 0 {
			if strings.HasPrefix(current.name, "~:{'") {
				// Extract content between quotes and handle escape sequences
				content := current.name[4:strings.LastIndex(current.name, "'")]
				unescaped := unescapeString(content)
				result.WriteString(unescaped)
			} else if strings.HasPrefix(current.name, "~:{[") {
				chars, err := parseCharClass(prng, current.name[3:1+strings.LastIndex(current.name, "]")])
				if err == nil && len(chars) > 0 {
					result.WriteString(chars[prng.RandomInt(0, len(chars))])
				}
			} else if strings.HasPrefix(current.name, "~:{") {
				name := current.name[3:strings.Index(current.name, "}")]
				jumpStack.Push(current.next[0].name)
				current = s.GetNode(name)
				continue // Skip the normal next node selection
			} else if current.name == "~:end:~" {
				nameInt := jumpStack.Pop()
				name, ok := nameInt.(string)
				if !ok {
					break
				}
				current = s.GetNode(name)
				continue // Skip the normal next node selection
			}
		}
		visited++
		// move to next (randomly selected if multiple options)
		if len(current.next) > 0 {
			current = current.next[prng.RandomInt(0, len(current.next))]
		} else {
			current = nil
		}
	}
	return result.String()
}

// Helper function to handle escape sequences
func unescapeString(s string) string {
	result := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '\\' && i+1 < len(s) {
			switch s[i+1] {
			case 'n':
				result = append(result, '\n')
				i++ // skip the next character
			case 't':
				result = append(result, '\t')
				i++
			case 'r':
				result = append(result, '\r')
				i++
			case '\\':
				result = append(result, '\\')
				i++
			case '\'':
				result = append(result, '\'')
				i++
			case '"':
				result = append(result, '"')
				i++
			default:
				// If it's not a recognized escape sequence, keep both characters
				result = append(result, s[i])
			}
		} else {
			result = append(result, s[i])
		}
	}
	return string(result)
}

// Predefined word list for alphanumeric patterns
var wordList = []string{
	"hello", "world", "test", "code", "function", "variable", "class", "method",
	"data", "user", "admin", "login", "password", "email", "server", "client",
	"database", "table", "query", "result", "error", "success", "failure", "debug",
}

func parseCharClass(prng *PRNG, charClass string) ([]string, error) {
	if len(charClass) < 2 || charClass[0] != '[' || charClass[len(charClass)-1] != ']' {
		return nil, fmt.Errorf("invalid format: %s", charClass)
	}

	content := charClass[1 : len(charClass)-1]

	// Check for common regex patterns
	switch content {
	case "0-9", "\\d":
		// Random number string
		numbers := []string{"42", "123", "7", "999", "256", "1024", "88", "13", "77", "101"}
		idx := prng.RandomInt(0, len(numbers))
		return []string{numbers[idx]}, nil

	case "a-zA-Z0-9", "\\w", "a-zA-Z", "A-Z", "a-z":
		// Random word from predefined list
		idx := prng.RandomInt(0, len(wordList))
		return []string{wordList[idx]}, nil
	}

	// Fall back to original character range expansion
	var chars []string
	runes := []rune(content)
	for i := 0; i < len(runes); i++ {
		if i+2 < len(runes) && runes[i+1] == '-' { // range like a-z
			start, end := runes[i], runes[i+2]
			if start > end {
				return nil, fmt.Errorf("invalid range %c-%c", start, end)
			}
			for r := start; r <= end; r++ {
				chars = append(chars, string(r))
			}
			i += 2 // skip past range
		} else {
			chars = append(chars, string(runes[i]))
		}
	}

	// Pick one random char from expanded set
	if len(chars) == 0 {
		return nil, fmt.Errorf("empty char class: %s", charClass)
	}
	idx := prng.RandomInt(0, len(chars))
	return []string{chars[idx]}, nil
}
