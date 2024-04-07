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

	stopCmd := flag.NewFlagSet("stop", flag.ExitOnError)
	stopName := stopCmd.String("name", "", "Container name")

	removeCmd := flag.NewFlagSet("remove", flag.ExitOnError)
	removeName := removeCmd.String("name", "", "Container name")

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
		err := discoverNode(*addressFlag, *nameFlag, client)
		if err != nil {
			fmt.Println("Error discovering node:", err)
			return
		}
		fmt.Println("Node discovered successfully")
	case "list":
		listCmd.Parse(os.Args[2:])
		err := listContainers()
		if err != nil {
			fmt.Println("Error listing containers:", err)
			return
		}
	case "stop":
		stopCmd.Parse(os.Args[2:])
		if *stopName == "" {
			stopCmd.PrintDefaults()
			return
		}
		err := stopContainer(*stopName)
		if err != nil {
			fmt.Println("Error stopping container:", err)
			return
		}
		fmt.Println("Container stopped successfully")
	case "remove":
		removeCmd.Parse(os.Args[2:])
		if *removeName == "" {
			removeCmd.PrintDefaults()
			return
		}
		err := removeContainer(*removeName)
		if err != nil {
			fmt.Println("Error removing container:", err)
			return
		}
		fmt.Println("Container removed successfully")
	default:
		fmt.Println("Invalid command")
		return
	}
}

func discoverNode(address, name string, client pb.ImmoServiceClient) error {

	// Call the DiscoverNode method on the server
	resp, err := client.DiscoverNode(context.Background(), &pb.NodeRequest{
		Name:    name,
		Address: address,
	})
	if err != nil {
		return err
	}
	// Print the response from the server
	fmt.Printf("Container deployed with ID: %s\n", resp.NodeId)
	return nil
}

func listContainers() error {
	// Implement listContainers logic here
	fmt.Println("Listing containers")
	return nil
}

func stopContainer(name string) error {
	// Implement stopContainer logic here
	fmt.Printf("Stopping container with name %s\n", name)
	return nil
}

func removeContainer(name string) error {
	// Implement removeContainer logic here
	fmt.Printf("Removing container with name %s\n", name)
	return nil
}
