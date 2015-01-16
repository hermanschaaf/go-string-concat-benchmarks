go-string-concat-benchmarks
===========================

Benchmarks to compare the different string concatenation methods in Go. For all the details, see [stringconcat_test.go](<stringconcat_test.go>).

There is also a [blog post detailing the methodology and results](http://herman.asia/efficient-string-concatenation-in-go).

In summary, for very few concatenations of short strings (fewer than hundred, length shorter than 10) using naive string appending is just fine. For more heavy-duty cases where efficiency is important, bytes.Buffer is the best choice out of the methods evaluated. strings.Join is a good choice when you already have a string slice that just needs to be concatenated into one string.

Here are how the methods tested stack up:

![Comparison of string concatenation methods in Go](http://img.svbtle.com/rlmmrxjtthkg.png)

Here is updated benchmark on dlintw's slow machine with corrected byte slices code:
```
BenchmarkNaiveConcat10	  200000	      9274 ns/op	     408 B/op	      14 allocs/op
BenchmarkNaiveConcat100	   10000	    178076 ns/op	   26936 B/op	     134 allocs/op
BenchmarkNaiveConcat1000	     100	  10297768 ns/op	 2699773 B/op	    1338 allocs/op
BenchmarkNaiveConcat10000	       2	 551974966 ns/op	271742592 B/op	   13674 allocs/op
BenchmarkByteSlice10	  200000	     10153 ns/op	     360 B/op	      16 allocs/op
BenchmarkByteSlice100	   20000	     81188 ns/op	    3144 B/op	     109 allocs/op
BenchmarkByteSlice1000	    2000	    810570 ns/op	   42121 B/op	    1015 allocs/op
BenchmarkByteSlice10000	     500	   6158915 ns/op	  443656 B/op	   10023 allocs/op
BenchmarkJoin10	  200000	     10440 ns/op	     455 B/op	      12 allocs/op
BenchmarkJoin100	   50000	     71594 ns/op	    3655 B/op	      45 allocs/op
BenchmarkJoin1000	    5000	    645059 ns/op	   32770 B/op	     350 allocs/op
BenchmarkJoin10000	     500	   6996372 ns/op	  554619 B/op	    3366 allocs/op
BenchmarkBufferString10	  200000	      9947 ns/op	     433 B/op	      10 allocs/op
BenchmarkBufferString100	   50000	     68339 ns/op	    2880 B/op	      43 allocs/op
BenchmarkBufferString1000	    5000	    605385 ns/op	   24194 B/op	     346 allocs/op
BenchmarkBufferString10000	     500	   4459925 ns/op	  223501 B/op	    3350 allocs/op
BenchmarkSliceString10	  500000	      6480 ns/op	     256 B/op	       9 allocs/op
BenchmarkSliceString100	   50000	     59868 ns/op	    2080 B/op	      42 allocs/op
BenchmarkSliceString1000	    5000	    586186 ns/op	   31458 B/op	     348 allocs/op
BenchmarkSliceString10000	     500	   4651593 ns/op	  337000 B/op	    3357 allocs/op
BenchmarkSliceStringA10	  200000	      6878 ns/op	     344 B/op	       7 allocs/op
BenchmarkSliceStringA100	   50000	     60004 ns/op	    3112 B/op	      37 allocs/op
BenchmarkSliceStringA1000	    5000	    561599 ns/op	   31208 B/op	     337 allocs/op
BenchmarkSliceStringA10000	     500	   4205116 ns/op	  315496 B/op	    3337 allocs/op
ok  	_/home/dlin/tmp/go-string-concat-benchmarks	65.784s
```
