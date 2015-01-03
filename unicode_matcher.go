/*
* Copyright ©2014 The nlpt Authors. All rights reserved.
* Use of this source code is governed by a BSD-style
* license that can be found in the LICENSE file.

* A Unicode Matcher tokenizer uses unicode code point ranges to segment text.
* It filters through strings leveraging standard library functions to parse and
* collect unicode segments into a 'bucket digest' of letter, number, punctuation, space, and symbol.
* Since this package uses runes the set is limited to characters of 4 bytes
* or less (the limit of the Rune type).
*
 */

package nlpt_tkz

import (
	"strings"
	"unicode"
)

//NewStateFunDigest initializes a new BucketDigest and explicitly allocates a length and cap on all stirng slices.
func NewUnicodeMatchDigest() *Digest {
	return &Digest{
		Letter:     make([]string, 0, 0),
		Title:      make([]string, 0, 0),
		Number:     make([]string, 0, 0),
		Punct:      make([]string, 0, 0),
		Space:      make([]string, 0, 0),
		Symbol:     make([]string, 0, 0),
		Bytes:      make([]byte, 0, 0),
		TokenBytes: make(map[string][]byte),
	}
}

//Tknz implements Tokenizer interface. Here it uses Unicode package to match runes for tokenization. It can be useful for really noisy data sets. For example, a sequence like 'expect0.7rant7!' will be tokenized into 3 buckets of LETTER: 'expectrant', NUMBER: '0 7 7', and PUNCT: '. !'.
//Caution should be used, however, as there is a great amount of information loss too. Date sequences, monetary sequences, urls, or any other complex combination of unicode sequences will be bucketized.
//One use of this tokenizer is to clean up naoisy data or for post-processing of already tokenized data for specific data-mining tasks. This is not a typical tokenizer. If you want basic tokenization see the Whist (whitespace), Lext (lexical scanner), Punkt (sentence segmenter) tokenizers.

func TknzUnicode(text string, digest *Digest) ([]string, *Digest) {
	for _, v := range text {
		switch true {
		case unicode.IsTitle(v):
			digest.Title = append(digest.Title, string(v))
		case unicode.IsLetter(v):
			digest.Letter = append(digest.Letter, string(v))
		case unicode.IsSpace(v):
			digest.Letter = append(digest.Letter, ", ")
		case unicode.IsNumber(v):
			digest.Number = append(digest.Number, string(v))
		case unicode.IsPunct(v):
			digest.Punct = append(digest.Punct, string(v))
		case unicode.IsSymbol(v):
			digest.Symbol = append(digest.Symbol, string(v))
		}
	}
	digest.Tokens = strings.Split(strings.Join(digest.Letter, ""), ", ")
	return digest.Tokens, digest
}

func TknzUnicodeBytes(byteSeq []byte, digest *Digest) *Digest {
	for _, b := range byteSeq {
		runeBytes := rune(b)
		stringedBytes := string(b)
		//lexBytesPadded := append([]byte{b}, BytesSpacePadding)
		switch true {
		case unicode.IsTitle(runeBytes):
			digest.Title = append(digest.Title, stringedBytes)
			//digest.TokenBytes[stringedBytes] = []byte{b}
			//digest.Bytes = ConcatByteSlice(digest.Bytes, lexBytesPadded)
		case unicode.IsLetter(runeBytes):
			digest.Letter = append(digest.Letter, stringedBytes)
			//digest.TokenBytes[stringedBytes] = []byte{b}
			//digest.Bytes = ConcatByteSlice(digest.Bytes, lexBytesPadded)
		case unicode.IsSpace(runeBytes):
			digest.Letter = append(digest.Letter, ", ")
			//digest.Bytes = ConcatByteSlice(digest.Bytes, lexBytesPadded)
			//digest.TokenBytes[stringedBytes] = []byte{b}
		case unicode.IsNumber(runeBytes):
			digest.Number = append(digest.Number, stringedBytes)
			//digest.TokenBytes[stringedBytes] = []byte{b}
			//digest.Bytes = ConcatByteSlice(digest.Bytes, lexBytesPadded)
		case unicode.IsPunct(runeBytes):
			digest.Punct = append(digest.Punct, stringedBytes)
			//digest.TokenBytes[stringedBytes] = []byte{b}
			//digest.Bytes = ConcatByteSlice(digest.Bytes, lexBytesPadded)
		case unicode.IsSymbol(runeBytes):
			digest.Symbol = append(digest.Symbol, stringedBytes)
			//digest.TokenBytes[stringedBytes] = []byte{b}
			//digest.Bytes = ConcatByteSlice(digest.Bytes, lexBytesPadded)
		}
	}
	digest.Tokens = strings.Split(strings.Join(digest.Letter, ""), ", ")
	for _, t := range digest.Tokens {
		digest.Bytes = ConcatByteSlice(digest.Bytes, []byte(t+"\n"))
		digest.TokenBytes[t] = []byte(t)
	}
	return digest
}
