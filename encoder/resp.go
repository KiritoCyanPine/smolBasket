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

func (enc RespEncoder) DecodeRESP(r io.Reader) ([]string, error) {
	reader := bufio.NewReader(r)
	prefix, err := reader.ReadByte()
	buffer := make([]string, 0)

	if err != nil {
		return buffer, fmt.Errorf("failed to read RESP prefix: %w", err)
	}

	switch prefix {
	case '+': // Simple String
		line, err := reader.ReadString('\n')
		if err != nil || !strings.HasSuffix(line, "\r\n") {
			return buffer, fmt.Errorf("malformed simple string: %w", err)
		}
		buffer = append(buffer, strings.TrimSuffix(line, "\r\n"))
		return buffer, nil

	case '-': // Error
		line, err := reader.ReadString('\n')
		if err != nil || !strings.HasSuffix(line, "\r\n") {
			return buffer, fmt.Errorf("malformed error string: %w", err)
		}
		return buffer, fmt.Errorf("RESP error: %s", strings.TrimSuffix(line, "\r\n"))

	case '$': // Bulk String
		line, err := reader.ReadString('\n')
		if err != nil || !strings.HasSuffix(line, "\r\n") {
			return buffer, fmt.Errorf("malformed bulk string length: %w", err)
		}
		length, err := strconv.Atoi(strings.TrimSuffix(line, "\r\n"))
		if err != nil || length < -1 {
			return buffer, fmt.Errorf("invalid bulk string length: %w", err)
		}
		if length == -1 {
			return buffer, nil // Null bulk string
		}
		data := make([]byte, length+2) // includes \r\n
		if _, err := io.ReadFull(reader, data); err != nil {
			return buffer, fmt.Errorf("incomplete bulk string data: %w", err)
		}
		if !strings.HasSuffix(string(data), "\r\n") {
			return buffer, fmt.Errorf("malformed bulk string data")
		}
		buffer = append(buffer, string(data[:length]))
		return buffer, nil

	case '*': // Array
		line, err := reader.ReadString('\n')
		if err != nil || !strings.HasSuffix(line, "\r\n") {
			return buffer, fmt.Errorf("malformed array length: %w", err)
		}
		n, err := strconv.Atoi(strings.TrimSuffix(line, "\r\n"))
		if err != nil || n < 0 {
			return buffer, fmt.Errorf("invalid array length: %w", err)
		}
		var result []string
		for i := 0; i < n; i++ {
			elem, err := enc.DecodeRESP(reader)
			if err != nil {
				return buffer, fmt.Errorf("malformed array element: %w", err)
			}
			result = append(result, elem...)
		}
		return result, nil

	default:
		return buffer, fmt.Errorf("unknown RESP prefix: %c", prefix)
	}
}

func (enc RespEncoder) EncodeRESPError(err error) []byte {
	if err == nil {
		return nil
	}
	return []byte(fmt.Sprintf("-%s\r\n", err.Error()))
}
