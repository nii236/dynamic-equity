package api

import (
	"context"
	pb "dynamic-equity/proto/dynamicequity"
)

type AuthController struct{}

// SignIn user
func (c *AuthController) SignIn(context.Context, *pb.SignInRequest) (*pb.SignInResponse, error) {
	return nil, nil
}

// SignOut user
func (c *AuthController) SignOut(context.Context, *pb.SignOutRequest) (*pb.SignOutResponse, error) {
	return nil, nil
}

// SignUp user
func (c *AuthController) SignUp(context.Context, *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	return nil, nil
}
