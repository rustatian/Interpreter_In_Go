package parser

import (
	"github.com/ValeryPiashchynski/InterpreterInGo/lexer"
	"github.com/ValeryPiashchynski/InterpreterInGo/token"
)

type Parser struct {
	//Pointer to the instance
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	//Read two tokens, so curToken and peekToke are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}
