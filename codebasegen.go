package resrap

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type file_type int

const (
	DIR file_type = iota
	FILE
)

// node describes either a directory or file entry.
type node struct {
	Name       string
	Type       file_type
	Pattern    string // E.g., "code_*.c"
	Count      int    // Number of files (e.g., 10)
	TokenCount int    // Tokens per file (e.g., 20)
	FileType   string // Language/type identifier (e.g., "C", "sql")
	Children   []*node
}

// ParseDSL parses the directory DSL into a tree.
func ParseDSL(input string) (*node, error) {
	lines := strings.Split(input, "\n")
	var root = &node{Name: "", Type: DIR}
	stack := []*node{root}
	indentStack := []int{-1}

	for lineNum, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		indent := len(line) - len(strings.TrimLeft(line, " "))
		name := strings.TrimSpace(line)

		var n node

		// Check if it's a file specification: name[count x tokens pattern]
		if i := strings.Index(name, "["); i != -1 && strings.HasSuffix(name, "]") {
			n.Type = FILE
			n.FileType = name[:i] // e.g., "C", "sql"

			rest := name[i+1 : len(name)-1] // e.g., "10x20 code_*.c"
			parts := strings.Fields(rest)

			if len(parts) < 2 {
				return nil, fmt.Errorf("line %d: invalid file spec format, expected 'type[countXtokens pattern]'", lineNum+1)
			}

			// Parse "10x20" format
			xparts := strings.Split(parts[0], "x")
			if len(xparts) != 2 {
				return nil, fmt.Errorf("line %d: invalid count format '%s', expected 'NUMxNUM'", lineNum+1, parts[0])
			}

			var err error
			n.Count, err = strconv.Atoi(xparts[0])
			if err != nil {
				return nil, fmt.Errorf("line %d: invalid file count '%s': %w", lineNum+1, xparts[0], err)
			}

			n.TokenCount, err = strconv.Atoi(xparts[1])
			if err != nil {
				return nil, fmt.Errorf("line %d: invalid token count '%s': %w", lineNum+1, xparts[1], err)
			}

			n.Pattern = parts[1] // e.g., "code_*.c"

			if n.Count <= 0 || n.TokenCount <= 0 {
				return nil, fmt.Errorf("line %d: count and token count must be positive numbers", lineNum+1)
			}
		} else {
			// It's a directory
			n.Type = DIR
			n.Name = name
		}

		// Pop from stack until correct indentation level
		for len(indentStack) > 0 && indentStack[len(indentStack)-1] >= indent {
			stack = stack[:len(stack)-1]
			indentStack = indentStack[:len(indentStack)-1]
		}

		if len(stack) == 0 {
			return nil, fmt.Errorf("line %d: indentation error", lineNum+1)
		}

		node := &n
		parent := stack[len(stack)-1]
		parent.Children = append(parent.Children, node)

		// Only directories can have children
		if n.Type == DIR {
			stack = append(stack, node)
			indentStack = append(indentStack, indent)
		}
	}

	return root, nil
}

// GenerateStructure creates the directory structure and files
func (n *node) GenerateStructure(r *Resrap, root string) error {
	return n.generate_node(r, root)
}

// Recursive function to create directories and generate files
func (n *node) generate_node(r *Resrap, root string) error {
	loc := root
	if n.Name != "" {
		loc = filepath.Join(root, n.Name)
	}

	if n.Type == DIR {
		// Create directory and parents if needed
		if err := os.MkdirAll(loc, 0755); err != nil {
			return fmt.Errorf("failed to create directory '%s': %w", loc, err)
		}

		// Recurse for children
		for _, child := range n.Children {
			if err := child.generate_node(r, loc); err != nil {
				return err
			}
		}
	} else {
		// FILE type - generate multiple files based on count
		if n.Count <= 0 {
			return fmt.Errorf("invalid file count %d for pattern '%s'", n.Count, n.Pattern)
		}

		// Ensure parent directory exists
		if err := os.MkdirAll(root, 0755); err != nil {
			return fmt.Errorf("failed to create parent directory '%s': %w", root, err)
		}

		for i := 1; i <= n.Count; i++ {
			// Replace * with the file number
			fileName := strings.ReplaceAll(n.Pattern, "*", strconv.Itoa(i))
			fullPath := filepath.Join(root, fileName)

			// Generate file content using the FileType (e.g., "C", "sql") and TokenCount
			content := r.GenerateRandom(n.FileType, "program", n.TokenCount)

			// Write to file
			if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
				return fmt.Errorf("failed to write file '%s': %w", fullPath, err)
			}

			fmt.Printf("Generated: %s (%d tokens)\n", fullPath, n.TokenCount)
		}
	}

	return nil
}
