merger
======

Playing around with golang merging overlapping time intervals.

# Requirements

* [Golang 1.13.8+](https://golang.org/dl/)
* [make](http://mirrors.ocf.berkeley.edu/gnu/make/) (optionally)

# Build

To build the program use the Makefile's `build` target
```
$ make build
```

# Run

To run the program using the provided Makefile's `run` target. Type following on your shell
```
$ make run
```

This will compile, build a binary `merger` in your CWD and is going to invoke it afterwards 
```
go build -o merger -v ./...
./merger
Input:  [[25 30] [2 19] [14 23] [4 8]]
Output: [[2 23] [25 30]]
```

# Test

To run the unit tests use the make target accordingly
```
$ make test
```

This will result in
```
go test -v ./...
=== RUN   TestMerger
--- PASS: TestMerger (0.00s)
PASS
ok  	github.com/denisneuling/evidence210531	(cached)

```

# Benchmark

To run a little benchmark use the make target accordingly 
```
$ make bench
```

This will result in something like
```
go test -bench=. -benchmem
goos: linux
goarch: amd64
pkg: github.com/denisneuling/evidence210531
BenchmarkMerge1-12          	 3860354	       327 ns/op	     152 B/op	       6 allocs/op
BenchmarkMerge10-12         	  913640	      1267 ns/op	     368 B/op	      11 allocs/op
BenchmarkMerge100-12        	  155721	      7052 ns/op	     592 B/op	      13 allocs/op
BenchmarkMerge1000-12       	    5529	    242111 ns/op	    3349 B/op	      39 allocs/op
BenchmarkMerge10000-12      	     553	   1858934 ns/op	    2922 B/op	      46 allocs/op
BenchmarkMerge100000-12     	      55	  23283858 ns/op	   75982 B/op	    1857 allocs/op
BenchmarkMerge1000000-12    	       3	 364327270 ns/op	13340416 B/op	  333388 allocs/op
PASS
ok  	github.com/denisneuling/evidence210531	11.492s
```

# Learnings

While the size of randomized input slice grows by the power of ten, the memory consuption per operation as well as the amount of memory allocations is rather growing exponentially.  
Merging one million intervals will take around 364ms depending on the computation platform and results in a memory consumption of around 13.3MB: This could certainly receive some optimization, but the algorythm solves the task still fast enough while not eating too much memory.