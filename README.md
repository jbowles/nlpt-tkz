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
On `master` branch:

``` sh
PASS
BenchmarkStateFnTknzGoodStr	    20000	      70,958 ns/op
BenchmarkStateFnTknzBadStr	    20000	      64,112 ns/op
BenchmarkLex	                  30000	      43,337 ns/op
BenchmarkBukt	                  100000	    17,656 ns/op
BenchmarkSpace	                1000000	    1,557 ns/op
BenchmarkUncdMatchTknzGoodStr	  50000	      31,710 ns/op
BenchmarkUncdMatchTnkzBadStr	  100000	    21,061 ns/op
BenchmarkWhiteSpaceTknzGoodStr	1000000	    1,290 ns/op
BenchmarkWhiteSpaceTknzBadStr	  1000000	    1,621 ns/op
ok  	github.com/jbowles/nlpt_tkz	16.500s
```

Compare with `remove_interface` branch:

``` sh
PASS
BenchmarkStateFnTknzGoodStr	        10000	    136,076 ns/op
BenchmarkStateFnTknzBadStr	        20000	    89,851 ns/op
BenchmarkStateFnTknzBytesGoodStr	  10000	    136,100 ns/op
BenchmarkStateFnTknzBytesBadStr	    20000	    88,691 ns/op
BenchmarkLexStr	                    30000	    56,678 ns/op
BenchmarkUnicodeStr	                100000	  17,344 ns/op
BenchmarkWhitespaceStr	            1000000	  1,620 ns/op
BenchmarkLexBytes	                  30000	    57,201 ns/op
BenchmarkUnicodeBytes	              50000	    34,919 ns/op
BenchmarkUncdMatchTknzGoodStr	      50000	    31,309 ns/op
BenchmarkUncdMatchTnkzBadStr	      100000	  2,0302 ns/op
BenchmarkUncdMatchTknzBytesGoodStr  20000	    72,872 ns/op
BenchmarkUncdMatchTnkzBytesBadStr	  30000	    40,730 ns/op
BenchmarkWhiteSpaceTknzGoodStr	    1000000	  1,134 ns/op
BenchmarkWhiteSpaceTknzBadStr	      1000000	  1,431 ns/op
ok  	github.com/jbowles/nlpt_tkz	28.868s
```
