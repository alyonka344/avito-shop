package pg

import (
	"avito-shop/internal/model"
	"github.com/gofrs/uuid/v5"

	//"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	//"github.com/gofrs/uuid/v5"
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

func (r *PgUserRepository) Update(user *model.User) error {
	return nil
}
func (r *PgUserRepository) GetById(userID uuid.UUID) (model.User, error) {
	return model.User{}, nil
}
func (r *PgUserRepository) GetByName(userName string) (model.User, error) {
	return model.User{}, nil
}
