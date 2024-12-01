package main

import "testing"

func BenchmarkDay12(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}
