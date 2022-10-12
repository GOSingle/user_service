package repo

import (
	pb "user_service/genproto/user"
)

// UserStorageI ...
type UserStorageI interface {
	CreateUser(*pb.User) (*pb.User, error)
	GetUserById(*pb.GetId) (*pb.User, error)
	UpdateUser(*pb.User) (*pb.User, error)
	DeleteUser(*pb.GetId) (*pb.Empty, error)
}
