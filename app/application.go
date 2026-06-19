package app

import (
	"Authingo/config/env"
	db "Authingo/db/repositories"
	"Authingo/router"
	"fmt"
	"net/http"
	"time"
	"Authingo/controllers"
	"Authingo/services"
)

// config holds the cofiguration for the server
type Config struct{
	Addr string
}
type Application struct{
    Config Config
	Store db.Storage
}
func NewConfig() Config{
	port:=env.GetString("PORT","8080")
	return Config{
		Addr:port,
	}
}
func NewApplication (cfg Config) *Application{
	return &Application{
		Config:cfg,
		Store:*db.NewStorage(),
	}
}

func (app *Application) Run() error{
	ur:=db.NewUserRepository()
	us:=services.NewUserService(ur)
	uc:=controllers.NewUserController(us)
	uRouter:=router.NewUserRouter(uc)
	server:=&http.Server{
			Addr:app.Config.Addr,
			Handler:router.SetupRouter(uRouter), //TODO setup a chi router and put it here
			ReadTimeout: 10*time.Second,
			WriteTimeout: 10*time.Second,
	}
	fmt.Println("Starting server on",app.Config.Addr)
	return server.ListenAndServe()
}