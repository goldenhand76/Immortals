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
	discoverCmd := flag.NewFlagSet("discover", flag.ExitOnError)
	dAddressFlag := discoverCmd.String("address", "", "node address")

	// Parse command-line flags
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	nameFlag := addCmd.String("name", "", "node name")
	aAddressFlag := addCmd.String("address", "", "node address")

	// Parse command-line flags
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	// Parse Remove command flags
	removeCmd := flag.NewFlagSet("remove", flag.ExitOnError)
	rAddressFlag := removeCmd.String("address", "", "Node Address")

	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  immo discover --address <address>")
		fmt.Println("  immo add --name <name> --address <address>")
		fmt.Println("  immo remove --address <address>")
		fmt.Println("  immo list")
		return
	}

	// Parse the subcommand and execute the corresponding action
	switch os.Args[1] {
	case "discover":
		discoverCmd.Parse(os.Args[2:])
		if *dAddressFlag == "" {
			discoverCmd.PrintDefaults()
			return
		}
		resp, err := client.DiscoverNode(context.Background(), &pb.NodeRequest{
			Address: *dAddressFlag,
		})
		if err != nil {
			fmt.Println("Error discovering node:", err)
			return
		}
		fmt.Printf("Node Successfully Discovered: %s\n", resp)

	case "add":
		addCmd.Parse(os.Args[2:])
		if *aAddressFlag == "" || *nameFlag == "" {
			addCmd.PrintDefaults()
			return
		}
		resp, err := client.AddNode(context.Background(), &pb.NodeRequest{
			Name:    *nameFlag,
			Address: *aAddressFlag,
		})
		if err != nil {
			fmt.Println("Error discovering node:", err)
			return
		}
		fmt.Printf("Node Successfully Added: %s\n", resp)

	case "remove":
		removeCmd.Parse(os.Args[2:])
		if *rAddressFlag == "" {
			removeCmd.PrintDefaults()
			return
		}
		resp, err := client.RemoveNode(context.Background(), &pb.NodeRequest{
			Address: *aAddressFlag,
		})
		if err != nil {
			fmt.Println("Error deleting node:", err)
			return
		}
		fmt.Printf("Node Successfully Removed: %s\n", resp)

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
