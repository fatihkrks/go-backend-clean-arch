package main

import (
	"example/go-backend-clean-architecture/bootstrap"
	"fmt"
)

func main() {
	app := bootstrap.App()
	env := app.Env
	fmt.Println(env.AppEnv, env.AccessTokenExpiryHour)

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	fmt.Println(db)
}
