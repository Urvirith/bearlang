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
			lex.next()
		default:
			tok.Type = token.ASSIGN
		}
	case '+':
		switch lex.peek() {
		case '=':
			tok.Type = token.ADD_ASSIGN
			lex.next()
		case '+':
			tok.Type = token.INC
			lex.next()
		default:
			tok.Type = token.ADD
		}
	case '-':
		switch lex.peek() {
		case '=':
			tok.Type = token.SUB_ASSIGN
			lex.next()
		case '-':
			tok.Type = token.DEC
			lex.next()
		default:
			tok.Type = token.SUB
		}
	case '*':
		switch lex.peek() {
		case '=':
			tok.Type = token.MUL_ASSIGN
			lex.next()
		default:
			tok.Type = token.MUL
		}
	case '/':
		switch lex.peek() {
		case '=':
			tok.Type = token.DIV_ASSIGN
			lex.next()
		case '/':
			lex.consumeComment()
			tok.Type = token.COMMENT
		case '*':
			lex.consumeMultiLineComment()
			tok.Type = token.COMMENT
		default:
			tok.Type = token.DIV
		}
	case '%':
		tok.Type = token.MOD
	case '&':
		switch lex.peek() {
		case '=':
			tok.Type = token.AND_ASSIGN
			lex.next()
		case '&':
			tok.Type = token.BAND
			lex.next()
		default:
			tok.Type = token.AND
		}
	case '|':
		switch lex.peek() {
		case '=':
			tok.Type = token.OR_ASSIGN
			lex.next()
		case '|':
			tok.Type = token.BOR
			lex.next()
		default:
			tok.Type = token.OR
		}
	case '!':
		switch lex.peek() {
		case '=':
			tok.Type = token.NEQ
			lex.next()
		default:
			tok.Type = token.NOT
		}
	case '<':
		switch lex.peek() {
		case '=':
			tok.Type = token.LEQ
			lex.next()
		case '<':
			tok.Type = token.LSHF
			lex.next()
		default:
			tok.Type = token.LES
		}
	case '>':
		switch lex.peek() {
		case '=':
			tok.Type = token.GEQ
			lex.next()
		case '>':
			tok.Type = token.RSHF
			lex.next()
		default:
			tok.Type = token.GRT
		}
	case '~':
		tok.Type = token.COMP
	case '^':
		switch lex.peek() {
		case '=':
			tok.Type = token.XOR_ASSIGN
			lex.next()
		default:
			tok.Type = token.XOR
		}
	case '(':
		tok.Type = token.LPAREN
	case ')':
		tok.Type = token.RPAREN
	case '{':
		tok.Type = token.LBRACE
	case '}':
		tok.Type = token.RBRACE
	case '[':
		tok.Type = token.LBRACK
	case ']':
		tok.Type = token.RBRACK
	case ',':
		tok.Type = token.COMMA
	case '.':
		tok.Type = token.FULLSTOP
	case ':':
		tok.Type = token.COLON
	case ';':
		tok.Type = token.SCOLON
	case 0:
		tok.Type = token.EOF
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

	if tok.Type == token.IDENT {
		tok.Literal = lex.buf[pos:lex.pos]
	}

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
