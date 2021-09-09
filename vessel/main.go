package main 

import (
	"net"
	"log"
	"errors"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	pb "github.com/karankumarshreds/GoMicroservices/vessel/proto"
)

const (
	port = ":50052"
)

// Dummy repository instead of a database for now 
type Repository struct {
	vessels []*pb.Vessel
}

type Service struct {
	repo Repository
}

func main() {
	// create repository instance 
	repo := Repository{}
	
	// setup grpc server 
	lis, err := net.Listen("tcp", port)
	logError("Error while listening", err)
	grpcServer := grpc.NewServer()
	// register our service on the grpc server, this will tie our application 
	// into the auto-generated 'interface' code for our protobuf definition
	pb.RegisterVesselServiceServer(grpcServer, &Service{repo})
	err = grpcServer.Serve(lis)
	logError("Error while serving as GRPC" ,err)

}

func (s * Service) FindAvailable(ctx context.Context, spec *pb.Specification) (*pb.Response, error) {
	vessel, err := s.repo.FindAvailableVessels(spec)
	if err != nil {
		return nil, err
	}
	return &pb.Response{ Vessel: vessel }, nil
}


/*
	These are helper methods, these have nothing to do with the protobuf interface
	The functions mentioned above are making use of these helper methods and those 
	above functions need to have the same types and the names as defined in the proto
*/

func (repo *Repository) FindAvailableVessels(requirement *pb.Specification) (*pb.Vessel, error) {
	for _, vessel := range repo.vessels {
		if requirement.Capacity <= vessel.Capacity && requirement.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}
	return nil, errors.New("no vessel found by that spec")
}

func logError(message string, err error) {
	if err != nil {
		log.Printf(message + "%v", err)
	}
}