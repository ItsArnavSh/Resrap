package resrap

import (
	"strings"
	"unicode"
)

type regexer struct {
	cached_rex map[string]cacheRexState
}

func newRegexer() regexer {
	return regexer{cached_rex: make(map[string]cacheRexState)}
}

type cacheRexState struct {
	cumu_freq []float32
	options   []rune
}

func (r *regexer) GenerateString(regex string, prn *prng) string {
	size := prn.RandomInt(3, 7) // generate a size between 3 and 7
	var sb strings.Builder

	for i := 0; i < size; i++ {
		x := prn.Random() // random float 0-1
		idx := closestIndex(r.cached_rex[regex].cumu_freq, float32(x))
		sb.WriteRune(r.cached_rex[regex].options[idx]) // use WriteRune instead of WriteByte
	}

	return sb.String()
}
func (r *regexer) CacheRegex(regex string) {
	tokens := r.ExpandClass(regex)
	var biasarr []float32
	var sum float32
	for _, token := range tokens {
		bias := float32(r.Bias(token))
		biasarr = append(biasarr, bias)
		sum += bias
	}
	for i := range biasarr {
		biasarr[i] /= sum
	}

	cdf := make([]float32, len(biasarr))
	cum := float32(0)
	for i, w := range biasarr {
		cum += w
		cdf[i] = cum
	}

	r.cached_rex[regex] = cacheRexState{cumu_freq: cdf, options: tokens}
}

// closestIndex finds the first index in cdf where cdf[idx] >= x
func closestIndex(cdf []float32, x float32) int {
	for i, val := range cdf {
		if x <= val {
			return i
		}
	}
	return len(cdf) - 1 // fallback
}
func (r *regexer) ExpandClass(class string) []rune {
	var chars []rune
	runes := []rune(class)

	for i := 0; i < len(runes); i++ {
		if i+2 < len(runes) && runes[i+1] == '-' {
			// range a-z
			for c := runes[i]; c <= runes[i+2]; c++ {
				chars = append(chars, c)
			}
			i += 2
		} else {
			chars = append(chars, runes[i])
		}
	}
	return chars
}

// Bias returns an integer weight for a rune based on usage frequency
func (k *regexer) Bias(r rune) int {
	// Lowercase letters: frequency in English words (roughly)
	switch unicode.ToLower(r) {
	case 'e':
		return 12
	case 'a', 'i', 'o':
		return 9
	case 'n', 'r', 't', 's', 'l':
		return 6
	case 'c', 'd', 'm', 'u', 'p', 'b', 'g':
		return 4
	case 'f', 'h', 'v', 'k', 'w', 'y':
		return 3
	case 'j', 'x', 'q', 'z':
		return 1
	}

	// Uppercase letters: slightly less likely than lowercase
	if unicode.IsUpper(r) {
		return k.Bias(unicode.ToLower(r)) / 2
	}

	// Digits: moderately common
	if unicode.IsDigit(r) {
		return 3
	}

	// Underscore: quite common in identifiers
	if r == '_' {
		return 5
	}

	// Everything else: low probability
	return 1
}
