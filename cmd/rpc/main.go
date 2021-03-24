package main

import (
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/plally/subscription_api/proto"
	"github.com/plally/subscription_api/types"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"gorm.io/gorm"
	"net"
)

type server struct {
	proto.UnimplementedSubscriptionApiServer
	database *gorm.DB
}
q
func main() {
	types.RegisterE621()
	lis, err := net.Listen("tcp", ":8181")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	creds, err := credentials.NewServerTLSFromFile("rpc-cert.pem", "rpc-key.pem")
	if err != nil {
		log.Fatal("error loading tls: ", err)
	}

	s := grpc.NewServer(
		grpc.Creds(creds),
		grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(authFunc)),
	)

	proto.RegisterSubscriptionApiServer(s, &server{
		database: connectToDatabase(),
	})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
