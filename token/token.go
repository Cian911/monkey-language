package token

// This defines as the TokenTypes for our language
const (
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"

	// Identifies & Literals
	IDENT = "IDENT" // add, foobar, x, y
	INT   = "INT"   // 12345

	// Operators
	PLUS   = "+"
	ASSIGN = "="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

// Specify language keywords
var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// Defining TokenType as a string allows us to use many different values as TokenTypes
type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
