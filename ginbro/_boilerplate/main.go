package main

import (
	"github.com/dejavuzhou/felix/ginbro/_boilerplate/config"
	"github.com/dejavuzhou/felix/ginbro/_boilerplate/handlers"
	"github.com/dejavuzhou/felix/ginbro/_boilerplate/tasks"
)

func main() {
	if config.GetBool("app.enable_cron") {
		go tasks.RunTasks()
	}
	defer handlers.Close()
	handlers.ServerRun()
}
