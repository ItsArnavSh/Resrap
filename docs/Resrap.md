
# Resrap

`Resrap` is a Go library for managing language grammars and generating random code or text based on them. It provides both **single-threaded** and **multithreaded** implementations (see [ResrapMT.md](docs/ResrapMT.md) for the multithreaded version).

---

## Overview

The `Resrap` object allows:

- Loading multiple grammars
- Generating random content from a grammar
- Deterministic generation using a seed

> **Note:** The grammar graphs are treated as immutable after parsing, allowing efficient traversal and generation.

---

## Structs

### `Resrap`

The main object of the library.

```go
type Resrap struct {
    languageGraph map[string]lang
}
````

* `languageGraph` — a map of grammar names to `lang` objects representing parsed grammars.

---

## Constructor

### `NewResrap() *Resrap`

Creates and returns a new `Resrap` instance.

```go
resrap := resrap.NewResrap()
```

* Returns a `Resrap` object with no grammars loaded.

---

## Grammar Parsing

### `ParseGrammar(name, grammar string)`

Parses a grammar from a string and stores it under the given name.

```go
resrap.ParseGrammar("C", "grammar content here")
```

* `name` — unique identifier for this grammar (e.g., `"C"`).
* `grammar` — the grammar definition as a string.
* The grammar is normalized internally after parsing.

---

### `ParseGrammarFile(name, location string)`

Parses a grammar from a file and stores it under the given name.

```go
resrap.ParseGrammarFile("C", "example/C.g4")
```

* `name` — unique identifier for this grammar.
* `location` — path to the grammar file.
* Internally, the grammar is normalized after parsing.

---

## Content Generation

### `GenerateRandom(name, starting_node string, tokens int) string`

Generates random content from the grammar identified by `name`.

```go
code := resrap.GenerateRandom("C", "program", 100)
```

* `name` — grammar identifier.
* `starting_node` — starting symbol in the grammar.
* `tokens` — number of tokens to generate.
* **Returns:** generated content as a string.
* Uses a non-deterministic random number generator.

---

### `GenerateWithSeeded(name, starting_node string, seed uint64, tokens int) string`

Generates deterministic content from the grammar using a numeric seed.

```go
code := resrap.GenerateWithSeeded("C", "program", 12345, 100)
```

* `name` — grammar identifier.
* `starting_node` — starting symbol in the grammar.
* `seed` — numeric seed for deterministic generation.
* `tokens` — number of tokens to generate.
* **Returns:** generated content as a string.

---

## Usage Example

```go
package main

import (
    "fmt"
    "yourmodule/resrap"
)

func main() {
    r := resrap.NewResrap()
    r.ParseGrammarFile("C", "example/C.g4")

    // Random generation
    code := r.GenerateRandom("C", "program", 100)
    fmt.Println(code)

    // Deterministic generation
    code2 := r.GenerateWithSeeded("C", "program", 12345, 100)
    fmt.Println(code2)
}
```

---

## Notes

* `Resrap` is **single-threaded**.
* For multithreaded generation, see the dedicated documentation: [ResrapMT.md](docs/ResrapMT.md)
* Grammar graphs are immutable after parsing — safe for concurrent reads if needed.

---

This file provides a **full overview** of the single-threaded `Resrap` API and its usage.
