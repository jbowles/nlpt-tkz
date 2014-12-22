/*
* Copyright ©2015 The nlpt Authors. All rights reserved.
* Use of this source code is governed by a BSD-style
* license that can be found in the LICENSE file.
 */

package nlpt_tkz

import "strings"

type WhiteSpaceDigest struct {
	Tokens     []string
	SpaceCount int
	CharCount  int
}

//NewWhiteSpaceDigest intitializes a digest for white space tokenization.
func NewWhiteSpaceDigest() *WhiteSpaceDigest {
	return &WhiteSpaceDigest{
		Tokens:     make([]string, 0, 0),
		CharCount:  0,
		SpaceCount: 0,
	}
}

//Tknz implements the Tokenizer interface. This uses the strings package Split() with a white space separator as well as collecting some other metadata for the digest.
func (digest *WhiteSpaceDigest) Tknz(text string) ([]string, *WhiteSpaceDigest) {
	digest.Tokens = strings.Split(text, " ")
	digest.SpaceCount = strings.Count(text, " ")
	digest.CharCount = len(text)

	return digest.Tokens, digest
}
