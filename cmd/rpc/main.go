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
	"os"
)

type server struct {
	proto.UnimplementedSubscriptionApiServer
	database *gorm.DB
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	args := []grpc.ServerOption{grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(authFunc))}
	if _, ok := os.LookupEnv("INSECURE"); !ok {
		creds, err := credentials.NewServerTLSFromFile("rpc-cert.pem", "rpc-key.pem")
		if err != nil {
			log.Fatal("error loading tls: ", err)
		}
		args = append(args, grpc.Creds(creds))
	}


	s := grpc.NewServer(args...)

	proto.RegisterSubscriptionApiServer(s, &server{
		database: connectToDatabase(),
	})

	types.RegisterE621()
	types.RegisterRSS()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
