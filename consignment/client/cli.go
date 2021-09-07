package main

// This is a command line utility that will take a JSON consignment file nad interact with our 
// GRPC service 


import (
	"os"
	"log"
	"io/ioutil"
	"encoding/json"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	pb "github.com/karankumarshreds/GoMicroservices/consignment/proto"
)

const (
	grpcServerAddr = "localhost:50051"
	defaultFileName = "consignment.json"
)

func main() {
	conn, err := grpc.Dial(grpcServerAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot make grpc connection %v", err)
	}
	defer conn.Close()

	// create client instance 
	client := pb.NewShippingServiceClient(conn)

	file := defaultFileName
	if len(os.Args) > 1 {
		file = os.Args[1]
	}
	// converted consignment 
	consignment, err := parseFile(file)
	if err != nil {
		log.Fatal("Could not parse")
	}
	// send it to the grpc server 
	response, err := client.CreateConsignment(
		context.Background(),
		consignment,
	)
	if err != nil {
		log.Fatal("Could not invoke remote procedure %v", err)
	}
	log.Printf("Created consignment as %v", response)

}

// this function will parse the json file and return the type of consignment 
// defined in the proto to send it to the grpc server
func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err !=nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, nil
}