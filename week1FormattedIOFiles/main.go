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
		var wordsNumber, sizeOftheFile, totalFiles, totalWords, totalSize int
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
		fmt.Println(wordsNumber, sizeOftheFile, fName)
		totalWords += wordsNumber
		totalSize += sizeOftheFile
		sizeOftheFile = 0
		wordsNumber = 0
		fs.Close()
		totalFiles += 1
		if totalFiles > 1 {
			fmt.Println(totalWords, totalSize, totalFiles)
		}
	}
}
