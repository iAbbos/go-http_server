package server

import (
	"fmt"
	"github.com/iAbbos/go-http_server/internal/delivery/tcp/handler"
	"github.com/iAbbos/go-http_server/internal/pkg/config"
	"net"
)

type Server struct {
	Config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		Config: config,
	}
}

func (s *Server) Run() error {
	l, err := net.Listen("tcp", s.Config.Server.Host+":"+s.Config.Server.Port)
	if err != nil {
		return fmt.Errorf("tsp server error on listen: %s %w", s.Config.Server.Port, err)
	}
	fmt.Printf("Server is running on %s:%s\n", s.Config.Server.Host, s.Config.Server.Port)
	for {
		conn, err := l.Accept()
		if err != nil {
			return fmt.Errorf("tcp server error on accept: %w", err)
		}
		go handler.HandleConnection(handler.HandleOption{
			Conn:   conn,
			Config: s.Config,
		})
	}
}
