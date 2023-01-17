package parser

import (
	"github.com/Urvirith/bearlang/src/lexer"
	"github.com/Urvirith/bearlang/src/token"
)

type Parser struct {
	lex *lexer.Lexer // Lexer pointer
	tok *token.Token // Current Token Available
}

func Init(buf string) *Parser {
	return &Parser{
		lex: lexer.Init(buf),
	}
}
