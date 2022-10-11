package app

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	"testing"
	"vending-machine/db"
	"vending-machine/domain"
)

func Test_UserDeposit_OK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSqlxClient := db.NewMockInitDBInterface(ctrl)
	mockAuth := domain.MockJwtClaimsInterface{}
	mockRepository := domain.NewMockUserCrud(ctrl)
	userService := NewUserCrudService(mockSqlxClient, mockRepository, &mockAuth)
	user := domain.User{
		ID:       1,
		Username: "user1",
		Password: "",
		Deposit:  10,
		Role:     "buyer",
	}
	updateUser := domain.User{
		ID:       1,
		Username: "user1",
		Password: "",
		Deposit:  20,
		Role:     "buyer",
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
	mockRepository.EXPECT().GetUserΒyID(user, &mockTX).Return(&user, nil)
	mockRepository.EXPECT().UpdateUser(updateUser, &mockTX).Return(&user, nil)
	_, err := userService.Deposit(user, requestUser)
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}
}

func Test_UserDeposit_DepositAmountInvalid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSqlxClient := db.NewMockInitDBInterface(ctrl)
	mockAuth := domain.MockJwtClaimsInterface{}
	mockRepository := domain.NewMockUserCrud(ctrl)
	userService := NewUserCrudService(mockSqlxClient, mockRepository, &mockAuth)
	user := domain.User{
		ID:       1,
		Username: "user1",
		Password: "",
		Deposit:  12,
		Role:     "buyer",
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
	_, err := userService.Deposit(user, requestUser)
	if err == nil {
		t.Errorf("Expected error")
	}
	expectedError := errors.New("Deposit value must be one of {5,10,20,50,100}")
	if err.Error() != expectedError.Error() {
		t.Errorf("Unexpected error %s", err)
	}
}
