package lexer

import (
	"testing"

	"github.com/Urvirith/bearlang/src/token"
)

func TestTokens(t *testing.T) {
	input := `= + - * / % | & ! ~ ^ += -= *= /= ++ -- |= &= ^= << >> == != > < >= <= || && ( ) { } [ ] , . : ; import fn var vol struct enum union const return if elif else match default for loop while true false i8 i16 i32 i64 i128 u8 u16 u32 u64 u128 f32 f64 bool`

	tests := []struct {
		expectType    token.TokenType
		expectLiteral string
	}{
		{token.ASSIGN, ""},
		{token.ADD, ""},
		{token.SUB, ""},
		{token.MUL, ""},
		{token.DIV, ""},
		{token.MOD, ""},
		{token.OR, ""},
		{token.AND, ""},
		{token.NOT, ""},
		{token.COMP, ""},
		{token.XOR, ""},
		{token.ADD_ASSIGN, ""},
		{token.SUB_ASSIGN, ""},
		{token.MUL_ASSIGN, ""},
		{token.DIV_ASSIGN, ""},
		{token.INC, ""},
		{token.DEC, ""},
		{token.OR_ASSIGN, ""},
		{token.AND_ASSIGN, ""},
		{token.XOR_ASSIGN, ""},
		{token.LSHF, ""},
		{token.RSHF, ""},
		{token.EQU, ""},
		{token.NEQ, ""},
		{token.GRT, ""},
		{token.LES, ""},
		{token.GEQ, ""},
		{token.LEQ, ""},
		{token.BOR, ""},
		{token.BAND, ""},
		{token.LPAREN, ""},
		{token.RPAREN, ""},
		{token.LBRACE, ""},
		{token.RBRACE, ""},
		{token.LBRACK, ""},
		{token.RBRACK, ""},
		{token.COMMA, ""},
		{token.FULLSTOP, ""},
		{token.COLON, ""},
		{token.SCOLON, ""},
		{token.IMPORT, ""},
		{token.FUNCTION, ""},
		{token.VAR, ""},
		{token.VOLITILE, ""},
		{token.STRUCT, ""},
		{token.ENUM, ""},
		{token.UNION, ""},
		{token.CONST, ""},
		{token.RETURN, ""},
		{token.IF, ""},
		{token.ELIF, ""},
		{token.ELSE, ""},
		{token.MATCH, ""},
		{token.DEFAULT, ""},
		{token.FOR, ""},
		{token.LOOP, ""},
		{token.WHILE, ""},
		{token.TRUE, ""},
		{token.FALSE, ""},
		{token.I8, ""},
		{token.I16, ""},
		{token.I32, ""},
		{token.I64, ""},
		{token.I128, ""},
		{token.U8, ""},
		{token.U16, ""},
		{token.U32, ""},
		{token.U64, ""},
		{token.U128, ""},
		{token.F32, ""},
		{token.F64, ""},
		{token.BOOL, ""},
		{token.EOF, ""},
	}

	l := Init(input)

	for i, tt := range tests {
		tok := l.Scan()

		if tok.Type != tt.expectType {
			t.Fatalf("tests[%d] - tokentype wrong. expected: %q, got: %q", i, tt.expectType, tok.Type)
		}

		if tok.Literal != tt.expectLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected: %q, got: %q", i, tt.expectLiteral, tok.Literal)
		}
	}
}

