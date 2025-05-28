package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Connected to KV server at localhost:9000")
	reader := bufio.NewReader(os.Stdin)
	serverReader := bufio.NewReader(conn)

	for {
		fmt.Print("> ")
		cmd, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Input error:", err)
			return
		}

		cmd = strings.TrimSpace(cmd)
		if cmd == "" {
			continue
		}

		_, err = conn.Write([]byte(cmd + "\n"))
		if err != nil {
			fmt.Println("Write error:", err)
			return
		}

		resp, err := serverReader.ReadString('\n')
		if err != nil {
			fmt.Println("Read error:", err)
			return
		}

		fmt.Printf("Response: %s", resp)
	}
}
