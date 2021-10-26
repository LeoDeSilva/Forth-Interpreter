package lexer

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) Lex() []Token {
	var tokens []Token

	for l.ch != 0 {
		tok := l.NextToken()
		tokens = append(tokens, tok)
	}

	tokens = append(tokens, Token{Type: EOF, Literal: ""})

	return tokens
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func newToken(tokenType string, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) NextToken() Token {
	var tok Token
	l.eatWhitespace()

	switch l.ch {
	case ':':
		tok = newToken(COLON, l.ch)
	case ';':
		tok = newToken(SEMICOLON, l.ch)
	case '.':
		tok = newToken(DOT, l.ch)
	case ',':
		tok = newToken(COMMA, l.ch)
	case '=':
		tok = l.readDouble(EQ, '=', EE)
	case '>':
		tok = l.readDouble(GT, '=', GTE)
	case '<':
		tok = l.readDouble(LT, '=', LTE)
	case '+':
		tok = l.readDouble(ADD, '!', ADD_EQ)
	case '-':
		tok = l.readDouble(SUB, '!', SUB_EQ)
	case '/':
		tok = l.readDouble(DIV, '!', DIV_EQ)
	case '*':
		tok = l.readDouble(MUL, '!', MUL_EQ)
	case '%':
		tok = newToken(MOD, l.ch)
	case '@':
		tok = newToken(AT, l.ch)
	case '$':
		tok = newToken(DOLLAR, l.ch)
	case '?':
		tok = newToken(QUESTION, l.ch)
	case '"':
		tok.Literal = l.readString()
		tok.Type = STRING
	case '!':
		tok = l.readDouble(NOT, '=', NE)
	case '(':
		tok = newToken(LPAREN, l.ch)
	case ')':
		tok = newToken(RPAREN, l.ch)
	case '{':
		tok = newToken(LBRACE, l.ch)
	case '}':
		tok = newToken(RBRACE, l.ch)
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = INT
			tok.Literal = l.readNumber()
			return tok

		} else {
			tok = newToken(ILLEGAL, l.ch)
		}

	}

	l.readChar()
	return tok
}

func (l *Lexer) eatWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '?' || ch == '-'
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	l.readChar()
	position := l.position

	for l.ch != '"' {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readDouble(firstType string, second byte, secondType string) Token {
	ch := l.ch
	if l.peekChar() == second {
		l.readChar()
		return Token{Type: secondType, Literal: string(ch) + string(l.ch)}
	} else {
		return Token{Type: firstType, Literal: string(ch)}
	}
}
