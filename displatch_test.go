package dispatchbenchmark

import "testing"

func Benchmark_Baseline(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n++
	}
}

func Benchmark_Static(b *testing.B) {
	f := &foo{}
	for i := 0; i < b.N; i++ {
		f.increment()
	}
}

func Benchmark_Dynamic(b *testing.B) {
	var f incrementor = &foo{}
	for i := 0; i < b.N; i++ {
		f.increment()
	}
}
