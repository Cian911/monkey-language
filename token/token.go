package token

// This defines as the TokenTypes for our language
const (
  EOF = "EOF"
  ILLEGAL = "ILLEGAL"

  // Identifies & Literals
  IDENT = "IDENT" // add, foobar, x, y
  INT = "INT" // 12345

  // Operators
  PLUS = "+"
  ASSIGN = "="

  // Delimiters
  COMMA = ","
  SEMICOLON = ";"

  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"

  // keywords
  FUNCTION = "FUNCTION"
  LET = "LET"
)

// Defining TokenType as a string allows us to use many different values as TokenTypes
type TokenType string

type Token struct {
  Type    TokenType
  Literal string
}