func TestCode(t *testing.T) {
	input := `
		fn main(x: u32, y: u32) u32 {
			var float: f32 = 50.1;
			var woof: u32 = x * y;

			woof++;
			woof--;
			
			var arf = 5;
			arf *= woof;

			return arf;
		}
	`

	tests := []struct {
		expectType    token.TokenType
		expectLiteral string
	}{
		{token.FUNCTION, ""},
		{token.IDENT, "main"},
		{token.LPAREN, ""},
		{token.IDENT, "x"},
		{token.COLON, ""},
		{token.U32, ""},
		{token.COMMA, ""},
		{token.IDENT, "y"},
		{token.COLON, ""},
		{token.U32, ""},
		{token.RPAREN, ""},
		{token.U32, ""},
		{token.LBRACE, ""},
		{token.VAR, ""},
		{token.IDENT, "float"},
		{token.COLON, ""},
		{token.F32, ""},
		{token.ASSIGN, ""},
		{token.NUM, "50.1"},
		{token.SCOLON, ""},
		{token.VAR, ""},
		{token.IDENT, "woof"},
		{token.COLON, ""},
		{token.U32, ""},
		{token.ASSIGN, ""},
		{token.IDENT, "x"},
		{token.MUL, ""},
		{token.IDENT, "y"},
		{token.SCOLON, ""},
		{token.IDENT, "woof"},
		{token.INC, ""},
		{token.SCOLON, ""},
		{token.IDENT, "woof"},
		{token.DEC, ""},
		{token.SCOLON, ""},
		{token.VAR, ""},
		{token.IDENT, "arf"},
		{token.ASSIGN, ""},
		{token.NUM, "5"},
		{token.SCOLON, ""},
		{token.IDENT, "arf"},
		{token.MUL_ASSIGN, ""},
		{token.IDENT, "woof"},
		{token.SCOLON, ""},
		{token.RETURN, ""},
		{token.IDENT, "arf"},
		{token.SCOLON, ""},
		{token.RBRACE, ""},
		{token.EOF, ""},
	}

	l := Init(input)

	for i, tt := range tests {
		tok := l.Scan()

		if tok.Type != tt.expectType {
			t.Fatalf("tests[%d] - tokentype wrong. expected: %q, got: %q", i, tt.expectType, tok.Type)
		}

		if tok.Literal != tt.expectLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected: %q, got: %q", i, tt.expectLiteral, tok.Literal)
		}
	}
}

func TestComments(t *testing.T) {
	input := `
		/*
			This is a multiline comment
		*/
		fn main(x: u32, y: u32) u32 {
			// Float Assigned
			var float: f32 = 50.1;
			// Woof Assigned x * y;
			var woof: u32 = x * y;

			// Increment Woof
			woof++;
			// Decrement Woof
			woof--;
			
			// Arf Assigned
			var arf: u32 = 5;
			// Mul Assign Arf by Woof
			arf *= woof;

			//Return Woof
			return arf;
		} // End Main
	`

	tests := []struct {
		expectType    token.TokenType
		expectLiteral string
	}{
		{token.COMMENT, ""},
		{token.FUNCTION, ""},
		{token.IDENT, "main"},
		{token.LPAREN, ""},
		{token.IDENT, "x"},
		{token.COLON, ""},
		{token.U32, ""},
		{token.COMMA, ""},
		{token.IDENT, "y"},
		{token.COLON, ""},
		{token.U32, ""},
		{token.RPAREN, ""},
		{token.U32, ""},
		{token.LBRACE, ""},
		{token.COMMENT, ""},
		{token.VAR, ""},
		{token.IDENT, "float"},
		{token.COLON, ""},
		{token.F32, ""},
		{token.ASSIGN, ""},
		{token.NUM, "50.1"},
		{token.SCOLON, ""},
		{token.COMMENT, ""},
		{token.VAR, ""},
		{token.IDENT, "woof"},
		{token.COLON, ""},
		{token.U32, ""},
		{token.ASSIGN, ""},
		{token.IDENT, "x"},
		{token.MUL, ""},
		{token.IDENT, "y"},
		{token.SCOLON, ""},
		{token.COMMENT, ""},
		{token.IDENT, "woof"},
		{token.INC, ""},
		{token.SCOLON, ""},
		{token.COMMENT, ""},
		{token.IDENT, "woof"},
		{token.DEC, ""},
		{token.SCOLON, ""},
		{token.COMMENT, ""},
		{token.VAR, ""},
		{token.IDENT, "arf"},
		{token.COLON, ""},
		{token.U32, ""},
		{token.ASSIGN, ""},
		{token.NUM, "5"},
		{token.SCOLON, ""},
		{token.COMMENT, ""},
		{token.IDENT, "arf"},
		{token.MUL_ASSIGN, ""},
		{token.IDENT, "woof"},
		{token.SCOLON, ""},
		{token.COMMENT, ""},
		{token.RETURN, ""},
		{token.IDENT, "arf"},
		{token.SCOLON, ""},
		{token.RBRACE, ""},
		{token.COMMENT, ""},
		{token.EOF, ""},
	}

	l := Init(input)

	for i, tt := range tests {
		tok := l.Scan()

		if tok.Type != tt.expectType {
			t.Fatalf("tests[%d] - tokentype wrong. expected: %q, got: %q", i, tt.expectType, tok.Type)
		}

		if tok.Literal != tt.expectLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected: %q, got: %q", i, tt.expectLiteral, tok.Literal)
		}
	}
}
