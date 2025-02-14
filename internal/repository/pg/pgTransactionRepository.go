package pg

import (
	"avito-shop/internal/model"
	"fmt"
	"github.com/Masterminds/squirrel"
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
		Columns("from_user", "to_user", "amount", "transaction_status").
		Values(transaction.FromUser, transaction.ToUser, transaction.Amount, transaction.TransactionStatus).
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

func (r PgTransactionRepository) GetAllSentByUserName(userName string) ([]model.Transaction, error) {
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
		Where(squirrel.Eq{"from_user": userName}).
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

func (r PgTransactionRepository) GetAllReceivedByUserName(userName string) ([]model.Transaction, error) {
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
		Where(squirrel.Eq{"to_user": userName}).
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
