package controller

import (
	"github.com/doddeeph/todo-app/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type TaskController struct {
	TaskService domain.TaskService
}

func (tc *TaskController) Create(ctx *gin.Context) {
	var task domain.Task
	err := ctx.ShouldBind(&task)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	userId := ctx.GetString("x-user-id")
	task.ID = primitive.NewObjectID()
	task.UserId, err = primitive.ObjectIDFromHex(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	err = tc.TaskService.Create(ctx, &task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, domain.SuccessResponse{Message: "Task created successfully"})
}

func (tc *TaskController) Fetch(ctx *gin.Context) {
	userId := ctx.GetString("x-user-id")
	tasks, err := tc.TaskService.FetchByUserId(ctx, userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}
