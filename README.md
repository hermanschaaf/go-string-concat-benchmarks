go-string-concat-benchmarks
===========================

Benchmarks to compare the different string concatenation methods in Go. For all the details, see [stringconcat_test.go](<stringconcat_test.go>).

There is also a [blog post detailing the methodology and results](http://herman.asia/efficient-string-concatenation-in-go).

In summary, for very few concatenations of short strings (fewer than hundred, length shorter than 10) using naive string appending is just fine. For more heavy-duty cases where efficiency is important, bytes.Buffer is the best choice out of the methods evaluated. strings.Join is a good choice when you already have a string slice that just needs to be concatenated into one string.

Here are how the methods tested stack up:

![Comparison of string concatenation methods in Go](http://img.svbtle.com/rlmmrxjtthkg.png)

And here are the raw results (also including a benchmark for byte slices):

```
BenchmarkNaiveConcat10	  500000	      6738 ns/op	     442 B/op	      21 allocs/op
BenchmarkNaiveConcat100	   10000	    115895 ns/op	   27736 B/op	     201 allocs/op
BenchmarkNaiveConcat1000	     500	   4802280 ns/op	 2684216 B/op	    2001 allocs/op
BenchmarkNaiveConcat10000	       5	 272311728 ns/op	264645752 B/op	   20001 allocs/op
BenchmarkByteSlice10	  200000	      8674 ns/op	     432 B/op	      28 allocs/op
BenchmarkByteSlice100	   50000	     62424 ns/op	    3440 B/op	     211 allocs/op
BenchmarkByteSlice1000	    5000	    607276 ns/op	   48688 B/op	    2019 allocs/op
BenchmarkByteSlice10000	     500	   5697016 ns/op	  509104 B/op	   20029 allocs/op
BenchmarkJoin10	  200000	      8750 ns/op	     740 B/op	      19 allocs/op
BenchmarkJoin100	   50000	     56398 ns/op	    6018 B/op	     113 allocs/op
BenchmarkJoin1000	    5000	    484793 ns/op	   51304 B/op	    1018 allocs/op
BenchmarkJoin10000	     500	   5896195 ns/op	 1112755 B/op	   10032 allocs/op
BenchmarkBufferString10	  200000	      9120 ns/op	     433 B/op	      18 allocs/op
BenchmarkBufferString100	   50000	     49411 ns/op	    2720 B/op	     111 allocs/op
BenchmarkBufferString1000	    5000	    422961 ns/op	   23488 B/op	    1014 allocs/op
BenchmarkBufferString10000	     500	   3842033 ns/op	  297216 B/op	   10018 allocs/op
```
