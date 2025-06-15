package userservice

import (
	"context"
	"instargam/protogen/userpb"
)

type UserServiceServer struct {
	userpb.UnimplementedUserServiceServer
}

func (s *UserServiceServer) GetUsers(ctx context.Context, req *userpb.GetUsersRequest) (*userpb.GetUsersResponse, error) {
	users := []*userpb.User{
		{
			Id:       "asd",
			Username: "zxc",
			Email:    "qe",
			Phone:    "czxs",
		},
	}
	return &userpb.GetUsersResponse{
		Users: users,
	}, nil
}
