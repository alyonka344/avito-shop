package seed

import (
	"embed"
	_ "embed"
	"fmt"
	"github.com/jmoiron/sqlx"
)

//go:embed *.sql
var seedFiles embed.FS

func ApplySeeds(db *sqlx.DB) error {
	seedData, err := seedFiles.ReadFile("init_merch.sql")
	if err != nil {
		return fmt.Errorf("reading seed file: %w", err)
	}

	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		var e error
		if err == nil {
			e = tx.Commit()
		} else {
			e = tx.Rollback()
		}

		if err == nil && e != nil {
			err = fmt.Errorf("finishing transaction: %w", e)
		}
	}()

	_, err = tx.Exec(string(seedData))
	if err != nil {
		return fmt.Errorf("applying seed data: %w", err)
	}

	return nil
}
