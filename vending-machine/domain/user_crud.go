package domain

import "github.com/jmoiron/sqlx"

type UserCrud interface {
	CreateUser(user User, tx *sqlx.Tx) (*User, error)
	GetUserΒyID(user User, tx *sqlx.Tx) (*User, error)
	GetUserByUsername(user User, tx *sqlx.Tx) (*User, error)
	UpdateUser(user User, tx *sqlx.Tx) (*User, error)
}
