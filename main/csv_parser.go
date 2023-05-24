package main

import (
	"bufio"
	"os"
)

func openFile(filename string) *bufio.Scanner {
	readFile, err := os.Open(filename)

	if err != nil {
		// do something
	}

	return bufio.NewScanner(readFile)

	//fileScaner.Scan()
	/*
		header := fileScaner.Text()

		fmt.Println(header)*/
}
