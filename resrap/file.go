package resrap

import (
	"bufio"
	"os"
	"strings"
)

// ParseFile reads a file and returns statements split by lines ending with ';',
// skipping lines starting with "//".
func ParseFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var statements []string
	var current strings.Builder

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// skip comment or empty lines
		if strings.HasPrefix(line, "//") || line == "" {
			continue
		}

		// accumulate this line
		current.WriteString(line + " ")

		// if line ends with ';', finalize statement
		if strings.HasSuffix(line, ";") {
			stmt := strings.TrimSpace(current.String())
			statements = append(statements, stmt)
			current.Reset()
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return statements, nil
}

type lang struct {
	graph *syntaxGraph
	nodes int
}

func newLang() lang {
	return lang{}
}
func (l *lang) GetGraph() *syntaxGraph {
	return l.graph
}
func (l *lang) ParserFile(filename string) error {
	lines, err := ParseFile(filename)
	if err != nil {
		return err
	}
	gb := newGraphBuilder()
	content := strings.Join(lines, "")
	err = gb.start_generation(content)
	l.graph = &gb.pars.graph
	return err
}

func (l *lang) ParserString(data string) error {
	gb := newGraphBuilder()
	err := gb.start_generation(data)
	l.graph = &gb.pars.graph
	return err
}
