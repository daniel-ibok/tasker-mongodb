package main

import (
	"go-mongodb/cli"
	"go-mongodb/database"
	"log"
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
