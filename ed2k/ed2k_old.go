package ed2k

import (
	"bufio"
	"encoding/hex"
	md4 "golang.org/x/crypto/md4"
	"io"
)

func HashOld(reader io.Reader) (string, error) {
	buffer := make([]byte, BLOCK_SIZE)

	var blocks []byte

	hasher := md4.New()
	writer := bufio.NewWriter(hasher)
	for {
		hasher.Reset()
		count, err := reader.Read(buffer)
		if count > 0 {
			writer.Write(buffer[:count])
			writer.Flush()
			blocks = hasher.Sum(blocks)
		}

		if err == io.EOF {
			break
		} else if err != nil {
			return "", err
		}
	}

	if len(blocks) > 16 {
		hasher.Reset()
		writer.Write(blocks)
		writer.Flush()
		blocks = hasher.Sum(nil)
	}

	return hex.EncodeToString(blocks), nil
}
