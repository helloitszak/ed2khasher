package main

import (
	"flag"
	"fmt"
	"github.com/ubercow/ed2khasher/ed2k"
	"io"
	"log"
	"os"
	"path"
	"sync"
)

var wait sync.WaitGroup

func usage() {
	log.Print("usage: ed2khasher --pure [files]\n")
	flag.PrintDefaults()
	os.Exit(42)
}

func hashFile(file io.Reader) (string, error) {
	str, err := ed2k.Hash(file)
	if err != nil {
		return "", err
	}
	return str, nil
}

func hashFileOld(file io.Reader) (string, error) {
	str, err := ed2k.HashOld(file)
	if err != nil {
		return "", err
	}
	return str, nil
}

var pure bool

func pipe(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("error reading file %s", file)
	}
	defer file.Close()

	str, err := hashFile(file)
	if err != nil {
		log.Fatalf("error hashing %s (%s)", filename, err)
	}

	if !pure {
		stat, err := file.Stat()
		if err != nil {
			log.Fatalf("couldn't stat %s (%s)", filename, err)
		}
		str = fmt.Sprintf("ed2k://|file|%s|%d|%s|", path.Base(filename), stat.Size(), str)
	}

	fmt.Println(str)

	wait.Done()
}

func main() {
	flag.BoolVar(&pure, "pure", false, "Only print ED2K Hash")
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		log.Fatal("Input file missing")
	}

	wait.Add(len(args))
	for _, file := range args {
		go pipe(file)
	}
	wait.Wait()
}
