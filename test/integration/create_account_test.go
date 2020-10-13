// +build integration
// +build !unit

package integration

import (
	"os"
	"testing"

	"github.com/jonatak/go-form3/pkg/account"
	"github.com/jonatak/go-form3/pkg/form3"
)

func assertNotEmpty(t *testing.T, v, msg string) {
	if v == "" {
		t.Error(msg)
	}
}

func TestEnvSetUP(t *testing.T) {

	form3Endpoint := os.Getenv("FORM3_ENDPOINT")
	form3OrdID := os.Getenv("FORM3_ORG_ID")

	assertNotEmpty(t, form3Endpoint, "FORM3_ENDPOINT env var should be present.")
	assertNotEmpty(t, form3OrdID, "FORM3_ORG_ID env var should be present.")
}

func TestCreateAccount(t *testing.T) {

	form3Endpoint := os.Getenv("FORM3_ENDPOINT")
	form3OrdID := os.Getenv("FORM3_ORG_ID")

	client := form3.New(form3OrdID, form3Endpoint)
	_, err := client.Account.Create("ad27e265-9605-4b4b-a0e5-3003ea9cc4dd", &account.Account{
		Country:      "GB",
		BaseCurrency: "GBP",
		BankID:       "400300",
		BankIDCode:   "GBDSC",
		BIC:          "NWBKGB22",
	})
	if err != nil {
		t.Error(err)
	}
	if client == nil {
		t.Error("Client point is empty")
	}
}
