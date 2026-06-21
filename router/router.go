package router

import (
	"Authingo/controllers"
	"Authingo/middlewares"

	"github.com/go-chi/chi/v5"
)
type Router interface{
	Register(r chi.Router)
}
func SetupRouter(UserRouter Router) *chi.Mux{
	chirouter:=chi.NewRouter()
	chirouter.Use(middlewares.RequestLogger)
	chirouter.Get("/ping",controllers.PingHandler)
	UserRouter.Register(chirouter)
	return chirouter
}