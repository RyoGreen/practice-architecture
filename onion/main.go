package main

import (
	"log"
	"onion-architecture/infrastructure/postgresql"
	"onion-architecture/presentation"
)

type test struct{}

func main() {
	if err := postgresql.Init(); err != nil {
		log.Println(err)
		return
	}
	defer postgresql.CloseDB()
	repo := postgresql.NewJobRepositoryPostgres()

	presentation.Run(repo)
}
