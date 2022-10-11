package app

import (
	"errors"
	"vending-machine/db"
	"vending-machine/domain"
)

type TransactionsCrudService struct {
	Repository domain.TransactionsCrud
	DBClient   db.InitDBInterface
}

func NewTransactionsCrudService(c db.InitDBInterface, sqlStore domain.TransactionsCrud) *TransactionsCrudService {
	return &TransactionsCrudService{
		Repository: sqlStore,
		DBClient:   c,
	}
}

func (s *TransactionsCrudService) Purchase(p domain.PurchaseRequestDTO) (*domain.PurchaseResponseDTO, error) {
	if p.Amount <= 0 || p.ProductID < 1 || p.BuyerID < 1 {
		return nil, errors.New("Invalid fields")
	}

	productToFind := domain.Product{
		ID: p.ProductID,
	}
	tx, err := s.DBClient.GetTransaction()
	if err != nil {
		return nil, err
	}
	productFromDB, err := s.Repository.GetProductByID(productToFind, tx)
	if err != nil {
		s.DBClient.Rollback(tx)
		return nil, err
	}
	if productFromDB.Amount < p.Amount {
		s.DBClient.Rollback(tx)
		return nil, errors.New("Not enough products available")
	}

	userToFind := domain.User{
		ID: p.BuyerID,
	}
	userFromDB, err := s.Repository.GetUserΒyID(userToFind, tx)
	if err != nil {
		s.DBClient.Rollback(tx)
		return nil, err
	}

	if userFromDB.Deposit < p.Amount*productFromDB.Cost {
		s.DBClient.Rollback(tx)
		return nil, errors.New("Not enough funds")
	}

	productFromDB.Amount = productFromDB.Amount - p.Amount
	updatedProduct, err := s.Repository.UpdateProduct(*productFromDB, tx)
	if err != nil {
		s.DBClient.Rollback(tx)
		return nil, err
	}

	userFromDB.Deposit = userFromDB.Deposit - (p.Amount * productFromDB.Cost)
	updatedUser, err := s.Repository.UpdateUser(*userFromDB, tx)
	if err != nil {
		s.DBClient.Rollback(tx)
		return nil, err
	}

	response := &domain.PurchaseResponseDTO{
		ProductID:   updatedProduct.ID,
		ProductName: updatedProduct.Name,
		BuyerID:     updatedUser.ID,
		Change:      coinsBreakdown(updatedUser.Deposit),
	}
	s.DBClient.Commit(tx)
	return response, nil
}
