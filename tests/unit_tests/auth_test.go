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

func TestAuthUsecase_ValidateOrCreateUserReturnsExistingUser(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockUserRepository(ctrl)
	mockAuthService := mocks.NewMockAuthService(ctrl)
	mockHashService := mocks.NewMockHashService(ctrl)

	existingUser := &model.User{
		Username: "testuser",
		Password: "hashedpass",
		Balance:  0,
	}

	// act
	mockUserRepo.EXPECT().ExistsByName("testuser").Return(true, nil)
	mockUserRepo.EXPECT().GetByName("testuser").Return(existingUser, nil)

	usecase := usecase_impl.NewAuthUsecase(mockUserRepo, mockAuthService, mockHashService)

	user, err := usecase.ValidateOrCreateUser("testuser", "password")

	// assert
	assert.NoError(t, err)
	assert.Equal(t, existingUser.Username, user.Username)
	assert.Equal(t, existingUser.Password, user.Password)
}

func TestAuthUsecase_ValidateOrCreateUserWithEmptyUsername(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockUserRepository(ctrl)
	mockAuthService := mocks.NewMockAuthService(ctrl)
	mockHashService := mocks.NewMockHashService(ctrl)

	// act
	mockUserRepo.EXPECT().ExistsByName("").Return(false, errors.New("username cannot be empty"))

	usecase := usecase_impl.NewAuthUsecase(mockUserRepo, mockAuthService, mockHashService)

	user, err := usecase.ValidateOrCreateUser("", "password")

	// assert
	assert.Error(t, err)
	assert.Equal(t, model.User{}, user)
}

func TestAuthUsecase_CreateNewUserIfNotExists(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockUserRepository(ctrl)
	mockAuthService := mocks.NewMockAuthService(ctrl)
	mockHashService := mocks.NewMockHashService(ctrl)

	newUser := &model.User{
		Username: "newuser",
		Password: "hashedpassword",
	}

	mockUserRepo.EXPECT().ExistsByName("newuser").Return(false, nil)
	mockHashService.EXPECT().HashPassword("password").Return("hashedpassword", nil)
	mockUserRepo.EXPECT().Create(newUser).Return(nil)

	usecase := usecase_impl.NewAuthUsecase(mockUserRepo, mockAuthService, mockHashService)

	// act
	user, err := usecase.ValidateOrCreateUser("newuser", "password")

	// assert
	assert.NoError(t, err)
	assert.Equal(t, newUser.Username, user.Username)
	assert.Equal(t, newUser.Password, user.Password)
}
