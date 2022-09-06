package server

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func StarServer() {
	listen, err := net.Listen("tcp", "3001")
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()

	if err := s.Serve(listen); err != nil {
		log.Fatalln("failed starting server", err)
	}
}
