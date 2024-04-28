package main

import (
	"architecture/onion/infrastructure/postgresql"
	"architecture/onion/presentation"
	"log"
)

func main() {
	if err := postgresql.Init(); err != nil {
		log.Println(err)
		return
	}
	defer postgresql.CloseDB()
	repo := postgresql.NewJobRepositoryPostgres()

	presentation.Run(repo)
}
