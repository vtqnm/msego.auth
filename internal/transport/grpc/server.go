package grpc

import (
	"context"
	"errors"

	"github.com/vtqnm/msego.auth/internal/services"
	sso "github.com/vtqnm/msego.proto/gen/go/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	sso.UnimplementedAuthServer
	authService services.AuthService
}

func Register(GRPCServer *grpc.Server, authService services.AuthService) {
	sso.RegisterAuthServer(GRPCServer, &Service{authService: authService})
}

func (s *Service) Register(ctx context.Context, req *sso.RegisterRequest) (*sso.RegisterResponse, error) {
	if req.GetEmail() == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	if req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	uid, err := s.authService.Register(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		if errors.Is(err, storage.ErrUserExists) {
			return nil, status.Error(codes.AlreadyExists, "user already exists")
		}

		return nil, status.Error(codes.Internal, "failed to register user")
	}

	return &sso.RegisterResponse{UserId: uid}, nil
}

func (s *Service) Login(ctx context.Context, req *sso.LoginRequest) (*sso.LoginResponse, error) {
	if req.GetEmail() == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	if req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	token, err := s.authService.Login(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		if errors.Is(err, auth.ErrInvalidCredentials) {
			return nil, status.Error(codes.InvalidArgument, "Invalid email or password")
		}

		return nil, status.Error(codes.Internal, "failed to login")
	}

	return &sso.LoginResponse{Token: token}, nil
}
