package ed2k

import (
	"bufio"
	"encoding/hex"
	"golang.org/x/crypto/md4"
	"io"
)

var BLOCK_SIZE int = 9500 * 1024

func Hash(rd io.Reader) (string, error) {
	reader := bufio.NewReader(rd)

	buffer := make([]byte, BLOCK_SIZE)

	hasher := md4.New()

	inner  := md4.New()

	last, err := reader.Read(buffer)
	if err != nil && err != io.EOF {
		return "", err
	}

	if last > 0 {
		hasher.Write(buffer[:last])
		inner.Write(hasher.Sum(nil))
		hasher.Reset()
	}

	if last < BLOCK_SIZE {
		return hex.EncodeToString(inner.Sum(nil)), nil
	}

	for {
		count, err := reader.Read(buffer)
		if count > 0 {
			hasher.Write(buffer[:count])
			inner.Write(hasher.Sum(nil))
			hasher.Reset()
		}


		if err == io.EOF {
			if last == BLOCK_SIZE {
				inner.Write(hasher.Sum(nil))
				hasher.Reset()
			}
			break;
		} else if err != nil {
			return "", err
		}

		last = count
	}

	return hex.EncodeToString(inner.Sum(nil)), nil
}
