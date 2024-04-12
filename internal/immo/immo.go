package immo

import (
	pb "Immortals/api/gRPC/immo/immo"
	"Immortals/pkg/node"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.ImmoServiceServer
}

func (s *server) DiscoverNode(ctx context.Context, req *pb.NodeRequest) (*pb.NodeResponse, error) {
	log.Printf("Discovering node with address %s\n", req.Address)
	nodeData, err := node.Discover(req.Address)
	fmt.Println("node Data:", nodeData)
	if err != nil {
		return nil, err
	}
	return &pb.NodeResponse{
		NodeId:   req.Address,
		Sensor:   nodeData.Sensor,
		Actuator: nodeData.Actuator,
		Error:    ""}, nil
}

func (s *server) AddNode(ctx context.Context, req *pb.NodeRequest) (*pb.NodeResponse, error) {
	log.Printf("Adding node with address %s and name %s\n", req.Address, req.Name)
	nodeData, err := node.Add(req.Address)
	if err != nil {
		return nil, err
	}
	return &pb.NodeResponse{
		NodeId:   req.Address,
		Sensor:   nodeData.Sensor,
		Actuator: nodeData.Actuator,
		Error:    ""}, nil
}

func (s *server) ListNodes(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	err := node.List()
	if err != nil {
		return nil, err
	}
	fmt.Println("Nodes : ")
	return &pb.ListResponse{}, nil
}

func SetupImmo() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterImmoServiceServer(s, &server{})
	log.Println("Immo server started")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
