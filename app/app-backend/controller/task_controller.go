package controller

import (
	"fmt"
	"github.com/wkrzyzanowski/notion-task-integrator/logging"
	"github.com/wkrzyzanowski/notion-task-integrator/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wkrzyzanowski/notion-task-integrator/server"
)

const (
	GetTasksUrl   = "/api/tasks"
	GetTaskUrl    = "/api/tasks/:taskId"
	CreateTaskUrl = "/api/tasks"
)

type TasksController struct {
	Endpoints []server.ApiEndpoint
}

type TaskApiInterface interface {
	GetTasks() []model.Task
	GetTask() model.Task
	CreateTask(task model.Task) string
}

func NewTasksController() *TasksController {
	return &TasksController{
		Endpoints: endpoints,
	}
}

func (ctrl *TasksController) GetEndpoints() []server.ApiEndpoint {
	return ctrl.Endpoints
}

var endpoints = []server.ApiEndpoint{
	{
		HttpMethod:   http.MethodGet,
		RelativePath: GetTasksUrl,
		HandlerFunc: []gin.HandlerFunc{
			GetTasks(),
		},
	},
	{
		HttpMethod:   http.MethodGet,
		RelativePath: GetTaskUrl,
		HandlerFunc: []gin.HandlerFunc{
			GetTask(),
		},
	},
	{
		HttpMethod:   http.MethodPost,
		RelativePath: CreateTaskUrl,
		HandlerFunc: []gin.HandlerFunc{
			CreateTask(),
		},
	},
}

func GetTasks() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Empty list.",
		})
	}
}

func GetTask() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		taskId := ctx.Param("taskId")
		if taskId != "" {
			message := fmt.Sprintf("Parameter value is: %v", string(taskId))
			ctx.JSON(200, gin.H{
				"message": message,
			})
		} else {
			ctx.JSON(404, gin.H{
				"message": "Bad Request",
			})
		}
	}
}

func CreateTask() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var jsonData, err = ctx.GetRawData()
		if err == nil {
			message := fmt.Sprintf("Your message is: %v", string(jsonData))
			ctx.JSON(200, gin.H{
				"message": message,
			})
		} else {
			logging.Logger.Error(err.Error())
			ctx.JSON(404, gin.H{
				"message": "Bad Request",
			})
		}
	}
}
