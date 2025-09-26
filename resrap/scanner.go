package resrap

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/google/uuid"
)

type scanner struct {
	synG syntaxGraph
}

type TokenType int8

const (
	word        TokenType = iota //Normal words
	character                    //'...'
	maybe                        //?
	oneormore                    //+
	anyno                        //*
	bracks                       //(...)
	option                       // |
	padding                      //just to account for my bad indexing skills
	regexrange                   //[...]
	infinite                     //^
	probability                  //<...>

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
	case infinite:
		return "Infinite ^"
	case probability:
		return "Probability <...>"
	default:
		return fmt.Sprintf("unknown(%d)", t)
	}
}

type Token struct {
	data string
	typ  TokenType
}

func (s *scanner) ScanLine(line string) {
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
	s.addStatement(heading, content, false)

}

func (s *scanner) SeperateTokens(content string) []Token {
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
		case ch == '<':
			flush()
			j := i + 1
			depth := 1
			for j < len(content) && depth > 0 {
				if content[j] == '<' {
					depth++
				} else if content[j] == '>' {
					depth--
				}
				j++
			}
			if depth == 0 {
				tokens = append(tokens, Token{content[i : j-1], probability})
				i = j - 1
			} else {
				// unmatched '<'
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
		case ch == '^':
			flush()
			tokens = append(tokens, Token{"^", infinite})
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

func NewScanner() scanner {
	return scanner{
		synG: NewSyntaxGraph(),
	}
}
func (s *scanner) addStatement(heading, content string, depth bool) (*syntaxNode, *syntaxNode) {
	//The definitive end block
	endNode := s.synG.GetNode("~:end:~")
	if depth {

		endNode = s.synG.GetNode(generator("end")) //One endnode per block
	}
	parentNode := s.synG.GetNode(heading)
	//Converting the content here to a block
	// Every block is a self containted block with references to other headers and directional nodes for | + * etc
	tokens := s.SeperateTokens(content)
	tokens = append(tokens, Token{"", padding})
	bufferNode := parentNode
	var startBuffer *syntaxNode //Stores the starts
	prob := func(index int) float32 {
		index++
		//If the function can find a probability next to it, cool, return its value
		// If its not, just return a .50 it will be normalized later on anyways
		if tokens[index].typ != probability {

			return 0.5
		}
		num := tokens[index].data[1:]
		numf, err := strconv.ParseFloat(num, 32)
		if err != nil {
			return 0
		}
		return float32(numf)
	}
	for i, token := range tokens {
		switch token.typ {
		case word:
			//Just leave it as is, we will assume its definition exists, here we will simply need to generate an exit case

			wordNode := s.synG.GetNode(generator("{" + token.data + "}"))
			bufferNode.AddEdgeNext(&s.synG, wordNode, prob(i))
			jumpNode := s.synG.GetNode(generator("jmp"))
			wordNode.AddEdgeNext(&s.synG, jumpNode, 1) //Prob does not matter
			startBuffer = bufferNode
			bufferNode = jumpNode
			//Basically just add the word and next to it its jump node
			// So when generating, the control will pass to the node at the location and save the exit in a stack
			// Then when it reached its local collapse node, then the control will automatically come back to default
		case character, regexrange:
			//True leaf nodes just add simply to next and update bufferNode
			leafNode := s.synG.GetNode(generator("{" + token.data + "}"))
			bufferNode.AddEdgeNext(&s.synG, leafNode, prob(i))
			jumpNode := s.synG.GetNode(generator("jmp"))
			leafNode.AddEdgeNext(&s.synG, jumpNode, 1)
			startBuffer = bufferNode
			bufferNode = jumpNode
		case maybe:
			startBuffer.AddEdgeNext(&s.synG, bufferNode, 1-prob(i)) //An option to skip to the end
		case oneormore:
			bufferNode.AddEdgeNext(&s.synG, startBuffer, prob(i)) //An option to go to the start
		case anyno:
			startBuffer.AddEdgeNext(&s.synG, bufferNode, 1-prob(i)) //Well both of them combined
			bufferNode.AddEdgeNext(&s.synG, startBuffer, prob(i))
		case option:
			//in Case of an option, no need to really do anything, just set the buffer settings back to the parent
			bufferNode.AddEdgeNext(&s.synG, endNode, prob(i))
			bufferNode = parentNode
			startBuffer = nil
		case padding:
			bufferNode.AddEdgeNext(&s.synG, endNode, 1)
		case bracks:
			//Get the final bracket
			startBuffer, bufferNode = s.addStatement(bufferNode.name, token.data[1:len(token.data)-1], true)
		case infinite:
			//Now at the end it will loop back to this case
			endNode.AddEdgeNext(&s.synG, startBuffer, prob(i))
		default:
			continue
		}
	}
	return parentNode, endNode
}
func generator(typ string) string {
	return "~:" + typ + ":~" + uuid.NewString()
}
