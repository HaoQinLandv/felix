package main

import (
	"{{.ProjectPackage}}/handlers"
	"{{.ProjectPackage}}/tasks"
	"{{.ProjectPackage}}/config"
)

func main() {
	if config.GetBool("app.enable_cron") {
		go tasks.RunTasks()
	}
	defer handlers.Close()
	handlers.ServerRun()
}
