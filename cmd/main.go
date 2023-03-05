package main

import (
	"github.com/adrianwudev/sqltest/db"
	"github.com/adrianwudev/sqltest/domain"
	member "github.com/adrianwudev/sqltest/repository"
)

func main() {
	db := db.New()
	memRepo := member.NewPostgresMemberRepository(db)
	defer db.Close()

	memRepo.FetchAll()

	member := *domain.NewMember("Carlos", 27)
	memRepo.Add(member)
}
