package main

import (
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/plally/subscription_api/proto"
	"github.com/plally/subscription_api/types"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"net"
)

type server struct {
	proto.UnimplementedSubscriptionApiServer
	database *gorm.DB
}


func main() {
	types.RegisterE621()
	lis, err := net.Listen("tcp", ":8181")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(authFunc)),
	)

	proto.RegisterSubscriptionApiServer(s, &server{
		database: connectToDatabase(),
	})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
