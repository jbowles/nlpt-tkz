package nlpt_tkz

import (
	"testing"
)

func BenchmarkWhiteSpaceTknzGoodStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wsdigestone = NewWhiteSpaceDigest()
		wsdigestone.Tknz(ThoreauOne)
	}
}

func BenchmarkWhiteSpaceTknzBadStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wsdigesttwo = NewWhiteSpaceDigest()
		wsdigesttwo.Tknz(BadStr)
	}
}

func TestWhiteSpaceTknz(t *testing.T) {
	var wsdigest = NewWhiteSpaceDigest()
	tok3, digest := wsdigest.Tknz(ThoreauThree)

	if len(tok3) != 19 {
		t.Log("Expected thoreauThree string length=19, got=", len(tok3))
		t.Fail()
	}

	if len(ThoreauThree) != digest.CharCount {
		t.Log("Expected string and digest character counts to be equal, got string length=", len(ThoreauThree), "CharCount=", digest.CharCount)
		t.Fail()
	}
}
