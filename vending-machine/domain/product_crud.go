package domain

import "github.com/jmoiron/sqlx"

type ProductCrud interface {
	GetAllProducts(tx *sqlx.Tx) ([]Product, error)
	GetProductByID(product Product, tx *sqlx.Tx) (*Product, error)
	CreateProduct(product Product, tx *sqlx.Tx) (*Product, error)
	UpdateProduct(product Product, tx *sqlx.Tx) (*Product, error)
	DeleteProduct(product Product, tx *sqlx.Tx) error
}
