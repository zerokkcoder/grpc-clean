package handler

import (
	"context"
	"errors"

	"github.com/zerokkcoder/grpc-clean/internal/models"
	interfaces "github.com/zerokkcoder/grpc-clean/pkg/v1"
	pb "github.com/zerokkcoder/grpc-clean/proto"
	"google.golang.org/grpc"
)

type UserService struct {
	useCase interfaces.UseCaseInterface
	pb.UnimplementedUserServiceServer
}

func NewServer(grpcService *grpc.Server, usecase interfaces.UseCaseInterface) {
	userGrpc := &UserService{
		useCase: usecase,
	}
	pb.RegisterUserServiceServer(grpcService, userGrpc)
}

// Create
//
// This function creates a user with the data supplied
func (srv *UserService) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserProfileResponse, error) {
	data := srv.transformUserRPC(req)
	if data.Email == "" || data.Name == "" {
		return &pb.UserProfileResponse{}, errors.New("please provide all fields")
	}

	user, err := srv.useCase.Create(data)
	if err != nil {
		return &pb.UserProfileResponse{}, err
	}
	return srv.transformUserModel(user), nil
}

// Get
//
// This function returns the user instance
func (srv *UserService) Read(ctx context.Context, req *pb.SingleUserRequest) (*pb.UserProfileResponse, error) {
	id := req.GetId()
	if id == "" {
		return &pb.UserProfileResponse{}, errors.New("id cannot be blank")
	}
	user, err := srv.useCase.Get(id)
	if err != nil {
		return &pb.UserProfileResponse{}, err
	}
	return srv.transformUserModel(user), nil
}

func (srv *UserService) transformUserRPC(req *pb.CreateUserRequest) models.User {
	return models.User{Name: req.GetName(), Email: req.GetEmail()}
}

func (srv *UserService) transformUserModel(user models.User) *pb.UserProfileResponse {
	return &pb.UserProfileResponse{Name: user.Name, Email: user.Email}
}
