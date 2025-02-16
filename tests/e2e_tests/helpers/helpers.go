package helpers

import (
	"avito-shop/cmd/app"
	"avito-shop/cmd/config"
	"avito-shop/cmd/initDB"
	"avito-shop/internal/model"
	"avito-shop/seed"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"log"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"
)

type TestConfig struct {
	Server *httptest.Server
}

func SetupTestEnvironment(t *testing.T) *TestConfig {
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
		Server: server,
	}
}

func PerformPurchase(t *testing.T, baseURL, merchName string, token string) *http.Response {
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

func PerformTransfer(t *testing.T, baseURL, recipient string, amount int, token string) *http.Response {
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

func GetInfo(t *testing.T, baseURL, token string) *model.UserInfo {
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

func AuthenticateAndGetToken(t *testing.T, env *TestConfig, username, password string) string {
	user := model.User{Username: username, Password: password}
	jsonBody, err := json.Marshal(user)
	require.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/auth", env.Server.URL), bytes.NewBuffer(jsonBody))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	var authResponse struct {
		Token string `json:"token"`
	}
	err = json.NewDecoder(resp.Body).Decode(&authResponse)
	require.NoError(t, err)

	return authResponse.Token
}
