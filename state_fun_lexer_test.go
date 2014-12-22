package nlpt_tkz

import (
	"testing"
)

func BenchmarkStateFnTknzGoodStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var sdigestone = NewStateFnDigest()
		sdigestone.Tknz(ThoreauTwo)
	}
}

func BenchmarkStateFnTknzBadStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var sdigesttwo = NewStateFnDigest()
		sdigesttwo.Tknz(BadStr)
	}
}

func TestStateFnGoodStr(t *testing.T) {
	var sdigest = NewStateFnDigest()
	_, lxd := sdigest.Tknz(ThoreauOne)

	if lxd.TokenCount != 44 {
		t.Log("Expected word count to be 44, but got", lxd.TokenCount)
		t.Fail()
	}

	if lxd.PunctCount != 6 {
		t.Log("Expected punct count to be 5, but got", lxd.PunctCount)
		t.Fail()
	}

	if lxd.SpaceCount != 43 {
		t.Log("Expected space count to be 43, but got", lxd.SpaceCount)
		t.Fail()
	}

	if lxd.LineCount != 1 {
		t.Log("Expected line count to be 1, but got", lxd.LineCount)
		t.Fail()
	}

	if lxd.CharCount != 212 {
		t.Log("Expected char count to be 212, but got", lxd.CharCount)
		t.Fail()
	}

	if lxd.LastTokenType != 4 {
		t.Log("Expected last token type to be 4, but got", lxd.LastTokenType)
		t.Fail()
	}

	if len(lxd.Tokens) != lxd.TokenCount {
		t.Log("Expected token and word count to be equal. Tokens=", len(lxd.Tokens), "Words=", lxd.TokenCount)
		t.Fail()
	}

}
