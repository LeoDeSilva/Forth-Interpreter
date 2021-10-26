package lexer 

import (
    "testing"
    "forth/token"
)

func TestNextToken(t *testing.T) {
    input := `:fizz? 3 % 0 = dup if "Fizz" . then ;
    :do-fizz-buzz 1 25 do fizz? if i . then loop ;
    `

    tests := []struct {
        expectedType string
        expectedLiteral string
    }{
        {token.COLON, ":"},
        {token.IDENTIFIER,"fizz?"},
        {token.INT, "3"},
        {token.MOD, "%"},
        {token.INT, "0"},
        {token.EQ, "="},
        {token.DUP, "dup"},
        {token.IF, "if"},
        {token.STRING, "Fizz"},
        {token.DOT, "."},
        {token.THEN, "then"},
        {token.SEMICOLON, ";"},
        {token.COLON, ":"},
        {token.IDENTIFIER, "do-fizz-buzz"},
        {token.INT, "1"},
        {token.INT, "25"},
        {token.DO, "do"},
        {token.IDENTIFIER, "fizz?"},
        {token.IF, "if"},
        {token.IDENTIFIER, "i"},
        {token.DOT, "."},
        {token.THEN, "then"},
        {token.LOOP, "loop"},
        {token.SEMICOLON, ";"},
    }

    l := New(input)
    for i,tt := range tests {
        tok := l.NextToken()

        if tok.Type != tt.expectedType {
            t.Fatalf("tests[%d] - tokentype wrong. expected %q, got %q",i,tt.expectedType,tok.Type)
        }

        if tok.Literal != tt.expectedLiteral {
            t.Fatalf("tests[%d] - literal wrong. expected %q, got %q", i,tt.expectedLiteral,tok.Literal)
        }
    }
 }



