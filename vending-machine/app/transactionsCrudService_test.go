package app

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	"testing"
	"vending-machine/db"
	"vending-machine/domain"
)

func Test_Transactions_Buy_OK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSqlxClient := db.NewMockInitDBInterface(ctrl)

	mockRepository := domain.NewMockTransactionsCrud(ctrl)
	transactionsService := NewTransactionsCrudService(mockSqlxClient, mockRepository)

	request := domain.PurchaseRequestDTO{
		ProductID: 1,
		Amount:    1,
		BuyerID:   1,
	}
	product := domain.Product{
		ID:     1,
		Amount: 1,
		Cost:   5,
		Name:   "MockProduct",
		Seller: 1,
	}
	user := domain.User{
		ID:       1,
		Username: "MockUser1",
		Password: "",
		Deposit:  10,
		Role:     "buyer",
	}

	mockTX := sqlx.Tx{
		Tx:     nil,
		Mapper: nil,
	}

	mockSqlxClient.EXPECT().GetTransaction().Return(&mockTX, nil).AnyTimes()
	mockSqlxClient.EXPECT().Rollback(&mockTX).Return(nil).AnyTimes()
	mockSqlxClient.EXPECT().Commit(&mockTX).Return(nil).AnyTimes()
	mockRepository.EXPECT().GetProductByID(domain.Product{ID: 1}, &mockTX).Return(&product, nil).AnyTimes()
	mockRepository.EXPECT().GetUserΒyID(gomock.Any(), &mockTX).Return(&user, nil)
	mockRepository.EXPECT().UpdateProduct(gomock.Any(), &mockTX).Return(&product, nil)
	mockRepository.EXPECT().UpdateUser(gomock.Any(), &mockTX).Return(&user, nil)
	_, err := transactionsService.Purchase(request)
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}

}

func Test_Transactions_Buy_NotEnoughProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSqlxClient := db.NewMockInitDBInterface(ctrl)

	mockRepository := domain.NewMockTransactionsCrud(ctrl)
	transactionsService := NewTransactionsCrudService(mockSqlxClient, mockRepository)

	request := domain.PurchaseRequestDTO{
		ProductID: 1,
		Amount:    10,
		BuyerID:   1,
	}
	product := domain.Product{
		ID:     1,
		Amount: 1,
		Cost:   5,
		Name:   "MockProduct",
		Seller: 1,
	}

	mockTX := sqlx.Tx{
		Tx:     nil,
		Mapper: nil,
	}

	mockSqlxClient.EXPECT().GetTransaction().Return(&mockTX, nil).AnyTimes()
	mockSqlxClient.EXPECT().Rollback(&mockTX).Return(nil).AnyTimes()
	mockSqlxClient.EXPECT().Commit(&mockTX).Return(nil).AnyTimes()
	mockRepository.EXPECT().GetProductByID(domain.Product{ID: 1}, &mockTX).Return(&product, nil).AnyTimes()
	_, err := transactionsService.Purchase(request)
	if err == nil {
		t.Errorf("Expected error")
	}
	expectedError := errors.New("Not enough products available")
	if err.Error() != expectedError.Error() {
		t.Errorf("Unexpected error %s", err)
	}
}
