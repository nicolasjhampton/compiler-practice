package main

import (
	t "compiler-practice/tokenizer"
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
	tokenizer.Tokenize()
	fmt.Println(tokenizer.Tokens)
}
