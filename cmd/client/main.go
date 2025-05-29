package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"strings"

	"github.com/KiritoCyanPine/smolBasket/encoder"
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

		enc := encoder.RespEncoder{}

		line := strings.TrimSpace(cmd)
		parts := strings.Fields(line)

		_, err = conn.Write(enc.EncodeRESPCommand(parts...))
		if err != nil {
			fmt.Println("Write error:", err)
			return
		}

		var b bytes.Buffer

		for {
			buf, err := serverReader.ReadString('\n')
			if err == io.EOF {
				if len(line) > 0 {
					// fmt.Printf("EOF : %s", buf)
					b.WriteString(buf)
				}
				break
			}

			if err != nil {
				fmt.Println("Read error:", err)
				break
			}

			// fmt.Printf("%s", buf)
			b.WriteString(buf)

			if serverReader.Buffered() == 0 {
				break
			}
		}

		fmt.Println("Server response:", b.String())

		reply, err := enc.DecodeRESP(bytes.NewReader(b.Bytes()))
		if err != nil {
			fmt.Println("Decode error:", err)
			continue
		}

		fmt.Println("Decoded response:", reply)
	}
}
