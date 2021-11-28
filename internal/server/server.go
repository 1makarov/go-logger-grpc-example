package server

import (
	"fmt"
	"github.com/1makarov/go-logger-rpc-example/internal/proto"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	server  *grpc.Server
	service proto.LoggerServer
}

func New(service proto.LoggerServer) *Server {
	return &Server{
		server:  grpc.NewServer(),
		service: service,
	}
}

func (s *Server) Start(port string) error {
	addr := fmt.Sprintf(":%s", port)

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	proto.RegisterLoggerServer(s.server, s.service)

	if err = s.server.Serve(listen); err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop() {
	s.server.GracefulStop()
}
