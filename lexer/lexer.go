package lexer 

import (
    "forth/token"
)

type Lexer struct {
    input string
    position int
    readPosition int
    ch byte
}


func New(input string) *Lexer {
    l := &Lexer{input: input}
    l.readChar()
    return l
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

func newToken(tokenType string, ch byte) token.Token {
    return token.Token{Type:tokenType,Literal:string(ch)}
}


func (l *Lexer) NextToken() token.Token {
    var tok token.Token
    l.eatWhitespace()

    switch l.ch {
        case ':':
            tok = newToken(token.COLON, l.ch)
        case ';':
            tok = newToken(token.SEMICOLON, l.ch)
        case '.':
            tok = newToken(token.DOT, l.ch)
        case ',':
            tok = newToken(token.COMMA, l.ch)
        case '=':
            tok = l.readDouble(token.EQ,'=',token.EE)
        case '>':
            tok = l.readDouble(token.GT,'=',token.GTE)
        case '<':
            tok = l.readDouble(token.LT,'=',token.LTE)
        case '+':
            tok = l.readDouble(token.ADD,'!',token.ADD_EQ)
        case '-':
            tok = l.readDouble(token.SUB,'!',token.SUB_EQ)
        case '/':
            tok = l.readDouble(token.DIV,'!',token.DIV_EQ)
        case '*':
            tok = l.readDouble(token.MUL,'!',token.MUL_EQ)
        case '%':
            tok = newToken(token.MOD, l.ch)
        case '@':
            tok = newToken(token.AT, l.ch)
        case '$':
            tok = newToken(token.DOLLAR, l.ch)
        case '?':
            tok = newToken(token.QUESTION, l.ch)
        case '"':
            tok.Literal = l.readString()
            tok.Type = token.STRING
        case '!':
            tok = l.readDouble(token.NOT,'=',token.NE)
        case '(':
            tok = newToken(token.LPAREN, l.ch)
        case ')':
            tok = newToken(token.RPAREN, l.ch)
        case '{':
            tok = newToken(token.LBRACE, l.ch)
        case '}':
            tok = newToken(token.RBRACE, l.ch)
        case 0:
            tok.Literal = ""
            tok.Type = token.EOF
        default:
            if isLetter(l.ch) {
                tok.Literal = l.readIdentifier()
                tok.Type = token.LookupIdentifier(tok.Literal)
                return tok
            } else if isDigit(l.ch) {
                tok.Type = token.INT
                tok.Literal = l.readNumber()
                return tok
                
            } else {
                tok = newToken(token.ILLEGAL, l.ch)
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
    for isDigit(l.ch){
        l.readChar()
    }

    return l.input[position:l.position]
}

func (l *Lexer) readString() string {
    l.readChar()
    position := l.position

    for l.ch != '"'{
        l.readChar()
    }

    return l.input[position:l.position]
}

func (l *Lexer) readDouble(firstType string, second byte, secondType string) token.Token {
    ch := l.ch
    if  l.peekChar() == second {
        l.readChar()
        return token.Token{Type:secondType,Literal:string(ch) + string(l.ch)}
    } else {
        return token.Token{Type:firstType,Literal:string(ch)}
    }
}
