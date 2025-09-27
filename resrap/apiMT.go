package resrap

type codeGenReq struct {
	name      string
	startnode string
	seed      uint64
	tokens    int
	id        string
}
type CodeGenRes struct {
	Code string
	Id   string
}

type ResrapMT struct {
	languageGraph map[string]lang
	poolsize      int //Number of threads in the pool
	waitqueuesize int
	pendingjobs   chan codeGenReq
	CodeChannel   chan CodeGenRes
}

// Since Graphs once made are treated as immutable, only traversed through, we will not add locks
// NewResrapMT creates and returns a new Resrap MultiThreaded instance.
// The returned instance starts with no loaded grammars.
func NewResrapMT(poolsize, waitqueuesize int) *ResrapMT {
	return &ResrapMT{
		languageGraph: make(map[string]lang),
		waitqueuesize: waitqueuesize,
		poolsize:      poolsize,
		pendingjobs:   make(chan codeGenReq, waitqueuesize),
		CodeChannel:   make(chan CodeGenRes),
	}
}

// ParseGrammar parses a grammar string and stores it under the given name.
// name: a unique identifier for this grammar (e.g., "C").
// grammar: the grammar definition as a string.
func (r *ResrapMT) ParseGrammar(name, grammar string) {
	lang := newLang()
	lang.ParserString(grammar)
	r.languageGraph[name] = lang
	r.languageGraph[name].graph.Normalize()
}

// ParseGrammarFile parses a grammar from a file and stores it under the given name.
// name: a unique identifier for this grammar (e.g., "C").
// location: path to the grammar file.
func (r *ResrapMT) ParseGrammarFile(name, location string) {
	lang := newLang()
	lang.ParserFile(location)
	r.languageGraph[name] = lang
	r.languageGraph[name].graph.Normalize()
}

// GenerateRandom generates content from the grammar identified by 'name'.
// starting_node: the starting symbol in the grammar for generation.
// Returns a string containing the generated content.
// The generation is non-deterministic (random).
func (r *ResrapMT) GenerateRandom(id, name, starting_node string, tokens int) {
	req := codeGenReq{name: name, startnode: starting_node, tokens: tokens, seed: 0, id: id}
	r.pendingjobs <- req
}

// GenerateWithSeeded generates content from the grammar identified by 'name'.
// starting_node: the starting symbol in the grammar for generation.
// seed: a numeric seed to make generation deterministic.
// Returns a string containing the generated content.
func (r *ResrapMT) GenerateWithSeeded(id, name, starting_node string, seed uint64, tokens int) {
	req := codeGenReq{name: name, startnode: starting_node, tokens: tokens, seed: seed, id: id}
	r.pendingjobs <- req
}
func (r *ResrapMT) mtparser() {
	for job := range r.pendingjobs {
		prng := newPRNG(job.seed)
		code := r.languageGraph[job.name].graph.GraphWalk(&prng, job.startnode, job.tokens)
		r.CodeChannel <- CodeGenRes{Code: code, Id: job.id}
	}
}

func (r *ResrapMT) StartResrap() {

	for _ = range r.poolsize {
		go r.mtparser()
	}
}
