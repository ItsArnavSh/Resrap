# Resrap
[![Landing](https://img.shields.io/badge/docs-resrap.osdc.dev-blue)](https://resrap.osdc.dev)
[![Go Reference](https://pkg.go.dev/badge/github.com/osdc/resrap.svg)](https://pkg.go.dev/github.com/osdc/resrap)

*Just a parser… in reverse.*

---

## What is Resrap?

Resrap is a **seedable, grammar-based code snippet generator**. Instead of parsing code, it **generates code** from formal grammars — producing endless, realistic-looking (or hilariously nonsensical) snippets.

It works with **any language** that can be described with a grammar (even English if you like!) and is perfect for:

* Typing practice with realistic-looking snippets
* Stress-testing parsers, syntax highlighters, or linters
* Fun exploration of procedural code generation

Resrap now also supports **probabilistic and infinitely repeatable grammars** via the **ABNF (Awesome BNF) format** — see `docs/ABNF.md` for full reference.

---

## How?

Resrap reads a grammar and builds a **graph of expansions**. It then randomly traverses the graph (or deterministically with a seed) to produce snippets that:

* Follow the grammar’s syntax rules
* Look structurally like real code
* Include probabilities for weighted choices (`<0.2>`)
* Support infinite loops via the `^` operator

**Example grammar snippet (simplified C):**

```abnf
program : (header+<0.4>) function^;
header:'#include<'identifier'.h>\n';
function:functionheader'{''\n'functioncontent'}';
functionheader:datatype ' ' identifier '(' ')' ;
...
...
````

---

## Generated code example

```c
#include<success.h>
#include<email.h>
double class(){
while(variable < query && password < variable){
int result = variable + (user / hello);
}
}double user(){
if(class > user && result < class){
float password = 1024.13 - (13.7);
}
}double hello(
```

---

## Installation

```bash
go get github.com/ItsArnavSh/Resrap@v0.1.0
```

---

## Single-Threaded Usage

```golang
package main

import (
    "fmt"
    "github.com/ItsArnavSh/Resrap/resrap"
)

func main() {
    r := resrap.NewResrap()
    r.ParseGrammarFile("C", "example/C.g4")

    code := r.GenerateRandom("C", "program", 400)
    fmt.Println(code)

    seeded := r.GenerateWithSeeded("C", "program", 20, 400) // 20 = seed, 400 tokens
    fmt.Println(seeded)
}
```

* **GenerateRandom** → Random traversal using internal PRNG
* **GenerateWithSeeded** → Deterministic generation using a fixed seed

---

## Multithreaded Usage (ResrapMT)

`ResrapMT` allows **parallel generation** for multiple jobs using a worker pool. This is ideal for **server environments** or **batch processing**.

### Example

```golang
package main

import (
    "fmt"
    "github.com/ItsArnavSh/Resrap/resrap"
)

func main() {
    rmt := resrap.NewResrapMT(10, 100) // 10 workers, queue size 100
    rmt.ParseGrammarFile("C", "example/C.g4")
    rmt.StartResrap() // spawn worker pool

    // Submit jobs with unique IDs
    for i := 0; i < 10; i++ {
        jobID := fmt.Sprintf("job-%d", i)
        rmt.GenerateRandom(jobID, "C", "program", 100)
    }

    // Collect results
    for i := 0; i < 10; i++ {
        res := <-rmt.CodeChannel
        fmt.Printf("Job %s generated code: %s\n", res.Id, res.Code[:50])
    }
}
```

### Notes

* **IDs:** You must create unique IDs for each job; results are returned with the ID.
* **CodeChannel:** A **blocking, unbounded channel** — handle results yourself.
* **Why multithreaded?** Efficiently handles **many concurrent jobs**, fully utilizing CPU cores while keeping grammar graphs immutable and lock-free.

> For benchmarks and performance comparisons, see [benchmark-results/Multithreading.md](benchmark-results/Multithreading.md).

---

## Roadmap

* Maintain generation sessions (e.g., generate snippets in chunks)
* Dynamic worker scaling and idle worker shutdown

---

## Motivation

Resrap was created to:

* Generate **unlimited, realistic code snippets**
* Avoid copyright issues from using real code
* Provide a **fun, deterministic, and probabilistic code generator**
* Give programmers a playground for **syntax, speed, and randomness**

“Just a parser… in reverse.”

---

## ABNF (Awesome BNF)

* `^` → Infinite generation (loops nodes without halting)
* `<prob>` → Weighted probabilities for branching
* Compatible with standard EBNF operators: `+`, `*`, `?`, `()`

See [docs/ABNF.md](docs/ABNF.md) for full syntax and examples.

