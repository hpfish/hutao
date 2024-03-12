package main

import "testing"

func BenchmarkTailRecur(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tailRecur(5, 0)
	}
}
