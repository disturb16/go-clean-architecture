package service

import "github.com/disturb16/go-sqlite-service/internal/persons"

// Service is a struct able to access all data required
// to perform business logic functions
type Service struct {
	repo persons.Repository
}

// New constructs and returns a Service struct
func New(repo persons.Repository) persons.Service {
	return &Service{
		repo: repo,
	}
}
