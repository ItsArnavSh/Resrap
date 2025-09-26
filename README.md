# Resrap
*Just a parser… in reverse.*

---

## What is Resrap?

Resrap is a **seedable, grammar-based code snippet generator**. Instead of parsing code, it **generates code** from formal grammars — producing endless, realistic-looking nonsense code.

It works with **any language** that can be described with a grammar (even English if you like!), and is perfect for:

- Typing practice with realistic-looking snippets
- Stress-testing parsers or syntax highlighters
- Fun exploration of procedural code generation

---

## How?

Resrap reads a grammar and builds a **graph of expansions**. It then randomly traverses the graph (or deterministically with a seed) to produce snippets that:

- Follow the grammar’s syntax rules
- Look structurally like real code
- Contain humorous or nonsensical variable names and expressions

**Example grammar snippet (simplified C):**
```text
program : header+ function+;
header:'#include<'identifier'.h>\n';
function:functionheader'{''\n'functioncontent'}';
functionheader:datatype ' ' identifier '(' ')' ;
...
````

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

## Roadmap

* **NextUP**
- Weighted nodes to bias generation (e.g., more variables, fewer numbers)
- Have infinite Code length so it can be generated upto any number of lines as needed.
- Maintain Generation sessions to allow you generate continuous snippets in breaks (Eg first get first 10 lines then next 10 and so on)
---

## Motivation

We wanted a way to:

* Generate **unlimited, realistic code snippets**
* Avoid copyright issues from using real code in samples or demos
* Make a **fun, shareable, and deterministic code generator**
* Give programmers and typists alike a playground for syntax, speed, and randomness

Resrap is “just a parser… in reverse” i.e. it turns grammars into chaos, one snippet at a time.

---
