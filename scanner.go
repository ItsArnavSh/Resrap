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
	word       TokenType = iota //Normal words
	character                   //'...'
	maybe                       //?
	oneormore                   //+
	anyno                       //*
	bracks                      //(...)
	option                      // |
	padding                     //just to account for my bad indexing skills
	regexrange                  //[...]

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
	case regexrange:
		return "regexrange [...]"
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
	_ = s.contentHandler(parentNode, content, false)

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
		case ch == '[':
			flush()
			j := i + 1
			depth := 1
			for j < len(content) && depth > 0 {
				if content[j] == '[' {
					depth++
				} else if content[j] == ']' {
					depth--
				}
				j++
			}
			if depth == 0 {
				tokens = append(tokens, Token{content[i:j], regexrange})
				i = j - 1
			} else {
				// unmatched '('
				tokens = append(tokens, Token{string(ch), word})
			}
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
func (s *Scanner) contentHandler(parentNode *SyntaxNode, content string, deep bool) *SyntaxNode {
	collapseNodeName := ":~exit~:" + uuid.NewString()
	collapseNode := s.synG.GetNode(collapseNodeName)
	bufferToken := parentNode
	bufferbufferToken := parentNode
	tokens := s.SeperateTokens(content)
	tokens = append(tokens, Token{"", padding})

	for i, token := range tokens {
		switch token.typ {
		case option:
			bufferToken.AddEdgeNext(&s.synG, collapseNode)
		case oneormore:
			if tokens[i-1].typ == bracks {
				continue
			}
			if tokens[i+1].typ == padding {
				bufferToken.AddEdgeDown(&s.synG, collapseNode)
			}
			bufferToken.AddEdgeDown(&s.synG, bufferToken)
		case anyno:
			if tokens[i-1].typ == bracks {
				continue
			}
			if tokens[i+1].typ == padding {
				bufferToken.AddEdgeDown(&s.synG, collapseNode)
				bufferbufferToken.AddEdgeDown(&s.synG, collapseNode)
			}
			bufferbufferToken.AddEdgeDown(&s.synG, collapseNode)
			bufferToken.AddEdgeDown(&s.synG, bufferToken)
		case maybe:
			if tokens[i-1].typ == bracks {
				continue
			}
			parentNode.AddEdgeDown(&s.synG, collapseNode)
			bufferToken.AddEdgeDown(&s.synG, collapseNode)
		case bracks:
			// Extract the content inside the brackets from the token data
			brackContent := token.data
			if len(brackContent) < 2 || brackContent[0] != '(' || brackContent[len(brackContent)-1] != ')' {
				// Invalid bracket token, skip it
				continue
			}

			// Extract content between the brackets
			substr := brackContent[1 : len(brackContent)-1]

			// Skip empty brackets
			if strings.TrimSpace(substr) == "" {
				continue
			}

			hollowNode := s.synG.GetNode(":~brac:~" + uuid.NewString()) //To replace the bracket
			bufferToken.AddEdgeNext(&s.synG, hollowNode)
			localCollapseNode := s.contentHandler(hollowNode, substr, true) //All nodes converge in this one for this brac
			bufferbufferToken = bufferToken
			bufferToken = localCollapseNode

			//Depending on what lies ahead modify these two nodes
			//Refer to diagrams..which I hope I make
			//Special case in case of brackets for the operations so will handle them right here
			if i+1 < len(tokens) {
				switch tokens[i+1].typ {
				case anyno:
					localCollapseNode.AddEdgeDown(&s.synG, hollowNode)
					hollowNode.AddEdgeDown(&s.synG, localCollapseNode)
				case oneormore:
					localCollapseNode.AddEdgeDown(&s.synG, hollowNode)
				case maybe:
					hollowNode.AddEdgeDown(&s.synG, localCollapseNode)
				default:
					continue
				}
			}

		case word, character, regexrange:
			newNode := s.synG.GetNode(token.data)
			if i > 0 && (tokens[i-1].typ != option) {
				bufferToken.AddEdgeNext(&s.synG, newNode)
			} else {
				parentNode.AddEdgeDown(&s.synG, newNode)
			}
			bufferbufferToken = bufferToken
			bufferToken = newNode
		case padding:
			if bufferToken != collapseNode && (deep || !strings.Contains(bufferToken.name, "~:brac:~")) {
				bufferToken.AddEdgeNext(&s.synG, collapseNode)
			}
		}
	}

	return collapseNode
}

func NewScanner() Scanner {
	return Scanner{
		synG: NewSyntaxGraph(),
	}
}
