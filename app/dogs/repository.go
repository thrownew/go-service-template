package dogs

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"pupa/databases"
)

const (
	selectDogSQL = `
SELECT
	d.id,
	d.created_at,
	d.updated_at,
	d.name,
	d.good_boy
FROM dogs d 
`
)

var (
	ErrorNotFound = errors.New("not found")
)

type (
	Repository struct {
		db *sql.DB
	}

	Dog struct {
		ID      int64
		Created time.Time
		Updated time.Time
		Name    string
		GoodBoy bool
	}
)

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) DogByName(ctx context.Context, name string) (Dog, error) {
	d, err := scanDog(r.db.QueryRowContext(ctx, selectDogSQL+" WHERE name = ?", name))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Dog{}, ErrorNotFound
		}
		return Dog{}, fmt.Errorf("query: %w", err)
	}
	return d, nil
}

func scanDog(s databases.SQLScanner) (Dog, error) {
	var d Dog
	err := s.Scan(
		&d.ID,
		&d.Created,
		&d.Updated,
		&d.Name,
		&d.GoodBoy,
	)
	if err != nil {
		return Dog{}, fmt.Errorf("scan: %w", err)
	}
	return d, nil
}
