package onion

import (
	"log"
	"onion-architecture/infrastructure/postgresql"
	"onion-architecture/presentation"
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
