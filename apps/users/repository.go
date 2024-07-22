package users

import (
	"context"
	"database/sql"
	"online-shop-ddd/infra/response"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func newRepository(db *sqlx.DB) repository {
	return repository{
		db: db,
	}
}

func (r repository) CreateUser(ctx context.Context, model UserEntity) (err error) {
	query := `insert into users (
		email, password, role, public_id, created_at, updated_at
	) values (
	 	:email, :password, :role, :public_id, :created_at, :updated_at
	)`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, model)
	return
}

func (r repository) GetUserByEmail(ctx context.Context, email string) (model UserEntity, err error) {
	query := `
			SELECT 
				id, public_id, email, password, role, created_at, updated_at
			FROM users 
			where email = $1
	`

	err = r.db.GetContext(ctx, &model, query, email)
	if err != nil {
		if err == sql.ErrNoRows {
			err = response.ErrNotFound
			return
		}
		return
	}
	return
}
