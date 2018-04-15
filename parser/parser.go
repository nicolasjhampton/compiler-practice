package parser

import (
	"fmt"
	t "compiler-practice/tokenizer"
	"errors"
)

type Parser struct {
	Tokens []t.Token
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// type AST struct {

// }

func (p *Parser) Parse() {
	p.ParseDef()
}

func (p *Parser) ParseDef() {
	name, err := p.Consume("def")
	identity, err2 := p.Consume("identifier")
	_, err3 := p.ParseArgNames()
	body, err4 := p.ParseExpr()
	end, err5 := p.Consume("end")
	check(err)
	check(err2)
	check(err3)
	check(err4)
	check(err5)
	fmt.Println(*name)
	fmt.Println(*identity)

	fmt.Println(*body)
	fmt.Println(*end)
}

func (p *Parser) ParseExpr() (*t.Token, error) {
	return p.ParseInteger()
}

func (p *Parser) ParseInteger() (*t.Token, error) {
	intr, err := p.Consume("integer")
	return intr, err
}

func (p *Parser) ParseArgNames() (*t.Token, error) {
	_, err := p.Consume("oparen")
	check(err)
	// Arguments here
	_, err2 := p.Consume("cparen")
	check(err2)
	return nil, nil
}

func (p *Parser) Consume(typeExpected string) (*t.Token, error) {
	token := p.Tokens[0]
	p.Tokens = p.Tokens[1:]
	if token.Type == typeExpected {
		return &token, nil
	} 
	return nil, errors.New(fmt.Sprintf("Expected token type %s but got %s", typeExpected, token.Type))
}