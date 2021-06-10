package service

import (
	"context"
	"errors"

	"github.com/disturb16/go-sqlite-service/internal/persons/entity"
)

var errNameEmpty = errors.New("name is empty")
var errAgeNotValid = errors.New("age is not valid")

func (s Service) SavePerson(ctx context.Context, name string, age int) (int64, error) {

	if name == "" {
		return -1, errNameEmpty
	}

	if age < 18 {
		return -1, errAgeNotValid
	}

	p := entity.Person{
		Name: name,
		Age:  age,
	}

	return s.repo.SavePerson(ctx, p)
}

func (s Service) Persons(ctx context.Context, limit int) ([]entity.Person, error) {
	return s.repo.Find(ctx, limit)
}

func (s Service) Person(ctx context.Context, id int64) (*entity.Person, error) {
	return s.repo.FindOne(ctx, id)
}
