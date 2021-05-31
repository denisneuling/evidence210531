package main

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestMerger(t *testing.T) {
	for _, c := range []struct {
		in   [][]int
		want [][]int
	}{
		{
			[][]int{{1, 1}, {2, 3}},
			[][]int{{1, 1}, {2, 3}},
		},
		{
			[][]int{{1, 2}, {2, 3}},
			[][]int{{1, 3}},
		},
		{
			[][]int{{2, 3}, {1, 2}, {1, 2}},
			[][]int{{1, 3}},
		},
		{
			[][]int{{25, 30}, {2, 19}, {14, 23}},
			[][]int{{2, 23}, {25, 30}},
		},
		{
			[][]int{{25, 30}, {2, 19}, {14, 23}, {4, 8}},
			[][]int{{2, 23}, {25, 30}},
		},
		{
			[][]int{{60, 70}, {50, 59}},
			[][]int{{50, 59}, {60, 70}},
		},

		{
			[][]int{{60, 70}, {50, 59}, {40, 48}},
			[][]int{{40, 48}, {50, 59}, {60, 70}},
		},
		//
		{
			[][]int{{60, 70}, {50, 59}, {40, 48}, {12, 18}},
			[][]int{{12, 18}, {40, 48}, {50, 59}, {60, 70}},
		},
		{
			[][]int{{60, 70}, {50, 69}},
			[][]int{{50, 70}},
		},

		{
			[][]int{},
			[][]int{},
		},

		{
			[][]int{
				{1, 2}, {3, 4}, {1, 4},
			},
			[][]int{{1, 4}},
		},
	} {
		got, e := Merge(c.in)

		if e != nil {
			t.Error(e)
		} else if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Merge(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

// prevent compiler from line erasure
var benchmarkResult [][]int

func randSlice(len int) [][]int {
	a := make([][]int, len)
	for e := 0; e < len; e++ {
		a[e] = []int{rand.Int(), rand.Int()}
	}
	return a
}

func benchmarkMerge(input [][]int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		benchmarkResult, _ = Merge(input)
	}
}

func BenchmarkMerge1(b *testing.B)       { benchmarkMerge(randSlice(1), b) }
func BenchmarkMerge10(b *testing.B)      { benchmarkMerge(randSlice(10), b) }
func BenchmarkMerge100(b *testing.B)     { benchmarkMerge(randSlice(100), b) }
func BenchmarkMerge1000(b *testing.B)    { benchmarkMerge(randSlice(1000), b) }
func BenchmarkMerge10000(b *testing.B)   { benchmarkMerge(randSlice(10000), b) }
func BenchmarkMerge100000(b *testing.B)  { benchmarkMerge(randSlice(100000), b) }
func BenchmarkMerge1000000(b *testing.B) { benchmarkMerge(randSlice(1000000), b) }
