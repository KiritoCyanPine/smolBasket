package main

import (
	"bytes"
	"log"

	"github.com/KiritoCyanPine/smolBasket/configuration"
	bae "github.com/KiritoCyanPine/smolBasket/encoder"
	"github.com/KiritoCyanPine/smolBasket/handler"
	"github.com/KiritoCyanPine/smolBasket/storage"
	"github.com/panjf2000/gnet"
)

type kvServer struct {
	*gnet.EventServer
	storageManager storage.Manager
	enc            bae.Encoder
}

func (s *kvServer) OnInitComplete(server gnet.Server) (action gnet.Action) {
	log.Printf("KV Server started on %s [multi-core: %v]", server.Addr.String(), server.Multicore)
	return
}

func (s *kvServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {

	if len(frame) == 0 {
		return s.enc.EncodeBAEError(handler.ErrEmptyFrame), gnet.None
	}

	commands, err := s.enc.DecodeBAE(bytes.NewReader(frame))
	if err != nil {
		return s.enc.EncodeBAEError(err), gnet.None
	}

	// route Handlers
	data, err := handler.Handler(s.enc, s.storageManager, commands)
	if err != nil {
		return handler.HandleError(s.enc, err)
	}

	return data, gnet.None
}

func main() {
	// Initialize the storage manager and BAE encoder
	storageManager := storage.NewStorageManager()
	encoder := bae.BaeEncoder{}

	// Create the server instance
	server := &kvServer{
		storageManager: storageManager,
		enc:            encoder,
	}

	config := configuration.GetConfiguraation()

	// Define the server address and options
	address := "tcp://:" + config.Port

	options := []gnet.Option{
		gnet.WithMulticore(config.MultiCore),
		gnet.WithReusePort(config.ReusePort),
		gnet.WithLoadBalancing(config.LoadBalancing),
	}

	// Start the server and handle errors
	log.Printf("Starting KV Server on %s...", address)
	if err := gnet.Serve(server, address, options...); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
