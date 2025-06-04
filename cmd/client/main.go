package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"strings"

	"github.com/KiritoCyanPine/smolBasket/configuration"
	"github.com/KiritoCyanPine/smolBasket/encoder"
	"github.com/KiritoCyanPine/smolBasket/handler"
	"github.com/chzyer/readline"
)

func main() {

	config := configuration.GetConfiguraation()
	conn, err := connectToServer("localhost:" + config.Port)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	rl, err := setupReadline()
	if err != nil {
		fmt.Println("Readline error:", err)
		return
	}
	defer rl.Close()

	fmt.Println("smolBasket CLI (with history). Type 'exit' to quit.")
	runClient(conn, rl)
}

func connectToServer(address string) (net.Conn, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func setupReadline() (*readline.Instance, error) {
	rl, err := readline.New("> ")
	if err != nil {
		return nil, err
	}
	return rl, nil
}

func runClient(conn net.Conn, rl *readline.Instance) {
	encoder := encoder.BaeEncoder{}

	for {
		line, err := rl.Readline()
		if err != nil {
			break // CTRL+C / CTRL+D
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if line == "exit" {
			break
		}

		if line == "help" || line == "HELP" {
			fmt.Println(handler.GetHelpText())
			continue
		}

		cmd := splitCommand(line)
		bae := encoder.EncodeBAECommand(cmd...)

		if err := sendCommand(conn, bae); err != nil {
			fmt.Println("Write error:", err)
			break
		}

		serverResponse, err := readServerResponse(conn)
		if err != nil {
			fmt.Println("Read error:", err)
			break
		}

		reply, err := encoder.DecodeBAE(bytes.NewReader([]byte(serverResponse)))
		if err != nil {
			fmt.Println("Decode error:", err)
			fmt.Println("Use the `HELP` command to know about smolBasket-cli commands")
			continue
		}

		fmt.Println("Decoded response:", reply, len(reply))
	}
}

func sendCommand(conn net.Conn, command []byte) error {
	_, err := conn.Write(command)
	return err
}

func readServerResponse(conn net.Conn) (string, error) {
	var b bytes.Buffer
	serverReader := bufio.NewReader(conn)

	for {
		buf, err := serverReader.ReadString('\n')
		if err == io.EOF {
			if len(buf) > 0 {
				b.WriteString(buf)
			}
			break
		}
		if err != nil {
			return "", err
		}

		b.WriteString(buf)

		if serverReader.Buffered() == 0 {
			break
		}
	}

	return b.String(), nil
}

func splitCommand(line string) []string {
	var parts []string
	var current strings.Builder
	inQuotes := false

	for _, char := range line {
		switch char {
		case '"':
			inQuotes = !inQuotes
		case ' ':
			if inQuotes {
				current.WriteRune(char)
			} else {
				if current.Len() > 0 {
					parts = append(parts, current.String())
					current.Reset()
				}
			}
		default:
			current.WriteRune(char)
		}
	}

	if current.Len() > 0 {
		parts = append(parts, current.String())
	}

	return parts
}
