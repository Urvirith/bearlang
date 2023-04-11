package lexer

import (
	"github.com/Urvirith/bearlang/src/token"
)

type Lexer struct {
	buf      string
	pos      int
	pos_read int
	ch       byte
}

func Init(buf string) *Lexer {
	lex := &Lexer{buf: buf}
	lex.next()
	return lex
}

// Scan the tokens
func (lex *Lexer) Scan() *token.Token {
	tok := &token.Token{Type: token.ILLEGAL, Literal: ""}
	lex.consume()

	// Skip whitepace
	switch lex.ch {
	case '=':
		switch lex.peek() {
		case '=':
			tok.Type = token.EQU
			tok.Literal = "=="
			lex.next()
		default:
			tok.Type = token.ASSIGN
			tok.Literal = "="
		}
	case '+':
		switch lex.peek() {
		case '=':
			tok.Type = token.ADD_ASSIGN
			tok.Literal = "+="
			lex.next()
		case '+':
			tok.Type = token.INC
			tok.Literal = "++"
			lex.next()
		default:
			tok.Type = token.ADD
			tok.Literal = "+"
		}
	case '-':
		switch lex.peek() {
		case '=':
			tok.Type = token.SUB_ASSIGN
			tok.Literal = "-="
			lex.next()
		case '-':
			tok.Type = token.DEC
			tok.Literal = "--"
			lex.next()
		default:
			tok.Type = token.SUB
			tok.Literal = "-"
		}
	case '*':
		switch lex.peek() {
		case '=':
			tok.Type = token.MUL_ASSIGN
			tok.Literal = "*="
			lex.next()
		default:
			tok.Type = token.MUL
			tok.Literal = "*"
		}
	case '/':
		switch lex.peek() {
		case '=':
			tok.Type = token.DIV_ASSIGN
			tok.Literal = "/="
			lex.next()
		case '/':
			lex.consumeComment()
			tok.Type = token.COMMENT
		case '*':
			lex.consumeMultiLineComment()
			tok.Type = token.COMMENT
		default:
			tok.Type = token.DIV
			tok.Literal = "/"
		}
	case '%':
		tok.Type = token.MOD
		tok.Literal = "%"
	case '&':
		switch lex.peek() {
		case '=':
			tok.Type = token.AND_ASSIGN
			tok.Literal = "&="
			lex.next()
		case '&':
			tok.Type = token.BAND
			tok.Literal = "&&"
			lex.next()
		default:
			tok.Type = token.AND
			tok.Literal = "&"
		}
	case '|':
		switch lex.peek() {
		case '=':
			tok.Type = token.OR_ASSIGN
			tok.Literal = "|="
			lex.next()
		case '|':
			tok.Type = token.BOR
			tok.Literal = "||"
			lex.next()
		default:
			tok.Type = token.OR
			tok.Literal = "|"
		}
	case '!':
		switch lex.peek() {
		case '=':
			tok.Type = token.NEQ
			tok.Literal = "!="
			lex.next()
		default:
			tok.Type = token.NOT
			tok.Literal = "!"
		}
	case '<':
		switch lex.peek() {
		case '=':
			tok.Type = token.LEQ
			tok.Literal = "<="
			lex.next()
		case '<':
			tok.Type = token.LSHF
			tok.Literal = "<<"
			lex.next()
		default:
			tok.Type = token.LES
			tok.Literal = "<"
		}
	case '>':
		switch lex.peek() {
		case '=':
			tok.Type = token.GEQ
			tok.Literal = ">="
			lex.next()
		case '>':
			tok.Type = token.RSHF
			tok.Literal = ">>"
			lex.next()
		default:
			tok.Type = token.GRT
			tok.Literal = ">"
		}
	case '~':
		tok.Type = token.COMP
		tok.Literal = "~"
	case '^':
		switch lex.peek() {
		case '=':
			tok.Type = token.XOR_ASSIGN
			tok.Literal = "^="
			lex.next()
		default:
			tok.Type = token.XOR
			tok.Literal = "^"
		}
	case '(':
		tok.Type = token.LPAREN
		tok.Literal = "("
	case ')':
		tok.Type = token.RPAREN
		tok.Literal = ")"
	case '{':
		tok.Type = token.LBRACE
		tok.Literal = "{"
	case '}':
		tok.Type = token.RBRACE
		tok.Literal = "}"
	case '[':
		tok.Type = token.LBRACK
		tok.Literal = "["
	case ']':
		tok.Type = token.RBRACK
		tok.Literal = "]"
	case ',':
		tok.Type = token.COMMA
		tok.Literal = ","
	case '.':
		tok.Type = token.FULLSTOP
		tok.Literal = "."
	case ':':
		tok.Type = token.COLON
		tok.Literal = ":"
	case ';':
		tok.Type = token.SCOLON
		tok.Literal = ";"
	case 0:
		tok.Type = token.EOF
		tok.Literal = "eof"
	default:
		if lex.isNum() {
			tok = lex.read_int()
			return tok
		} else if lex.isLetter() {
			tok = lex.read_id()
			return tok
		} else { // Unknown Token, Relay That Back
			println("Unknown Token at pos: {:?}, token is: {}", lex.pos, lex.ch)
			tok.Type = token.ILLEGAL
		}
	}

	lex.next()
	return tok
}

// Consume Unwanted Characters
func (lex *Lexer) consume() {
	for lex.ch == ' ' || lex.ch == '\t' || lex.ch == '\n' || lex.ch == '\r' || lex.ch == '\f' {
		lex.next()
	}
}

func (lex *Lexer) consumeComment() {
	for lex.ch != '\n' {
		lex.next()
	}
}

func (lex *Lexer) consumeMultiLineComment() {
	for lex.ch != 0 {
		if lex.ch == '*' && lex.peek() == '/' {
			lex.next()
			return
		}

		lex.next()
	}
}

// Get the next character from the input file
func (lex *Lexer) next() {
	// Get next character or at of buffer return 0
	if lex.pos_read < len(lex.buf) {
		lex.ch = lex.buf[lex.pos_read]
	} else {
		lex.ch = 0
	}

	lex.pos = lex.pos_read
	lex.pos_read += 1
}

// Peek Function - look at the next character
func (lex *Lexer) peek() byte {
	// Get next character or at of buffer return 0
	if lex.pos_read < len(lex.buf) {
		return lex.buf[lex.pos_read]
	} else {
		return 0
	}
}

// Read the identifier of the input string
func (lex *Lexer) read_id() *token.Token {
	pos := lex.pos
	tok := &token.Token{Type: token.ILLEGAL, Literal: ""}

	for lex.isLetter() || lex.isDigit() {
		lex.next()
	}

	tok.Type = token.LookupID(lex.buf[pos:lex.pos])
	tok.Literal = lex.buf[pos:lex.pos]

	return tok
}

// Scan and return an integer literal or float literal
func (lex *Lexer) read_int() *token.Token {
	pos := lex.pos

	for lex.isDigit() {
		lex.next()
	}

	return &token.Token{Type: token.NUM, Literal: lex.buf[pos:lex.pos]}
}

// Verify is letter
func (lex *Lexer) isLetter() bool {
	return 'a' <= lex.ch && lex.ch <= 'z' || 'A' <= lex.ch && lex.ch <= 'Z' || lex.ch == '_'
}

// Verify is number
func (lex *Lexer) isDigit() bool {
	return '0' <= lex.ch && lex.ch <= '9' || lex.ch == '.'
}

// Verify is number
func (lex *Lexer) isNum() bool {
	return '0' <= lex.ch && lex.ch <= '9'
}
