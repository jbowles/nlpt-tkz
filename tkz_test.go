package nlpt_tkz

import (
	"fmt"
	"testing"
)

/*
BENCHMARKS: go test -bench=.
*/

func BenchmarkLex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Tokenize("lex", ThoreauOne)
	}
}

func BenchmarkBukt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Tokenize("bukt", ThoreauOne)
	}
}

func BenchmarkSpace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Tokenize("space", ThoreauOne)
	}
}

func TestTokenizeLexOption(t *testing.T) {
	tokens, digest := Tokenize("lex", ThoreauThree)
	//fmt.Printf("Tokens = %v\n DigestType = %T\n", tokens, digest)
	//fmt.Printf("DIGEST %v", digest)

	if len(tokens) != 19 {
		t.Log("Expected thoreauThree to be length=19, got=", len(tokens))
		t.Fail()
	}

	typ := fmt.Sprintf("%T", digest)
	if typ != "*nlpt_tkz.StateFnDigest" {
		t.Log("Expected digest to be *nlpt_tkz.StateFnDigest", typ)
		t.Fail()
	}
}

func TestTokenizeDefaultOption(t *testing.T) {
	tokens, digest := Tokenize("lexer", ThoreauThree)
	//fmt.Printf("Tokens = %v\n DigestType = %T\n", tokens, digest)
	//fmt.Printf("DIGEST %v", digest)

	if len(tokens) != 19 {
		t.Log("Expected thoreauThree to be length=19, got=", len(tokens))
		t.Fail()
	}

	typ := fmt.Sprintf("%T", digest)
	if typ != "*nlpt_tkz.WhiteSpaceDigest" {
		t.Log("Expected digest to be *nlpt_tkz.WhiteSpaceDigest", typ)
		t.Fail()
	}
}
