package main

import (
	"Authingo/app"

	config "Authingo/config/env"
	dbConfig "Authingo/config/env/db"
)
func main(){
	config.Load()
	cfg:=app.NewConfig()
	app:=app.NewApplication(cfg)
	dbConfig.SetupDB()
	app.Run()
}