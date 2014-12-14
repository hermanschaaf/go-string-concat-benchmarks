// Package stringconcat exists only to provide benchmarks for the different approaches
// to string concatenation in Go.
package stringconcat

import (
	"bytes"
	"strconv"
	"strings"
	"testing"
)

var s []string = []string{}

// nextString is an iterator we use to represent a process
// that returns strings that we want to concatenate in order.
func nextString() func() string {
	n := 10000
	// closure captures variable n
	return func() string {
		n += 1
		return strconv.Itoa(n)
	}
}

var global string

// benchmarkNaiveConcat provides a benchmark for basic built-in
// Go string concatenation. Because strings are immutable in Go,
// it performs the worst of the tested methods. The time taken to
// set up the array that is appended is not counted towards the
// time for naive concatenation.
func benchmarkNaiveConcat(b *testing.B, numConcat int) {
	// Reports memory allocations
	b.ReportAllocs()

	var ns string
	for i := 0; i < b.N; i++ {
		next := nextString()
		ns = ""
		for u := 0; u < numConcat; u++ {
			ns += next()
		}
	}
	// we assign to a global variable to make sure compiler
	// or runtime optimizations don't skip over the operations
	// we were benchmarking. This might be unnecessary, but it's
	// safe.
	global = ns
}

func BenchmarkNaiveConcat10(b *testing.B) {
	benchmarkNaiveConcat(b, 10)
}

func BenchmarkNaiveConcat100(b *testing.B) {
	benchmarkNaiveConcat(b, 100)
}

func BenchmarkNaiveConcat1000(b *testing.B) {
	benchmarkNaiveConcat(b, 1000)
}

func BenchmarkNaiveConcat10000(b *testing.B) {
	benchmarkNaiveConcat(b, 10000)
}

// benchmarkByteSlice provides a benchmark for the time it takes
// to repeatedly append returned strings to a byte slice, and
// finally casting the byte slice to string type.
func benchmarkByteSlice(b *testing.B, numConcat int) {
	// Reports memory allocations
	b.ReportAllocs()

	var ns string
	for i := 0; i < b.N; i++ {
		next := nextString()
		b := []byte{}
		for u := 0; u < numConcat; u++ {
			b = append(b, []byte(next())...)
		}
		ns = string(b)
	}
	global = ns
}

func BenchmarkByteSlice10(b *testing.B) {
	benchmarkByteSlice(b, 10)
}

func BenchmarkByteSlice100(b *testing.B) {
	benchmarkByteSlice(b, 100)
}

func BenchmarkByteSlice1000(b *testing.B) {
	benchmarkByteSlice(b, 1000)
}

func BenchmarkByteSlice10000(b *testing.B) {
	benchmarkByteSlice(b, 10000)
}

// benchmarkJoin provides a benchmark for the time it takes to set
// up an array with strings, and calling strings.Join on that array
// to get a fully concatenated string.
func benchmarkJoin(b *testing.B, numConcat int) {
	// Reports memory allocations
	b.ReportAllocs()

	var ns string
	for i := 0; i < b.N; i++ {
		next := nextString()
		a := []string{}
		for u := 0; u < numConcat; u++ {
			a = append(a, next())
		}
		ns = strings.Join(a, "")
	}
	global = ns
}

func BenchmarkJoin10(b *testing.B) {
	benchmarkJoin(b, 10)
}

func BenchmarkJoin100(b *testing.B) {
	benchmarkJoin(b, 100)
}

func BenchmarkJoin1000(b *testing.B) {
	benchmarkJoin(b, 1000)
}

func BenchmarkJoin10000(b *testing.B) {
	benchmarkJoin(b, 10000)
}

// benchmarkJoinSize provides a benchmark for the time it takes to set
// up an array with strings, and calling strings.Join on that array
// to get a fully concatenated string – when the (approximate) number of
// strings is known in advance.
//
// This is identical to benchmarkJoin, except numConcat is used to size
// the []string slice's initial capacity to avoid needless reallocation.
func benchmarkJoinSize(b *testing.B, numConcat int) {
	// Reports memory allocations
	b.ReportAllocs()

	var ns string
	for i := 0; i < b.N; i++ {
		next := nextString()
		a := make([]string, 0, numConcat)
		for u := 0; u < numConcat; u++ {
			a = append(a, next())
		}
		ns = strings.Join(a, "")
	}
	global = ns
}

func BenchmarkJoinSize10(b *testing.B) {
	benchmarkJoinSize(b, 10)
}

func BenchmarkJoinSize100(b *testing.B) {
	benchmarkJoinSize(b, 100)
}

func BenchmarkJoinSize1000(b *testing.B) {
	benchmarkJoinSize(b, 1000)
}

func BenchmarkJoinSize10000(b *testing.B) {
	benchmarkJoinSize(b, 10000)
}

// benchmarkBufferString
func benchmarkBufferString(b *testing.B, numConcat int) {
	// Reports memory allocations
	b.ReportAllocs()

	var ns string
	for i := 0; i < b.N; i++ {
		next := nextString()
		buffer := bytes.NewBufferString("")
		for u := 0; u < numConcat; u++ {
			buffer.WriteString(next())
		}
		ns = buffer.String()
	}
	global = ns
}

func BenchmarkBufferString10(b *testing.B) {
	benchmarkBufferString(b, 10)
}

func BenchmarkBufferString100(b *testing.B) {
	benchmarkBufferString(b, 100)
}

func BenchmarkBufferString1000(b *testing.B) {
	benchmarkBufferString(b, 1000)
}

func BenchmarkBufferString10000(b *testing.B) {
	benchmarkBufferString(b, 10000)
}
