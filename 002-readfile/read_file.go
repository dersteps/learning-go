package main

import (
	"bufio"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Open a file, check for errors (might not exist, for example)
	file, err := os.Open("/tmp/go-test.dat")
	check(err)

	// Defer the closing of file
	defer file.Close()

	bReader := bufio.NewReader(file)

}
