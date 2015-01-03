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

Compare with `remove_interface` branch:

``` sh
..
```
