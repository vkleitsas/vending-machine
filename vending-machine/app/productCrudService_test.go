package app

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	"testing"
	"vending-machine/db"
	"vending-machine/domain"
)

func Test_ProductDelete_OK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSqlxClient := db.NewMockInitDBInterface(ctrl)
	mockRepository := domain.NewMockProductCrud(ctrl)
	productService := NewProductCrudService(mockSqlxClient, mockRepository)

	product := domain.Product{
		ID:     1,
		Amount: 1,
		Cost:   10,
		Name:   "MockProduct",
		Seller: 1,
	}

	requestUser := domain.User{
		ID:       1,
		Username: "user1",
		Password: "",
		Deposit:  20,
		Role:     "buyer",
	}
	mockTX := sqlx.Tx{
		Tx:     nil,
		Mapper: nil,
	}
	mockSqlxClient.EXPECT().GetTransaction().Return(&mockTX, nil).AnyTimes()
	mockSqlxClient.EXPECT().Rollback(&mockTX).Return(nil).AnyTimes()
	mockSqlxClient.EXPECT().Commit(&mockTX).Return(nil).AnyTimes()
	mockRepository.EXPECT().GetProductByID(product, &mockTX).Return(&product, nil)
	mockRepository.EXPECT().DeleteProduct(product, &mockTX).Return(nil)
	err := productService.DeleteProduct(product, requestUser)
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}
}

func Test_ProductDelete_ProductDoesNotBelongToUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSqlxClient := db.NewMockInitDBInterface(ctrl)
	mockRepository := domain.NewMockProductCrud(ctrl)
	productService := NewProductCrudService(mockSqlxClient, mockRepository)

	product := domain.Product{
		ID:     1,
		Amount: 1,
		Cost:   10,
		Name:   "MockProduct",
		Seller: 2,
	}

	requestUser := domain.User{
		ID:       1,
		Username: "user1",
		Password: "",
		Deposit:  20,
		Role:     "buyer",
	}
	mockTX := sqlx.Tx{
		Tx:     nil,
		Mapper: nil,
	}
	mockSqlxClient.EXPECT().GetTransaction().Return(&mockTX, nil).AnyTimes()
	mockSqlxClient.EXPECT().Rollback(&mockTX).Return(nil).AnyTimes()
	mockSqlxClient.EXPECT().Commit(&mockTX).Return(nil).AnyTimes()
	mockRepository.EXPECT().GetProductByID(product, &mockTX).Return(&product, nil)
	err := productService.DeleteProduct(product, requestUser)
	if err == nil {
		t.Errorf("Expected error")
	}
	expectedError := errors.New("Invalid operation - product belogs to another user")

	if err.Error() != expectedError.Error() {
		t.Errorf("Unexpected error %s", err)
	}
}
