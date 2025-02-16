package unit_tests

import (
	"avito-shop/internal/mocks"
	"avito-shop/internal/model"
	"avito-shop/internal/usecase/usecase_impl"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetInfoReturnsCompleteUserInfoForExistingUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockUserRepository(ctrl)
	mockTxRepo := mocks.NewMockTransactionRepository(ctrl)
	mockPurchaseRepo := mocks.NewMockPurchaseRepository(ctrl)

	userName := "testUser"
	expectedUser := &model.User{Username: userName, Balance: 100}
	expectedSentTx := []model.Transaction{{FromUser: userName, ToUser: "user2", Amount: 50}}
	expectedReceivedTx := []model.Transaction{{FromUser: "user3", ToUser: userName, Amount: 30}}
	expectedPurchases := []model.Purchase{{UserName: userName, MerchName: "item1"}}

	mockUserRepo.EXPECT().GetByName(userName).Return(expectedUser, nil)
	mockTxRepo.EXPECT().GetAllSentByUserName(userName).Return(expectedSentTx, nil)
	mockTxRepo.EXPECT().GetAllReceivedByUserName(userName).Return(expectedReceivedTx, nil)
	mockPurchaseRepo.EXPECT().GetAllByUserName(userName).Return(expectedPurchases, nil)

	useCase := usecase_impl.NewUserUsecase(mockUserRepo, mockTxRepo, mockPurchaseRepo)

	result, err := useCase.GetInfo(userName)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser.Balance, result.Coins)
	assert.Len(t, result.Inventory, 1)
	assert.Equal(t, "item1", result.Inventory[0].Type)
	assert.Equal(t, 1, result.Inventory[0].Quantity)
	assert.Len(t, result.CoinHistory.Sent, 1)
	assert.Len(t, result.CoinHistory.Received, 1)
}

func TestGetInfoReturnsErrorForNonExistentUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockUserRepository(ctrl)
	mockTxRepo := mocks.NewMockTransactionRepository(ctrl)
	mockPurchaseRepo := mocks.NewMockPurchaseRepository(ctrl)

	userName := "nonExistentUser"
	expectedError := errors.New("user not found")

	mockUserRepo.EXPECT().GetByName(userName).Return(nil, expectedError)

	useCase := usecase_impl.NewUserUsecase(mockUserRepo, mockTxRepo, mockPurchaseRepo)

	result, err := useCase.GetInfo(userName)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to get user")
}
