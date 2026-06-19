package db

import (
	
	"fmt"
)
type UserRepository interface{
		Create() error
}
type UsserRepositoryImpl struct{
		// db *sql.DB
}
func NewUserRepository() UserRepository{
	return &UsserRepositoryImpl{
			
	}
}
func(u *UsserRepositoryImpl) Create() error{
	fmt.Println("Creating user in UserRepository")
	return nil
}