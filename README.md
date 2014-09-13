go-string-concat-benchmarks
===========================

Benchmarks to compare the different string concatenation methods in Go. For all the details, see [stringconcat_test.go](<stringconcat_test.go>).

There is also a [blog post detailing the methodology and results](http://herman.asia/efficient-string-concatenation-in-go).

In summary, for very few concatenations of short strings (fewer than hundred, length shorter than 10) using naive string appending is just fine. For more heavy-duty cases where efficiency is important, bytes.Buffer is the best choice out of the methods evaluated. strings.Join is a good choice when you already have a string slice that just needs to be concatenated into one string.

Here are how the methods tested stack up:

![Comparison of string concatenation methods in Go](http://img.svbtle.com/rlmmrxjtthkg.png)

Here is updated benchmark on dlintw's slow machine with corrected byte slices code:
```
BenchmarkNaiveConcat10	  200000	      9506 ns/op	     408 B/op	      14 allocs/op
BenchmarkNaiveConcat100	   10000	    208609 ns/op	   26936 B/op	     134 allocs/op
BenchmarkNaiveConcat1000	     100	  11728684 ns/op	 2699833 B/op	    1342 allocs/op
BenchmarkNaiveConcat10000	       2	 555584549 ns/op	271742592 B/op	   13674 allocs/op
BenchmarkByteSlice10	  200000	     10077 ns/op	     360 B/op	      16 allocs/op
BenchmarkByteSlice100	   20000	     82046 ns/op	    3144 B/op	     109 allocs/op
BenchmarkByteSlice1000	    2000	    947426 ns/op	   42121 B/op	    1015 allocs/op
BenchmarkByteSlice10000	     200	  11628370 ns/op	  443656 B/op	   10023 allocs/op
BenchmarkJoin10	  100000	     17296 ns/op	     455 B/op	      12 allocs/op
BenchmarkJoin100	   20000	     77366 ns/op	    3655 B/op	      45 allocs/op
BenchmarkJoin1000	    5000	    806334 ns/op	   32770 B/op	     350 allocs/op
BenchmarkJoin10000	     200	   9872287 ns/op	  554637 B/op	    3366 allocs/op
BenchmarkBufferString10	  200000	     11895 ns/op	     433 B/op	      10 allocs/op
BenchmarkBufferString100	   10000	    111954 ns/op	    2880 B/op	      43 allocs/op
BenchmarkBufferString1000	    2000	    699708 ns/op	   24194 B/op	     346 allocs/op
BenchmarkBufferString10000	     500	   4558195 ns/op	  223501 B/op	    3350 allocs/op
BenchmarkSliceString10	  500000	      7634 ns/op	     256 B/op	       9 allocs/op
BenchmarkSliceString100	   50000	     58240 ns/op	    2080 B/op	      42 allocs/op
BenchmarkSliceString1000	    5000	    684792 ns/op	   31458 B/op	     348 allocs/op
BenchmarkSliceString10000	     500	   5263568 ns/op	  337000 B/op	    3357 allocs/op
BenchmarkSliceStringA10	  500000	      9986 ns/op	     376 B/op	       8 allocs/op
BenchmarkSliceStringA100	   20000	     69826 ns/op	    3624 B/op	      38 allocs/op
BenchmarkSliceStringA1000	    5000	    795028 ns/op	   56040 B/op	     340 allocs/op
BenchmarkSliceStringA10000	     500	   5689369 ns/op	  561256 B/op	    3340 allocs/op
```
