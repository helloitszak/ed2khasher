package main

import (
	"flag"
	"fmt"
	"github.com/ubercow/ed2khasher/ed2k"
	"os"
	"path"
	"sync"
)

var wait sync.WaitGroup

func usage() {
	fmt.Fprintf(os.Stderr, "usage: ed2khasher [files]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func hashFile(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading file %s", file)
		os.Exit(1)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't stat %s (%s)", filename, err)
		os.Exit(1)
	}

	str, err := ed2k.Hash(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error hashing %s (%s)", filename, err)
		os.Exit(1)
	}

	return fmt.Sprintf("ed2k://|file|%s|%d|%s|\n", path.Base(filename), stat.Size(), str)
}

func hashFileOld(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading file %s", file)
		os.Exit(1)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't stat %s (%s)", filename, err)
		os.Exit(1)
	}

	str, err := ed2k.HashOld(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error hashing %s (%s)", filename, err)
		os.Exit(1)
	}

	return fmt.Sprintf("ed2k://|file|%s|%d|%s|\n", path.Base(filename), stat.Size(), str)
}

func pipe(filename string) {
	fmt.Printf(hashFile(filename))
	wait.Done()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Input file missing")
		os.Exit(1)
	}

	wait.Add(len(args))
	for _, file := range args {
		go pipe(file)
	}
	wait.Wait()
}
