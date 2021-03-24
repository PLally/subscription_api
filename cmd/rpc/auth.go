package main

import (
	"context"
	"crypto/rsa"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/plally/vulpes_authenticator/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var jwtPublicKey *rsa.PublicKey

func init() {
	var err error
	jwtPublicKey, err = auth.ReadPublicKey("jwt.key.pub")
	if err != nil {
		panic(err)
	}
}

func authFunc(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	success, err := auth.ValidateToken(jwtPublicKey, token)
	if !success || err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}

	return ctx, nil
}
