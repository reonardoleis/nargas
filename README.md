
![](https://i.imgur.com/ix6cbaA.png)
### [N]ext-[A]dvanced [R]emodeler for [G]enerating [A]ltered [S]yntax

Transpiler implemented using Go with gocc (for lexical & syntactic analysis) for "Rinha de Compilers" compiler (or interpreter) building challenge.

Given any AST file on the challenge's proposed structure, it will transpile it to Go or JS.

### Usage
1. Install go
2. Build with `make build`
3. Run with `./nargas INPUT_FILE_PATH OUTPUT_LANGUAGE [OUTPUT_FILE_PATH]` where **OUTPUT_LANGUAGE** can be either `go` or `js` and **OUTPUT_FILE_PATH** is optional