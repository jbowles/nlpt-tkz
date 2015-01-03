# nlpt_tkz
Natural language tokenizer; supports many types of tokenizer including a simple white space tokenizer, a unicode pattern-matcher tokenizer, and state function tokenizing lexer.

There are 2 toplevel functions: `TokenizeStr(string, string) []string *Digest` and `TokenizeBytes([]byte, string) *Digest`. Both functions require you to select which tokenizer type to be used (the second argument). The `TokenizeBytes` function only supports `lex` and 'uicode` options and will only return a `Digest` struct with a `Bytes` field; it is for use in instances where dealing with strings in not preferred. The former function is for dealing direclty with strings and supports `lex`, `unicode`, `whitespace`; it returns a slice of tokens as well as digest, and depending on the tokenizer used the digest will contain different data fields.

## Get it

```sh
go get github.com/jbowles/nlpt_tkz
```

## Use it

```go
package main

import (
	"fmt"
	tkz "github.com/jbowles/nlpt_tkz"
)

func main() {
	s := "From: mathew <mathew@mantis.co.uk> \nSubject: Alt.Atheism FAQ: Atheist Resources\n\nArchive-name: atheism/resources\nAlt-atheism-archive-name: resources\nLast-modified: 11 December 1992\nVersion: 1.0"
	b := []byte(s)
	digest1 := tkz.TokenizeBytes(b, "lex")
	digest2 := tkz.TokenizeBytes(b, "unicode")
	_, digest3 := tkz.TokenizeStr(s, "lex")
	_, digest4 := tkz.TokenizeStr(s, "unicode")
	_, digest5 := tkz.TokenizeStr(s, "whitespace")

	fmt.Printf("-----printed digest-----")
	fmt.Printf("\n\n")
	fmt.Printf("LexBytes \n %+v\n\n", digest1)
	fmt.Printf("UnicodeBytes \n %+v\n\n", digest2)
	fmt.Printf("LexStr \n %+v\n\n", digest3)
	fmt.Printf("UnicodeStr \n %+v\n\n", digest4)
	fmt.Printf("WhitespaceStr \n %+v\n\n", digest5)
	fmt.Printf("---------------------")
	fmt.Printf("\n\n\n")
	fmt.Printf("-----printed bytes-----")
	fmt.Printf("\n\n")

	fmt.Printf("++++++ LexBytes Printed +++++++ \n\n %s\n", digest1.Bytes)
	fmt.Printf("\n\n")
	fmt.Printf("+++++ UnicodeBytes Printed ++++++ \n\n %s\n", digest2.Bytes)
}
```

## Tokenizers so far...
The white space tokenizer is merely a wrapper around `strings.Split(" ")` with some digest content for counts of tokens and such.

The `lex` (technically a State Function Lexer) and `unicode` (technically using go's `unicode` package for matching unicode code points) will return very differnt tokens. The rule of thumb is that if you just need "words" and have a text that is pretty clean use `unicode`; it is the fastest of the two.

### Example output of lex and unicode bytes parsers
An excerpt from a corpus of news and emails shows differences between the two tokenizers in parsing the email headers.
##### Original text

```sh
From: mathew <mathew@mantis.co.uk>
Subject: Alt.Atheism FAQ: Atheist Resources

Archive-name: atheism/resources
Alt-atheism-archive-name: resources
Last-modified: 11 December 1992
Version: 1.0
```
#### lex

```sh
From: mathew <mathew@mantis.co.uk>Subject: Alt.Atheism FAQ: Atheist ResourcesArchive-name: atheism/resourcesAlt-atheism-archive-name: resourcesLast-modified: 11 December 1992Version: 1.0
```

#### unicode
From mathew mathewmantiscoukSubject AltAtheism FAQ Atheist ResourcesArchivename atheismresourcesAltatheismarchivename resourcesLastmodified 11 December 1992Version 10


## Benchmarks
On old `master` branch:

``` sh
PASS
BenchmarkStateFnTknzGoodStr	    20000	     93640 ns/op
BenchmarkStateFnTknzBadStr	    20000	     63169 ns/op
BenchmarkLex	                  20000	     94277 ns/op
BenchmarkBukt	                  50000	     31131 ns/op
BenchmarkSpace	                500000	      2743 ns/op
BenchmarkUncdMatchTknzGoodStr	  50000	     30771 ns/op
BenchmarkUncdMatchTnkzBadStr	  100000	     20315 ns/op
BenchmarkWhiteSpaceTknzGoodStr	500000	      2599 ns/op
BenchmarkWhiteSpaceTknzBadStr	  1000000	      1497 ns/op
ok  	github.com/jbowles/nlpt_tkz	17.760s
```

On old `master` branch:

``` sh
PASS
BenchmarkStateFnTknzGoodStr	   10000	    142893 ns/op
BenchmarkStateFnTknzBadStr	   20000	     98818 ns/op
BenchmarkStateFnTknzBytesGoodStr	   20000	     94985 ns/op
BenchmarkStateFnTknzBytesBadStr	   30000	     59880 ns/op
BenchmarkLexStrGood	   10000	    144635 ns/op
BenchmarkUnicodeStrGood	   50000	     30375 ns/op
BenchmarkWhitespaceStrGood	  500000	      2872 ns/op
BenchmarkLexBytesGood	   20000	     95068 ns/op
BenchmarkUnicodeBytesGood	  200000	      6493 ns/op
BenchmarkLexStrBad	   20000	     97813 ns/op
BenchmarkUnicodeStrBad	  100000	     20998 ns/op
BenchmarkWhitespaceStrBad	 1000000	      1999 ns/op
BenchmarkLexBytesBad	   20000	     64115 ns/op
BenchmarkUnicodeBytesBad	  300000	      5101 ns/op
BenchmarkUncdMatchTknzGoodStr	   50000	     31779 ns/op
BenchmarkUncdMatchTnkzBadStr	  100000	     20875 ns/op
BenchmarkUncdMatchTknzBytesGoodStr	  200000	      6570 ns/op
BenchmarkUncdMatchTnkzBytesBadStr	  300000	      5106 ns/op
BenchmarkWhiteSpaceTknzGoodStr	 1000000	      1211 ns/op
BenchmarkWhiteSpaceTknzBadStr	 1000000	      1517 ns/op
ok  	github.com/jbowles/nlpt_tkz	39.393s
```
