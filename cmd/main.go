package main

import (
	"github.com/doddeeph/todo-app/api/route"
	"github.com/doddeeph/todo-app/bootstrap"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	app := bootstrap.App()
	env := app.Env
	db := app.Mongo.Database(env.DBName)
	defer app.CloseDbConnection()
	timeout := time.Duration(env.ContextTimeout) * time.Second
	g := gin.Default()
	route.Setup(timeout, db, g)
	g.Run(env.ServerAddress)
}
