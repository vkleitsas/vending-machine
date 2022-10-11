package domain

import "github.com/jmoiron/sqlx"

type TransactionsCrud interface {
	GetProductByID(p Product, tx *sqlx.Tx) (*Product, error)
	GetUserΒyID(u User, tx *sqlx.Tx) (*User, error)
	UpdateProduct(p Product, tx *sqlx.Tx) (*Product, error)
	UpdateUser(u User, tx *sqlx.Tx) (*User, error)
}
