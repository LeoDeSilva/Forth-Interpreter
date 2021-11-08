package evaluator

import (
	"fmt"
	"forth/lexer"
	"forth/parser"
	"strconv"
)

type Environment struct {
	Variables map[int]int
	Identifiers map[string]int
	Stack []interface{}
}

type Evaluator struct{
	ast parser.ProgramNode
	environment *Environment
	position int 
	node interface{}
}

func New(ast parser.ProgramNode, environment *Environment) *Evaluator {
	e := &Evaluator{ast: ast, position: 0, node:ast.Statements[0], environment: environment}
	return e
}

func (e *Evaluator) advance(){
	e.position++

	if e.position < len(e.ast.Statements) {
		e.node = e.ast.Statements[e.position]
	}
}

func pop(array []interface{}) ([]interface{}, interface{}) {
	node := array[len(array) - 1]
	return array[:len(array) - 1], node
}

func (e *Evaluator) Evaluate() bool{
	for e.position < len(e.ast.Statements) {
		err := e.Eval()
		if err {return true}
		e.advance()
	}
	return false
}

func (e *Evaluator) Eval() bool {
	switch node := e.node.(type) {
	case parser.ProgramNode:
		for e.position < len(e.ast.Statements) {
			err := e.Eval()
			if err {return true}
			e.advance()
		}
		return false
	case parser.WhileNode:
		fmt.Println(node,"WHILE")
	case parser.IfNode:
		fmt.Println(node,"IF")
	case parser.FunctionNode:
		fmt.Println(node,"FUNCTION")
	case parser.VariableNode:
		id := len(e.environment.Variables)
		e.environment.Variables[id] = 0
		e.environment.Identifiers[node.Identifier] = id
	case lexer.Token:
		err := e.EvalToken(node)
		if err { return true }
		return false
	}
  return false
}

func (e *Evaluator) EvalToken(token lexer.Token) bool {
	switch token.Type {
	case lexer.INT:
		value, _ := strconv.Atoi(token.Literal)
		e.environment.Stack = append(e.environment.Stack, value)
	case lexer.STRING:
		e.environment.Stack = append(e.environment.Stack, token.Literal)
	case lexer.IDENTIFIER:
		e.environment.Stack = append(e.environment.Stack, e.environment.Identifiers[token.Literal])

	case lexer.AT:
		node := e.environment.Stack[len(e.environment.Stack) - 1]
		e.environment.Stack,_ = pop(e.environment.Stack)
		e.environment.Stack = append(e.environment.Stack, e.environment.Variables[node.(int)])

	case lexer.DOT:
		fmt.Println(e.environment.Stack[len(e.environment.Stack) - 1])
		e.environment.Stack, _ = pop(e.environment.Stack)

	case lexer.NOT:
		var id, value interface{}
		e.environment.Stack, id = pop(e.environment.Stack)
		e.environment.Stack, value = pop(e.environment.Stack)
		e.environment.Variables[id.(int)] = value.(int)

	case lexer.DROP:
		if len(e.environment.Stack) < 1 {
			fmt.Println("Underflow Error")
			return true
		}

		e.environment.Stack, _ = pop(e.environment.Stack)

	case lexer.NIP:
		if len(e.environment.Stack) < 2 {
			fmt.Println("Underflow Error")
			return true
		}

		e.environment.Stack = append(
			e.environment.Stack[:len(e.environment.Stack) - 2], 
			e.environment.Stack[len(e.environment.Stack) -1:]...
		)

	case lexer.SWAP:
		if len(e.environment.Stack) < 2 {
			fmt.Println("Underflow Error")
			return true
		}

		node := e.environment.Stack[len(e.environment.Stack) - 2]
		e.environment.Stack = append(
			e.environment.Stack[:len(e.environment.Stack) - 2], 
			e.environment.Stack[len(e.environment.Stack) -1:]...
		)
		e.environment.Stack = append(e.environment.Stack, node)

	case lexer.ROT:
		if len(e.environment.Stack) < 3 {
			fmt.Println("Underflow Error")
			return true
		}
		node := e.environment.Stack[len(e.environment.Stack) - 3]
		e.environment.Stack = append(
			e.environment.Stack[:len(e.environment.Stack) - 3], 
			e.environment.Stack[len(e.environment.Stack) -2:]...
		)
		e.environment.Stack = append(e.environment.Stack, node)

	case lexer.OVER:
		if len(e.environment.Stack) < 2 {
			fmt.Println("Underflow Error")
			return true
		}
		node := e.environment.Stack[len(e.environment.Stack) - 2]
		e.environment.Stack = append(e.environment.Stack, node)

	case lexer.DUP:
		e.environment.Stack = append(e.environment.Stack, e.environment.Stack[len(e.environment.Stack) - 1])

	case lexer.ADD,lexer.SUB,lexer.MUL,lexer.DIV,lexer.MOD:
		err := e.arith()
		if err {return true}

	case lexer.ADD_EQ:
		err := e.EQarith()
		if err{return true}	
	}

	return false
}

func (e *Evaluator) arith() bool {
	node := e.node.(lexer.Token)
	var first,second interface{}

	e.environment.Stack, second = pop(e.environment.Stack)
	e.environment.Stack, first = pop(e.environment.Stack)

	switch node.Type{
	case lexer.ADD:
		e.environment.Stack = append(e.environment.Stack, first.(int)+second.(int))
	case lexer.SUB:
		e.environment.Stack = append(e.environment.Stack, first.(int)-second.(int))
	case lexer.MUL:
		e.environment.Stack = append(e.environment.Stack, first.(int)*second.(int))
	case lexer.DIV:
		e.environment.Stack = append(e.environment.Stack, first.(int)/second.(int))
	case lexer.MOD:
		e.environment.Stack = append(e.environment.Stack, first.(int)%second.(int))
	}
	return false
}

func (e *Evaluator) EQarith() bool {
	node := e.node.(lexer.Token)

	var id, value interface{}
	var result int
	e.environment.Stack, id = pop(e.environment.Stack)
	e.environment.Stack, value = pop(e.environment.Stack)

	switch node.Type{
	case lexer.ADD_EQ:
		result = value.(int) + e.environment.Variables[id.(int)]
	case lexer.SUB_EQ:
		result = value.(int) - e.environment.Variables[id.(int)]
	case lexer.MUL_EQ:
		result = value.(int) * e.environment.Variables[id.(int)]
	case lexer.DIV_EQ:
		result = value.(int) / e.environment.Variables[id.(int)]
	}
	e.environment.Variables[id.(int)] = result
	return false
}