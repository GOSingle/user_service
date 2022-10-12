package service

import (
	"context"
	pb "user_service/genproto/user"
	l "user_service/pkg/logger"
	"user_service/storage"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	storage storage.IStorage
	logger  l.Logger
}

// NewProductService ...

func NewUserService(db *sqlx.DB, log l.Logger) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	user, err := s.storage.User().CreateUser(req)
	if err != nil {
		s.logger.Error("Error while insert", l.Any("Error insert user", err))
		return &pb.User{}, status.Error(codes.Internal, "something went wrong,please check product infto")
	}
	return user, nil
}

func (s *UserService) GetUserById(ctx context.Context, req *pb.GetId) (*pb.User, error) {
	user, err := s.storage.User().GetUserById(req)
	if err != nil {
		s.logger.Error("Error while insert", l.Any("Error insert user", err))
		return &pb.User{}, status.Error(codes.Internal, "something went wrong,please check product infto")
	}
	return user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	user, err := s.storage.User().UpdateUser(req)
	if err != nil {
		s.logger.Error("Error while updating", l.Any("Update", err))
		return &pb.User{}, status.Error(codes.InvalidArgument, "Please recheck user info")
	}
	return user, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.GetId) (*pb.Empty, error) {
	user, err := s.storage.User().DeleteUser(req)
	if err != nil {
		s.logger.Error("Error while updating", l.Any("Delete", err))
		return &pb.Empty{}, status.Error(codes.InvalidArgument, "Please recheck user info")
	}
	return user, nil
}
