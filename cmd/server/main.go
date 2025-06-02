package main

import (
	"bytes"
	"log"

	resp "github.com/KiritoCyanPine/smolBasket/encoder"
	"github.com/KiritoCyanPine/smolBasket/handler"
	"github.com/KiritoCyanPine/smolBasket/storage"
	"github.com/panjf2000/gnet"
)

type kvServer struct {
	*gnet.EventServer
	storageManager storage.Manager
	enc            resp.Encoder
}

func (s *kvServer) OnInitComplete(server gnet.Server) (action gnet.Action) {
	log.Printf("KV Server started on %s [multi-core: %v]", server.Addr.String(), server.Multicore)
	return
}

func (s *kvServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {

	if len(frame) == 0 {
		return []byte("ERR Empty Frame\n"), gnet.None
	}

	commands, err := s.enc.DecodeRESP(bytes.NewReader(frame))
	if err != nil {
		return s.enc.EncodeRESPError(err), gnet.None
	}

	// route Handlers
	data, err := handler.Handler(s.enc, s.storageManager, commands)
	if err != nil {
		return handler.HandleError(s.enc, err)
	}

	return data, gnet.None
}

func main() {
	// Initialize the storage manager and RESP encoder
	storageManager := storage.NewStorageManager()
	encoder := resp.RespEncoder{}

	// Create the server instance
	server := &kvServer{
		storageManager: storageManager,
		enc:            encoder,
	}

	// Define the server address and options
	address := "tcp://:9000"
	options := []gnet.Option{
		gnet.WithMulticore(true),
		gnet.WithReusePort(true),
	}

	// Start the server and handle errors
	log.Printf("Starting KV Server on %s...", address)
	if err := gnet.Serve(server, address, options...); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
