package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	. "github.com/wkrzyzanowski/notion-task-integrator/logging"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

const (
	FrontendAppSourcesEnv    = "FRONTEND_APP_SOURCES"
	AppServerPortEnv         = "SERVER_PORT"
	DefaultWebappSourcesPath = "../webapp/dist/public"
	AppServerPortDefault     = "8080"
)

type ApiEndpoint struct {
	RelativePath string
	HttpMethod   string
	HandlerFunc  []gin.HandlerFunc
}

type ApiController interface {
	GetEndpoints() []ApiEndpoint
}

type ApiMiddleware struct {
	Name     string
	Function gin.HandlerFunc
}

type Server struct {
	instance *gin.Engine
}

var serverInstance = Server{
	instance: gin.Default(),
}

func GetServerInstance() *Server {
	return &serverInstance
}

func (s *Server) ServeWebApp() *Server {
	webappSources := os.Getenv(FrontendAppSourcesEnv)

	if webappSources == "" {
		webappSources = DefaultWebappSourcesPath
	}

	Logger.Sugar().Infof("Serving webapp from: %s", webappSources)
	serverInstance.instance.Use(static.Serve("/", static.LocalFile(webappSources, false)))
	return s
}

func (s *Server) RegisterGlobalHandlers(chain []ApiMiddleware) *Server {
	for _, element := range chain {
		Logger.Sugar().Infof("Register middleware: %v", element.Name)
		serverInstance.instance.Use(element.Function)
	}
	return s
}

func (s *Server) RegisterControllers(apiController []ApiController) *Server {
	for _, controller := range apiController {
		for _, endpoint := range controller.GetEndpoints() {

			switch x := endpoint.HttpMethod; x {
			case http.MethodGet:
				serverInstance.instance.GET(endpoint.RelativePath, endpoint.HandlerFunc...)
			case http.MethodPost:
				serverInstance.instance.POST(endpoint.RelativePath, endpoint.HandlerFunc...)
			case http.MethodPut:
				serverInstance.instance.PUT(endpoint.RelativePath, endpoint.HandlerFunc...)
			case http.MethodPatch:
				serverInstance.instance.PATCH(endpoint.RelativePath, endpoint.HandlerFunc...)
			case http.MethodDelete:
				serverInstance.instance.DELETE(endpoint.RelativePath, endpoint.HandlerFunc...)
			default:
				msg := fmt.Sprintf("Http Method misconfigured or not supported: %v", controller)
				log.Fatalln(msg)
			}

		}
	}
	return s
}

func (s *Server) Run() {

	serverPort := os.Getenv(AppServerPortEnv)

	if serverPort == "" {
		Logger.Sugar().Infof("Set app port to default value: %v", AppServerPortDefault)
		serverPort = fmt.Sprintf(":%s", AppServerPortDefault)
	} else {
		serverPort = fmt.Sprintf(":%s", serverPort)
	}

	Logger.Sugar().Infof("Frontend: http://localhost%s/", serverPort)
	Logger.Sugar().Infof("Backend: http://localhost%s/api/", serverPort)

	err := serverInstance.instance.Run(serverPort)

	if err != nil {
		log.Fatalln(err)
	}
}
