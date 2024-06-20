package libpsql

import (
	"database/sql"
	"errors"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

func UpMigration(dbname string, db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		return err
	}

	const migrationURL = "file://exercise4/db/migration/psql"

	var migration *migrate.Migrate

	migration, err = migrate.NewWithDatabaseInstance(migrationURL, dbname, driver)

	if err != nil {
		return err
	}

	if err = migration.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
