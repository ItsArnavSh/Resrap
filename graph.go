package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SyntaxNode struct {
	down []*SyntaxNode
	next []*SyntaxNode
	name string
}

func (s *SyntaxNode) AddEdgeNext(g *SyntaxGraph, node *SyntaxNode) {
	s.next = append(s.next, node)
	g.nodeRef[node.name] = node

}
func (s *SyntaxNode) AddEdgeDown(g *SyntaxGraph, node *SyntaxNode) {
	s.down = append(s.down, node)
	g.nodeRef[node.name] = node
}

type SyntaxGraph struct {
	nodeRef map[string]*SyntaxNode
}

func (s *SyntaxGraph) GetNode(name string) *SyntaxNode {
	if s.nodeRef[name] != nil {
		return s.nodeRef[name]
	}
	newNode := &SyntaxNode{nil, nil, name}
	s.nodeRef[name] = newNode
	return newNode
}

func NewSyntaxGraph() SyntaxGraph {
	return SyntaxGraph{
		nodeRef: make(map[string]*SyntaxNode),
	}
}

// Interactive traversal
func (s *SyntaxGraph) TraverseInteractive(startName string) {
	node := s.GetNode(startName)
	if node == nil {
		fmt.Println("Node not found:", startName)
		return
	}

	stack := []*SyntaxNode{} // for "go back" functionality
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nCurrent node:", node.name)

		edges := []*SyntaxNode{}
		fmt.Println("Edges:")
		idx := 0
		for _, n := range node.next {
			fmt.Printf("%d) next -> %s\n", idx, n.name)
			edges = append(edges, n)
			idx++
		}
		for _, d := range node.down {
			fmt.Printf("%d) down -> %s\n", idx, d.name)
			edges = append(edges, d)
			idx++
		}
		fmt.Println("-1) go back")
		fmt.Print("Choice: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input, try again")
			continue
		}

		if choice == -1 {
			if len(stack) == 0 {
				fmt.Println("Already at the top, exiting.")
				return
			}
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			continue
		}

		if choice < 0 || choice >= len(edges) {
			fmt.Println("Invalid choice, try again")
			continue
		}

		stack = append(stack, node)
		node = edges[choice]
	}
}

func (s *SyntaxGraph) ExportDOTFile(filename string) error {
	// Open file for writing
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create DOT file: %w", err)
	}
	defer f.Close()

	// Write the DOT content
	_, err = f.WriteString("digraph G {\n  rankdir=LR;\n")
	if err != nil {
		return fmt.Errorf("failed to write to DOT file: %w", err)
	}

	seen := map[string]bool{}
	var visit func(node *SyntaxNode) error
	visit = func(node *SyntaxNode) error {
		if seen[node.name] {
			return nil
		}
		seen[node.name] = true

		// Declare node
		_, err := fmt.Fprintf(f, "  \"%s\" [shape=box];\n", node.name)
		if err != nil {
			return err
		}

		// next edges (solid blue)
		for _, nxt := range node.next {
			_, err := fmt.Fprintf(f, "  \"%s\" -> \"%s\" [color=blue];\n", node.name, nxt.name)
			if err != nil {
				return err
			}
			if err := visit(nxt); err != nil {
				return err
			}
		}

		// down edges (dashed red)
		for _, dn := range node.down {
			_, err := fmt.Fprintf(f, "  \"%s\" -> \"%s\" [style=dashed color=red];\n", node.name, dn.name)
			if err != nil {
				return err
			}
			if err := visit(dn); err != nil {
				return err
			}
		}

		return nil
	}

	// Traverse all nodes from map (to catch unreferenced nodes)
	for _, node := range s.nodeRef {
		if err := visit(node); err != nil {
			return err
		}
	}

	_, err = f.WriteString("}\n")
	if err != nil {
		return fmt.Errorf("failed to finalize DOT file: %w", err)
	}

	return nil
}

func (s *SyntaxGraph) Clean() {
	// Step 1: Clean duplicates and prioritize down over next
	for _, node := range s.nodeRef {
		// Deduplicate next edges
		nextSeen := map[string]*SyntaxNode{}
		cleanNext := []*SyntaxNode{}
		for _, n := range node.next {
			if _, exists := nextSeen[n.name]; !exists {
				nextSeen[n.name] = n
				cleanNext = append(cleanNext, n)
			}
		}
		node.next = cleanNext

		// Deduplicate down edges
		downSeen := map[string]*SyntaxNode{}
		cleanDown := []*SyntaxNode{}
		for _, d := range node.down {
			if _, exists := downSeen[d.name]; !exists {
				downSeen[d.name] = d
				cleanDown = append(cleanDown, d)
			}
		}
		node.down = cleanDown

		// Prioritize down over next
		filteredNext := []*SyntaxNode{}
		for _, n := range node.next {
			if _, exists := downSeen[n.name]; !exists {
				filteredNext = append(filteredNext, n)
			}
		}
		node.next = filteredNext
	}

	// Step 2: Remove isolated nodes
	// Mark all nodes that are referenced by edges
	referenced := make(map[string]bool)
	for _, node := range s.nodeRef {
		for _, n := range node.next {
			referenced[n.name] = true
		}
		for _, d := range node.down {
			referenced[d.name] = true
		}
	}

	// Remove nodes that have no edges pointing to them and don’t point anywhere
	for name, node := range s.nodeRef {
		if len(node.next) == 0 && len(node.down) == 0 && !referenced[name] {
			delete(s.nodeRef, name)
		}
	}
}

func (s *SyntaxGraph) RandomWalker(start string, no int32) {
	//
}
