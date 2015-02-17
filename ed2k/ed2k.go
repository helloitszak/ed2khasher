package ed2k

import (
	"encoding/hex"
	"golang.org/x/crypto/md4"
	"io"
)

const BLOCK_SIZE int = 9500 * 1024

func Hash(reader io.Reader) (string, error) {
	buffer := make([]byte, BLOCK_SIZE)

	blocks := make([]byte, 0)

	hasher := md4.New()
	for {
		count, err := reader.Read(buffer)

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

	if len(blocks) > 16 {
		hasher.Reset()
		hasher.Write(blocks)
		blocks = hasher.Sum(nil)
	}

	return hex.EncodeToString(blocks), nil
}
