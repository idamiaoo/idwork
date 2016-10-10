package main

import (
	"go/idwork/models"
	pb "go/idwork/uidgenerator"

	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	port = ":45454"
)

type server struct {
}

func newserver() *server {
	s := new(server)
	return s
}

func (s *server) NextUid(ctx context.Context, in *pb.Request) (*pb.Reply, error) {
	uid := UIDGenerator.GetID(in.Game)
	return &pb.Reply{
		Uid: uid,
	}, nil
}

func main() {
	models.Init()
	UIDInit()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUIDGenneratorServer(grpcServer, &server{})
	grpcServer.Serve(lis)
}
