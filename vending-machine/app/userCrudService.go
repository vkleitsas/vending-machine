package app

import (
	"errors"
	"vending-machine/db"
	"vending-machine/domain"
)

var coins = []int{5, 10, 20, 50, 100}

type UserCrudService struct {
	Repository domain.UserCrud
	DBClient   db.InitDBInterface
	Auth       domain.JwtClaimsInterface
}

func NewUserCrudService(c db.InitDBInterface, sqlStore domain.UserCrud, auth domain.JwtClaimsInterface) *UserCrudService {
	return &UserCrudService{
		Repository: sqlStore,
		DBClient:   c,
		Auth:       auth,
	}
}

func (s *UserCrudService) AuthenticateUser(u domain.User) (*string, error) {
	if u.Username == "" {
		return nil, errors.New("Username cannot be empty")
	}
	if u.Password == "" {
		return nil, errors.New("Password cannot be empty")
	}
	tx, err := s.DBClient.GetTransaction()
	if err != nil {
		return nil, err
	}
	user, err := s.Repository.GetUserByUsername(u, tx)
	if err != nil {
		s.DBClient.Rollback(tx)
		return nil, err
	}
	s.DBClient.Commit(tx)

	passwordCheck := CheckPasswordHash(u.Password, user.Password)
	if !passwordCheck {
		return nil, errors.New("Authentication failed")
	}
	user.Password = ""
	token, err := s.Auth.CreateToken(user.Username, *user)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (s *UserCrudService) GetUser(u domain.User) (*domain.User, error) {
	if u.ID <= 0 {
		return nil, errors.New("Invalid user ID")
	}
	tx, err := s.DBClient.GetTransaction()
	if err != nil {
		return nil, err
	}
	user, err := s.Repository.GetUserΒyID(u, tx)
	if err != nil {
		s.DBClient.Rollback(tx)
		return nil, err
	}
	s.DBClient.Commit(tx)
	return user, err
}

func (s *UserCrudService) CreateUser(u domain.User) (*domain.User, error) {
	if u.Username == "" || u.Password == "" || u.Deposit < 0 || u.Role == "" {
		return nil, errors.New("User fields cannot be empty")
	}
	if u.ID != 0 {
		return nil, errors.New("User ID must be empty")
	}

	if !containsInt(u.Deposit, coins) {
		return nil, errors.New("Deposit value must be one of {5,10,20,50,100}")
	}
	tx, err := s.DBClient.GetTransaction()
	if err != nil {
		return nil, err
	}
	retrievedUser, err := s.Repository.GetUserByUsername(u, tx)
	if err != nil {
		s.DBClient.Rollback(tx)
		return nil, err
	}
	if retrievedUser != nil {
		s.DBClient.Rollback(tx)
		return nil, errors.New("Username already exists")
	}
	u.Password, err = HashPassword(u.Password)
	if err != nil {
		s.DBClient.Rollback(tx)
		return nil, err
	}
	savedUser, err := s.Repository.CreateUser(u, tx)
	if err != nil {
		s.DBClient.Rollback(tx)
		return nil, err
	}
	s.DBClient.Commit(tx)
	savedUser.Password = ""
	return savedUser, nil

}

func (s *UserCrudService) Deposit(u domain.User, requestUser domain.User) (*domain.User, error) {
	if u.ID < 1 {
		return nil, errors.New("Invalid user ID")
	}
	if !containsInt(u.Deposit, coins) {
		return nil, errors.New("Deposit value must be one of {5,10,20,50,100}")
	}
	tx, err := s.DBClient.GetTransaction()
	if err != nil {
		return nil, err
	}
	retrievedUser, err := s.Repository.GetUserΒyID(u, tx)
	if err != nil {
		s.DBClient.Rollback(tx)
		return nil, err
	}
	if retrievedUser.ID != requestUser.ID {
		s.DBClient.Rollback(tx)
		return nil, errors.New("User can deposit only in their own account")
	}
	retrievedUser.Deposit = retrievedUser.Deposit + u.Deposit

	updatedUser, err := s.Repository.UpdateUser(*retrievedUser, tx)
	if err != nil {
		s.DBClient.Rollback(tx)
		return nil, err
	}
	s.DBClient.Commit(tx)
	return updatedUser, nil
}

func (s *UserCrudService) Reset(u domain.User, requestUser domain.User) (*domain.User, error) {
	if u.ID < 1 {
		return nil, errors.New("Invalid user ID")
	}
	tx, err := s.DBClient.GetTransaction()
	if err != nil {
		return nil, err
	}
	retrievedUser, err := s.Repository.GetUserΒyID(u, tx)
	if err != nil {
		s.DBClient.Rollback(tx)
		return nil, err
	}
	if retrievedUser.ID != requestUser.ID {
		s.DBClient.Rollback(tx)
		return nil, errors.New("User can reset only their own account")
	}
	retrievedUser.Deposit = 0

	updatedUser, err := s.Repository.UpdateUser(*retrievedUser, tx)
	if err != nil {
		s.DBClient.Rollback(tx)
		return nil, err
	}
	s.DBClient.Commit(tx)
	return updatedUser, nil
}
