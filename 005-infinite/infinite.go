package main

import (
	"fmt"
	"time"
)

func doSomething() string {
	return "Test"
}

func main() {
	/*counter := 1
	for {
		if counter > 300 {
			break
		}

		fmt.Printf("This is printed in a loop (% 3d of 300 executions)\n", counter)
		counter++
		time.Sleep(666 * time.Millisecond)
	}*/
	c := time.Tick(5 * time.Second)
	for now := range c {
		fmt.Printf("%v %s\n", now, func() string {
			return "This is an inline function, bitches"

		}())
	}
}
