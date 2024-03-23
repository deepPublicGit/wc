package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func bytesCounter(filePath string) {
	fi, err := os.Stat(filePath)
	if err != nil {
		panic(err)
	}
	fmt.Println(fi.Size())
}
func readFile(filePath string) *bufio.Scanner {
	fi, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	return bufio.NewScanner(fi)
}
func linesCounter(fileScanner bufio.Scanner) {
	/*fi, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(fi)*/
	var count int

	for fileScanner.Scan() {
		count++
	}
	fmt.Println(count)
}

func wordsCounter(fileScanner bufio.Scanner) {
	/*fi, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(fi)*/
	fileScanner.Split(bufio.ScanWords)
	var count int

	for fileScanner.Scan() {
		count++
	}
	fmt.Println(count)
}

func charsCounter(fileScanner bufio.Scanner) {

	fileScanner.Split(bufio.ScanRunes)
	var count int

	for fileScanner.Scan() {
		count++
	}
	fmt.Println(count)
}

func main() {
	//defaultFlag := flag.String("", "", "return the total bytes in a file")
	var flagVar string

	flag.StringVar(&flagVar, "c", "", "return the total bytes in a file")
	flag.StringVar(&flagVar, "l", "", "return the total lines in a file")
	flag.StringVar(&flagVar, "w", "", "return the total words in a file")
	flag.StringVar(&flagVar, "m", "", "return the total characters in a file")

	/*bytesCountFlag := flag.String("c", "", "return the total bytes in a file")
	linesCountFlag := flag.String("l", "", "return the total lines in a file")
	wordsCountFlag := flag.String("w", "", "return the total words in a file")
	charsCountFlag := flag.String("m", "", "return the total characters in a file")*/

	flag.Parse()

	if flagVar == "" {
		flagVar = flag.Arg(0)
		if flagVar == "" {
			fmt.Println("No file specified!")
			return
		}
	}

	fileBytes, err := os.ReadFile(flagVar)
	flag := findFlag()

	switch flag {
	case "c":

	}
	/*	if len(os.Args) == 2 {
			fileScanner := readFile(os.Args[1])
			charsCounter(*fileScanner)
			linesCounter(*fileScanner)
			wordsCounter(*fileScanner)
		} else {
			if *bytesCountFlag != "" {
				bytesCounter(*bytesCountFlag)
			} else {
				if *linesCountFlag != "" {
					linesCounter()
				} else if *wordsCountFlag != "" {
					wordsCounter(*wordsCountFlag)
				} else if *charsCountFlag != "" {
					//charsCounter(*charsCountFlag)
				}
			}
		}

		fmt.Println("Done!")*/
}
