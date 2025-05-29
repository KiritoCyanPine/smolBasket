package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	resp "github.com/KiritoCyanPine/smolBasket/encoder"
	"github.com/KiritoCyanPine/smolBasket/storage"
	"github.com/panjf2000/gnet"
)

type kvServer struct {
	*gnet.EventServer
	storageManager *storage.StorageManager
}

func (s *kvServer) OnInitComplete(server gnet.Server) (action gnet.Action) {
	log.Printf("KV Server started on %s [multi-core: %v]", server.Addr.String(), server.Multicore)
	return
}

func (s *kvServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {

	if len(frame) == 0 {
		return []byte("ERR Empty Frame\n"), gnet.None
	}

	// pass the frame to RESP decoder
	// Here we can use the RESP encoder/decoder to parse the command
	enc := resp.RespEncoder{}
	reply, err := enc.DecodeRESP(bytes.NewReader(frame))
	if err != nil {
		return enc.EncodeRESPError(err), gnet.None
	}

	fmt.Printf("Decoded frame: %v \n\t %T", reply, reply)

	replyArray := strings.Fields(reply)

	fmt.Println("Received command:", replyArray)
	// Check if the command is valid and has the correct number of arguments
	// Here we can implement a simple command parser
	// For simplicity, we will just check the first part of the command
	// and route it to the appropriate handler
	// In a real-world scenario, you would have a more complex command parser
	// and a command router that handles different commands
	// For now, we will just handle SET and GET commands

	// based on the command level Pass to appropriate Router

	return enc.EncodeRESPCommand("Value", " 1a awdawdjnaiwun %n %v \n\n\t asdawdaw", "12312312312"), gnet.None
}

func main2() {

	enc := resp.RespEncoder{}

	cmd := enc.EncodeRESPCommand("PING", "foo", "bar", "val2")
	fmt.Println("Encoded RESP Command:\n", string(cmd))

	// Simulate server response (like "+OK\r\n")
	// serverResp := []byte("+OK\r\n")
	reply, err := enc.DecodeRESP(bytes.NewReader(cmd))
	if err != nil {
		panic(err)
	}
	fmt.Println("Decoded Response:", reply)
	// s := &kvServer{
	// 	store: make(map[string]string),
	// }

	// log.Fatal(gnet.Serve(s, "tcp://:9000", gnet.WithMulticore(true), gnet.WithReusePort(true)))
}

func main() {
	s := &kvServer{
		storageManager: storage.NewStorageManager(),
	}

	log.Fatal(gnet.Serve(s, "tcp://:9000", gnet.WithMulticore(true), gnet.WithReusePort(true)))
}
