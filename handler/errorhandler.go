package handler

import (
	"github.com/KiritoCyanPine/smolBasket/encoder"
	"github.com/panjf2000/gnet"
)

func HandleError(enc encoder.Encoder, err error) ([]byte, gnet.Action) {
	switch err {
	case ErrConnectionClosed:
		return enc.EncodeBAECommand("OK"), gnet.Close
	case ErrInvalidCommand:
		return enc.EncodeBAEError(err), gnet.None
	case ErrNoBasketFound:
		return enc.EncodeBAEError(err), gnet.None
	case ErrKeyNotFound:
		return enc.EncodeBAEError(err), gnet.None
	case ErrInvalidKeyFormat:
		return enc.EncodeBAEError(err), gnet.None
	default:
		return enc.EncodeBAEError(err), gnet.None
	}
}
