package rest

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"todolist/config"
	"todolist/rest/handler/todoHandler"
	"todolist/rest/middlewares"
)

type Server struct {
	config      *config.Config
	todohandler *todoHandler.Handler
}

func NewServer(
	config *config.Config,
	todoHandler *todoHandler.Handler,
) *Server {
	return &Server{
		config:      config,
		todohandler: todoHandler,
	}
}

func (server *Server) Start() {
	manager := middlewares.NewManager()

	manager.Use(
		middlewares.Cors,
		middlewares.Preflight,
		middlewares.Logger,
	)

	mux := http.NewServeMux()

	wrappedmux := manager.WrapMux(mux)

	server.todohandler.RegisterRoutes(mux, manager)

	addr := ":" + strconv.Itoa(int(server.config.HttpPort))

	fmt.Println("Server is running at the Port", addr)

	err := http.ListenAndServe(addr, wrappedmux)
	if err != nil {
		fmt.Println("Error loading the server", err)
		os.Exit(1)
	}
}
