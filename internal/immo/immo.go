package immo

import (
	pb "Immortals/api/gRPC/immo/immo"
	db "Immortals/internal/database/sqlite/sqlc"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	pb.ImmoServiceServer
	store *db.Store
}

var ErrNodeExists = errors.New("node already exists")

func (s *Server) DiscoverNode(ctx context.Context, req *pb.NodeRequest) (*pb.NodeResponse, error) {
	log.Printf("Discovering node with address %s\n", req.Address)
	nodeData, err := db.Discover(req.Name, req.Address)
	if err != nil {
		log.Printf("Error Discovering node with address %s\n : %s", req.Address, err)
		return nil, err
	}
	nodeResponse := &pb.NodeResponse{}
	nodeResponse.Name = nodeData.Name.String
	// nodeResponse.Sensor = DiscoverSensor(nodeData.Sensor)
	// nodeResponse.Actuator = DiscoverActuator(nodeData.Actuator)
	fmt.Println("node Data:", nodeData)

	return nodeResponse, nil
}

func (s *Server) AddNode(ctx context.Context, req *pb.NodeRequest) (*pb.NodeResponse, error) {
	log.Printf("Adding node with address %s and name %s\n", req.Address, req.Name)
	node, err := s.store.GetNodeByClientID(context.Background(), sql.NullString{String: req.Address, Valid: true})

	if err != nil {
		fmt.Printf("Error Adding Node: %s", err)
	}

	if node.ID > 0 {
		fmt.Println("Node is Already Exists")
		return nil, ErrNodeExists
	}

	_, err = s.DiscoverNode(context.Background(), req)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	arg := db.CreateNodeParams{
		Name:     sql.NullString{String: req.Name, Valid: true},
		ClientID: sql.NullString{String: req.Address, Valid: true},
	}

	node, err = s.store.CreateNode(context.Background(), arg)
	if err != nil {
		return nil, err
	}

	return &pb.NodeResponse{
		Id:    node.ID,
		Error: ""}, nil
}

func (s *Server) ListNodes(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	listResponse := &pb.ListResponse{}

	nodes, err := s.store.ListNodes(context.Background())
	if err != nil {
		return nil, err
	}
	for _, node := range nodes {
		nodeResponse := &pb.NodeResponse{}
		nodeResponse.Id = node.ID
		nodeResponse.Name = node.Name.String
		listResponse.Nodes = append(listResponse.Nodes, nodeResponse)
	}
	fmt.Printf("Nodes : %v\n", nodes)
	return listResponse, nil
}

func (s *Server) RemoveNode(ctx context.Context, req *pb.NodeRequest) (*pb.RemoveResponse, error) {
	removeResponse := &pb.RemoveResponse{}

	err := s.store.DeleteNodeByClientID(context.Background(), sql.NullString{String: req.Address, Valid: true})
	if err != nil {
		log.Printf("Error in deleting node : %s", err)
		return nil, err
	}
	fmt.Printf("Node Removed %s\n", req.Address)
	removeResponse.Status = true
	return removeResponse, nil
}

func SetupImmo(store *db.Store) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterImmoServiceServer(s, &Server{store: store})

	log.Println("Immo server started: localhost:50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
