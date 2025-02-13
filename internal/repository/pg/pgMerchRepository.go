package pg

import (
	"avito-shop/internal/model"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type PgMerchRepository struct {
	db *sqlx.DB
}

func NewPgMerchRepository(db *sqlx.DB) *PgMerchRepository {
	return &PgMerchRepository{db: db}
}

func (p PgMerchRepository) GetByName(merchName string) (*model.MerchItem, error) {
	tx, err := p.db.Beginx()
	if err != nil {
		return nil, err
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

	query, args, err := squirrel.
		Select("*").
		From("merch").
		Where(squirrel.Eq{"name": merchName}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	var merch []model.MerchItem

	err = tx.Select(&merch, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	if len(merch) == 0 {
		return nil, fmt.Errorf("this merch item is not exist: %w", err)
	}

	return &merch[0], nil
}
