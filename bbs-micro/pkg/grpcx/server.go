package grpcx

import (
	"google.golang.org/grpc"
	"net"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-11 11:56

type ServerX struct {
	Addr string
	*grpc.Server
}

func (s *ServerX) Serve() error {
	lis, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	return s.Server.Serve(lis)
}
