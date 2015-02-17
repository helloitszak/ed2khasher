package main

import (
	"io"
	"os"
	"testing"
)

func open(filename string) io.Reader {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return file
}

func BenchmarkHashFileLegacy(b *testing.B) {
	for n := 0; n < b.N; n++ {
		file := open("test.mp4")
		hashFileOld(file)
	}
}

func BenchmarkHashFile(b *testing.B) {
	for n := 0; n < b.N; n++ {
		file := open("test.mp4")
		hashFile(file)
	}
}
