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
func (uc *UserController) GetUserByID(w http.ResponseWriter,r *http.Request){
	fmt.Println("user fetching end point")
	uc.UserService.GetUserByID()
	w.Write([]byte("user registerd endpoint"))
}