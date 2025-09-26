
# ABNF – Awesome BNF

**ABNF** (Awesome BNF) is a lightweight, custom grammar format designed for the **Resrap parser**. It extends standard EBNF with additional operators and conventions to make grammar definitions both expressive and efficient for high-throughput code generation.

This document explains the syntax, operators, and regex usage in ABNF.

---

## 1. Basic Structure

A **statement** in ABNF has the form:

```abnf
function : rules ;
```

* **function** → The name of the non-terminal rule.
* **rules** → A sequence of **functions, characters, tokens, or terminal definitions**.

Rules can reference other functions, terminals, or characters, forming a **directed graph** of grammar rules.

---

## 2. Operators

ABNF supports the following operators for defining complex structures:

| Operator | Meaning                                                                                              |
| -------- | ---------------------------------------------------------------------------------------------------- |
| `+`      | One or more repetitions of the preceding element.                                                    |
| `*`      | Zero or more repetitions of the preceding element.                                                   |
| `?`      | Optional element (zero or one occurrence).                                                           |
| `()`     | Grouping; treat enclosed sequence as a single element.                                               |
| `^`      | Infinite cycle; connects the current rule’s end back to a target node, enabling infinite generation. |

**Examples:**

```abnf
program : header+ function^;   # program = one or more headers followed by infinitely repeatable functions
header  : '#include<' identifier '.h>\n';  # header includes a file
functioncontent : block+;      # function content is one or more blocks
```

* `header+` → generate **one or more headers**
* `function^` → after generating a function, **loop back to the function node** (infinite repetition)
* `block+` → one or more blocks inside a function

---

## 3. Sample ABNF Grammar

```abnf
program : header+ function^;

header : '#include<' identifier '.h>\n';

function : functionheader '{' '\n' functioncontent '}';

functionheader : datatype ' ' identifier '(' ')';

functioncontent : block+;

block : statement+ | ifblock | whileblock;

ifblock : 'if(' conditionalexpression '){\n' statement+ '}\n';

whileblock : 'while(' conditionalexpression '){\n' statement+ '}\n';

conditionalexpression : conditionalexpone (conditionaljoin conditionalexpone)*;

conditionalexpone : (operand conditionaloperation operand);

conditionaloperation : ' < ' | ' > ';

conditionaljoin : ' && ' | ' || ';

statement : assignment;

assignment : datatype ' ' identifier ' = ' expression ';\n';

expression : operand (operator (operand | '(' expression ')'))*;

operand : identifier | integer | float;

operator : ' + ' | ' - ' | ' * ' | ' / ';

datatype : 'int' | 'float' | 'double';
```

---

## 4. Terminals & Regex

ABNF allows defining terminals using **regex-like patterns**:

| Terminal     | Pattern        | Description                 |
| ------------ | -------------- | --------------------------- |
| `identifier` | `[A-Z]`        | Single uppercase letter     |
| `integer`    | `[0-9]`        | Single digit                |
| `float`      | `[0-9]\.[0-9]` | Decimal number with one dot |

* Regex in ABNF is **always enclosed in square brackets**.
* The parser treats the regex as a **terminal matcher**, generating tokens that match the pattern.
* Only simple character classes are supported: `[A-Z]`, `[0-9]`, `[a-z]`, etc.

**Example:**

```abnf
identifier : [A-Z];   # generates A, B, C ... Z
integer    : [0-9];   # generates 0-9
```

---

## 5. Grouping and Nesting

* Use parentheses `()` to group sequences:

```abnf
expression : operand (operator operand)*;  # operand followed by zero or more operator-operand pairs
```

* Supports **nested structures**, e.g., `((A+B)*C)`
* Combined with `+`, `*`, `?` operators, it allows defining complex grammars concisely.

---

## 6. Infinite Generation (`^`)

* The `^` operator creates a **loop in the grammar graph**, enabling infinite traversal:

```abnf
function^  # after generating a function, loop back to its start
```

* Combined with **stack-based traversal**, this ensures infinite but safe generation of code blocks.

---

## 7. Summary

ABNF extends standard BNF with:

* **Repetition operators**: `+`, `*`
* **Optionality**: `?`
* **Grouping**: `()`
* **Infinite cycles**: `^`
* **Regex terminals**: `[A-Z]`, `[0-9]`, `[a-z]`

This allows **structured, efficient, and infinite code generation** for the Resrap parser, making it ideal for grammar-based fuzzing, testing, or code synthesis.

---
