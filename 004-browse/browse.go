package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func listDir(dir string, yields chan string, done chan bool) {
	err := filepath.Walk(dir, func(path string, file os.FileInfo, err error) error {
		yields <- path
		return err
	})
	if err != nil {
		yields <- err.Error()
	}
	close(yields)
	done <- true
}

func main() {
	tmp := "/data/local"

	yielded := make(chan string)
	done := make(chan bool)

	go listDir(tmp, yielded, done)

	go func() {
		for {
			item, stillOpen := <-yielded
			if stillOpen {
				fmt.Println("Entry: " + item)
			}
		}
	}()
	<-done
}
