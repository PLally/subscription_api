package main

import (
	"context"
	"fmt"
	"github.com/plally/subscription_api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		"subrpc.foxorsomething.net",
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

	sub, err := c.Subscribe(context.Background(), &proto.Subscription{
		Destination:        &proto.Destination{
			Identifier: "509187820980797470",
			Type:       "discord",
		},
		SubscriptionSource: &proto.SubscriptionSource{
			Tags: "letodoesart",
			Type: "e621",
		},
	})
	if err != nil {
		fmt.Println(err)
		fmt.Println(status.Code(err) == codes.AlreadyExists)
	}

	fmt.Println(sub.Id)

}
