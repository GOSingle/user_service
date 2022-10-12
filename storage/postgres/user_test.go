package postgres

import (
	"user_service/config"
	pb "user_service/genproto/user"

	"user_service/pkg/db"
	"user_service/storage/repo"

	"github.com/stretchr/testify/suite"
)

type UserSuiteTest struct {
	suite.Suite
	ClenUpFunc func()
	Repository repo.UserStorageI
}

func (s *UserSuiteTest) SetupSuite() {
	pgPool, cleanUpfunc := db.ConnectTODBForSuite(config.Load())

	s.Repository = NewUserRepo(pgPool)
	s.ClenUpFunc = cleanUpfunc
}

func (s *UserSuiteTest) TestUserCrud() {
	//create ...
	userCreate := pb.User{
		Name:     "new name",
		LastName: "new last name",
	}
	user, err := s.Repository.CreateUser(&userCreate)
	s.Nil(err)
	s.NotNil(user)
	//update ...
	updateUser := pb.User{
		Id:       user.Id,
		Name:     "new name2",
		LastName: "new lastname2",
	}

	updatedUser, err := s.Repository.UpdateUser(&updateUser)
	s.Nil(err)
	s.NotEmpty(updatedUser)
	//get ...
	getUser, err := s.Repository.GetUserById(&pb.GetId{Id: updatedUser.Id})
	s.Nil(err)
	s.NotEmpty(getUser)
	//delete ...
	deletedUser, err := s.Repository.DeleteUser(&pb.GetId{Id: getUser.Id})
	s.Nil(err)
	s.NotEmpty(deletedUser)
}
