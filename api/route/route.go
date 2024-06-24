package route

import (
	"github.com/doddeeph/todo-app/mongo"
	"github.com/gin-gonic/gin"
	"time"
)

func Setup(timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")
	NewTaskRouter(db, timeout, publicRouter)

	//protectedRouter := gin.Group("")
	//protectedRouter.Use(middleware.JwtAuthMiddleware(""))
	//NewTaskRouter(db, timeout, protectedRouter)
}
