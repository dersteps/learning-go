package main

import "fmt"
import "os"
import "github.com/akamensky/argparse"

func main() {
	parser := argparse.NewParser("Test", "Testing GO argparse")
	s := parser.String("s", "string", &argparse.Options{Required: true, Help: "Something"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	} else {
		fmt.Println(*s)
	}

}
