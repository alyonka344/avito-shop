package pg

import (
	"avito-shop/internal/model"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type PgPurchaseRepository struct {
	db *sqlx.DB
}

func NewPgPurchaseRepository(db *sqlx.DB) *PgPurchaseRepository {
	return &PgPurchaseRepository{db: db}
}

func (r PgPurchaseRepository) Create(purchase *model.Purchase) error {
	tx, err := r.db.Beginx()
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

	query, args, err := squirrel.
		Insert("purchases").
		Columns("username", "merch_name", "created_at").
		Values(purchase.UserName, purchase.MerchName, purchase.CreatedAt).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	err = tx.QueryRowx(query, args...).Scan(&purchase.ID)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}

func (r PgPurchaseRepository) GetAllByUserName(userName string) ([]model.Purchase, error) {
	tx, err := r.db.Beginx()
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
		From("purchases").
		Where(squirrel.Eq{"username": userName}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	var purchases []model.Purchase

	err = tx.Select(&purchases, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	return purchases, nil
}
