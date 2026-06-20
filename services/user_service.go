package services

import (
	db "Authingo/db/repositories"
	"Authingo/models"
	"Authingo/utils"
	"fmt"
)
type UserService interface {
	GetUserByID() error
	GetAllUsers() ([]*models.User,error)
	DeleteUserById(id int64) error
	CreateUser() error
	LoginUser() error
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
func (u *UserServiceImpl) GetAllUsers() ([]*models.User,error){
	fmt.Println("fetching all users in userservice")
	users,err:=u.userRepository.GetAll()
	if err!=nil{
		fmt.Println("error fetching all users in userservice",err)
		return nil,err
	}
	return users,nil
}
func (u *UserServiceImpl) DeleteUserById(id int64) error{
	fmt.Println("deleting user in userservice, id:",id)
	err:=u.userRepository.DeleteById(id)
	if err!=nil{
		fmt.Println("error deleting user in userservice",err)
		return err
	}
	return nil
}
func (u *UserServiceImpl) CreateUser() error{
	fmt.Println("creating user in userservice")
	password:="password_example"
	hashedPassword,err:=utils.HashPassword(password)
	if err!=nil{
		fmt.Println("error hashing password",err)
		return err
	}
	u.userRepository.Create(
		"username_example",
		"user@example.com",
		hashedPassword,
	)
	return nil
}
func (u *UserServiceImpl) LoginUser() error{
	response:=utils.CheckPasswordHash("password_example","$2a$10$ftaFjZRUn05STnDBQ8.aK.G1U8dPC8iUzfF94NrmVou1REa4NVuwy")
	fmt.Println("Login response:",response)
	return nil
}