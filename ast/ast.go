package ast

import "github.com/ValeryPiashchynski/InterpreterInGo/token"

//All Nodes connected to each other
type Node interface {
	TokenLiteral() string
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

//The root node of every AST our parser produces
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token
	//To hold the identifier
	Name *Identifier
	//For the expression that produces value
	Value Expression
}

func (ls *LetStatement) statementNode() {

}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token //The token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {

}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
