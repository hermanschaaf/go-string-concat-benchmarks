go-string-concat-benchmarks
===========================

Benchmarks to compare the different string concatenation methods in Go. For all the details, see [stringconcat_test.go](<stringconcat_test.go>).

There is also a [blog post detailing the methodology and results](http://herman.asia/efficient-string-concatenation-in-go).

In summary, for very few concatenations of short strings (fewer than hundred, length shorter than 10) using naive string appending is just fine. For more heavy-duty cases where efficiency is important, bytes.Buffer is the best choice out of the methods evaluated. strings.Join is a good choice when you already have a string slice that just needs to be concatenated into one string.

Here are how the methods tested stack up:

![Comparison of string concatenation methods in Go](http://img.svbtle.com/rlmmrxjtthkg.png)

And here are the raw results (also including a benchmark for byte slices):

```
BenchmarkNaiveConcat10-3     	 2000000	      1192 ns/op	     360 B/op	      11 allocs/op
BenchmarkNaiveConcat100-3    	   50000	     34117 ns/op	   26408 B/op	     101 allocs/op
BenchmarkNaiveConcat1000-3   	     500	   2641900 ns/op	 2694414 B/op	    1004 allocs/op
BenchmarkNaiveConcat10000-3  	      10	 188733914 ns/op	271630262 B/op	   10339 allocs/op
BenchmarkByteSlice10-3       	 2000000	       717 ns/op	     208 B/op	       7 allocs/op
BenchmarkByteSlice100-3      	  300000	      4000 ns/op	    1552 B/op	      10 allocs/op
BenchmarkByteSlice1000-3     	   30000	     53874 ns/op	   26128 B/op	      16 allocs/op
BenchmarkByteSlice10000-3    	    2000	    701131 ns/op	  283668 B/op	      24 allocs/op
BenchmarkByteSliceSize10-3   	 3000000	       536 ns/op	     200 B/op	       4 allocs/op
BenchmarkByteSliceSize100-3  	  300000	      3396 ns/op	    1560 B/op	       4 allocs/op
BenchmarkByteSliceSize1000-3 	   30000	     40858 ns/op	   15896 B/op	       4 allocs/op
BenchmarkByteSliceSize10000-3	    2000	    575724 ns/op	  163866 B/op	       4 allocs/op
BenchmarkJoin10-3            	 1000000	      1724 ns/op	     648 B/op	       9 allocs/op
BenchmarkJoin100-3           	  200000	      9183 ns/op	    5128 B/op	      12 allocs/op
BenchmarkJoin1000-3          	   20000	     82654 ns/op	   43528 B/op	      15 allocs/op
BenchmarkJoin10000-3         	    1000	   1540246 ns/op	  941844 B/op	      24 allocs/op
BenchmarkJoinSize10-3        	 2000000	       887 ns/op	     312 B/op	       5 allocs/op
BenchmarkJoinSize100-3       	  200000	      5981 ns/op	    2712 B/op	       5 allocs/op
BenchmarkJoinSize1000-3      	   20000	     66210 ns/op	   27160 B/op	       5 allocs/op
BenchmarkJoinSize10000-3     	    2000	    814003 ns/op	  278555 B/op	       5 allocs/op
BenchmarkBufferString10-3    	 1000000	      1354 ns/op	     400 B/op	       8 allocs/op
BenchmarkBufferString100-3   	  200000	      6410 ns/op	    2368 B/op	      11 allocs/op
BenchmarkBufferString1000-3  	   20000	     56510 ns/op	   18880 B/op	      14 allocs/op
BenchmarkBufferString10000-3 	    2000	    699816 ns/op	  170178 B/op	      17 allocs/op
BenchmarkBufferSize10-3      	 2000000	       766 ns/op	     312 B/op	       5 allocs/op
BenchmarkBufferSize100-3     	  300000	      4548 ns/op	    1672 B/op	       5 allocs/op
BenchmarkBufferSize1000-3    	   30000	     50437 ns/op	   16008 B/op	       5 allocs/op
BenchmarkBufferSize10000-3   	    2000	    670182 ns/op	  163978 B/op	       5 allocs/op
```
