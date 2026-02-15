package cmd

import (
	"todolist/config"
	"todolist/repo"
	"todolist/rest"
	"todolist/rest/handler/todoHandler"
	"todolist/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()

	m := middlewares.NewMiddlewares(cnf)

	// create a task repository and pass it to the handler
	repo := repo.NewTaskRepo()
	todoHandler := todoHandler.NewHandler(m, repo)
	server := rest.NewServer(cnf, todoHandler)

	server.Start()
}
