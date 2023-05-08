package route

type ObGeneratedColumnSimpleFunc interface {
	String() string
	getRefColumnNames() []string
}

type ObGeneratedColumnExpressParser struct {
	lexer ObGeneratedColumnLexer
}

type ObGeneratedColumnLexer struct {
	text      string
	ch        byte
	pos       int
	mark      int
	genToken  ObGeneratedColumnSimpleToken
	buf       []byte
	bufPos    int
	stringVal string
}

const (
	GenTokenComma  = ","
	GenTokenLparen = "("
	GenTokenRparen = ")"
	GenTokenSub    = "-"
	GenTokenPlus   = "+"
	// LITERAL_HEX, LITERAL_FLOAT, LITERAL_INT, IDENTIFIER, ERROR, EOF are no use
)

type ObGeneratedColumnSimpleToken struct {
	name string
}

func newEmptyObGeneratedColumnSimpleToken() ObGeneratedColumnSimpleToken {
	return ObGeneratedColumnSimpleToken{""}
}

func newObGeneratedColumnSimpleToken(name string) ObGeneratedColumnSimpleToken {
	return ObGeneratedColumnSimpleToken{name: name}
}
