package main

import (
	"time"

	"github.com/gin-gonic/gin"
	route "github.com/khoahase151217/go-clean-api-architecture/api/route"
	"github.com/khoahase151217/go-clean-api-architecture/bootstrap"
)

func main() {

	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	route.Setup(env, timeout, db, gin)

	gin.Run(env.ServerAddress)
}
