package encoder

import "io"

type Encoder interface {
	// EncodeBAECommand builds a RESP like command like
	// `SET my
	// key myvalue`.
	EncodeBAECommand(cmd ...string) []byte
	// DecodeBAE parses a single BAE reply (very basic, handles simple types).
	DecodeBAE(r io.Reader) ([]string, error)

	// EncodeBAEError encodes an error message in BAE format.
	EncodeBAEError(err error) []byte
}
