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

type Repository struct {
	consignments []*pb.Consignment
}

