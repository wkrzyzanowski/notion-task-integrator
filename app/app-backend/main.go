package main

import (
	"github.com/wkrzyzanowski/notion-task-integrator/controller"
	"github.com/wkrzyzanowski/notion-task-integrator/logging"
	"github.com/wkrzyzanowski/notion-task-integrator/middleware"
	"github.com/wkrzyzanowski/notion-task-integrator/server"
	"os"
)

const (
	DefaultNotionApiUrl = "https://api.notion.com"
	NotionApiKey        = "NOTION_API_KEY"
)

func main() {
	if os.Getenv(NotionApiKey) == "" {
		logging.Logger.Error("NOTION_API_KEY is required!")
		os.Exit(1)
	}

	startServer()
}

func startServer() {
	server.
		GetServerInstance().
		RegisterGlobalHandlers(getGlobalMiddleware()).
		ServeWebApp().
		RegisterControllers(getControllers()).
		Run()
}

func getControllers() []server.ApiController {
	return []server.ApiController{
		controller.NewTasksController(),
	}
}

func getGlobalMiddleware() []server.ApiMiddleware {
	return []server.ApiMiddleware{
		middleware.NewRequestLoggerMiddleware(),
	}
}
