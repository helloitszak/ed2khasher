package main

import (
	"flag"
	"fmt"
	"github.com/ubercow/ed2khasher/ed2k"
	"io"
	"log"
	"os"
	"path"
	"runtime"
)

func hashFile(file io.Reader, old bool) (string, error) {
	str, err := ed2k.Hash(file, old)
	if err != nil {
		return "", err
	}
	return str, nil
}

func pipe(filename string, pure bool, old bool) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("error reading file %s", file)
	}
	defer file.Close()

	str, err := hashFile(file, old)
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
}

func usage() {
	log.Print("usage: ed2khasher --pure [files]\n")
	flag.PrintDefaults()
	os.Exit(42)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var pure bool
	var old bool
	flag.BoolVar(&pure, "pure", false, "Only print ED2K Hash")
	flag.BoolVar(&old, "old", false, "Use old method of ed2k hashing")
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		log.Fatal("Input file missing")
	}

	for _, file := range args {
		pipe(file, pure, old)
	}
}
