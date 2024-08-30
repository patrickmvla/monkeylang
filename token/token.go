package token

type TokenType string

type Token struct {
	Type	TokenType
	Literal	string
}

const (
	Illegal = "Illegal"
	EOF = "EOF"
	
	// Identifiers + literals
	Identifier = "Identifier" // add, x, y...
	Int = "Int" // 12345
	String = "String" // "x", "y"

	// Operators
	Assign = "="
	Plus = "+"
	Minus = "-"
	Bang = "!"
	Asterisk = "*"
	Slash = "/"
	Equal = "=="
	NotEqual = "!="

	LessThan = "<"
	GreaterThan = ">"

	// Delimiters
	Comma = ","
	SemiColon = ";"
	Colon = ":"

	LeftParenthesis = "("
	RightParenthesis = ")"
	LeftBrace = "{"
	RightBrace = "}"
	LeftBracket = "["
	RightBracket = "]"

	// KeyWords
	Function = "Function"
	Let = "Let"
	True = "True"
	False = "False"
	If = "If"
	Else = "Else"
	Return = "Return"
)

var keywords = map[string]TokenType {
	"fn": Function,
	"let": Let,
	"true": True,
	"false": False,
	"if": If,
	"else": Else,
	"return": Return,
}

func LookUpIdentifierType(identifier string) TokenType {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}

	return Identifier
}