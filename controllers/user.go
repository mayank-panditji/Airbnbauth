package controllers

import (
	"Authingo/services"
	"fmt"
	"net/http"
)
type UserController struct{
	UserService services.UserService
}
func NewUserController(_userService services.UserService) *UserController{
	return &UserController{
		UserService:_userService,
	}
}
func (uc *UserController) RegisterUser(w http.ResponseWriter,r *http.Request){
	fmt.Println("user registration end point")
	uc.UserService.CreateUser()
	w.Write([]byte("user registerd endpoint"))
}