package resrap

import (
	"fmt"
	"math"
	"slices"
	"strconv"
)

type parser struct {
	func_ptr     uint32
	print_ptr    uint32
	name_map     map[string]uint32  //Maps func names to their ids
	rev_name_map map[uint32]string  //Help in debugging
	def_check    map[uint32]bool    //To check if a function exists
	charmap      map[uint32]string  //To store the print values corresponding to ids
	inter_rep    map[uint32][]token //Intermediate Representation
	tokens       []token
	errors       []error
	index        int
	graph        syntaxGraph
	regexhandler regexer
}

func new_parser() parser {
	return parser{
		print_ptr: math.MaxInt32, // grows downward
		func_ptr:  1000,          // IDs below 1000 reserved for core graph nodes

		// allocate maps to avoid nil panics
		name_map:     make(map[string]uint32),
		def_check:    make(map[uint32]bool),
		charmap:      make(map[uint32]string),
		inter_rep:    make(map[uint32][]token),
		rev_name_map: make(map[uint32]string),
		tokens:       []token{},
		errors:       []error{},
		graph:        newSyntaxGraph(),
		regexhandler: newRegexer(),
	}
}

func (i *parser) get_print_ptr() uint32 {
	i.print_ptr--
	return i.print_ptr
}
func (i *parser) get_func_ptr() uint32 {
	i.func_ptr++
	return i.func_ptr
}
func (i *parser) curr() token {
	return i.tokens[i.index]
}
func (i *parser) match(word tokenType, expec []tokenType) bool {
	return slices.Contains(expec, word)
}
func (i *parser) expect(expected []tokenType, errmsg string) {
	if !i.match(i.curr().typ, expected) {
		i.errors = append(i.errors, fmt.Errorf("%s", errmsg))
	}
	i.index++

}
func (i *parser) get_index(name string) uint32 {
	value, ok := i.name_map[name]
	if ok {
		return value
	}
	value = i.get_func_ptr()
	i.def_check[value] = false //tb set back to true in the subject, else remain false
	i.name_map[name] = value
	i.rev_name_map[value] = name
	return value
}
func (i *parser) parse_grammar() {
	for i.index < len(i.tokens) {
		i.parse_subject()
	}
}
func (i *parser) parse_subject() {
	subject := i.curr()

	i.expect([]tokenType{identifier}, "Expected Subject at start of statement")
	i.expect([]tokenType{colon}, "Expected Colon after Subject")
	id := i.get_index(subject.text)
	if i.def_check[id] { //If map is already set to true
		i.errors = append(i.errors, fmt.Errorf("Multiple definitions for %s", subject.text))
	}

	i.def_check[id] = true
	startnode := i.graph.GetNode(uint32(start), start)
	startnode.AddEdgeNext(&i.graph, i.graph.GetNode(id, header), 1)
	//Send here only if current is col else crash code
	if i.match(i.tokens[i.index-1].typ, []tokenType{colon}) {
		i.parse_rules(id, false)
	} else {
		return
	}

}

func (i *parser) parse_rules(root uint32, isDeep bool) (*syntaxNode, *syntaxNode) {

	rootnode := i.graph.GetNode(root, idk)
	bufferNode := rootnode
	endNode := i.graph.GetNode(uint32(end), end)
	var startBuffer *syntaxNode
	if isDeep { //Means called from a backet so a pseudo end branch
		endNode = i.graph.GetNode(i.get_func_ptr(), end)
	}
	for {
		switch i.curr().typ {
		case identifier:
			//Means its a reference to a different Subject(presumably)
			pointerid := i.get_index(i.tokens[i.index].text)
			node := i.graph.GetNode(i.get_func_ptr(), pointer)
			node.pointer = pointerid
			bufferNode.AddEdgeNext(&i.graph, node, i.get_probability())
			jumpNode := i.graph.GetNode(i.get_func_ptr(), jump)
			node.AddEdgeNext(&i.graph, jumpNode, 1)
			startBuffer = bufferNode
			bufferNode = jumpNode
			//Basically just add the word and next to it its jump node
			// So when generating, the control will pass to the node at the location and save the exit in a stack
			// Then when it reached its local collapse node, then the control will automatically come back to default
		case character, regex:
			index := i.get_print_ptr()
			i.charmap[index] = i.tokens[i.index].text
			var leafnode *syntaxNode
			if i.tokens[i.index].typ == character {
				leafnode = i.graph.GetNode(index, ch)
			} else {
				leafnode = i.graph.GetNode(index, rx)
				i.regexhandler.CacheRegex(i.curr().text)
			}
			bufferNode.AddEdgeNext(&i.graph, leafnode, i.get_probability())
			jumpNode := i.graph.GetNode(i.get_func_ptr(), jump)
			leafnode.AddEdgeNext(&i.graph, jumpNode, 1)
			startBuffer = bufferNode
			bufferNode = jumpNode
		case colon:
			//Colon is not allowed here
			return nil, nil
		case maybe:
			startBuffer.AddEdgeNext(&i.graph, bufferNode, 1-i.get_probability()) //An option to skip to the end
		case oneormore:
			bufferNode.AddEdgeNext(&i.graph, startBuffer, i.get_probability()) //An option to go to the start
		case anyno:
			startBuffer.AddEdgeNext(&i.graph, bufferNode, 1-i.get_probability()) //Well both of them combined
			bufferNode.AddEdgeNext(&i.graph, startBuffer, i.get_probability())
		case option:
			//in Case of an option, no need to really do anything, just set the buffer settings back to the parent
			bufferNode.AddEdgeNext(&i.graph, endNode, i.get_probability())
			bufferNode = rootnode
			startBuffer = nil
		case padding:
			bufferNode.AddEdgeNext(&i.graph, endNode, 1)
			if isDeep {
				i.errors = append(i.errors, fmt.Errorf("Stray '('"))
			}
			i.index++
			return nil, nil //End of this statement
		case bracopen:
			i.index++
			startBuffer, bufferNode = i.parse_rules(bufferNode.id, true)
		case bracclose:
			if isDeep {
				bufferNode.AddEdgeNext(&i.graph, endNode, 1)
				return rootnode, endNode
			}
			i.errors = append(i.errors, fmt.Errorf("Stray ')' found"))
		case infinite:
			//Now at the end it will loop back to this case
			endNode.AddEdgeNext(&i.graph, startBuffer, 1)
		default:
			continue
		}
		i.index++
	}
}
func (i *parser) get_probability() float32 {
	i.index++
	if i.tokens[i.index].typ == probability {
		num := i.tokens[i.index].text
		numf, err := strconv.ParseFloat(num, 32)
		if err != nil {
			i.index--
			i.errors = append(i.errors, err)
			return 0
		}
		if numf < 0 {
			i.errors = append(i.errors, fmt.Errorf("Negative Probability Found"))
			return 0
		}
		return float32(numf)
	}
	i.index-- //Reverting
	return 0.5
}
func (p *parser) ValidateGraph() []error {
	var errors []error
	for key, val := range p.def_check {
		if !val {
			errors = append(errors, fmt.Errorf("Definition of '%s' not found", p.rev_name_map[key]))
		}
	}
	return errors
}
