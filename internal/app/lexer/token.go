package lexer

import (
	"fmt"
	"strconv"
)

type TokenType string

const (
	COMMENT    TokenType = "COMMENT"
	WORD       TokenType = "WORD"
	NUMBER     TokenType = "NUMBER"
	NULL       TokenType = "NULL"
	WHITESPACE TokenType = "WHITESPACE"
	NEXTLINE   TokenType = "NEXTLINE"
	TAB        TokenType = "TAB"
	EQUALS     TokenType = "EQUALS"
	START      TokenType = "START"
	END        TokenType = "END"
	COMPARISON TokenType = "COMPARISON"
)

var Spec = map[string]TokenType{
	`^[\<\>]=?`:                        COMPARISON,
	`^#.+`:                             COMMENT,
	`^"?[a-zA-Z_]+:?(\w+)?(\.?\d+)?"?`: WORD,
	`^-?\d+[\.,]?(\d?)+`:               NUMBER,
	`^ +`:                              NULL,
	`^\n+`:                             NULL,
	`^\t+`:                             NULL,
	`^==?`:                             EQUALS,
	`^{`:                               START,
	`^}`:                               END,
}

type Token struct {
	Type  TokenType `json:"type"`
	Value []byte    `json:"value"`
}

func (t Token) String() string {
	return fmt.Sprintf("type:\t%v,\tvalue:\t%v", t.Type, strconv.Quote(string(t.Value)))
}
