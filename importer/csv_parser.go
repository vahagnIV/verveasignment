package main

import (
	"bufio"
	"os"
)

func openFile(filename string) (*bufio.Scanner, error) {
	readFile, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	return bufio.NewScanner(readFile), nil

}
