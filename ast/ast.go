package ast

import (
	"bytes"
	"github.com/ValeryPiashchynski/InterpreterInGo/token"
)

//All Nodes connected to each other
type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	//To provide a token literal
	Node
	//To parse statements
	statementNode()
}

type Expression interface {
	Node
	//To parse expressions
	expressionNode()
}

type Program struct {
	Statements []Statement
}

type ReturnStatement struct {
	Token       token.Token //the 'return' token
	ReturnValue Expression
}

type LetStatement struct {
	Token token.Token
	//To hold the identifier
	Name *Identifier
	//For the expression that produces value
	Value Expression
}

type Identifier struct {
	Token token.Token //The token.IDENT token
	Value string
}

type ExpressionStatement struct {
	Token      token.Token //the first token of the expression
	Expression Expression
}

//Expression statements
func (es *ExpressionStatement) statementNode() {

}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

//Return statements
func (rs *ReturnStatement) statementNode() {

}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

//Let statements
func (ls *LetStatement) statementNode() {

}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

//Identifier
func (i *Identifier) expressionNode() {

}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

//The root node of every AST our parser produces
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

//Create a buffer and write the return value of each statement to it, return buffer as s string
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}
