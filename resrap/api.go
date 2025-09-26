package resrap

// The Resrap Object allows for multiple graphs
type Resrap struct {
	languageGraph map[string]lang
}

func (r *Resrap) ParseGrammar(name, grammar string) {
	lang := newLang()
	lang.ParserString(grammar)
	r.languageGraph[name] = lang
}
func (r *Resrap) ParseGrammarFile(name, location string) {
	lang := newLang()
	lang.ParserFile(location)
	r.languageGraph[name] = lang
}
func (r *Resrap) GenerateRandom(name, starting_node string) string {
	prng := newPRNG(0)
	return r.languageGraph[name].graph.GraphWalk(&prng, starting_node)
}
func (r *Resrap) GenerateWithSeeded(name, starting_node string, seed uint64) string {
	prng := newPRNG(seed)
	return r.languageGraph[name].graph.GraphWalk(&prng, starting_node)
}
func NewResrap() *Resrap {
	return &Resrap{
		languageGraph: make(map[string]lang),
	}
}
