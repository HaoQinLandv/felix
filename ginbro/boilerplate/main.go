package main

import (
	"github.com/dejavuzhou/felix/ginbro/boilerplate/config"
	"github.com/dejavuzhou/felix/ginbro/boilerplate/handlers"
	"github.com/dejavuzhou/felix/ginbro/boilerplate/tasks"
)

func main() {
	if config.GetBool("app.enable_cron") {
		go tasks.RunTasks()
	}
	defer handlers.Close()
	handlers.ServerRun()
}
