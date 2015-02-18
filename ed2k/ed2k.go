package ed2k

import (
	"encoding/hex"
	"golang.org/x/crypto/md4"
	"io"
)

const BLOCK_SIZE int = 9500 * 1024

func Hash(reader io.Reader, oldMethod bool) (string, error) {
	buffer := make([]byte, BLOCK_SIZE)
	blocks := make([]byte, 0)
	hasher := md4.New()
	totalBytes := 0

	for {
		count, err := reader.Read(buffer)
		totalBytes += count

		if err == io.EOF {
			break
		} else if err != nil {
			return "", err
		} else if count > 0 {
			hasher.Reset()
			hasher.Write(buffer[:count])
			blocks = hasher.Sum(blocks)
		}
	}

	if totalBytes%BLOCK_SIZE == 0 && oldMethod == true {
		hasher.Reset()
		hasher.Write([]byte{})
		blocks = hasher.Sum(blocks)
	}

	if len(blocks) > 16 {
		hasher.Reset()
		hasher.Write(blocks)
		blocks = hasher.Sum(nil)
	}

	return hex.EncodeToString(blocks), nil
}
