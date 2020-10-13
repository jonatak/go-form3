// +build integration
// +build !unit

package integration

import (
	"os"
	"testing"

	"github.com/jonatak/go-form3/pkg/account"
	"github.com/jonatak/go-form3/pkg/form3"
	"github.com/stretchr/testify/assert"
)

func TestEnvSetUP(t *testing.T) {

	form3Endpoint := os.Getenv("FORM3_ENDPOINT")
	form3OrdID := os.Getenv("FORM3_ORG_ID")

	assert.NotEmpty(t, form3Endpoint, "FORM3_ENDPOINT env var should be present.")
	assert.NotEmpty(t, form3OrdID, "FORM3_ORG_ID env var should be present.")
}

func getAccountResource() (string, *account.Account) {
	accountID := "ad27e265-9605-4b4b-a0e5-3003ea9cc4dd"
	account := &account.Account{
		Country:      "GB",
		BaseCurrency: "GBP",
		BankID:       "400300",
		BankIDCode:   "GBDSC",
		BIC:          "NWBKGB22",
	}
	return accountID, account
}

func TestCreateAccount(t *testing.T) {

	form3Endpoint := os.Getenv("FORM3_ENDPOINT")
	form3OrdID := os.Getenv("FORM3_ORG_ID")

	client := form3.New(form3OrdID, form3Endpoint)

	accountID, account := getAccountResource()

	createResponse, err := client.Account.Create(accountID, account)

	assert.Nil(t, err)

	assert.NotEmpty(t, createResponse.Links.Self)

	assert.NotEmpty(t, createResponse.Data)
	assert.Equal(t, createResponse.Data.OrganisationID, form3OrdID)
	assert.Equal(t, createResponse.Data.ID, accountID)
	assert.Equal(t, createResponse.Data.Version, 0)

	assert.Equal(t, createResponse.Data.Attributes, account)
}

func TestFetchAccount(t *testing.T) {

	form3Endpoint := os.Getenv("FORM3_ENDPOINT")
	form3OrdID := os.Getenv("FORM3_ORG_ID")

	client := form3.New(form3OrdID, form3Endpoint)

	accountID, account := getAccountResource()

	response, err := client.Account.Fetch(accountID)

	assert.Nil(t, err)

	assert.NotEmpty(t, response.Links.Self)

	assert.NotEmpty(t, response.Data)
	assert.Equal(t, response.Data.OrganisationID, form3OrdID)
	assert.Equal(t, response.Data.ID, accountID)
	assert.Equal(t, response.Data.Version, 0)

	assert.Equal(t, response.Data.Attributes, account)
}

func TestDeleteAccount(t *testing.T) {

	form3Endpoint := os.Getenv("FORM3_ENDPOINT")
	form3OrdID := os.Getenv("FORM3_ORG_ID")

	client := form3.New(form3OrdID, form3Endpoint)

	accountID, _ := getAccountResource()

	responseCode, err := client.Account.Delete(accountID, 0)

	assert.Nil(t, err)
	assert.Equal(t, responseCode, 204)
}
