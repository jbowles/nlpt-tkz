/*
* Top level public API for using tokenizers.
 */
package nlpt_tkz

import "github.com/jbowles/go_lexer"

const (
	BytesSpacePadding = byte('\n')
)

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

//TokenizeStr is a top-level function to choose tokenizer types: TknzStateFun, TknzUnicode, TknzWhiteSpace.
//It creates an abstraction around all the tokenizer implementations for a simple API, facilitating an easy call from the client and for binary installations.
func TokenizeStr(text, typ string) (tokens []string, digest *Digest) {

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

//TokenizeBytes is a top-level function to choose tokenizer types: TknzStateFunBytes, TknzUnicodeBytes.
//It creates an abstraction around tokenizer implementations that accept byte slices for a simple API, facilitating an easy call from the client and for binary installations.
func TokenizeBytes(textBytes []byte, typ string) (digest *Digest) {

	switch typ {
	case "lex":
		digest = TknzStateFunBytes(textBytes, NewStateFnDigest())
	case "unicode":
		digest = TknzUnicodeBytes(textBytes, NewUnicodeMatchDigest())
	default:
		panic("Tokenizer type not supported")
	}
	return
}

func ConcatByteSlice(slice1, slice2 []byte) []byte {
	new_slice := make([]byte, len(slice1)+len(slice2))
	copy(new_slice, slice1)
	copy(new_slice[len(slice1):], slice2)
	return new_slice
}
