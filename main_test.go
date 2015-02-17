package main

import (
	"testing"
)

func BenchmarkHashFile100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		hashFile("test.mp4")
	}
}

func BenchmarkHashFileOld100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		hashFileOld("test.mp4")
	}
}
