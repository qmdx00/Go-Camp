package server

import (
	"context"
	"log"
	pb "week04/api/book/service/v1"
)

type bookServer struct {
	pb.UnimplementedBookServer
}

func (*bookServer) GetBooks(_ context.Context, in *pb.Request) (*pb.Response, error) {
	log.Printf("Received: %s", in.String())
	return &pb.Response{
		Items: nil,
	}, nil
}
