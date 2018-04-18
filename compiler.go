package main

import (
	t "compiler-practice/tokenizer"
	p "compiler-practice/parser"
	g "compiler-practice/generator"
	// "reflect"
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	tokenizer := t.Tokenizer{}
	tokenizer.Initialize("./file.lang")
	tokens := tokenizer.Tokenize()
	parser := p.Parser{ Tokens: tokens }
	tree := parser.Parse()
	
	// val := reflect.Indirect(reflect.ValueOf(tree))
	// fmt.Println(*tree)
	//generator := g.Generator{ Tree: tree }
	text := g.Generate(tree)
	fmt.Println(text)
	
}
