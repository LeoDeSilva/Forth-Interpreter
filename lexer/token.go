package lexer

type Token struct {
	Type    string
	Literal string
}

var keywords = map[string]string{
	"if":       IF,
	"then":     THEN,
	"do":       DO,
	"loop":     LOOP,
	"dup":      DUP,
	"swap":     SWAP,
	"rot":      ROT,
	"drop":     DROP,
	"nip":      NIP,
	"invert":   INVERT,
	"and":      AND,
	"or":       OR,
	"variable": VARIABLE,
	"over":     OVER,
}

func LookupIdentifier(identifier string) string {
	if token, ok := keywords[identifier]; ok {
		return token
	}

	return IDENTIFIER
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENTIFIER = "IDENTIFIER"
	STRING     = "STRING"
	INT        = "INT"

	ADD = "ADD"
	SUB = "SUB"
	MUL = "MUL"
	DIV = "DIV"
	MOD = "MOD"

	ADD_EQ = "PLUS_EQ"
	SUB_EQ = "SUB_EQ"
	MUL_EQ = "MUL_EQ"
	DIV_EQ = "DIV_EQ"

	EQ  = "EQ"
	EE  = "EE"
	NOT = "NOT"
	NE  = "NE"
	GT  = "GT"
	GTE = "GTE"
	LT  = "LT"
	LTE = "LTE"
	AT  = "AT"

	COMMA     = "COMMA"
	DOT       = "DOT"
	COLON     = "COLON"
	SEMICOLON = "SEMICOLON"
	QUESTION  = "QUESTION"
	DOLLAR    = "DOLLAR"

	LPAREN = "LPAREN"
	RPAREN = "RPAREN"
	LBRACE = "LBRACE"
	RBRACE = "RBRACE"

	// KEYWORDS
	IF   = "IF"
	THEN = "THEN"
	DO   = "DO"
	LOOP = "LOOP"

	DUP  = "DUP"
	SWAP = "SWAP"
	ROT  = "ROT"
	DROP = "DROP"
	NIP  = "NIP"
	OVER = "OVER"

	INVERT   = "INVERT"
	AND      = "AND"
	OR       = "OR"
	VARIABLE = "VARIABLE"

	// Parser Nodes
	PROGRAM  = "PROGRAM"
	WHILE    = "WHILE"
	FUNCTION = "FUNCTION"
)
