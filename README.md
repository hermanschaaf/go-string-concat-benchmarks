go-string-concat-benchmarks
===========================

Benchmarks to compare the different string concatenation methods in Go. For all the details, see [stringconcat_test.go](<stringconcat_test.go>). 

There is also a [blog post detailing the methodology and results](http://herman.asia/efficient-string-concatenation-in-go). 

In summary, for very few concatenations of short strings (fewer than hundred, length shorter than 10) using naive string appending is just fine. For more heavy-duty cases where efficiency is important, bytes.Buffer is the best choice out of the methods evaluated. strings.Join is a good choice when you already have a string slice that just needs to be concatenated into one string.

Here are the raw results:

```
BenchmarkNaiveConcat10    500000          6403 ns/op
BenchmarkNaiveConcat100    20000         77001 ns/op
BenchmarkNaiveConcat1000         500       3105361 ns/op
BenchmarkNaiveConcat10000          5     236872134 ns/op
BenchmarkJoin10   200000          8129 ns/op
BenchmarkJoin100       50000         51891 ns/op
BenchmarkJoin1000       5000        465179 ns/op
BenchmarkJoin10000       500       5911317 ns/op
BenchmarkBufferString10   200000          7688 ns/op
BenchmarkBufferString100       50000         45027 ns/op
BenchmarkBufferString1000       5000        398311 ns/op
BenchmarkBufferString10000       500       3702906 ns/op
```
