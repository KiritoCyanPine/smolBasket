package encoder

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type RespEncoder struct{}

// EncodeCommandToRESP builds a RESP-compliant command like `SET mykey myvalue`.
func (enc RespEncoder) EncodeRESPCommand(cmd ...string) []byte {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("*%d\r\n", len(cmd)))
	for _, arg := range cmd {
		b.WriteString(fmt.Sprintf("$%d\r\n%s\r\n", len(arg), arg))
	}
	return b.Bytes()
}

// DecodeRESP parses a single RESP reply (very basic, handles simple types).
func (enc RespEncoder) DecodeRESP(r io.Reader) (string, error) {
	reader := bufio.NewReader(r)
	prefix, err := reader.ReadByte()
	if err != nil {
		return "", err
	}

	switch prefix {
	case '+': // Simple String
		line, _ := reader.ReadString('\n')
		return strings.TrimSuffix(line, "\r\n"), nil

	case '-': // Error
		line, _ := reader.ReadString('\n')
		return "", fmt.Errorf("RESP error: %s", strings.TrimSuffix(line, "\r\n"))

	case '$': // Bulk String
		line, _ := reader.ReadString('\n')
		length, _ := strconv.Atoi(strings.TrimSuffix(line, "\r\n"))
		if length == -1 {
			return "", nil
		}
		data := make([]byte, length+2) // includes \r\n
		io.ReadFull(reader, data)
		return string(data[:length]), nil

	case '*': // Array
		line, _ := reader.ReadString('\n')
		n, _ := strconv.Atoi(strings.TrimSuffix(line, "\r\n"))
		var result []string
		for i := 0; i < n; i++ {
			elem, err := enc.DecodeRESP(reader)
			if err != nil {
				return "", err
			}
			result = append(result, elem)
		}
		return strings.Join(result, " "), nil

	default:
		return "", fmt.Errorf("unknown RESP prefix: %c", prefix)
	}
}

func (enc RespEncoder) EncodeRESPError(err error) []byte {
	if err == nil {
		return nil
	}
	return []byte(fmt.Sprintf("-%s\r\n", err.Error()))
}
