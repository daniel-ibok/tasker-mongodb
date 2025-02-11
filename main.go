package main

import (
	"log"
	"task-manager/cli"
	"task-manager/database"
)

func main() {

	// mongodb database instance
	err := database.NewDBInstance()
	if err != nil {
		log.Fatal(err)
	}

	// start task cli
	cli.TaskCLI()
	if err := cli.Run(); err != nil {
		log.Fatal(err)
	}
}
