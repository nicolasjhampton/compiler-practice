package parser

import (
	"fmt"
	t "compiler-practice/tokenizer"
	"strconv"
)

type Parser struct {
	Tokens []t.Token
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type DefNode struct {
	Name string
	ArgNames []string
	Body IntegerNode
}

type IntegerNode struct {
	Value int
}

func (p *Parser) Parse() *DefNode {
	return p.ParseDef()
}

func (p *Parser) ParseDef() *DefNode {
	name := p.Consume("def")
	identity := p.Consume("identifier")
	args := p.ParseArgNames()
	body := p.ParseExpr()
	end := p.Consume("end")
	dnode := DefNode{ Name: identity.Value, ArgNames: args, Body: *body}
	return &dnode
}

func (p *Parser) ParseExpr() *IntegerNode {
	return p.ParseInteger()
}

func (p *Parser) ParseInteger() *IntegerNode {
	intr := p.Consume("integer")
	num, err := strconv.Atoi(intr.Value)
	check(err)
	iNode := IntegerNode{ Value: num }
	return &iNode
}

func (p *Parser) ParseArgNames() []string {
	_ = p.Consume("oparen")
	// Arguments here
	_ = p.Consume("cparen")
	return []string{}
}

func (p *Parser) Consume(typeExpected string) *t.Token {
	token := p.Tokens[0]
	p.Tokens = p.Tokens[1:]
	if token.Type != typeExpected {
		panic(fmt.Sprintf("Expected token type %s but got %s", typeExpected, token.Type))
	} 
	return &token
}