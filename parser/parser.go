package parser

import (
	"fmt"
	"forth/lexer"
)

type ProgramNode struct {
	Type string 
	Statements []interface{}
}

type IfNode struct {
	Type string
	Consequence ProgramNode
}

type WhileNode struct {
	Type string 
	Consequence ProgramNode
}

type FunctionNode struct {
	Type string 
	Identifier string
	Program ProgramNode
}

type VariableNode struct {
	Type string 
	Identifier string
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
		ast.Statements = append(ast.Statements, token)
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
		case lexer.COLON:
			node, err = p.parseFunctionNode()
			if err {return node, true}
		case lexer.VARIABLE:
			node,err = p.parseVariableNode()
			if err {return node, true}
		case lexer.ILLEGAL:
			fmt.Println("UNEXPECTED TOKEN")
			return node, true
		default:
			node = p.token
	}

	return node, false
}

func (p *Parser) parseVariableNode() (VariableNode, bool){
	p.advance()
	if p.token.Type != lexer.IDENTIFIER{ return VariableNode{}, true}
	identifier := p.token.Literal
	return VariableNode{lexer.VARIABLE, identifier}, false
	
}

func (p *Parser) parseIfNode() (IfNode, bool) {
	node := IfNode{Type:lexer.IF}
	p.advance()

	for p.token.Type != lexer.THEN {
		subNode, err := p.ParseToken()
		if err {return node, true}
		node.Consequence.Statements = append(node.Consequence.Statements, subNode)

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
		node.Consequence.Statements = append(node.Consequence.Statements, subNode)

		if p.position >= len(p.tokens) {
			fmt.Println("Expected 'loop' DO statement")
			return node, true
		}

		p.advance()
	}
	
	return node, false
}

func (p *Parser) parseFunctionNode() (FunctionNode, bool) {
	node := FunctionNode{Type:lexer.FUNCTION, Program: ProgramNode{}}	
	p.advance()

	node.Identifier = p.token.Literal
	p.advance()

	for p.token.Type != lexer.SEMICOLON {
		subNode, err := p.ParseToken()
		if err {return node, true}

		node.Program.Statements = append(node.Program.Statements, subNode)

		if p.position >= len(p.tokens) {
			return node, true
		}
		p.advance()
	}

	return node, false
}