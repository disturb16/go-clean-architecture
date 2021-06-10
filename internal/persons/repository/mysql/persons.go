package mysql

import (
	"context"

	"github.com/disturb16/go-clean-architecture/internal/persons/entity"
)

func (m Mysql) Find(ctx context.Context, limit int) ([]entity.Person, error) {
	return nil, nil
}

func (m Mysql) FindOne(ctx context.Context, id int64) (*entity.Person, error) {
	return nil, nil
}

func (m Mysql) SavePerson(ctx context.Context, person entity.Person) (int64, error) {
	return -1, nil
}
