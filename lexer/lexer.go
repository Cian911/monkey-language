package lexer

import "github.com/cian911/monkey-language/token"

type Lexer struct {
	input        string
	position     int  // Current position in input (current char)
	readPosition int  // Current reading position in input (after current char)
	ch           byte // Current char being inspected
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readChar()

	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		tok = newToken(token.BANG, l.ch)
	case '*':
		tok = newToken(token.ASTERIX, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	// Advance position
	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	// If we encounter whitespace or newLine, move position to next ch
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	// If we come across a letter, continue to read it
	// In order to check if it's an identifier
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	// If we come across a number, continue to read it
	// In order to check if it's an int
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	// Remember these are bytes, so we're dealing with ascii here
	// We're basically checking if we're in range of the human alphabet
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokType,
		Literal: string(ch),
	}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// If we've reached the end of the input, set ch to 0
		// 0 is the ASCII code for NUL
		l.ch = 0
	} else {
		// Else set ch to next value in input
		l.ch = l.input[l.readPosition]
	}
	// Set position to point to next value
	l.position = l.readPosition
	l.readPosition += 1
}
