// wordcount is a word counter command line utilitary
// usage: wordcount [pathtofile]
// writes a wordcount.txt file

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// usage print usage string to standard output
func usage() {
	fmt.Println("Usage: wordcount some_file")
}

// source return the source from an arg string
func source(arg string) io.Reader {
	switch arg {
	case "-":
		return os.Stdin
	default:
		f, err := os.Open(arg)
		if err != nil {
			log.Fatalln(err)
		}
		return f
	}
}

// tokeniseFile scan r io.Reader and tokenise each lines
func tokeniseFile(r io.Reader, signal chan string) {

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		tokeniseLine(scanner.Text(), signal)
	}

	close(signal)
}

// tokeniseLine tokenises string `line` into chan `signal`
func tokeniseLine(line string, signal chan string) {
	for _, word := range strings.Fields(line) {
		signal <- word
	}
}

// main checks command line arguments and compute line count from them
func main() {

	if len(os.Args) < 2 {
		usage()
		return
	}

	signal := make([]chan string, len(os.Args[1:]))

	for i, arg := range os.Args[1:] {
		signal[i] = make(chan string)
		go tokeniseFile(source(arg), signal[i])
	}

	wordMap := map[string]int64{}
	for _, channel := range signal {
		for word := range channel {
			wordMap[word]++
		}
	}

	resultFile, err := os.Create("/tmp/wordcount.txt")
	if err != nil {
		log.Fatalln(err)
	}

	for word, count := range wordMap {
		resultFile.WriteString(fmt.Sprintln(word, count))
	}
}
