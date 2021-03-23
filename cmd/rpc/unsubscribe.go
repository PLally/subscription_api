package main

import (
	"context"
	"errors"
	"github.com/plally/subscription_api/database"
	"github.com/plally/subscription_api/proto"
	"github.com/plally/subscription_api/subscription"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) UnSubscribe(ctx context.Context, subscriptionToDelete *proto.Subscription) (*proto.Success, error) {
	handler := subscription.GetSubTypeHandler(subscriptionToDelete.GetSubscriptionSource().GetType())
	if handler == nil {
		return &proto.Success{Success: false}, errors.New("handler does not exist")
	}
	tags, err := handler.Validate(subscriptionToDelete.GetSubscriptionSource().GetTags())
	if err != nil {
		return &proto.Success{Success: false}, err
	}

	subtype := database.SubscriptionType{
		Type: subscriptionToDelete.GetSubscriptionSource().GetType(),
		Tags: tags,
	}
	dest := database.Destination{
		DestinationType:    subscriptionToDelete.GetDestination().GetType(),
		ExternalIdentifier: subscriptionToDelete.GetDestination().GetIdentifier(),
	}
	s.database.First(&subtype, subtype)
	s.database.First(&dest, dest)

	sub := database.Subscription{
		SubscriptionTypeID: subtype.ID,
		DestinationID:      dest.ID,
	}

	if s.database.Delete(&sub, sub).RowsAffected == 0 {
		return &proto.Success{Success: false}, status.Error(codes.NotFound, "not found")
	}

	sub.Destination = dest
	sub.SubscriptionType = subtype

	return &proto.Success{Success: true}, nil

}