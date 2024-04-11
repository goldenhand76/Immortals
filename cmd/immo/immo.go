// immo.go

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	pb "Immortals/api/gRPC/immo/immo"

	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("server_addr", "localhost:50051", "The server address in the format of host:port")
)

func main() {
	flag.Parse()

	// Set up a connection to the server
	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a client for the Immo service
	client := pb.NewImmoServiceClient(conn)

	// Parse command-line flags
	discoverCmd := flag.NewFlagSet("deploy", flag.ExitOnError)
	nameFlag := discoverCmd.String("name", "", "node name")
	addressFlag := discoverCmd.String("address", "", "node address")

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  immo discover --name <name> --address <address>")
		fmt.Println("  immo list")
		fmt.Println("  immo stop --name <name>")
		fmt.Println("  immo remove --name <name>")
		return
	}

	// Parse the subcommand and execute the corresponding action
	switch os.Args[1] {
	case "discover":
		discoverCmd.Parse(os.Args[2:])
		if *addressFlag == "" || *nameFlag == "" {
			discoverCmd.PrintDefaults()
			return
		}
		resp, err := client.DiscoverNode(context.Background(), &pb.NodeRequest{
			Name:    *nameFlag,
			Address: *addressFlag,
		})
		if err != nil {
			fmt.Println("Error discovering node:", err)
			return
		}
		fmt.Printf("Container deployed with ID: %s\n", resp.NodeId)

	case "list":
		listCmd.Parse(os.Args[2:])
		resp, err := client.ListNodes(context.Background(), &pb.ListRequest{})
		if err != nil {
			fmt.Println("Error listing containers:", err)
			return
		}
		fmt.Printf("Discovered Nodes: %s\n", resp.Nodes)

	default:
		fmt.Println("Invalid command")
		return
	}
}
