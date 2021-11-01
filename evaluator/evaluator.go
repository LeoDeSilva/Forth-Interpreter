package evaluator

import (
	"fmt"
	"forth/lexer"
	"forth/parser"
)

type Environment struct {
	Variables map[string]interface{}
}

type Evaluator struct{
	ast parser.ProgramNode
	environment *Environment
	position int 
	node interface{}
}

func New(ast parser.ProgramNode, environment *Environment) *Evaluator {
	e := &Evaluator{ast: ast, position: 0, node:ast.Statemets[0], environment: environment}
	return e
}

func (e *Evaluator) advance(){
	e.position++

	if e.position < len(e.ast.Statemets) {
		e.node = e.ast.Statemets[e.position]
	}
}

func (e *Evaluator) Evaluate() bool{
	for e.position < len(e.ast.Statemets) {
		e.Eval()
		e.advance()
	}
	return false
}

func (e *Evaluator) Eval() {
	switch node := e.node.(type) {
	case parser.WhileNode:
		fmt.Println(node,"WHILE")
	case parser.IfNode:
		fmt.Println(node,"IF")
	case lexer.Token:
		fmt.Println(node,"Token")
			 
	}
}

