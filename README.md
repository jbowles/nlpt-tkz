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
On `remove_interface` branch:

``` sh
PASS
BenchmarkStateFnTknzGoodStr	   10000	    136076 ns/op
BenchmarkStateFnTknzBadStr	   20000	     89851 ns/op
BenchmarkStateFnTknzBytesGoodStr	   10000	    136100 ns/op
BenchmarkStateFnTknzBytesBadStr	   20000	     88691 ns/op
BenchmarkLexStr	   30000	     56678 ns/op
BenchmarkUnicodeStr	  100000	     17344 ns/op
BenchmarkWhitespaceStr	 1000000	      1620 ns/op
BenchmarkLexBytes	   30000	     57201 ns/op
BenchmarkUnicodeBytes	   50000	     34919 ns/op
BenchmarkUncdMatchTknzGoodStr	   50000	     31309 ns/op
BenchmarkUncdMatchTnkzBadStr	  100000	     20302 ns/op
BenchmarkUncdMatchTknzBytesGoodStr	   20000	     72872 ns/op
BenchmarkUncdMatchTnkzBytesBadStr	   30000	     40730 ns/op
BenchmarkWhiteSpaceTknzGoodStr	 1000000	      1134 ns/op
BenchmarkWhiteSpaceTknzBadStr	 1000000	      1431 ns/op
ok  	github.com/jbowles/nlpt_tkz	28.868s
```
