package tokenizer

import (
	io "io/ioutil"
	r "regexp"
	s "strings"
	"errors"
)

var tokenTypes = [...][2]string{
	[2]string{"def", "\\bdef\\b"},
	[2]string{"end", "\\bend\\b"},
	[2]string{"identifier", "\\b[a-zA-Z]+\\b"},
	[2]string{"integer", "\\b[0-9]+\\b"},
	[2]string{"oparen", "\\("},
	[2]string{"cparen", "\\)"},
	[2]string{"comma", "\\,"},
}

type Tokenizer struct {
	file   string
	Tokens []Token
}

type Token struct {
	Type  string
	Value string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (t *Tokenizer) Initialize(fileName string) {
	dat, err := io.ReadFile(fileName)
	check(err)
	t.file = s.TrimSpace(string(dat))
}

func (t *Tokenizer) Tokenize() []Token {
	for ;len(t.file) != 0; {
		token, err := t.FindNextToken()
		check(err)
		t.Tokens = append(t.Tokens, *token)
	}
	return t.Tokens
}

func (t *Tokenizer) FindNextToken() (*Token, error) {
	for _, reg := range tokenTypes {
		pattern := "\\A(" + reg[1] + ")"
		exp := r.MustCompile(pattern)
		loc := exp.FindStringIndex(t.file)
		if loc != nil {
			value := t.file[loc[0]:loc[1]]
			t.file = s.TrimSpace(t.file[loc[1]:])
			token := Token{Type: reg[0], Value: value}
			return &token, nil
		}
	}
	return nil, errors.New("Couldn't match token on current text in file")
}
