/*
* Top level public API for using tokenizers.
 */
package nlpt_tkz

type Tokenizer interface {
	Tknz(text string) ([]string, interface{})
}

//Tokenize is a top-level function to delegate Tokenizer implementation of Tknz().
//It creates an abstraction around all the tokenizer implementations for a simple API, facilitating an easy call from the client and for binary installations (i.e., tokens, digest := Tokenize("lext", "Simple sentence")).
func Tokenize(typ, text string) (tokens []string, digest interface{}) {
	switch typ {
	case "lex":
		sdigest := NewStateFnDigest()
		tokens, digest = sdigest.Tknz(text)
	case "bukt":
		udigest := NewUnicodeMatchDigest()
		tokens, digest = udigest.Tknz(text)
	default:
		wdigest := NewWhiteSpaceDigest()
		tokens, digest = wdigest.Tknz(text)
	}
	return
}
