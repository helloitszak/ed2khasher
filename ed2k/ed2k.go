package ed2k

import (
	"encoding/hex"
	"golang.org/x/crypto/md4"
	"io"
	"log"
	"runtime"
	"sync"
)

const BLOCK_SIZE int = 9500 * 1024

type Chunk struct {
	bytes    []byte
	position int
	hash     []byte
}

type CorrelatedChunks struct {
	chunkMap   ChunkMap
	totalBytes int
}

type ChunkMap map[int][]byte

func readChunksToChannel(reader io.Reader, channel chan Chunk) {
	for chunkCount := 0; ; chunkCount++ {
		buffer := make([]byte, BLOCK_SIZE)
		count, err := reader.Read(buffer)
		if err == io.EOF {
			close(channel)
			break
		} else if err != nil {
			log.Fatalf("error reading %s", err)
		} else if count > 0 {
			channel <- Chunk{bytes: buffer[:count], position: chunkCount}
		}
	}
}

func correlateChunksToMap(channel <-chan Chunk, chunkMap chan<- CorrelatedChunks) {
	correlatedChunks := CorrelatedChunks{chunkMap: make(ChunkMap), totalBytes: 0}
	for {
		chunk, ok := <-channel
		if !ok {
			chunkMap <- correlatedChunks
			return
		}

		correlatedChunks.chunkMap[chunk.position] = chunk.hash
		correlatedChunks.totalBytes += len(chunk.bytes)
	}
}

func hashChunks(input <-chan Chunk, output chan<- Chunk, wg *sync.WaitGroup) {
	defer wg.Done()
	hasher := md4.New()
	for {
		chunk, ok := <-input
		if !ok {
			return
		}

		hasher.Reset()
		hasher.Write(chunk.bytes)
		chunk.hash = hasher.Sum(nil)
		output <- chunk
	}
}

func Hash(reader io.Reader, oldMethod bool) (string, error) {
	chunks := make(chan Chunk, 50)
	hashes := make(chan Chunk, 50)

	go readChunksToChannel(reader, chunks)

	var wg sync.WaitGroup
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go hashChunks(chunks, hashes, &wg)
	}

	processedChan := make(chan CorrelatedChunks)
	go correlateChunksToMap(hashes, processedChan)
	wg.Wait()
	close(hashes)

	correlatedChunks := <-processedChan
	hashList := make([]byte, 16*len(correlatedChunks.chunkMap))

	for position, hash := range correlatedChunks.chunkMap {
		copy(hashList[position*16:(position+1)*16], hash)
	}

	hasher := md4.New()
	if correlatedChunks.totalBytes%BLOCK_SIZE == 0 && oldMethod == true {
		hasher.Reset()
		hasher.Write([]byte{})
		hashList = hasher.Sum(hashList)
	}

	var finalSum []byte
	if len(hashList) > 16 {
		hasher.Reset()
		hasher.Write(hashList)
		finalSum = hasher.Sum(nil)
	} else {
		finalSum = hashList
	}

	return hex.EncodeToString(finalSum), nil

}
