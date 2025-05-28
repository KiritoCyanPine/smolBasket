package main

import (
	"log"
	"strings"
	"sync"

	"github.com/panjf2000/gnet"
)

type kvServer struct {
	*gnet.EventServer
	store map[string]string
	mutex sync.RWMutex
}

func (s *kvServer) OnInitComplete(server gnet.Server) (action gnet.Action) {
	log.Printf("KV Server started on %s [multi-core: %v]", server.Addr.String(), server.Multicore)
	return
}

func (s *kvServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	line := strings.TrimSpace(string(frame))
	parts := strings.Fields(line)

	if len(parts) == 0 {
		return []byte("ERR Empty Command\n"), gnet.None
	}

	switch strings.ToUpper(parts[0]) {
	case "SET":
		if len(parts) < 3 {
			return []byte("ERR Usage: SET key value\n"), gnet.None
		}
		key, value := parts[1], strings.Join(parts[2:], " ")
		s.mutex.Lock()
		s.store[key] = value
		s.mutex.Unlock()
		return []byte("OK\n"), gnet.None

	case "GET":
		if len(parts) != 2 {
			return []byte("ERR Usage: GET key\n"), gnet.None
		}
		key := parts[1]
		s.mutex.RLock()
		value, ok := s.store[key]
		s.mutex.RUnlock()
		if !ok {
			return []byte("(nil)\n"), gnet.None
		}
		return []byte(value + "\n"), gnet.None

	default:
		return []byte("ERR Unknown Command\n"), gnet.None
	}
}

func main() {
	s := &kvServer{
		store: make(map[string]string),
	}

	log.Fatal(gnet.Serve(s, "tcp://:9000", gnet.WithMulticore(true), gnet.WithReusePort(true)))
}
