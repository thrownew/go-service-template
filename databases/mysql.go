package databases

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	// DefaultReadTimeout for db connections
	DefaultReadTimeout = 3 * time.Second
	// DefaultWriteTimeout for db connections
	DefaultWriteTimeout = 5 * time.Second
)

type (
	SQLScanner interface {
		Scan(dest ...any) error
	}

	SQLExecutor interface {
		ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
		QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
		QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	}
)

func NewDB() (*sql.DB, error) {
	if err := upMigration(); err != nil {
		return nil, fmt.Errorf("migration: %w", err)
	}
	db, err := sql.Open("mysql", buildDSN())
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}
	return db, nil
}

func buildDSN() string {
	opts := fmt.Sprintf(
		"%s?charset=%s&timeout=%s&readTimeout=%s&writeTimeout=%s&transaction_isolation='%s'&rejectReadOnly=true&parseTime=true&loc=UTC&interpolateParams=true",
		withDefault(os.Getenv("DB_NAME"), "pupa"), "utf8mb4", DefaultWriteTimeout, DefaultReadTimeout, DefaultWriteTimeout, "READ-COMMITTED",
	)
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		withDefault(os.Getenv("DB_USER"), "pupa"),
		withDefault(os.Getenv("DB_PASSWORD"), "pupa"),
		withDefault(os.Getenv("DB_HOST"), "127.0.0.1"),
		withDefault(os.Getenv("DB_PORT"), "3306"),
		opts,
	)
}

func withDefault[T comparable](v, d T) T {
	var r T
	if v == r {
		return d
	}
	return v
}
