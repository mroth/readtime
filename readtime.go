package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var wpm = flag.Int("r", 200, "average reading rate in words per minute")

func usage() {
	fmt.Printf("Usage : %s [-options] [inputfiles...]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	flag.Parse()

	if len(flag.Args()) < 1 {
		usage()
	}

	for _, f := range flag.Args() {
		rt := readtime(wordsInFile(f))

		if len(flag.Args()) > 1 {
			fmt.Printf("%3d min read\t%s\n", rt, f)
		} else {
			fmt.Printf("%d min read\n", rt)
		}
	}
}

// readtime converts number of words to estimate of read time in minutes
func readtime(words int) int {
	return words / *wpm + 1
}

// wordsInFile takes a filepath and returns the number of words
func wordsInFile(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	return countWords(scanner)
}

// countWords simply scans a buffer and returns the number of words
func countWords(s *bufio.Scanner) (words int) {
	s.Split(bufio.ScanWords)
	for s.Scan() {
		words++
	}
	return
}
