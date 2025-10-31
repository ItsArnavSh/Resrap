# Codebase generation

V1.1 introduces a format to define a directory structure, complete with no of files, token length etc for each file and generate entire codebases, not just single files


An example structure format can be seen below:
```dsl
src/
  code/
    c[10x20 code_*.c]
    core/
      c[5x1000 *.c]
  more/
    sql[10x100 *.sql]
```
`sql[10x100 *.sql]`: access the sql grammar, generate 10 files, 100 tokens each with name [unique_identifier].sql

Then simply call GenerateCodebase
```go
	r := resrap.NewResrap()
	r.ParseGrammarFile("sql", "../example/sql.g4")
	r.ParseGrammarFile("c", "../example/c.g4")
	r.GenerateCodebase("../example/format.dsl", ".")
//Errors are ignored for simplicity
// Be responsible DON'T EVER IGNORE THEM
```
