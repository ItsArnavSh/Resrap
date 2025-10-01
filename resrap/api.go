package resrap

// Resrap is the main accesspoint for singlethreaded uses
// Pretty Much Collection of graphs which can be generated using parsing grammar
type Resrap struct {
	languageGraph map[string]lang
}

// NewResrap creates and returns a new Resrap instance.
// The returned instance starts with no loaded grammars.
func NewResrap() *Resrap {
	return &Resrap{
		languageGraph: make(map[string]lang),
	}
}

// ParseGrammar parses a grammar string and stores it under the given name.
// name: a unique identifier for this grammar (e.g., "C"), should be in ABNF format(Check osdc/resrap for more info on that).
// Returns error generated while parsing
func (r *Resrap) ParseGrammar(name, grammar string) error {
	lang := newLang()
	err := lang.ParserString(grammar)
	r.languageGraph[name] = lang
	r.languageGraph[name].graph.Normalize()
	return err
}

// ParseGrammarFile parses a grammar from a file and stores it under the given name.
// name: a unique identifier for this grammar (e.g., "C"), should be in ABNF format(Check osdc/resrap for more info on that).
// location: path to the grammar file.
// Returns error generated while parsing
func (r *Resrap) ParseGrammarFile(name, location string) error {
	lang := newLang()
	err := lang.ParserFile(location)
	r.languageGraph[name] = lang
	r.languageGraph[name].graph.Normalize()

	return err
}

// GenerateRandom generates content from the grammar identified by 'name'.
// starting_node: the starting heading in the grammar for generation.
// Returns a string containing the generated content.
// The generation is non-deterministic (random).
func (r *Resrap) GenerateRandom(name, starting_node string, tokens int) string {
	prng := newPRNG(0)
	return r.languageGraph[name].graph.GraphWalk(&prng, starting_node, tokens)
}

// GenerateWithSeeded generates content from the grammar identified by 'name'.
// starting_node: the starting symbol in the grammar for generation.
// seed: a numeric seed to make generation deterministic.
// Returns a string containing the generated content.
func (r *Resrap) GenerateWithSeeded(name, starting_node string, seed uint64, tokens int) string {
	prng := newPRNG(seed)
	return r.languageGraph[name].graph.GraphWalk(&prng, starting_node, tokens)
}
