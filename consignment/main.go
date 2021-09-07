package main


import (
	"log"
	"net"
	"github.com/karankumarshreds/GoMicroservices/consignment"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
)

const (
	port = ":50051"
)

type Repository struct {
	consignments []*pb.consignment.pb.go
}

