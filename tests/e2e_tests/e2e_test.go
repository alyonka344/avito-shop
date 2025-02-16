package e2e_tests

import (
	"avito-shop/tests/e2e_tests/helpers"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestPurchaseMerch(t *testing.T) {
	cfg := helpers.SetupTestEnvironment(t)
	defer cfg.Server.Close()

	testCases := []struct {
		name            string
		userName        string
		password        string
		merchName       string
		merchPrice      int
		expectedStatus  int
		expectedBalance int
		numOfPurchases  int
	}{
		{
			name:            "Успешная покупка мерча",
			userName:        "testUser1",
			password:        "05052004",
			merchName:       "pink-hoody",
			expectedStatus:  http.StatusOK,
			expectedBalance: 500,
			numOfPurchases:  1,
		},
		{
			name:            "Недостаточно средств",
			userName:        "testUser2",
			password:        "password",
			merchName:       "hoody",
			expectedStatus:  http.StatusBadRequest,
			expectedBalance: 100,
			numOfPurchases:  4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			token := helpers.AuthenticateAndGetToken(t, cfg, tc.userName, tc.password)

			var resp *http.Response
			for i := 0; i < tc.numOfPurchases; i++ {
				resp = helpers.PerformPurchase(t, cfg.Server.URL, tc.merchName, token)
			}
			assert.Equal(t, tc.expectedStatus, resp.StatusCode)

			info := helpers.GetInfo(t, cfg.Server.URL, token)
			assert.Equal(t, tc.expectedBalance, info.Coins)
		})
	}
}

func TestTransferCoins(t *testing.T) {
	cfg := helpers.SetupTestEnvironment(t)
	defer cfg.Server.Close()

	testCases := []struct {
		name                     string
		sender                   string
		recipient                string
		senderPassword           string
		recipientPassword        string
		amount                   int
		expectedStatus           int
		expectedSenderBalance    int
		expectedRecipientBalance int
	}{
		{
			name:                     "Успешный перевод монет",
			sender:                   "sender1",
			recipient:                "recipient1",
			senderPassword:           "password",
			recipientPassword:        "password",
			amount:                   500,
			expectedStatus:           http.StatusOK,
			expectedSenderBalance:    500,
			expectedRecipientBalance: 1500,
		},
		{
			name:                     "Недостаточно средств для перевода",
			sender:                   "sender2",
			recipient:                "recipient2",
			senderPassword:           "password",
			recipientPassword:        "password",
			amount:                   1500,
			expectedStatus:           http.StatusBadRequest,
			expectedSenderBalance:    1000,
			expectedRecipientBalance: 1000,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			senderToken := helpers.AuthenticateAndGetToken(t, cfg, tc.sender, tc.senderPassword)
			recipientToken := helpers.AuthenticateAndGetToken(t, cfg, tc.recipient, tc.recipientPassword)

			resp := helpers.PerformTransfer(t, cfg.Server.URL, tc.recipient, tc.amount, senderToken)
			assert.Equal(t, tc.expectedStatus, resp.StatusCode)

			senderUser := helpers.GetInfo(t, cfg.Server.URL, senderToken)
			recipientUser := helpers.GetInfo(t, cfg.Server.URL, recipientToken)

			assert.Equal(t, tc.expectedSenderBalance, senderUser.Coins)
			assert.Equal(t, tc.expectedRecipientBalance, recipientUser.Coins)
		})
	}
}
