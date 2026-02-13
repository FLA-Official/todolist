package cmd

import (
	"todolist/config"
	"todolist/rest"
	"todolist/rest/handler/todoHandler"
	"todolist/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()

	m := middlewares.NewMiddlewares(cnf)

	todoHandler := todoHandler.Newhandler(m)
	server := rest.NewServer(cnf, todoHandler)

	server.Start()
}
