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
var pure bool
var old bool
var anchor bool

func hashFile(file io.Reader) (string, error) {
	str, err := ed2k.Hash(file, old)
	if err != nil {
		return "", err
	}
	return str, nil
}

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
		basename := path.Base(filename)
		str = fmt.Sprintf("ed2k://|file|%s|%d|%s|", basename, stat.Size(), str)
		if anchor {
			str = fmt.Sprintf("<a href=\"%s\">%s</a>", str, basename)
		}
	}

	fmt.Println(str)

	wait.Done()
}

func usage() {
	log.Print("usage: ed2khasher [options] [files]\n")
	flag.PrintDefaults()
	os.Exit(42)
}

func main() {
	flag.BoolVar(&pure, "pure", false, "Only print ED2K Hash")
	flag.BoolVar(&old, "old", false, "Use old method of ed2k hashing")
	flag.BoolVar(&anchor, "anchor", false, "Wrap HTML Link around ED2K Link")
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
