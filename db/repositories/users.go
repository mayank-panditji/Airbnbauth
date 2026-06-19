package db

import (
	"database/sql"
	"fmt"
	"Authingo/models"
)
type UserRepository interface{
		GetByID() (*models.User,error) 
		Create() (error)
		GetAll() ([]*models.User,error)
		DeleteById(id int64) error

}
type UsserRepositoryImpl struct{
		db *sql.DB
}
func NewUserRepository(_db *sql.DB) UserRepository{
	return &UsserRepositoryImpl{
			db:_db,
	}
}
func(u *UsserRepositoryImpl) GetAll() ([]*models.User,error){
	return nil,nil
}
func(u *UsserRepositoryImpl) DeleteById(id int64) error{
	return nil
}
func(u *UsserRepositoryImpl) Create() error{
	query:="INSERT INTO users(username,email,password) VALUES(?,?,?)"
	result,err:=u.db.Exec(query,"testusser","test@test.com","password123")
	if err!=nil{
		fmt.Println("error creating user",err)
		return err
	}
	rowsaffected,rowerr:=result.RowsAffected()
	if rowerr!=nil{
		fmt.Println("error getting rows affected",rowerr)
		return rowerr
	}
	if rowsaffected==0{
		fmt.Println("no rows affected,user are not created")
		return nil
	}
	fmt.Println("user created succesfully rows affected",rowsaffected)
	return nil
}
func(u *UsserRepositoryImpl) GetByID () (*models.User,error) {
	fmt.Println("fetching user in UserRepository")
	//prepare quesy
	query:="SELECT id,username,email,password,created_at,updated_at FROM users WHERE id=?"
//wexecute query
	row:=u.db.QueryRow(query,1)
	//process the result
	user:=&models.User{}
	err:=row.Scan(&user.Id,&user.Username,&user.Email,&user.Password,&user.CreatedAt,&user.UpdatedAt)
	if err!=nil{
		if err==sql.ErrNoRows{
			fmt.Println("user not found")
			return nil,err
		}else{
			fmt.Println("error scanning user",err)
			return nil,err
		}
	}
	//print the user detail
	fmt.Println("user fetched succesfully",user)
	return user,nil
}