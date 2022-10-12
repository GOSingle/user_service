package postgres

import (
	"fmt"
	pb "user_service/genproto/user"

	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

// NewProductRepo ...

func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) CreateUser(user *pb.User) (*pb.User, error) {
	userResp := pb.User{}
	err := r.db.QueryRow(`insert into users (name,last_name) values ($1,$2) returning id,name, last_name`, user.Name, user.LastName).Scan(&userResp.Id, &userResp.Name, &userResp.LastName)
	if err != nil {
		return &pb.User{}, err
	}
	return &userResp, nil
}

func (r *userRepo) GetUserById(id *pb.GetId) (*pb.User, error) {
	userResp := &pb.User{}
	err := r.db.QueryRow(`select id, name, last_name from users where id=$1`, id.Id).Scan(&userResp.Id, &userResp.Name, &userResp.LastName)
	if err != nil {
		return &pb.User{}, err
	}
	fmt.Println(id.Id)

	return userResp, nil
}

func (r *userRepo) UpdateUser(usr *pb.User) (*pb.User, error) {
	_, err := r.db.Exec(`UPDATE users SET name=$1, last_name=$2 where id=$3`, usr.Name, usr.LastName, usr.Id)
	if err != nil {
		return &pb.User{}, err
	}
	return &pb.User{}, nil
}

func (r *userRepo) DeleteUser(user *pb.GetId) (*pb.Empty, error) {
	_, err := r.db.Exec(`DELETE FROM users WHERE id = $1`, user.Id)
	if err != nil {
		return &pb.Empty{}, nil
	}
	return &pb.Empty{}, nil
}
