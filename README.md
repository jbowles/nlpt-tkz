# nlpt_tkz
Natural language tokenizer; supports many types of tokenizer including a simple white space tokenizer, a unicode pattern-matcher tokenizer, and state function tokenizing lexer.

## Get it

```sh
go get 
```

## Use it

```go
```

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
