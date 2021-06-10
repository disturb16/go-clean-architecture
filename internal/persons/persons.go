package persons

import (
	"context"

	"github.com/disturb16/go-sqlite-service/internal/persons/entity"
)

// Service declares and summarizes the functionality a
// service in the containing package will implement
type Service interface {
	Persons(ctx context.Context, limit int) ([]entity.Person, error)
	Person(ctx context.Context, id int64) (*entity.Person, error)
	SavePerson(ctx context.Context, name string, age int) (int64, error)
}

// Repository declares and summarizes the functionality a
// repository in the containing package will implement
type Repository interface {
	Close() error
	Find(ctx context.Context, limit int) ([]entity.Person, error)
	FindOne(ctx context.Context, id int64) (*entity.Person, error)
	SavePerson(ctx context.Context, person entity.Person) (int64, error)
}
