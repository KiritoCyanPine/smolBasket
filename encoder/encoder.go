package encoder

import "io"

type Encoder interface {
	// EncodeRESPCommand builds a RESP-compliant command like
	// `SET my
	// key myvalue`.
	EncodeRESPCommand(cmd ...string) []byte
	// DecodeRESP parses a single RESP reply (very basic, handles simple types).
	DecodeRESP(r io.Reader) ([]string, error)

	// EncodeRESPError encodes an error message in RESP format.
	EncodeRESPError(err error) []byte
}
