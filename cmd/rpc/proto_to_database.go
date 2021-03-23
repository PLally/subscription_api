package main

import (
	"github.com/plally/subscription_api/database"
	"github.com/plally/subscription_api/proto"
)

func SubscriptionDatabaseToProto(sub *database.Subscription) *proto.Subscription {
	return &proto.Subscription{
		Destination: &proto.Destination{
			Identifier: sub.Destination.ExternalIdentifier,
			Type:       sub.Destination.DestinationType,
			Id:         uint32(sub.Destination.ID),
		},
		SubscriptionSource: &proto.SubscriptionSource{
			Tags: sub.SubscriptionType.Tags,
			Type: sub.SubscriptionType.Type,
			Id:   uint32(sub.SubscriptionType.ID),
		},
		Id: uint32(sub.ID),
	}
}
