package ed2k

import (
	"bytes"
	"testing"
)

func HashWrapper(t *testing.T, size int, hash string, old bool) {
	test := bytes.NewReader(make([]byte, size))
	testHash, err := Hash(test, old)

	if err != nil {
		t.Fatalf("Got error %s\n", err)
	}

	if testHash != hash {
		t.Errorf("Expected %s got %s\n", hash, testHash)
	}
}

func TestSmallFile(t *testing.T) {
	HashWrapper(t, 600, "a5b489c18c5bdc1f711a8edff22c13ff", false)
}

func TestSmallFileOld(t *testing.T) {
	HashWrapper(t, 600, "a5b489c18c5bdc1f711a8edff22c13ff", true)
}

func TestEqualFile(t *testing.T) {
	HashWrapper(t, 9728000, "d7def262a127cd79096a108e7a9fc138", false)
}

func TestEqualFileOld(t *testing.T) {
	HashWrapper(t, 9728000, "fc21d9af828f92a8df64beac3357425d", true)
}

func TestMultipleFile(t *testing.T) {
	HashWrapper(t, 19456000, "194ee9e4fa79b2ee9f8829284c466051", false)
}

func TestMultipleFileOld(t *testing.T) {
	HashWrapper(t, 19456000, "114b21c63a74b6ca922291a11177dd5c", true)
}

func TestLargeFile(t *testing.T) {
	HashWrapper(t, 19457000, "345da2ffa0f63eae5638b908f187bfb1", false)
}

func TestLargeFileOld(t *testing.T) {
	HashWrapper(t, 19457000, "345da2ffa0f63eae5638b908f187bfb1", true)
}

func HashWrapperBench(b *testing.B, size int, old bool, parallel bool) {
	test := bytes.NewReader(make([]byte, size))
	if(parallel) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				Hash(test, old)
			}
		})
		return
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Hash(test, old)
	}
}

func BenchmarkS(b *testing.B) {
	HashWrapperBench(b, 600, false, false)
}

func BenchmarkSO(b *testing.B) {
	HashWrapperBench(b, 600, true, false)
}

func BenchmarkSP(b *testing.B) {
	HashWrapperBench(b, 600, false, true)
}

func BenchmarkSOP(b *testing.B) {
	HashWrapperBench(b, 600, true, true)
}

func BenchmarkE(b *testing.B) {
	HashWrapperBench(b, 9728000, false, false)
}

func BenchmarkEO(b *testing.B) {
	HashWrapperBench(b, 9728000, true, false)
}

func BenchmarkEP(b *testing.B) {
	HashWrapperBench(b, 9728000, false, true)
}

func BenchmarkEOP(b *testing.B) {
	HashWrapperBench(b, 9728000, true, true)
}

func BenchmarkM(b *testing.B) {
	HashWrapperBench(b, 19456000, false, false)
}

func BenchmarkMO(b *testing.B) {
	HashWrapperBench(b, 19456000, true, false)
}

func BenchmarkMP(b *testing.B) {
	HashWrapperBench(b, 19456000, false, true)
}

func BenchmarkMOP(b *testing.B) {
	HashWrapperBench(b, 19456000, true, true)
}

func BenchmarkL(b *testing.B) {
	HashWrapperBench(b, 19457000, false, false)
}

func BenchmarkLO(b *testing.B) {
	HashWrapperBench(b, 19457000, true, false)
}

func BenchmarkLP(b *testing.B) {
	HashWrapperBench(b, 19457000, false, true)
}

func BenchmarkLOP(b *testing.B) {
	HashWrapperBench(b, 19457000, true, true)
}
