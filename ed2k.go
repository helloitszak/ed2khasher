package ed2k

import (
	"bufio"
	"encoding/hex"
	md4 "golang.org/x/crypto/md4"
	"io"
)

var BLOCK_SIZE uint = 9500 * 1024

func Hash(reader Reader) (string, error) {
	buffer := make([]byte, BLOCK_SIZE)

	// read BLOCK_SIZE
	var blocks string

	hasher := md4.New()
	writer := bufio.NewWriter(hasher)
	for {
		hasher.Reset()
		num, err := reader.Read(buffer)
		writer.Write(buffer)
		blocks += hex.EncodeToString(hasher.Sum(nil))

		if err == io.EOF {
			break
		}
	}

	return blocks, nil
}
