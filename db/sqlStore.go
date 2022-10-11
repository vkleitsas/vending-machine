package db

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"log"
	"vending-machine/domain"
)

type SQLStore struct {
}

func NewSQLStore() *SQLStore {
	return &SQLStore{}
}

func (s *SQLStore) CreateUser(user domain.User, tx *sqlx.Tx) (*domain.User, error) {

	query, err := tx.Preparex(`INSERT INTO "user" (username, password, deposit, role) 
	VALUES ($1, $2, $3, $4) RETURNING ID`)
	if err != nil {
		log.Println("CreateUser get prepare query error: ", err)
		return nil, err
	}
	var insertedID uint
	err = query.Get(&insertedID, user.Username, user.Password, user.Deposit, user.Role)
	if err != nil {
		log.Println("CreateUser store error: ", err)
		return nil, err
	}

	user.ID = int(insertedID)

	return &user, nil
}

func (s *SQLStore) GetUserByUsername(user domain.User, tx *sqlx.Tx) (*domain.User, error) {
	query := `SELECT id, username, password, deposit, role FROM "user" WHERE username = $1`

	retrievedUsers := []domain.User{}

	err := tx.Select(&retrievedUsers, query, user.Username)

	if err != nil {
		log.Println("GetUserByUsername error: ", err)
		return nil, err
	}
	if len(retrievedUsers) > 1 {
		return nil, errors.New("GetUserByUsername unexpected error")
	}

	return &retrievedUsers[0], nil
}
func (s *SQLStore) GetUserΒyID(user domain.User, tx *sqlx.Tx) (*domain.User, error) {
	query := `SELECT id, username, password, deposit, role FROM "user" WHERE id = $1`

	retrievedUsers := []domain.User{}

	err := tx.Select(&retrievedUsers, query, user.ID)

	if err != nil {
		log.Println("GetUserByID error: ", err)
		return nil, err
	}
	if len(retrievedUsers) > 1 {
		return nil, errors.New("GetUserByID unexpected error")
	}

	return &retrievedUsers[0], nil
}

func (s *SQLStore) UpdateUser(user domain.User, tx *sqlx.Tx) (*domain.User, error) {
	query := `UPDATE "user" SET username=$1, password=$2, deposit=$3, role=$4 WHERE id=$5`

	_, err := tx.Exec(query, user.Username, user.Password, user.Deposit, user.Role, user.ID)
	if err != nil {
		log.Println("UpdateUser error: ", err)
		return nil, err
	}

	return &user, nil
}

func (s *SQLStore) GetAllProducts(tx *sqlx.Tx) ([]domain.Product, error) {
	query := `SELECT id, product_name, amount_available, cost, seller_id FROM "product"`

	products := []domain.Product{}

	err := tx.Select(&products, query)

	if err != nil {
		log.Println("GetAllProducts error: ", err)
		return nil, err
	}

	return products, nil
}

func (s *SQLStore) GetProductByID(product domain.Product, tx *sqlx.Tx) (*domain.Product, error) {
	query := `SELECT id, product_name, amount_available, cost, seller_id FROM "product" WHERE id = $1`

	retrievedProducts := []domain.Product{}

	err := tx.Select(&retrievedProducts, query, product.ID)

	if err != nil {
		log.Println("GetProductByID error: ", err)
		return nil, err
	}
	if len(retrievedProducts) > 1 {
		return nil, errors.New("GetProductByID unexpected error")
	}

	return &retrievedProducts[0], nil
}

func (s *SQLStore) CreateProduct(product domain.Product, tx *sqlx.Tx) (*domain.Product, error) {
	query, err := tx.Preparex(`INSERT INTO "user" (product_name, amount_available, cost, seller_id) 
	VALUES ($1, $2, $3, $4) RETURNING ID`)
	if err != nil {
		log.Println("CreateProduct get prepare query error: ", err)
		return nil, err
	}
	var insertedID uint
	err = query.Get(&insertedID, product.Name, product.AmountAvailable, product.Cost, product.SellerID)
	if err != nil {
		log.Println("CreateProduct store error: ", err)
		return nil, err
	}

	product.ID = int(insertedID)

	return &product, nil
}

func (s *SQLStore) UpdateProduct(product domain.Product, tx *sqlx.Tx) (*domain.Product, error) {
	query := `UPDATE "user" SET product_name=$1, amount_available=$2, cost=$3, seller_id=$4 WHERE id=$5`

	_, err := tx.Exec(query, product.Name, product.AmountAvailable, product.Cost, product.SellerID, product.ID)
	if err != nil {
		log.Println("UpdateProduct error: ", err)
		return nil, err
	}

	return &product, nil
}

func (s *SQLStore) DeleteProduct(product domain.Product, tx *sqlx.Tx) error {
	query := `DELETE from "user" where id=$1`

	_, err := tx.Exec(query, product.ID)
	if err != nil {
		log.Println("DeleteProduct error: ", err)
		return err
	}

	return nil
}
