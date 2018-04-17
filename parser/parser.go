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
	Name Node
	ArgNames []Node
	Body Node
}

func (d DefNode) isNode() int {
	return 1
}

type CallNode struct {
	Name Node
	ArgExprs []Node
}

func (c CallNode) isNode() int {
	return 1
}

type IdentifierNode struct {
	Name string
	Ref Node
}

func (i IdentifierNode) isNode() int {
	return 1
}

type IntegerNode struct {
	Value int
}

func (i IntegerNode) isNode() int {
	return 1
}

type Node interface {
	isNode() int
}

type nodeFunc func(*Parser) Node

func (p *Parser) Parse() *DefNode {
	return p.ParseDef()
}

func (p *Parser) ParseDef() *DefNode {
	_ = p.Consume("def")
	identity := ParseIdentifier(p)
	fmt.Println(p.Tokens)
	args := p.ParseArgs(ParseIdentifier) //p.ParseArgNames()
	body := ParseExpr(p)
	_ = p.Consume("end")
	dnode := DefNode{ Name: identity, ArgNames: args, Body: body }
	return &dnode
}

func ParseExpr(p *Parser) Node {
	if p.Peek("integer", 0) {
		return ParseInteger(p)
	} else if p.Peek("identifier", 0) && p.Peek("oparen", 1) {
		fmt.Println(p.Tokens)
		return ParseCall(p)
	} else if p.Peek("identifier", 0) {
		return ParseIdentifier(p);
	}
	panic("No known token found in function body.")
}

func ParseCall(p *Parser) CallNode {
	identity := ParseIdentifier(p)
	argExprs := p.ParseArgs(ParseExpr) //p.ParseArgExprs();
	cNode := CallNode{ Name: identity, ArgExprs: argExprs }
	return cNode
}

func (p *Parser) ParseArgs(fn nodeFunc) []Node {
	args := []Node{}
	fmt.Println(p.Tokens)
	_ = p.Consume("oparen")
	for ;!p.Peek("cparen", 0); {
		args = append(args, fn(p))
		if p.Peek("comma", 0) {
			_ = p.Consume("comma")
		}
	}
	_ = p.Consume("cparen")
	return args
}

// func (p *Parser) ParseArgExprs(fn nodeFunc) []Node {
// 	args := []Node{}
// 	_ = p.Consume("oparen")
// 	for ;p.Peek("identifier", 0); {
// 		args = append(args, p.ParseExpr())
// 		if p.Peek("comma", 0) {
// 			_ = p.Consume("comma")
// 		}
// 	}
// 	_ = p.Consume("cparen")
// 	return args
// }

// func (p *Parser) ParseArgNames() []IdentifierNode {
// 	args := []IdentifierNode{}
// 	_ = p.Consume("oparen")
// 	for ;p.Peek("identifier", 0); {
// 		args = append(args, p.ParseIdentifier())
// 		if p.Peek("comma", 0) {
// 			_ = p.Consume("comma")
// 		}
// 	}
// 	_ = p.Consume("cparen")
// 	return args
// }

func ParseIdentifier(p *Parser) Node {
	identity := p.Consume("identifier")
	iNode := IdentifierNode{ Name: identity.Value }
	return iNode
}

func ParseInteger(p *Parser) IntegerNode {
	intr := p.Consume("integer")
	num, err := strconv.Atoi(intr.Value)
	check(err)
	iNode := IntegerNode{ Value: num }
	return iNode
}

func (p *Parser) Peek(typeExpected string, offset int) bool {
	return p.Tokens[offset].Type == typeExpected
}

func (p *Parser) Consume(typeExpected string) *t.Token {
	token := p.Tokens[0]
	p.Tokens = p.Tokens[1:]
	if token.Type != typeExpected {
		panic(fmt.Sprintf("Expected token type %s but got %s", typeExpected, token.Type))
	} 
	return &token
}