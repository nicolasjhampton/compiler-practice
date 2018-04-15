package main

import (
	t "compiler-practice/tokenizer"
	p "compiler-practice/parser"
	//"fmt"
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
	parser.Parse()
	//fmt.Println(tokens)
}
