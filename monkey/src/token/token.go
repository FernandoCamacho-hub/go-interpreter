package token

type TokenType string

type Token struct {
    Type TokenType
    Literal string
}

const (
    ILLEGAL = "ILLEGAL"
    EOF = "EOF"
    
    // Identifiers + literals
    IDENT = "IDENT"
    INT = "INT"
    
    // Operators
    ASSIGN = "ASSIGN"
    PLUS = "PLUS"
    
    // Delimiters
    COMA = "COMA"
    SEMICOLON = "SEMICOLON"
    
    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"

    // Keywords
    FUNCTION = "FUNCTION"
    LET = "LET"
)
