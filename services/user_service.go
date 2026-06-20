package services

import (
	"Authingo/config/env"
	db "Authingo/db/repositories"
	"Authingo/models"
	"Authingo/utils"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)
type UserService interface {
	GetUserByID() error
	GetAllUsers() ([]*models.User,error)
	DeleteUserById(id int64) error
	CreateUser() error
	LoginUser() (string,error)
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
func (u *UserServiceImpl) LoginUser() (string,error){
	email:="user@example.com"
	password:="password_example"
user, err := u.userRepository.GetByEmail(email)

	if err != nil {
		fmt.Println("Error fetching user by email:", err)
		return "", err
	}

	// Step 2. If user exists, or not. If not exists, return error
	if user == nil {
		fmt.Println("No user found with the given email")
		return "", fmt.Errorf("no user found with email: %s", email)
	}

	// Step 3. If user exists, check the password using utils.CheckPasswordHash
	isPasswordValid := utils.CheckPasswordHash(password, user.Password)

	if !isPasswordValid {
		fmt.Println("Password does not match")
		return "", nil
	}

	// Step 4. If password matches, print a JWT token, else return error saying password does not match
	jwtPayload := jwt.MapClaims{
		"email": user.Email,
		"id":    user.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtPayload)

	tokenString, err := token.SignedString([]byte(env.GetString("JWT_SECRET", "TOKEN")))

	if err != nil {
		fmt.Println("Error signing token:", err)
		return "", err
	}

	fmt.Println("JWT Token:", tokenString)

	return tokenString, nil
}