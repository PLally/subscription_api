package main

import (
	"context"
	"github.com/plally/subscription_api/database"
	"github.com/plally/subscription_api/proto"
	"github.com/plally/subscription_api/subscription"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)



func (s *server) Subscribe(ctx context.Context, newSubscription *proto.Subscription) (*proto.Subscription, error) {
	handler := subscription.GetSubTypeHandler(newSubscription.GetSubscriptionSource().GetType())
	if handler == nil {
		return nil, ErrTypeDoesNotExist
	}

	tags, err := handler.Validate(newSubscription.GetSubscriptionSource().GetTags())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
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
		return nil, status.Error(codes.AlreadyExists, "subscription already created")
	}
	sub.SubscriptionType = subtype
	sub.Destination = dest

	return SubscriptionDatabaseToProto(&sub), nil
}
