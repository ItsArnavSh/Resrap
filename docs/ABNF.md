# ABNF – Awesome BNF

**ABNF (Awesome BNF)** is a lightweight, custom grammar format designed for the **Resrap code generation tool**. It extends standard EBNF/BNF with **probabilities** and **infinite generation**, making grammar definitions **expressive, flexible, and efficient** for high-throughput code generation.

---

## 1. Quick Revision: EBNF and Grammar

ABNF inherits the basics of EBNF:

```abnf
function : rules ;
```

* **function** → Name of the non-terminal rule.
* **rules** → Sequence of non-terminals, terminals, or characters.

### Standard EBNF Operators

| Operator | Meaning                                               |
| -------- | ----------------------------------------------------- |
| `+`      | One or more repetitions of the preceding element      |
| `*`      | Zero or more repetitions                              |
| `?`      | Optional element (zero or one occurrence)             |
| `()`     | Grouping; treat enclosed sequence as a single element |

**Example:**

```abnf
expression : operand (operator operand)* ;  # operand followed by zero or more operator-operand pairs
```

Statements form a **directed graph**, allowing nodes to reference others and create structured generation flows.

---

## 2. Infinite Generation (`^`)

Meet the **Infinity Operator `^`**:

* Connects the **end of a rule back to the current node**, allowing **infinite generation**.
* The generator **does not halt** when reaching this node; it loops back to continue generation.

**Example (C Grammar):**

```abnf
program : function^;
```

* Generates **functions endlessly** until a token limit is reached.
* Can be combined with other operators (`+`, `*`, `?`) for **complex, repeatable structures**.

---

## 3. Probabilities (`<...>`)

By default, choices in grammar rules are **equally likely**. ABNF allows specifying **weighted probabilities** for finer control:

```abnf
char : a<0.2> | b<0.8>;
a    : 'A';
b    : 'B';
```

* `a` is chosen **20% of the time**, `b` **80%**.
* Probabilities are **normalized automatically** if they don’t sum to 1.
* If no probability is specified, choices are **assumed equal**.

**Looping with probability:**

```abnf
a : b+<0.4>;
```

* 40% chance to **loop back to `b`**, 60% to **move ahead**.
* Same logic applies to `*`, `?`, and `^`.

---

## 4. Summary

ABNF extends standard BNF/EBNF with:

* **Repetition operators**: `+`, `*`, `?`
* **Grouping**: `()`
* **Infinite generation**: `^` → loop nodes infinitely
* **Weighted choices**: `<prob>` → specify probabilities for branches
* **Default equal probability** → backwards compatibility

ABNF is ideal for **stochastic code generation**, fuzzing, and testing, giving users full control over **structure, randomness, and recursion** in generated code.
