package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"
)

func decode(input string) string {
	result, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "<ERROR WHILE DECODING SERVER'S ANSWER> " + err.Error()
	}
	return string(result)
}

func main() {
	connection, err := net.Dial("tcp", "127.0.0.1:9999")

	if err != nil {
		panic(err)
	}
	defer connection.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		one := []byte{}
		connection.SetReadDeadline(time.Now())
		if _, closeErr := connection.Read(one); closeErr == io.EOF {
			fmt.Println("Server closed the connection, exiting")
			connection.Close()
			return
		} else {
			var zero time.Time
			connection.SetReadDeadline(zero)
		}

		fmt.Print("Enter command: ")

		txt, e := reader.ReadString('\n')

		if e != nil {
			fmt.Printf("Sorry, that did not work: %s\n", e.Error())
			continue
		}

		// Need to get in here to enable upload command at least

		fmt.Printf("Sending command '%s' to server...\n", strings.TrimSpace(txt))
		fmt.Fprintf(connection, txt)
		fmt.Fprintf(connection, "\n")

		// Read server's answer
		answer, rErr := bufio.NewReader(connection).ReadString('\n')
		if rErr != nil {
			fmt.Printf("Something went wrong with server at '%s': '%s'", connection.RemoteAddr().String(), rErr.Error())
			return
		}

		fmt.Printf("Server's answer:\n %s\n", decode(strings.TrimSpace(answer)))
	}

}
