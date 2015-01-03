package nlpt_tkz

import (
	"fmt"
	"testing"
)

/*
BENCHMARKS: go test -bench=.
*/

func BenchmarkLexStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TokenizeStr(ThoreauThree, "lex")
	}
}

func BenchmarkUnicodeStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TokenizeStr(ThoreauThree, "unicode")
	}
}

func BenchmarkWhitespaceStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TokenizeStr(ThoreauThree, "whitespace")
	}
}

func BenchmarkLexBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TokenizeBytes(ththreeeByte, "lex")
	}
}

func BenchmarkUnicodeBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TokenizeBytes(ththreeeByte, "unicode")
	}
}

func TestTokenizeLexOption(t *testing.T) {
	tokens, digest := TokenizeStr(ThoreauThree, "lex")
	//fmt.Printf("LEX token bytes %v\n", digest.TokenBytes)
	//fmt.Printf("LEX bytes %v\n", digest.Bytes)
	//fmt.Printf("LEX bytes stringified %v\n", string(digest.Bytes))

	if len(tokens) != 19 {
		t.Log("Expected thoreauThree to be length=19, got=", len(tokens))
		t.Fail()
	}

	typ := fmt.Sprintf("%T", digest)
	if typ != "*nlpt_tkz.Digest" {
		t.Log("Expected digest to be *nlpt_tkz.StateFnDigest", typ)
		t.Fail()
	}
}

func TestTokenizeLexOptionForBytes(t *testing.T) {
	tokens, digest := TokenizeStr(ThoreauThree, "lex")
	//fmt.Printf("LEX token bytes %v\n", digest.TokenBytes)
	//fmt.Printf("LEX bytes %v\n", digest.Bytes)
	//fmt.Printf("LEX bytes stringified %v\n", string(digest.Bytes))

	if len(tokens) != 19 {
		t.Log("Expected thoreauThree to be length=19, got=", len(tokens))
		t.Fail()
	}

	typ := fmt.Sprintf("%T", digest)
	if typ != "*nlpt_tkz.Digest" {
		t.Log("Expected digest to be *nlpt_tkz.StateFnDigest", typ)
		t.Fail()
	}
}

func TestTokenizeUnicodeMatchOption(t *testing.T) {
	tokens, digest := TokenizeStr(ThoreauThree, "unicode")

	if len(tokens) != 19 {
		t.Log("Expected thoreauThree to be length=19, got=", len(tokens))
		t.Fail()
	}

	typ := fmt.Sprintf("%T", digest)
	if typ != "*nlpt_tkz.Digest" {
		t.Log("Expected digest to be *nlpt_tkz.StateFnDigest", typ)
		t.Fail()
	}
}

func TestTokenizeDefaultOption(t *testing.T) {
	tokens, digest := TokenizeStr(ThoreauThree, "whitespace")
	//fmt.Printf("Tokens = %v\n DigestType = %T\n", tokens, digest)
	//fmt.Printf("DIGEST %v", digest)

	if len(tokens) != 19 {
		t.Log("Expected thoreauThree to be length=19, got=", len(tokens))
		t.Fail()
	}

	typ := fmt.Sprintf("%T", digest)
	if typ != "*nlpt_tkz.Digest" {
		t.Log("Expected digest to be *nlpt_tkz.WhiteSpaceDigest", typ)
		t.Fail()
	}
}
