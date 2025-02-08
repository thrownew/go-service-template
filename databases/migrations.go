package databases

import (
	"embed"
	"errors"
	"fmt"
	"log/slog"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed migrations/*.sql
var migrationFS embed.FS

func upMigration() error {
	sourceDriver, err := iofs.New(migrationFS, "migrations")
	if err != nil {
		return fmt.Errorf("source driver: %w", err)
	}

	m, err := migrate.NewWithSourceInstance(
		"embed",
		sourceDriver,
		fmt.Sprintf("mysql://%s", buildDSN()),
	)
	if err != nil {
		return fmt.Errorf("source instance: %w", err)
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("up: %w", err)
	}

	slog.Info("Migrations: Up")

	return nil
}
