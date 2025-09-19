package main

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
