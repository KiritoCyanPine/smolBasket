package handler

import "errors"

var (
	ErrInvalidCommand   = errors.New("ERR Invalid Command")
	ErrNoBasketFound    = errors.New("ERR No Basket Found")
	ErrConnectionClosed = errors.New("ERR Connection Closed")
	ErrKeyNotFound      = errors.New("ERR Key Not Found")
	ErrInvalidKeyFormat = errors.New("ERR Invalid Key Format")
)
