package lexer

import (
	"regexp"
	"strconv"
)

type Lexer struct {
	Text   []byte
	cursor int
}

func New(text []byte) *Lexer {
	return &Lexer{
		Text:   text,
		cursor: 0,
	}
}

func (l *Lexer) hasMoreTokens() bool {
	return l.cursor < len(l.Text)
}

func (l *Lexer) isEOF() bool {
	return l.cursor == len(l.Text)
}

func (l *Lexer) _match(reg *regexp.Regexp, text []byte) []byte {
	if match := reg.Find(text); match != nil {
		l.cursor += len(match)
		return match
	}
	return nil
}

func (l *Lexer) GetNextToken() (*Token, error) {
	if !l.hasMoreTokens() {
		return nil, nil
	}

	_string := l.Text[l.cursor:]

	for k, token_type := range Spec {
		// todo: implement less greedy matching
		// fmt.Println("try:", k)
		reg := regexp.MustCompile(k)
		token_value := l._match(reg, _string)
		if token_value == nil {
			// fmt.Println("continue")
			continue
		}
		if token_type == NULL {
			// fmt.Println("null")
			return l.GetNextToken()
		}
		// fmt.Println("return")
		return &Token{
			Type:  token_type,
			Value: token_value,
		}, nil
	}

	panic("Unexpected token: " + strconv.Quote(string(_string[0])))
}
