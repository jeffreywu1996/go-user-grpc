package handler

import (
	"context"

	"github.com/jeffreywu1996/go-user/internal/model"
	"github.com/jeffreywu1996/go-user/internal/service"
	pb "github.com/jeffreywu1996/go-user/proto/user"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := &model.User{
		Name:  req.Name,
		Email: req.Email,
	}

	if err := h.userService.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{
		User: &pb.User{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}

func (h *UserHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := h.userService.GetUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{
		User: &pb.User{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}
