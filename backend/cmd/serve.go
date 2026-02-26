package cmd

import (
	"fmt"
	"os"
	"todolist/config"
	"todolist/infra/db"
	"todolist/repo"
	"todolist/rest"
	"todolist/rest/handler/projectHandler"
	"todolist/rest/handler/taskHandler"
	"todolist/rest/handler/userHandler"
	"todolist/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewConnection(cnf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	m := middlewares.NewMiddlewares(cnf)

	// create a task repository and pass it to the handler
	taskrepo := repo.NewTaskRepo(dbCon)
	userrepo := repo.NewUserRepo(dbCon)
	projectrepo := repo.NewProjectRepo(dbCon)
	// projectmemberrepo := repo.NewProjectMemberRepo(dbCon)

	taskhandler := taskHandler.NewHandler(m, taskrepo)
	userhandler := userHandler.NewHandler(m, userrepo)
	projectHandler := projectHandler.NewHandler(m, projectrepo)

	server := rest.NewServer(cnf, taskhandler, userhandler, projectHandler)

	server.Start()
}
