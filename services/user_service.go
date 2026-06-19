package services

import (
	db "Authingo/db/repositories"
	"fmt"
)
type UserService interface {
	GetUserByID() error
}
type UserServiceImpl struct {
	userRepository db.UserRepository
}
func NewUserService(_userRepository db.UserRepository) UserService{
	return &UserServiceImpl{
		userRepository:_userRepository,
	}
}
func (u *UserServiceImpl) GetUserByID() error{
	fmt.Println("creating user in userservice")
	u.userRepository.GetByID()
	return nil
}