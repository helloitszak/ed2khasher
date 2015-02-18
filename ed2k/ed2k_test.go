package ed2k

import (
	"bytes"
	"testing"
)

func TestSmallFile(t *testing.T) {
	test := bytes.NewReader(make([]byte, 600))
	validHash := "a5b489c18c5bdc1f711a8edff22c13ff"
	testHash, _ := Hash(test, false)
	if testHash != validHash {
		t.Error("Expected ", validHash, " got ", testHash)
	}
}

func TestSmallFileOld(t *testing.T) {
	test := bytes.NewReader(make([]byte, 600))
	validHash := "a5b489c18c5bdc1f711a8edff22c13ff"
	testHash, _ := Hash(test, true)
	if testHash != validHash {
		t.Error("Expected ", validHash, " got ", testHash)
	}
}

func TestEqualFile(t *testing.T) {
	test := bytes.NewReader(make([]byte, 9728000))
	validHash := "d7def262a127cd79096a108e7a9fc138"
	testHash, _ := Hash(test, false)
	if testHash != validHash {
		t.Error("Expected ", validHash, " got ", testHash)
	}
}

func TestEqualFileOld(t *testing.T) {
	test := bytes.NewReader(make([]byte, 9728000))
	validHash := "fc21d9af828f92a8df64beac3357425d"
	testHash, _ := Hash(test, true)
	if testHash != validHash {
		t.Error("Expected ", validHash, " got ", testHash)
	}
}

func TestMultipleFile(t *testing.T) {
	test := bytes.NewReader(make([]byte, 19456000))
	validHash := "194ee9e4fa79b2ee9f8829284c466051"
	testHash, _ := Hash(test, false)
	if testHash != validHash {
		t.Error("Expected ", validHash, " got ", testHash)
	}
}

func TestMultipleFileOld(t *testing.T) {
	test := bytes.NewReader(make([]byte, 19456000))
	validHash := "114b21c63a74b6ca922291a11177dd5c"
	testHash, _ := Hash(test, true)
	if testHash != validHash {
		t.Error("Expected ", validHash, " got ", testHash)
	}
}

func TestLargeFile(t *testing.T) {
	test := bytes.NewReader(make([]byte, 19457000))
	validHash := "345da2ffa0f63eae5638b908f187bfb1"
	testHash, _ := Hash(test, false)
	if testHash != validHash {
		t.Error("Expected ", validHash, " got ", testHash)
	}
}

func TestLargeFileOld(t *testing.T) {
	test := bytes.NewReader(make([]byte, 19457000))
	validHash := "345da2ffa0f63eae5638b908f187bfb1"
	testHash, _ := Hash(test, true)
	if testHash != validHash {
		t.Error("Expected ", validHash, " got ", testHash)
	}
}
