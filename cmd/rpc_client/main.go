package main

import (
	"context"
	"fmt"
	"github.com/plally/subscription_api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"log"
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
	jwtToken, _ := ioutil.ReadFile("jwt.key")
	creds := tokenAuth{string(jwtToken)}

	conn, err := grpc.Dial(
		"subrpc.foxorsomething.net:443",
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")),
		grpc.WithPerRPCCredentials(creds),
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
