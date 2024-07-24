package products

import (
	"context"

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

func (r repository) CreateProduct(ctx context.Context, model Product) (err error) {
	query := `
		INSERT INTO products (
			sku, name, price, stock, created_at, updated_at
		) Values (
			:sku, :name, :price, :stock, :created_at, :updated_at
		)
	`
	// stmt ini akan membuka connection pool baru, jadi jangan lupa di close lagi
	stmt, err := r.db.PrepareNamedContext(ctx, query)

	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, model)
	if err != nil {
		return
	}
	return
}
