package main

import (
	"fmt"
	t "compiler-practice/tokenizer"
)

func main() {
	tokenizer := t.Tokenizer{}
	tokenizer.Initialize("./file.lang")
	tokenizer.Tokenize()
	fmt.Println(tokenizer.Tokens)
	
}