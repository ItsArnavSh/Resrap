package resrap

import (
	"fmt"
	"strings"
)

// PrintGraph prints the entire syntax graph structure
func (s *syntaxGraph) PrintGraph() {
	fmt.Println("=== SYNTAX GRAPH ===")
	fmt.Println("\n--- Named Nodes ---")
	for name, id := range s.namemap {
		fmt.Printf("%s -> Node ID: %d\n", name, id)
	}

	fmt.Println("\n--- Character/Regex Map ---")
	for id, content := range s.charmap {
		fmt.Printf("Node %d: %q\n", id, content)
	}

	fmt.Println("\n--- Node Structure ---")
	// Sort node IDs for consistent output
	var nodeIDs []uint32
	for id := range s.nodeRef {
		nodeIDs = append(nodeIDs, id)
	}

	for _, id := range nodeIDs {
		node := s.nodeRef[id]
		s.printNode(node)
	}
}

// printNode prints a single node with all its details
func (s *syntaxGraph) printNode(node *syntaxNode) {
	if node == nil {
		return
	}

	fmt.Printf("\n┌─ Node ID: %d\n", node.id)
	fmt.Printf("│  Type: %s\n", nodeTypeToString(node.typ))

	if node.pointer != 0 {
		fmt.Printf("│  Pointer: -> %d\n", node.pointer)
	}

	// Find if this node has a name
	nodeName := ""
	for name, id := range s.namemap {
		if id == node.id {
			nodeName = name
			break
		}
	}
	if nodeName != "" {
		fmt.Printf("│  Name: %s\n", nodeName)
	}

	// Print content if it's a char or regex node
	if content, exists := s.charmap[node.id]; exists {
		fmt.Printf("│  Content: %q\n", content)
	}

	// Print edges
	if len(node.next) > 0 {
		fmt.Println("│  Edges:")
		for i, next := range node.next {
			prob := next.probability
			cf := ""
			if i < len(node.cf) {
				cf = fmt.Sprintf(" [CF: %.3f]", node.cf[i])
			}
			fmt.Printf("│    %d. -> Node %d (prob: %.3f)%s\n",
				i+1, next.node.id, prob, cf)
		}
	} else {
		fmt.Println("│  Edges: (none)")
	}
	fmt.Println("└─")
}

// PrintPath prints a specific path through the graph
func (s *syntaxGraph) PrintPath(startName string, maxSteps int) {
	startNode := s.nodeRef[s.namemap[startName]]
	if startNode == nil {
		fmt.Printf("Start node %q not found\n", startName)
		return
	}

	fmt.Printf("\n=== PATH FROM %s ===\n", startName)
	current := startNode
	step := 0

	for current != nil && step < maxSteps {
		fmt.Printf("\nStep %d: Node %d (%s)", step, current.id, nodeTypeToString(current.typ))

		if content, exists := s.charmap[current.id]; exists {
			fmt.Printf(" -> %q", content)
		}

		if current.pointer != 0 {
			fmt.Printf(" [points to %d]", current.pointer)
		}

		fmt.Println()

		if len(current.next) > 0 {
			fmt.Printf("  Options: ")
			for i, next := range current.next {
				if i > 0 {
					fmt.Print(", ")
				}
				fmt.Printf("Node %d (%.1f%%)", next.node.id, next.probability*100)
			}
			fmt.Println()
			current = current.next[0].node // Follow first option
		} else {
			current = nil
		}
		step++
	}
}

// PrintDOT generates a Graphviz DOT format representation
func (s *syntaxGraph) PrintDOT() string {
	var sb strings.Builder
	sb.WriteString("digraph SyntaxGraph {\n")
	sb.WriteString("  rankdir=LR;\n")
	sb.WriteString("  node [shape=box];\n\n")

	for id, node := range s.nodeRef {
		label := fmt.Sprintf("ID:%d\\n%s", id, nodeTypeToString(node.typ))

		// Add name if exists
		for name, nid := range s.namemap {
			if nid == id {
				label = fmt.Sprintf("%s\\n%s", name, label)
				break
			}
		}

		// Add content if exists
		if content, exists := s.charmap[id]; exists {
			// Escape quotes for DOT format
			escaped := strings.ReplaceAll(content, "\"", "\\\"")
			escaped = strings.ReplaceAll(escaped, "'", "\\'")
			label += fmt.Sprintf("\\n'%s'", escaped)
		}

		if node.pointer != 0 {
			label += fmt.Sprintf("\\nptr→%d", node.pointer)
		}

		sb.WriteString(fmt.Sprintf("  n%d [label=\"%s\"];\n", id, label))
	}

	sb.WriteString("\n")

	for _, node := range s.nodeRef {
		for _, next := range node.next {
			sb.WriteString(fmt.Sprintf("  n%d -> n%d [label=\"%.2f\"];\n",
				node.id, next.node.id, next.probability))
		}
	}

	sb.WriteString("}\n")
	return sb.String()
}

func nodeTypeToString(typ NodeType) string {
	switch typ {
	case start:
		return "START"
	case header:
		return "HEADER"
	case jump:
		return "JUMP"
	case end:
		return "END"
	case ch:
		return "CHAR"
	case rx:
		return "REGEX"
	case pointer:
		return "POINTER"
	case idk:
		return "IDK"
	default:
		return "UNKNOWN"
	}
}
