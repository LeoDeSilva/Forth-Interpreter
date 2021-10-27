package lexer

import (
	"testing"
)

func TestNextToken(t *testing.T) {
    input := `:fizz? 3 % 0 = dup if "Fizz" . then ;
    :do-fizz-buzz 1 25 do fizz? if i . then loop ;
    `

    tests := []struct {
        expectedType string
        expectedLiteral string
    }{
        {COLON, ":"},
        {IDENTIFIER,"fizz?"},
        {INT, "3"},
        {MOD, "%"},
        {INT, "0"},
        {EQ, "="},
        {DUP, "dup"},
        {IF, "if"},
        {STRING, "Fizz"},
        {DOT, "."},
        {THEN, "then"},
        {SEMICOLON, ";"},
        {COLON, ":"},
        {IDENTIFIER, "do-fizz-buzz"},
        {INT, "1"},
        {INT, "25"},
        {DO, "do"},
        {IDENTIFIER, "fizz?"},
        {IF, "if"},
        {IDENTIFIER, "i"},
        {DOT, "."},
        {THEN, "then"},
        {LOOP, "loop"},
        {SEMICOLON, ";"},
    }

    l := New(input)
    for i,tt := range tests {
        tok := l.NextToken()

        if tok.Type != tt.expectedType {
            t.Fatalf("tests[%d] - ype wrong. expected %q, got %q",i,tt.expectedType,tok.Type)
        }

        if tok.Literal != tt.expectedLiteral {
            t.Fatalf("tests[%d] - literal wrong. expected %q, got %q", i,tt.expectedLiteral,tok.Literal)
        }
    }
 }

func TestLookupIdentifier(t *testing.T) {
	tests := []struct {
		Type string 
		Literal string
	}{
		{IF,"if"},
		{IDENTIFIER,"hello"},
		{THEN, "then"},
		{LOOP,"loop"},
		{IDENTIFIER,"lood"},
		{IDENTIFIER,"add"},
		{VARIABLE,"variable"},
		{IDENTIFIER,"Variable"},
	}
	for i, tt := range tests {
		if LookupIdentifier(tt.Literal) != tt.Type {
			t.Fatalf("tests[%d] - type wrong. expected %q, got %q",i,LookupIdentifier(tt.Literal),tt.Type)
		}
	}
}
