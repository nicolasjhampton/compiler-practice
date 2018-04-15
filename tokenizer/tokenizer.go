package tokenizer

import (
	"fmt"
	io "io/ioutil"
	r "regexp"
	s "strings"
)


var tokenTypes = [...]string{
	"\bdef\b",
	"\bend\b",
	"[a-zA-Z]+",
	"[0-9]+",
	"\\(",
	"\\)",
}

type Tokenizer struct {
	file string
	Tokens []string
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func (t *Tokenizer) Initialize(fileName string) {
	dat, err := io.ReadFile(fileName)
    check(err)
	t.file = string(dat)
}

func (t *Tokenizer) Tokenize() {
	if len(t.file) == 0 {
		return 
	}
	for _, reg := range tokenTypes {
		exp := r.MustCompile(reg)
		t.file = s.TrimSpace(t.file)
		loc := exp.FindStringIndex(t.file)
		if loc != nil && loc[0] == 0 {
			fmt.Println(t.file[loc[0]:loc[1]])
			t.Tokens = append(t.Tokens, t.file[loc[0]:loc[1]])
			t.file = t.file[loc[1]:]
			t.Tokenize()
		}
	}
}
