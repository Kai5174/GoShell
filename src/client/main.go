package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func deleteLastNewLine(msg string) string {
	var newString string
	if strings.Contains(msg, "\r") {
		newString = strings.ReplaceAll(msg, "\r\n", "")
	} else if strings.Contains(msg, "\n") {
		newString = strings.ReplaceAll(msg, "\n", "")
	} else {
		newString = msg
	}
	return newString
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("Dial target failed, err", err)
	}

	var command string
	var inputReader *bufio.Reader
	inputReader = bufio.NewReader(os.Stdin)
	var tmp [4096 * 1024]byte

	fmt.Print("Please input the authKey: ")
	command, err = inputReader.ReadString('\n')
	command = deleteLastNewLine(command)
	conn.Write([]byte(command))
	n, err := conn.Read(tmp[:])
	msg := string(tmp[:n])

	if msg == "Success" {
		for {
			fmt.Print("cmd> ")
			command, err = inputReader.ReadString('\n')
			command = deleteLastNewLine(command)
			if err != nil {
				fmt.Println("Readio failed with command, err:", err)
			}
			conn.Write([]byte(command))
			if command == "exit" {
				fmt.Println("User exit.")
				break
			}
			n, err := conn.Read(tmp[:])
			msg = string(tmp[:n])
			if err != nil {
				fmt.Println("Read result failed from remote side, err:", err)
			}
			fmt.Println(msg)
		}
	} else {
		fmt.Println("Auth failed")
	}
	fmt.Println("Bye :3")
	conn.Close()

}
