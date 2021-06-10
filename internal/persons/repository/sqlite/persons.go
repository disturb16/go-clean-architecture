package sqlite

import (
	"context"
	"database/sql"

	"github.com/disturb16/go-clean-architecture/internal/persons/entity"
)

func (sl Sqlite) Find(ctx context.Context, limit int) ([]entity.Person, error) {

	var qry string
	if limit > 0 {
		qry = "select id, name, age from persons limit $1"
	} else {
		qry = "select id, name, age from persons"
	}

	persons := []entity.Person{}

	err := sl.db.SelectContext(ctx, &persons, qry, limit)
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func (sl Sqlite) FindOne(ctx context.Context, id int64) (*entity.Person, error) {

	const qry string = "SELECT id, name, age FROM persons WHERE id = $1"
	p := &entity.Person{}

	err := sl.db.GetContext(ctx, p, qry, id)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return p, nil
}

func (sl Sqlite) SavePerson(ctx context.Context, person entity.Person) (int64, error) {

	const qry string = "insert into persons (name, age) values (:name, :age)"

	result, err := sl.db.NamedExecContext(ctx, qry, person)

	if err != nil {
		return -1, err
	}

	insertedId, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return insertedId, nil
}
