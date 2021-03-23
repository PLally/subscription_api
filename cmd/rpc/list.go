package main

import (
	"context"
	"github.com/plally/subscription_api/database"
	"github.com/plally/subscription_api/proto"
)

func (s *server) ListSubscriptions(ctx context.Context, destination *proto.Destination) (*proto.SubscriptionList, error){
	var subscriptions []database.Subscription
	dest := database.Destination{
		ExternalIdentifier: destination.Identifier,
		DestinationType:    destination.Type,
	}
	s.database.First(&dest, dest)
	db := s.database.Model(database.Subscription{}).Where("destination_id = ?", dest.ID)

	db = database.Subscription{}.DoJoins(db)
	db.Find(&subscriptions)
	
	protoSubscriptions := make([]*proto.Subscription, len(subscriptions), len(subscriptions))
	for i, sub := range subscriptions {
		protoSubscriptions[i] = SubscriptionDatabaseToProto(&sub)
	}
	return &proto.SubscriptionList{Subscriptions: protoSubscriptions}, nil
}