/*
* Top level public API for using tokenizers.
 */
package nlpt_tkz

import "github.com/jbowles/go_lexer"

type Digest struct {
	Tokens        []string
	TokenBytes    map[string][]byte
	Bytes         []byte
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

// ConcatByteSlice copy slice 2 into a slice 1 and returns a new slice.
// Used in cases where 'append' will not work. Also, want to pad the new slice
// with a space (' ') so that Digest.Bytes has padding between the concatenated bytes.
func ConcatByteSlice(slice1, slice2 []byte) []byte {
	new_slice := make([]byte, len(slice1)+len(slice2))
	//new_slice := []byte{32}
	copy(new_slice, slice1)
	copy(new_slice[len(slice1):], slice2)
	return new_slice
}
