package main

import (
	"context"
	"errors"
	"github.com/plally/subscription_api/database"
	"github.com/plally/subscription_api/proto"
	"github.com/plally/subscription_api/subscription"
)

func (s *server) Subscribe(ctx context.Context, newSubscription *proto.Subscription) (*proto.Subscription, error) {
	handler := subscription.GetSubTypeHandler(newSubscription.GetSubscriptionSource().GetType())
	if handler == nil {
		return nil, errors.New("handler does not exist")
	}
	tags, err := handler.Validate(newSubscription.GetSubscriptionSource().GetTags())
	if err != nil {
		return nil, err
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
	sub.SubscriptionType = subtype
	sub.Destination = dest

	return SubscriptionDatabaseToProto(&sub), nil
}
