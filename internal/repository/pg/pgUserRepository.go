package pg

import (
	"avito-shop/internal/model"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/gofrs/uuid/v5"
	"github.com/jmoiron/sqlx"
)

const initialBalance = 1000

type PgUserRepository struct {
	db *sqlx.DB
}

func NewPgUserRepository(db *sqlx.DB) *PgUserRepository {
	return &PgUserRepository{db: db}
}

func (r *PgUserRepository) Create(user *model.User) error {
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
		Insert("users").
		Columns("username", "password", "balance").
		Values(user.Username, user.Password, initialBalance).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	err = tx.QueryRowx(query, args...).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}

func (r *PgUserRepository) GetById(userID uuid.UUID) (*model.User, error) {
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
		From("users").
		Where(squirrel.Eq{"id": userID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	var users []model.User

	err = tx.Select(&users, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	if len(users) == 0 {
		return nil, fmt.Errorf("this user is not exist: %w", err)
	}

	return &users[0], nil
}

func (r *PgUserRepository) GetByName(userName string) (*model.User, error) {
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
		From("users").
		Where(squirrel.Eq{"username": userName}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	var users []model.User

	err = tx.Select(&users, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	if len(users) == 0 {
		return nil, fmt.Errorf("this user is not exist: %w", err)
	}

	return &users[0], nil
}

func (r *PgUserRepository) Transfer(senderID uuid.UUID, recipientID uuid.UUID, amount int) error {
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

	deductQuery, args, err := squirrel.
		Update("users").
		Set("balance", squirrel.Expr("balance - ?", amount)).
		Where("id = ? AND balance >= ?", senderID, amount).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	res, err := tx.Exec(deductQuery, args...)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("insufficient funds or user not found")
	}

	addQuery, args, err := squirrel.
		Update("users").
		Set("balance", squirrel.Expr("balance + ?", amount)).
		Where(squirrel.Eq{"id": recipientID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	_, err = tx.Exec(addQuery, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *PgUserRepository) UpdateBalance(userID uuid.UUID, amount int) error {
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
		Update("users").
		Set("balance", squirrel.Expr("balance + ?", amount)).
		Where("id = ? AND balance >= ?", userID, -amount).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *PgUserRepository) ExistsByName(userName string) (bool, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return false, err
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

	var exists bool
	query := squirrel.Select("count(1) > 0").
		From("users").
		Where(squirrel.Eq{"username": userName})
	sql, args, _ := query.ToSql()

	err = r.db.QueryRow(sql, args...).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
