package main


import (
	"log"
	"net"
	pb "github.com/karankumarshreds/GoMicroservices/consignment/proto"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
)

const (
	port = ":50051"
)

type RepositoryInterface interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll(*pb.GetRequest) []*pb.Consignment
}

// Dummy repository instead of a database for now 
type Repository struct {
	consignments []*pb.Consignment
}

type Service struct {
	repo RepositoryInterface
}

func main() {
	// create repository instance (db instance in future)
	repo := &Repository{}

	// set up the grpc server 
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error while listening %v", err)
	}
	grpcServer := grpc.NewServer()
	// register our service on the grpc server, this will tie our implementation
	// into the auto-generated 'interface' code for our protobuf definition
	pb.RegisterShippingServiceServer(grpcServer, &Service{repo})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error while serving as GRPC %v", err)
	}

}

func (s *Service) CreateConsignment(ctx context.Context, in *pb.Consignment) (*pb.Response, error){
	// business logic to add the consignment and save it to our database (using helper method for this)
	c, err := s.repo.Create(in)
	if err != nil {
		return nil, err
	}
	return &pb.Response{ Created : true, Consignment : c }, nil
}

func (s *Service) GetConsignments(ctx context.Context, req *pb.GetRequest) (*pb.Response, error){
	// assuming there is no error for now
	return &pb.Response{ Consignments: s.repo.GetAll() }, nil
}

/*
	These are helper methods, these have nothing to do with the protobuf interface
	The functions mentioned above are making use of these helper methods and those 
	above functions need to have the same types and the names as defined in the proto
*/

// function to create the consignment and save it to the database 
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.consignments = append(repo.consignments, consignment)
	return consignment, nil
}

// function to get all the consignments saved in the repository
func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}
