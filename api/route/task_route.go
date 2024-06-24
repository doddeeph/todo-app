package route

import (
	"github.com/doddeeph/todo-app/api/controller"
	"github.com/doddeeph/todo-app/domain"
	"github.com/doddeeph/todo-app/mongo"
	"github.com/doddeeph/todo-app/repository"
	"github.com/doddeeph/todo-app/service"
	"github.com/gin-gonic/gin"
	"time"
)

func NewTaskRouter(db mongo.Database, timeout time.Duration, rg *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db, domain.CollectionTask)
	tc := &controller.TaskController{TaskService: service.NewTaskService(tr, timeout)}
	rg.POST("/tasks", tc.Create)
	rg.GET("/tasks", tc.Fetch)
}
