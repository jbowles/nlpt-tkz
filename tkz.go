/*
* Top level public API for using tokenizers.
 */
package nlpt_tkz

import "github.com/jbowles/go_lexer"

type Digest struct {
	Tokens        []string
	TokenBytes    map[string][]byte
	SpaceCount    int
	CharCount     int
	Letter        []string
	Title         []string
	Number        []string
	Punct         []string
	Space         []string
	Symbol        []string
	TokenCount    int
	PunctCount    int
	LineCount     int
	EmptyLine     bool
	LastTokenType lexer.TokenType
}

//Tokenize is a top-level function to delegate Tokenizer implementation of Tknz().
//It creates an abstraction around all the tokenizer implementations for a simple API, facilitating an easy call from the client and for binary installations (i.e., tokens, digest := Tokenize("lext", "Simple sentence")).
func Tokenize(text, typ string) (tokens []string, digest *Digest) {

	switch typ {
	case "lex":
		tokens, digest = TknzStateFun(text, NewStateFnDigest())
	case "unicode":
		tokens, digest = TknzUnicode(text, NewUnicodeMatchDigest())
	case "whitespace":
		tokens, digest = TknzWhiteSpace(text, NewWhiteSpaceDigest())
	default:
		panic("Tokenizer type not supported")
	}
	return
}
