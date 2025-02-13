package pg

import (
	"avito-shop/internal/model"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/gofrs/uuid/v5"
	"github.com/jmoiron/sqlx"
)

type PgTransactionRepository struct {
	db *sqlx.DB
}

func NewPgTransactionRepository(db *sqlx.DB) *PgTransactionRepository {
	return &PgTransactionRepository{db: db}
}

func (r PgTransactionRepository) Create(transaction *model.Transaction) error {
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
		Insert("transactions").
		Columns("from_user_id", "to_user_id", "amount", "transaction_status").
		Values(transaction.FromUserID, transaction.ToUserID, transaction.Amount, transaction.TransactionStatus).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	err = tx.QueryRowx(query, args...).Scan(&transaction.ID)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}

func (r PgTransactionRepository) GetAllSentByUserId(userID uuid.UUID) ([]model.Transaction, error) {
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
		From("transactions").
		Where(squirrel.Eq{"from_user_id": userID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	var transactions []model.Transaction

	err = tx.Select(&transactions, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	return transactions, nil
}

func (r PgTransactionRepository) GetAllReceivedByUserId(userID uuid.UUID) ([]model.Transaction, error) {
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
		From("transactions").
		Where(squirrel.Eq{"to_user_id": userID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	var transactions []model.Transaction

	err = tx.Select(&transactions, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	return transactions, nil
}
