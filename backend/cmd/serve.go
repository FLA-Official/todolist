package cmd

import (
	"fmt"
	"os"
	"todolist/config"
	"todolist/infra/db"
	"todolist/repo"
	"todolist/rest"
	"todolist/rest/handler/projectHandler"
	"todolist/rest/handler/projectMemberHandler"
	"todolist/rest/handler/taskHandler"
	"todolist/rest/handler/userHandler"
	"todolist/rest/middlewares"
	"todolist/service"
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
	projectmemberrepo := repo.NewProjectMemberRepo(dbCon)

	userService := service.NewUserService(userrepo)
	projectService := service.NewProjectService(projectrepo, projectmemberrepo)
	projectMemberService := service.NewProjectMemberService(projectmemberrepo, projectrepo)
	taskService := service.NewTaskService(taskrepo, projectrepo, projectmemberrepo)
	// projectmemberrepo := repo.NewProjectMemberRepo(dbCon)

	taskhandler := taskHandler.NewHandler(m, *taskService, *projectService, *projectMemberService)
	userhandler := userHandler.NewHandler(m, *userService)
	projectHandler := projectHandler.NewHandler(m, *projectService)
	projectMemberHandler := projectMemberHandler.NewHandler(m, *projectService, *projectMemberService)

	server := rest.NewServer(cnf, taskhandler, userhandler, projectHandler, projectMemberHandler)

	server.Start()
}
