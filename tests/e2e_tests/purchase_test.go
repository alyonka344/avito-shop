package e2e_tests

import (
	"avito-shop/cmd/app"
	"avito-shop/cmd/config"
	"avito-shop/cmd/initDB"
	"avito-shop/internal/model"
	"avito-shop/seed"
	"bytes"
	"encoding/json"
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"
)

// TestConfig содержит тестовые данные
type TestConfig struct {
	server *httptest.Server
}

func setupTestEnvironment(t *testing.T) *TestConfig {
	cfg := config.New()
	db, err := initDB.InitDatabase(cfg)
	require.NoError(t, err)
	_, b, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(b), "../..")
	migrationsPath := fmt.Sprintf("file://%s/migrations", projectRoot)
	if err := initDB.RunMigrations(db, cfg.Database.Name, migrationsPath); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	if err := seed.ApplySeeds(db); err != nil {
		log.Fatalf("Failed to apply seeds: %v", err)
	}

	application := app.New(db, cfg.SecretKey)
	server := httptest.NewServer(application.Router)

	return &TestConfig{
		server: server,
	}
}

func TestPurchaseMerch(t *testing.T) {
	cfg := setupTestEnvironment(t)
	defer cfg.server.Close()

	// Подготовка тестовых данных
	testCases := []struct {
		name            string
		userName        string
		password        string
		merchName       string
		merchPrice      int
		expectedStatus  int
		expectedBalance int
	}{
		{
			name:            "Успешная покупка мерча",
			userName:        "testUser1",
			password:        "05052004",
			merchName:       "pink-hoody",
			expectedStatus:  http.StatusOK,
			expectedBalance: 500,
		},
		{
			name:            "Недостаточно средств",
			userName:        "testUser2",
			password:        "password",
			merchName:       "test",
			expectedStatus:  http.StatusBadRequest,
			expectedBalance: 1000,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Создаем тестового пользователя
			token := authenticateAndGetToken(t, cfg, tc.userName, tc.password)

			// Выполняем запрос на покупку
			resp := performPurchase(t, cfg.server.URL, tc.userName, tc.merchName, token)
			assert.Equal(t, tc.expectedStatus, resp.StatusCode)

			// Проверяем баланс после покупки
			info := getInfo(t, cfg.server.URL, token)
			assert.Equal(t, tc.expectedBalance, info.Coins)
		})
	}
}

//func TestTransferCoins(t *testing.T) {
//	cfg := setupTestEnvironment(t)
//	defer cfg.server.Close()
//
//	testCases := []struct {
//		name                     string
//		sender                   string
//		recipient                string
//		senderPassword           string
//		recipientPassword        string
//		amount                   int
//		expectedStatus           int
//		expectedSenderBalance    int
//		expectedRecipientBalance int
//	}{
//		{
//			name:                     "Успешный перевод монет",
//			sender:                   "sender1",
//			recipient:                "recipient1",
//			senderPassword:           "password",
//			recipientPassword:        "password",
//			amount:                   500,
//			expectedStatus:           http.StatusOK,
//			expectedSenderBalance:    500,
//			expectedRecipientBalance: 500,
//		},
//		{
//			name:                     "Недостаточно средств для перевода",
//			sender:                   "sender2",
//			recipient:                "recipient2",
//			senderPassword:           "password",
//			recipientPassword:        "password",
//			amount:                   500,
//			expectedStatus:           http.StatusBadRequest,
//			expectedSenderBalance:    100,
//			expectedRecipientBalance: 0,
//		},
//	}
//
//	for _, tc := range testCases {
//		t.Run(tc.name, func(t *testing.T) {
//			// Создаем отправителя и получателя
//			createTestUser(t, cfg.server.URL, tc.sender, tc.senderPassword)
//			createTestUser(t, cfg.server.URL, tc.recipient, tc.recipientPassword)
//
//			// Выполняем перевод
//			resp := performTransfer(t, cfg.server.URL, tc.sender, tc.recipient, tc.amount)
//			assert.Equal(t, tc.expectedStatus, resp.StatusCode)
//
//			// Проверяем балансы после перевода
//			senderUser := getInfo(t, cfg.server.URL, token)
//			recipientUser := getInfo(t, cfg.server.URL, token)
//
//			assert.Equal(t, tc.expectedSenderBalance, senderUser.Balance)
//			assert.Equal(t, tc.expectedRecipientBalance, recipientUser.Balance)
//		})
//	}
//}

// Вспомогательные функции

//func createTestUser(t *testing.T, baseURL, username string, password string) {
//	user := model.User{
//		Username: username,
//		Password: password,
//	}
//
//	jsonBody, err := json.Marshal(user)
//	require.NoError(t, err)
//
//	resp, err := http.Post(baseURL+"/api/auth", "application/json", bytes.NewBuffer(jsonBody))
//	require.NoError(t, err)
//	require.Equal(t, http.StatusOK, resp.StatusCode)
//}

func performPurchase(t *testing.T, baseURL, username, merchName string, token string) *http.Response {
	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("%s/api/buy/%s", baseURL, merchName),
		nil)
	require.NoError(t, err)

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	require.NoError(t, err)

	return resp
}

func performTransfer(t *testing.T, baseURL, recipient string, amount int, token string) *http.Response {
	transferReq := struct {
		Recipient string `json:"recipient"`
		Amount    int    `json:"amount"`
	}{
		Recipient: recipient,
		Amount:    amount,
	}

	jsonBody, err := json.Marshal(transferReq)
	require.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost,
		baseURL+"/api/sendCoin",
		bytes.NewBuffer(jsonBody))
	require.NoError(t, err)

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	require.NoError(t, err)

	return resp
}

func getInfo(t *testing.T, baseURL, token string) *model.UserInfo {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/info", baseURL), nil)
	require.NoError(t, err)
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	require.NoError(t, err)

	var userInfo model.UserInfo
	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	require.NoError(t, err)

	return &userInfo
}

//func getTestToken(username string) string {
//	// В реальном приложении здесь должна быть логика получения токена
//	return "test_token_" + username
//}

func authenticateAndGetToken(t *testing.T, env *TestConfig, username, password string) string {
	// Create user credentials
	user := model.User{Username: username, Password: password}
	jsonBody, err := json.Marshal(user)
	require.NoError(t, err)

	// Send authentication request
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/auth", env.server.URL), bytes.NewBuffer(jsonBody))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	// Parse response to get token
	var authResponse struct {
		Token string `json:"token"`
	}
	err = json.NewDecoder(resp.Body).Decode(&authResponse)
	require.NoError(t, err)

	return authResponse.Token
}
