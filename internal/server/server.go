package server

import (
	"log"
	"net"

	"github.com/ratheeshkumar/restaurant_admin_serviceV1/internal/handlers"
	pb "github.com/ratheeshkumar/restaurant_admin_serviceV1/internal/pb"
	"google.golang.org/grpc"
)

func NewGrpcServer(handler *handlers.AdminHandler) {
	log.Println("connecting to gRPC server")
	lis, err := net.Listen("tcp", ":8084")
	if err != nil {
		log.Fatal("error creating listener on port 8084")
	}
	grp := grpc.NewServer()
	pb.RegisterAdminServiceServer(grp, handler)

	log.Printf("listening on gRPC server 8084")
	if err := grp.Serve(lis); err != nil {
		log.Fatal("error connecting to gRPC server")

	}
}
