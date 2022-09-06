package client

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	rcpConn *grpc.ClientConn
}

func NewClient() (*Client, error) {
	conn, err := grpc.Dial("", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &Client{
		rcpConn: conn,
	}, nil
}
