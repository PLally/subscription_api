package main

import (
	"context"
	"fmt"
	"github.com/plally/subscription_api/proto"
	"google.golang.org/grpc"
	"log"
	"os"
)


type tokenAuth struct {
	token string
}

func (t tokenAuth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "Bearer " + t.token,
	}, nil
}

func (tokenAuth) RequireTransportSecurity() bool {
	return false
}


func main() {
	jwtToken := os.Getenv("subscription_api_token")
	credentials := tokenAuth{jwtToken}

	conn, err := grpc.Dial(
		"127.0.0.1:8181",
		grpc.WithInsecure(),
		grpc.WithPerRPCCredentials(credentials),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewSubscriptionApiClient(conn)

	r, err := c.ListSubscriptions(context.Background(), &proto.Destination{
		Identifier: "509187820980797470",
		Type:       "discord",
	})
	if err != nil {
		log.Fatalf("eeeeeee: %v", err)
	}
	for _, sub := range r.GetSubscriptions() {
		fmt.Println(sub.SubscriptionSource.Tags)
	}
}
