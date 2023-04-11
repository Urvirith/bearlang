package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// Token Enumeration
const (
	EOF     = "EOF"     // End Of File
	ILLEGAL = "ILLEGAL" // "ILLEGAL"
	COMMENT = "COMMENT" // "COMMENT"

	// Identifiers & Literals
	IDENT = "IDENTIFIER" // "IDENTIFIER" add x y etc...
	NUM   = "NUMBER"     // "GENERIC" Number Datatype

	// Number Declarations
	I8   = "I8"   // Signed Integer 8 Bit
	I16  = "I16"  // Signed Integer 16 Bit
	I32  = "I32"  // Signed Integer 32 Bit
	I64  = "I64"  // Signed Integer 64 Bit
	I128 = "I128" // Signed Integer 128 Bit
	U8   = "U8"   // Unsigned Integer 8 Bit
	U16  = "U16"  // Unsigned Integer 16 Bit
	U32  = "U32"  // Unsigned Integer 32 Bit
	U64  = "U64"  // Unsigned Integer 64 Bit
	U128 = "U128" // Unsigned Integer 128 Bit
	F32  = "F32"  // Float 32 Bit
	F64  = "F64"  // Float 64 Bit
	BOOL = "BOOL" // Boolean Datatype

	// Operators
	ASSIGN     = "="
	ADD        = "+"
	SUB        = "-"
	MUL        = "*"
	DIV        = "/"
	MOD        = "%"
	INC        = "++"
	DEC        = "--"
	ADD_ASSIGN = "+="
	SUB_ASSIGN = "-="
	MUL_ASSIGN = "*="
	DIV_ASSIGN = "/="

	// Bitwise Operators
	OR         = "|"
	OR_ASSIGN  = "|="
	AND        = "&"
	AND_ASSIGN = "&="
	XOR        = "^"
	XOR_ASSIGN = "^="
	NOT        = "!"
	COMP       = "~"
	LSHF       = "<<"
	RSHF       = ">>"

	// Comparators
	EQU  = "=="
	NEQ  = "!="
	GRT  = ">"
	LES  = "<"
	GEQ  = ">="
	LEQ  = "<="
	BOR  = "||"
	BAND = "&&"

	// Delimiters
	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACK   = "["
	RBRACK   = "]"
	COMMA    = ","
	FULLSTOP = "."
	COLON    = ":"
	SCOLON   = ";"

	// Keywords
	IMPORT   = "IMPORT" // Import
	FUNCTION = "FN"     // Function
	LET      = "LET"    // Variable
	VOLITILE = "VOL"    // Volitile
	STRUCT   = "STRUCT" // Structure
	ENUM     = "ENUM"   // Enumeration
	UNION    = "UNION"  // Union
	CONST    = "CONST"  // Constant
	RETURN   = "RETURN" // Return

	// Flow Control
	IF      = "IF"
	ELIF    = "ELIF"
	ELSE    = "ELSE"
	MATCH   = "MATCH"
	DEFAULT = "DEFAULT"
	FOR     = "FOR"
	LOOP    = "LOOP"
	WHILE   = "WHILE"

	// BINARY
	TRUE  = "TRUE"
	FALSE = "FALSE"
)

var keywords = map[string]TokenType{
	"fn":      FUNCTION,
	"let":     LET,
	"vol":     VOLITILE,
	"struct":  STRUCT,
	"enum":    ENUM,
	"union":   UNION,
	"const":   CONST,
	"return":  RETURN,
	"import":  IMPORT,
	"if":      IF,
	"elif":    ELIF,
	"else":    ELSE,
	"match":   MATCH,
	"default": DEFAULT,
	"for":     FOR,
	"loop":    LOOP,
	"while":   WHILE,
	"true":    TRUE,
	"false":   FALSE,
	"i8":      I8,
	"i16":     I16,
	"i32":     I32,
	"i64":     I64,
	"i128":    I128,
	"u8":      U8,
	"u16":     U16,
	"u32":     U32,
	"u64":     U64,
	"u128":    U128,
	"f32":     F32,
	"f64":     F64,
	"bool":    BOOL,
}

func LookupID(id string) TokenType {
	if tok, ok := keywords[id]; ok {
		return tok
	}

	return IDENT
}
