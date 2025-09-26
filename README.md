# Resrap

*Just a parser… in reverse.*

---

## What is Resrap?

Resrap is a **seedable, grammar-based code snippet generator**. Instead of parsing code, it **generates code** from formal grammars — producing endless, realistic-looking (or hilariously nonsensical) code snippets.

It works with **any language** that can be described with a grammar (even English if you like!), and is perfect for:

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
program : header+ function^;

header : '#include<' identifier '.h>\n';
function : functionheader '{' '\n' functioncontent '}';
functionheader : datatype ' ' identifier '(' ')' ;
functioncontent : block+;
block : statement+ | ifblock | whileblock;
ifblock : 'if(' conditionalexpression '){\n' statement+ '}\n';
whileblock : 'while(' conditionalexpression '){\n' statement+ '}\n';
statement : assignment;
assignment : datatype ' ' identifier ' = ' expression ';\n';
expression : operand (operator (operand | '(' expression ')'))*;
operand : identifier | integer | float;
operator : ' + ' | ' - ' | ' * ' | ' / ';
datatype : 'int' | 'float' | 'double';
```

**ABNF new features in action:**

```abnf
char : a<0.2> | b<0.8>;
a    : 'A';
b    : 'B';

program : function^;  # infinite generation of functions
```

* Probabilities (`<...>`) control **branch selection weights**.
* Infinite generation (`^`) allows looping nodes without halting generation.

---

**Generated code example:**

```c
#include<hello.h>
#include<password.h>
float database(){
    int error = login + 7.77 - (database);
}
float test(){
    double failure = 42 * 7;
    if(start > 999.77 && code > 101){
        float method = 77 + (table * 42.88);
        float database = 7.256 / 42 * function;
    }
    double data = 999.13 - (table + (101.13));
    float test = hello / 88.13 + method + (101 / table);
}
```

---

## Installation

```bash
go get github.com/ItsArnavSh/Resrap@v0.1.0
```

---

## Usage

```golang
func examplecode() {
	graphs := resrap.NewResrap()
	graphs.ParseGrammarFile("C", "example/C.g4")
	
	random_content := graphs.GenerateRandom("C", "program",400)
	fmt.Println(random_content)
	
	seeded_content := graphs.GenerateWithSeeded("C", "program", 20,400)//20 is the seed 400 are the tokens 
	fmt.Println(seeded_content)
}
```

* **GenerateRandom** → Random traversal using internal PRNG
* **GenerateWithSeeded** → Deterministic generation using a fixed seed

---

## Roadmap

- Adding Multithreading Capabilities to serve multiple users
  
- Maintain generation sessions: generate snippets in chunks (e.g., first 10 lines, then next 10, etc.)

---

## Motivation

We wanted a way to:

* Generate **unlimited, realistic code snippets**
* Avoid copyright issues from using real code in samples or demos
* Make a **fun, shareable, and deterministic code generator**
* Give programmers and typists a playground for **syntax, speed, and randomness**

Resrap is “just a parser… in reverse” — it turns grammars into chaos, one snippet at a time.

---

## ABNF (Awesome BNF)

Resrap uses **ABNF**, a lightweight grammar format extending standard EBNF:

* **`^`** → Infinite generation (loops nodes without halting)
* **`<prob>`** → Weighted probabilities for branching
* Default behavior maintains equal probability for unspecified choices
* Fully compatible with standard EBNF operators: `+`, `*`, `?`, `()`

See [ABNF documentation](docs/ABNF.md) for detailed syntax, examples, and probability usage.

---

