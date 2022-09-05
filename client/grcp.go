package client

import (
	"google.golang.org/grpc"
)

type Client struct {
	rcpConn *grpc.ClientConn
}

func NewClient() (*Client, error) {
	conn, err := grpc.Dial("")
	if err != nil {
		return nil, err
	}

	return &Client{
		rcpConn: conn,
	}, nil
}
