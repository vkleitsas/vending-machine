package app

import (
	"errors"
	"vending-machine/db"
	"vending-machine/domain"
)

type ProductCrudService struct {
	Repository domain.ProductCrud
	DBClient   db.InitDBInterface
}

func NewProductCrudService(c db.InitDBInterface, sqlStore domain.ProductCrud) *ProductCrudService {
	return &ProductCrudService{
		Repository: sqlStore,
		DBClient:   c,
	}
}

func (s *ProductCrudService) GetAllProducts() ([]domain.Product, error) {
	tx, err := s.DBClient.GetTransaction()
	if err != nil {
		return nil, err
	}
	products, err := s.Repository.GetAllProducts(tx)
	if err != nil {
		s.DBClient.Rollback(tx)
		return nil, err
	}
	s.DBClient.Commit(tx)
	return products, nil
}

func (s *ProductCrudService) CreateProduct(p domain.Product, requestUser domain.User) (*domain.Product, error) {
	if p.Name == "" || p.Cost <= 0 || p.Seller <= 1 && p.Amount < 0 {
		return nil, errors.New("Invalid product fields")
	}

	if p.Seller != requestUser.ID {
		return nil, errors.New("Product's seller cannot be another user")
	}
	if p.Cost%5 != 0 {
		return nil, errors.New("Product cost value must be a multiple of 5}")
	}

	tx, err := s.DBClient.GetTransaction()
	if err != nil {
		return nil, err
	}
	createdProduct, err := s.Repository.CreateProduct(p, tx)
	if err != nil {
		return nil, err
		s.DBClient.Rollback(tx)
	}
	s.DBClient.Commit(tx)
	return createdProduct, nil
}

func (s *ProductCrudService) UpdateProduct(p domain.Product, requestUser domain.User) (*domain.Product, error) {
	if p.Seller <= 1 || p.Cost <= 0 || p.Amount < 0 || p.Name == "" {
		return nil, errors.New("Invalid product fields")
	}
	if !containsInt(p.Cost, costValues) {
		return nil, errors.New("Product cost value must be one of {5,10,20,50,100}")
	}
	tx, err := s.DBClient.GetTransaction()
	if err != nil {
		return nil, err
	}
	retrievedProduct, err := s.Repository.GetProductByID(p, tx)
	if err != nil {
		s.DBClient.Rollback(tx)
		return nil, err
	}
	if retrievedProduct.Seller != requestUser.ID {
		s.DBClient.Rollback(tx)
		return nil, errors.New("Invalid operation - product belogs to another user")
	}

	updatedProduct, err := s.Repository.UpdateProduct(p, tx)
	if err != nil {
		s.DBClient.Rollback(tx)
		return nil, err
	}
	s.DBClient.Commit(tx)
	return updatedProduct, err
}

func (s *ProductCrudService) DeleteProduct(p domain.Product, requestUser domain.User) error {
	if p.ID < 1 {
		return errors.New("Invalid product ID")
	}
	tx, err := s.DBClient.GetTransaction()
	if err != nil {
		return err
	}
	retrievedProduct, err := s.Repository.GetProductByID(p, tx)
	if err != nil {
		s.DBClient.Rollback(tx)
		return err
	}
	if retrievedProduct.Seller != requestUser.ID {
		s.DBClient.Rollback(tx)
		return errors.New("Invalid operation - product belogs to another user")
	}

	err = s.Repository.DeleteProduct(p, tx)
	if err != nil {
		s.DBClient.Rollback(tx)
		return err
	}
	s.DBClient.Commit(tx)
	return nil

}
