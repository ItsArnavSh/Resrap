package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/google/uuid"
)

type Scanner struct {
	synG SyntaxGraph
}

type TokenType int8

const (
	word      TokenType = iota //Normal words
	character                  //'...'
	maybe                      //?
	oneormore                  //+
	anyno                      //*
	bracks                     //(...)
	option                     // |
	padding                    //just to account for my bad indexing skills

)

func (t TokenType) String() string {
	switch t {
	case word:
		return "word"
	case character:
		return "character"
	case maybe:
		return "maybe (?)"
	case oneormore:
		return "one or more (+)"
	case anyno:
		return "zero or more (*)"
	case bracks:
		return "brackets ( ... )"
	case option:
		return "option (|)"
	default:
		return fmt.Sprintf("unknown(%d)", t)
	}
}

type Token struct {
	data string
	typ  TokenType
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
	_ = s.contentHandler(parentNode, content)

	fmt.Println(heading)
}

func (s *Scanner) SeperateTokens(content string) []Token {
	var tokens []Token
	var buffer strings.Builder

	flush := func() {
		if buffer.Len() > 0 {
			tokens = append(tokens, Token{buffer.String(), word})
			buffer.Reset()
		}
	}

	for i := 0; i < len(content); i++ {
		ch := rune(content[i])

		// Handle // comments
		if ch == '/' && i+1 < len(content) && content[i+1] == '/' {
			flush()
			// Skip until newline
			for i < len(content) && content[i] != '\n' {
				i++
			}
			continue
		}

		switch {
		case unicode.IsSpace(ch):
			flush()

		case ch == '(':
			flush()
			j := i + 1
			depth := 1
			for j < len(content) && depth > 0 {
				if content[j] == '(' {
					depth++
				} else if content[j] == ')' {
					depth--
				}
				j++
			}
			if depth == 0 {
				tokens = append(tokens, Token{content[i:j], bracks})
				i = j - 1
			} else {
				// unmatched '('
				tokens = append(tokens, Token{string(ch), word})
			}

		case ch == '\'':
			flush()
			j := i + 1
			for j < len(content) && rune(content[j]) != '\'' {
				j++
			}
			if j < len(content) {
				tokens = append(tokens, Token{content[i : j+1], character})
				i = j
			} else {
				tokens = append(tokens, Token{string(ch), character})
			}

		case ch == '?':
			flush()
			tokens = append(tokens, Token{"?", maybe})

		case ch == '+':
			flush()
			tokens = append(tokens, Token{"+", oneormore})

		case ch == '*':
			flush()
			tokens = append(tokens, Token{"*", anyno})
		case ch == '|':
			flush()
			tokens = append(tokens, Token{"|", option})
		case ch == ';':
			flush()
			// skip raw semicolon

		default:
			buffer.WriteRune(ch)
		}
	}
	flush()

	return tokens
}
func (s *Scanner) contentHandler(parentNode *SyntaxNode, content string) *SyntaxNode {
	collapseNodeName := ":~exit~:" + uuid.NewString()
	collapseNode := s.synG.GetNode(collapseNodeName)
	bufferToken := parentNode

	tokens := s.SeperateTokens(content)
	tokens = append(tokens, Token{"", padding})

	for i, token := range tokens {
		switch token.typ {
		case option:
			parentNode.AddEdgeDown(bufferToken)
			bufferToken.AddEdgeDown(collapseNode)
			bufferToken = parentNode
		case oneormore:
			parentNode.AddEdgeDown(bufferToken)
			bufferToken.AddEdgeDown(collapseNode)
			bufferToken.AddEdgeDown(bufferToken)
			bufferToken = parentNode
		case anyno:
			parentNode.AddEdgeDown(collapseNode)
			bufferToken.AddEdgeDown(collapseNode)
			bufferToken.AddEdgeDown(bufferToken)
			bufferToken = parentNode
		case maybe:
			parentNode.AddEdgeDown(collapseNode)
			parentNode.AddEdgeDown(bufferToken)
			bufferToken.AddEdgeDown(collapseNode)
			bufferToken = parentNode
		case bracks:
			//find last index of )
			closeIndex := strings.LastIndex(content, ")")
			substr := content[i+1 : closeIndex]
			hollowNode := s.synG.GetNode(":~brac:~" + uuid.NewString()) //To replace the bracket
			bufferToken.AddEdgeNext(hollowNode)
			localCollapseNode := s.contentHandler(hollowNode, substr) //All nodes converge in this one for this brac
			bufferToken = hollowNode
			//Depending on what lies ahead modify these two nodes
			//Refer to diagrams..which I hope I make
			//Special case in case of brackets for the operations so will handle them right here
			switch tokens[i+1].typ {
			case anyno:
				localCollapseNode.AddEdgeDown(hollowNode)
				hollowNode.AddEdgeDown(collapseNode)
				i++
			case oneormore:
				localCollapseNode.AddEdgeDown(hollowNode)
				localCollapseNode.AddEdgeDown(collapseNode)
				i++
			case maybe:
				hollowNode.AddEdgeDown(collapseNode)
				localCollapseNode.AddEdgeDown(collapseNode)
				i++
			default:
				continue
			}

		case word, character:

			newNode := s.synG.GetNode(token.data)
			bufferToken.AddEdgeNext(newNode)
			bufferToken = newNode
		case padding:
			continue
		}
	}

	return collapseNode
}

func NewScanner() Scanner {
	return Scanner{
		synG: SyntaxGraph{},
	}
}
