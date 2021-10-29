package parser

import (
	"fmt"
	"forth/lexer"
)

type ProgramNode struct {
	Type string 
	Statemets []interface{}
}

type IfNode struct {
	Type string
	Consequence []interface{}
}

type WhileNode struct {
	Type string 
	Consequence []interface{}
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
	ast := ProgramNode{Type:lexer.PROGRAM}

	for p.position < len(p.tokens) {
		token,err := p.ParseToken()
		if err {return ast, true}
		ast.Statemets = append(ast.Statemets, token)
		p.advance()
	}

	return ast, false
}

func (p *Parser) ParseToken() (interface{},bool) {
	var node interface{}
	var err bool

	switch p.token.Type {
		case lexer.IF:
			node, err = p.parseIfNode()
			if err {return node, true}
		case lexer.DO:
			node, err = p.parseWhileNode()
			if err {return node, true}
		default:
			node = p.token
	}

	return node, false
}

func (p *Parser) parseIfNode() (IfNode, bool) {
	node := IfNode{Type:lexer.IF}
	p.advance()

	for p.token.Type != lexer.THEN {
		subNode, err := p.ParseToken()
		if err {return node, true}
		node.Consequence = append(node.Consequence, subNode)

		if p.position >= len(p.tokens) {
			fmt.Println("Expected 'then' If statement")
			return node, true
		}

		p.advance()
	}
	
	return node, false
}

func (p *Parser) parseWhileNode() (WhileNode, bool) {
	node := WhileNode{Type:lexer.WHILE}
	p.advance()

	for p.token.Type != lexer.LOOP {
		subNode, err := p.ParseToken()
		if err {return node, true}
		node.Consequence = append(node.Consequence, subNode)

		if p.position >= len(p.tokens) {
			fmt.Println("Expected 'loop' DO statement")
			return node, true
		}

		p.advance()
	}
	
	return node, false
}