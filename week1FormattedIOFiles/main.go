package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Take n input files as arguments from console and return the fileSize,numberOfWords
func main() {

	for _, fName := range os.Args[1:] {
		var wordsNumber, sizeOftheFile int
		fs, err := os.Open(fName)
		if err != nil {
			fmt.Println(err)
		}
		scan := bufio.NewScanner(fs)

		for scan.Scan() {
			text := scan.Text()
			sizeOftheFile += len(text)
			wordsNumber += len(strings.Split(text, " "))
		}
		fmt.Println(wordsNumber, sizeOftheFile)
		sizeOftheFile = 0
		wordsNumber = 0
	}
}
