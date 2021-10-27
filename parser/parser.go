package parser

import (
	"forth/lexer"
)

type ProgramNode struct {
	statemets []interface{}
}

type IfNode struct {
	consequence []interface{}
}

type Parser struct {
	tokens []lexer.Token
	position int
	token lexer.Token
}

func (p *Parser) advance(){
	p.position++

	if p.position < len(p.tokens) {
		p.token = p.tokens[p.position]
	}
}

func New(tokens []lexer.Token) *Parser{
	p := &Parser{tokens: tokens, position: 0, token:tokens[0]}
	return p
}

func (p *Parser) Parse() (ProgramNode, bool) {
	var ast ProgramNode

	for p.position < len(p.tokens) {
		token,err := p.ParseToken()
		if err {return ast, true}
		ast.statemets = append(ast.statemets, token)
		p.advance()
	}

	return ast, false
}

func (p *Parser) ParseToken() (interface{},bool) {
	var node interface{}

	switch p.token.Type {
		case lexer.IF:
			node = IfNode{}
		default:
			node = p.token
	}

	return node, false
}