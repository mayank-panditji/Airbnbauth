package controllers

import (
	"Authingo/services"
	"Authingo/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"Authingo/dto"
	"github.com/go-chi/chi/v5"
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
func (uc *UserController) GetAllUsers(w http.ResponseWriter,r *http.Request){
	fmt.Println("get all users endpoint")
	users,err:=uc.UserService.GetAllUsers()
	if err!=nil{
		http.Error(w,"error fetching users",http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(users)
}
func (uc *UserController) DeleteUserById(w http.ResponseWriter,r *http.Request){
	idParam:=chi.URLParam(r,"id")
	id,err:=strconv.ParseInt(idParam,10,64)
	if err!=nil{
		http.Error(w,"invalid id",http.StatusBadRequest)
		return
	}
	fmt.Println("delete user endpoint, id:",id)
	err=uc.UserService.DeleteUserById(id)
	if err!=nil{
		http.Error(w,"error deleting user",http.StatusInternalServerError)
		return
	}
	w.Write([]byte("user deleted succesfully"))
}
func (uc *UserController) CreateUser(w http.ResponseWriter,r *http.Request){
	fmt.Println("create user endpoint")
	uc.UserService.CreateUser()
	
	w.Write([]byte("user created succesfully"))
}
func (uc *UserController) LoginUser(w http.ResponseWriter,r *http.Request){
	fmt.Println("login user endpoint")
	var payload dto.LoginRequestDTO
	if jsonErr:=utils.ReadJson(r,&payload); jsonErr!=nil{
		utils.WriteJsonErrorResponse(w,http.StatusBadRequest,"Something went wrong while logging",jsonErr)
		return
	}
	if validationErr:=utils.Validator.Struct(payload);validationErr!=nil{
		utils.WriteJsonErrorResponse(w,http.StatusBadRequest,"invalid input data",validationErr)
		return
	}
	jwtToken,err:=uc.UserService.LoginUser(&payload)
	if err!=nil{
		utils.WriteJsonErrorResponse(w,http.StatusInternalServerError,"error logging in user",err)
		return
	}
	
	utils.WriteJsonSuccessResponse(w,http.StatusOK,"user logged in succesfully",jwtToken,)
}