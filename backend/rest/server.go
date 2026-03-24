package rest

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"todolist/config"
	"todolist/rest/handler/projectHandler"
	"todolist/rest/handler/projectMemberHandler"
	"todolist/rest/handler/taskHandler"
	"todolist/rest/handler/userHandler"
	"todolist/rest/middlewares"
	"todolist/utils"
)

type Server struct {
	config               *config.Config
	taskHandler          *taskHandler.Handler
	userHandler          *userHandler.Handler
	projectHandler       *projectHandler.Handler
	projectMemberHandler *projectMemberHandler.Handler
}

func NewServer(
	config *config.Config,
	taskHandler *taskHandler.Handler,
	userHandler *userHandler.Handler,
	projectHandler *projectHandler.Handler,
	projectMemberHandler *projectMemberHandler.Handler,
) *Server {
	return &Server{
		config:               config,
		taskHandler:          taskHandler,
		userHandler:          userHandler,
		projectHandler:       projectHandler,
		projectMemberHandler: projectMemberHandler,
	}
}

func (server *Server) Start() {
	manager := middlewares.NewManager()
	logger := utils.NewLogger()

	manager.Use(
		middlewares.Preflight,
		middlewares.Cors,
		middlewares.Recovery(logger),
		middlewares.Logger(logger),
		middlewares.RequestID,
	)

	mux := http.NewServeMux()

	wrappedmux := manager.WrapMux(mux)

	server.taskHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)
	server.projectHandler.RegisterRoutes(mux, manager)
	server.projectMemberHandler.RegisterRoutes(mux, manager)

	addr := ":" + strconv.Itoa(int(server.config.HttpPort))

	fmt.Println("Server is running at the Port", addr)

	err := http.ListenAndServe(addr, wrappedmux)
	if err != nil {
		fmt.Println("Error loading the server", err)
		os.Exit(1)
	}
}
