package immo

import (
	pb "Immortals/api/gRPC/immo/immo"
	db "Immortals/internal/database"
	"Immortals/pkg/models"
	"Immortals/pkg/node"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.ImmoServiceServer
	db.DbContext
}

func (s *server) DiscoverNode(ctx context.Context, req *pb.NodeRequest) (*pb.NodeResponse, error) {
	log.Printf("Discovering node with address %s\n", req.Address)
	nodeData, err := node.Discover(req.Name, req.Address)
	nodeResponse := &pb.NodeResponse{}
	nodeResponse.Id = nodeData.NodeID
	nodeResponse.Sensor = DiscoverSensor(nodeData.Sensor)
	nodeResponse.Actuator = DiscoverActuator(nodeData.Actuator)
	fmt.Println("node Data:", nodeData)
	if err != nil {
		return nil, err
	}
	return nodeResponse, nil
}

func (s *server) AddNode(ctx context.Context, req *pb.NodeRequest) (*pb.NodeResponse, error) {
	log.Printf("Adding node with address %s and name %s\n", req.Address, req.Name)
	nodeData, err := node.Add(s.DbContext, req.Name, req.Address)
	if err != nil {
		return nil, err
	}
	return &pb.NodeResponse{
		Id:       nodeData.NodeID,
		Sensor:   DiscoverSensor(nodeData.Sensor),
		Actuator: DiscoverActuator(nodeData.Actuator),
		Error:    ""}, nil
}

func (s *server) ListNodes(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	listResponse := &pb.ListResponse{}

	nodes, err := node.List(s.DbContext)
	if err != nil {
		return nil, err
	}
	for _, m := range nodes {
		nodeResponse := &pb.NodeResponse{}
		nodeResponse.Name = m.NodeName
		nodeResponse.Id = m.NodeID
		nodeResponse.Name = m.NodeName
		nodeResponse.Sensor = DiscoverSensor(m.Sensor)
		nodeResponse.Actuator = DiscoverActuator(m.Actuator)
		listResponse.Nodes = append(listResponse.Nodes, nodeResponse)
	}
	fmt.Println("Nodes : ")
	return listResponse, nil
}

func DiscoverSensor(sensors []models.Sensor) []*pb.SensorResponse {
	var sensorList []*pb.SensorResponse
	for _, sensor := range sensors {
		sensorResponse := &pb.SensorResponse{}
		sensorResponse.Name = sensor.Name
		sensorResponse.Topic = sensor.Topic
		sensorList = append(sensorList, sensorResponse)
	}
	return sensorList
}

func DiscoverActuator(actuators []models.Actuator) []*pb.ActuatorResponse {
	var actuatorList []*pb.ActuatorResponse
	for _, sensor := range actuators {
		actuatorResponse := &pb.ActuatorResponse{}
		actuatorResponse.Name = sensor.Name
		actuatorResponse.Topic = sensor.Topic
		actuatorList = append(actuatorList, actuatorResponse)
	}
	return actuatorList
}

func SetupImmo(dbContext db.DbContext) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterImmoServiceServer(s, &server{DbContext: dbContext})

	log.Println("Immo server started")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
