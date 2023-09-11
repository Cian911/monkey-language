### Monkey Programming Language

Based on the book - Writing an Interpreter in Go.

##### What we will build

- An AST (Abstract Syntax Tree) or a "tree-walking interpreter".
- Our own "lexer"
- Our own "parser"
- Our own tree representation and an "evaluator"

##### Lexical Analysis

In order to work with our source code, we first need to transform it into something our interpreter is going to understand. We can do this by performing two transformations on our source code. `tokens` & `abstract syntax tree`.

`tokens` act like small interpretable data structures that are then fed to a parser which turns these tokens into an `abstract syntax tree`.

Example:

```bash
source:
"let x = 5 + 5;"

token:
[
  LET,
  IDENTIFIER(x),
  EQUALS_SIGN,
  INTEGER(5),
  PLUS_SIGN,
  INTEGER(5),
  SEMICOLON
]
```

##### Writing the Lexer

The lexer will take source code as input and output the tokens that represt the source code. It will go through it's input, and output the next token it recognizes. This will be done using a `NextToken()` method.



