package resrap

import (
	"fmt"
	"unicode/utf8"
)

// scanner will take all the components and convert them to a series of tokens

type scanner struct {
	Input  string
	pos    int // current byte offset
	width  int // width of last rune
	currR  rune
	lineno int
	tokens []token
}

func extracttokens(inp string) ([]token, []ScanError) {
	sc := scanner{}
	sc.Input = inp
	return sc.scan()
}

type ScanError struct {
	Pos int
	Msg string
}

// advance and return the next rune
func (s *scanner) next() rune {
	if s.pos >= len(s.Input) {
		s.width = 0
		s.currR = -1 // EOF
		return -1
	}
	r, w := utf8.DecodeRuneInString(s.Input[s.pos:])
	s.width = w
	s.pos += w
	s.currR = r
	return r
}

// look at the next rune without consuming it
func (s *scanner) peek() rune {
	r := s.next()
	s.backup()
	return r
}

// go back one rune
func (s *scanner) backup() {
	if s.width > 0 {
		s.pos -= s.width
	}
}

// return the current rune (last consumed by next)
func (s *scanner) curr() rune {
	return s.currR
}

func (s *scanner) scanDelimited(open, close rune, allowEscapes bool) (string, *ScanError) {
	start := s.pos
	var buf string

	for {
		r := s.next()
		if r == -1 {
			return "", &ScanError{start, fmt.Sprintf("unterminated %q", open)}
		}
		if r == close {
			return buf, nil
		}

		buf += string(r)
	}
}
func isAlpha(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func isIdentStart(r rune) bool {
	return isAlpha(r) || r == '_'
}

func isIdentPart(r rune) bool {
	return isAlpha(r) || isDigit(r) || r == '_'
}
func (s *scanner) scanIdentifier() string {
	buf := string(s.currR)

	// Keep reading while next runes are valid identifier parts
	for {
		r := s.peek()
		if !isIdentPart(r) {
			break
		}
		s.next()
		buf += string(s.currR)
	}
	return buf
}
func (s *scanner) scan() ([]token, []ScanError) {
	var errs []ScanError
	for s.next() != -1 {
		switch s.currR {
		case '+':
			s.tokens = append(s.tokens, token{0, oneormore, ""})
		case '*':
			s.tokens = append(s.tokens, token{0, anyno, ""})
		case '^':
			s.tokens = append(s.tokens, token{0, infinite, ""})
		case '?':
			s.tokens = append(s.tokens, token{0, maybe, ""})
		case '|':
			s.tokens = append(s.tokens, token{0, option, ""})
		case ';':
			s.tokens = append(s.tokens, token{0, padding, ""})
		case '(':
			s.tokens = append(s.tokens, token{0, bracopen, ""})
		case ')':
			s.tokens = append(s.tokens, token{0, bracclose, ""})
		case ':':
			s.tokens = append(s.tokens, token{0, colon, ""})
		case '\'':
			val, err := s.scanDelimited('\'', '\'', false)
			if err != nil {
				errs = append(errs, *err)
			} else {
				s.tokens = append(s.tokens, token{0, character, val})
			}

		case '<':
			val, err := s.scanDelimited('<', '>', false)
			if err != nil {
				errs = append(errs, *err)
			} else {
				s.tokens = append(s.tokens, token{0, probability, val})
			}

		case '[':
			val, err := s.scanDelimited('[', ']', false)
			if err != nil {
				errs = append(errs, *err)
			} else {
				s.tokens = append(s.tokens, token{0, regex, val})
			}
		default:
			if isIdentStart(s.currR) {
				buff := s.scanIdentifier()
				if buff != "" {
					s.tokens = append(s.tokens, token{0, identifier, buff})
				}

			}
		}
	}
	return s.tokens, errs
}

func (s *scanner) PrintTokens() {
	for _, token := range s.tokens {
		fmt.Println(token.text)
	}
}
