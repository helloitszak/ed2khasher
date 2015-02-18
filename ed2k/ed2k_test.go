package ed2k

import (
	"bytes"
	"testing"
)

func HashWrapper(t *testing.T, size int, hash string, old bool) {
	test := bytes.NewReader(make([]byte, size))
	testHash, err := Hash(test, old)
	
	if err != nil {
		t.Error("Got error ", err)
	}
	
	if testHash != hash {
		t.Error("Expected ", hash, " got ", testHash)
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
