package main

import (
	"context"
	"errors"
	"github.com/plally/subscription_api/database"
	"github.com/plally/subscription_api/proto"
	"github.com/plally/subscription_api/subscription"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"net"
	"os"
)

type server struct {
	proto.UnimplementedSubscriptionApiServer
	database *gorm.DB
}

func (s *server) Subscribe(ctx context.Context, newSubscription *proto.Subscription) (*proto.Subscription, error) {
	handler := subscription.GetSubTypeHandler(newSubscription.GetSubscriptionSource().GetType())

	tags, err := handler.Validate(newSubscription.GetSubscriptionSource().GetTags())
	if err != nil {

	}

	subtype := database.SubscriptionType{
		Type: newSubscription.GetSubscriptionSource().GetType(),
		Tags: tags,
	}
	dest := database.Destination{
		DestinationType:    newSubscription.GetDestination().GetType(),
		ExternalIdentifier: newSubscription.GetDestination().GetIdentifier(),
	}

	s.database.FirstOrCreate(&dest, dest)
	s.database.FirstOrCreate(&subtype, subtype)

	sub := database.Subscription{
		SubscriptionTypeID: subtype.ID,
		DestinationID:      dest.ID,
	}
	if s.database.FirstOrCreate(&sub, sub).RowsAffected == 0 {
		return nil, errors.New("subscription already created")
	}
	return &proto.Subscription{
		Destination:        &proto.Destination{
			Identifier: dest.ExternalIdentifier,
			Type:       dest.DestinationType,
			Id:         uint32(dest.ID),
		},
		SubscriptionSource: proto.SubscriptionSource{
			Tags: subtype.Tags,
			Type: subtype.Type,
			Id:   uint32(subtype.ID),
		},
		Id: uint32(sub.ID),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", os.Getenv("RPC_ADDR"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterSubscriptionApiServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
