package account_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/jonatak/go-form3/pkg/account"
	"github.com/stretchr/testify/assert"
)

func assertReturnedError(t *testing.T, got, want error) {
	assert.True(t, errors.Is(got, want), fmt.Sprintf("got %s, want %s", got, want))
}

func TestAccountValidation(t *testing.T) {

	t.Run("Account validation without country should be invalid", func(t *testing.T) {
		ar := &account.Account{}
		assertReturnedError(t, ar.IsValid(), account.ErrFieldMandatory)
	})

	t.Run("Account validation without country of 3 char should be invalid", func(t *testing.T) {
		ar := &account.Account{
			Country: "FRA",
		}
		assertReturnedError(t, ar.IsValid(), account.ErrInvalidCountry)
	})

	t.Run("Account validation with valid country shouldn't return error", func(t *testing.T) {
		ar := &account.Account{
			Country: "FR",
		}
		assertReturnedError(t, ar.IsValid(), nil)
	})

	t.Run("Base currency with more than 3 char should be invalid", func(t *testing.T) {
		ar := &account.Account{
			Country:      "FR",
			BaseCurrency: "EURO",
		}
		assertReturnedError(t, ar.IsValid(), account.ErrFieldInvalid)
	})

	t.Run("Base currency EUR should be valid", func(t *testing.T) {
		ar := &account.Account{
			Country:      "FR",
			BaseCurrency: "EUR",
		}
		assertReturnedError(t, ar.IsValid(), nil)
	})

	t.Run("Base currency with more than 3 char should be invalid", func(t *testing.T) {
		ar := &account.Account{
			Country:      "FR",
			BaseCurrency: "EURO",
		}
		assertReturnedError(t, ar.IsValid(), account.ErrFieldInvalid)
	})

	t.Run("Base currency EUR should be valid", func(t *testing.T) {
		ar := &account.Account{
			Country:      "FR",
			BaseCurrency: "EUR",
		}
		assertReturnedError(t, ar.IsValid(), nil)
	})

	t.Run("BankID with more than 11 char is invalid", func(t *testing.T) {
		ar := &account.Account{
			Country:      "FR",
			BaseCurrency: "EUR",
			BankID:       "1234567890av",
		}
		assertReturnedError(t, ar.IsValid(), account.ErrFieldBankIDInvalidLength)
	})

	t.Run("BankID with less than 11 char is invalid", func(t *testing.T) {
		ar := &account.Account{
			Country:      "FR",
			BaseCurrency: "EUR",
			BankID:       "1234567890",
		}
		assertReturnedError(t, ar.IsValid(), nil)
	})
}

func TestClassificationTypeValidation(t *testing.T) {
	t.Run("Personal Classification type is valid", func(t *testing.T) {
		ar := &account.Account{
			Country:               "FR",
			BaseCurrency:          "EUR",
			BankID:                "1234567890",
			AccountClassification: "Personal",
		}
		assertReturnedError(t, ar.IsValid(), nil)
	})

	t.Run("Business Classification type is valid", func(t *testing.T) {
		ar := &account.Account{
			Country:               "FR",
			BaseCurrency:          "EUR",
			BankID:                "1234567890",
			AccountClassification: "Business",
		}
		assertReturnedError(t, ar.IsValid(), nil)
	})

	t.Run("Other Classification type is valid", func(t *testing.T) {
		ar := &account.Account{
			Country:               "FR",
			BaseCurrency:          "EUR",
			BankID:                "1234567890",
			AccountClassification: "Other",
		}
		assertReturnedError(t, ar.IsValid(), account.ErrInvalidAccountClassificationType)
	})
}

func TestStatusValidation(t *testing.T) {
	t.Run("pending status is valid", func(t *testing.T) {
		ar := &account.Account{
			Country:               "FR",
			BaseCurrency:          "EUR",
			BankID:                "1234567890",
			AccountClassification: "Personal",
			Status:                "pending",
		}
		assertReturnedError(t, ar.IsValid(), nil)
	})

	t.Run("confirmed status is valid", func(t *testing.T) {
		ar := &account.Account{
			Country:               "FR",
			BaseCurrency:          "EUR",
			BankID:                "1234567890",
			AccountClassification: "Personal",
			Status:                "confirmed",
		}
		assertReturnedError(t, ar.IsValid(), nil)
	})

	t.Run("failed status is valid", func(t *testing.T) {
		ar := &account.Account{
			Country:               "FR",
			BaseCurrency:          "EUR",
			BankID:                "1234567890",
			AccountClassification: "Personal",
			Status:                "failed",
		}
		assertReturnedError(t, ar.IsValid(), nil)
	})

	t.Run("other status is valid", func(t *testing.T) {
		ar := &account.Account{
			Country:               "FR",
			BaseCurrency:          "EUR",
			BankID:                "1234567890",
			AccountClassification: "Personal",
			Status:                "other",
		}
		assertReturnedError(t, ar.IsValid(), account.ErrInvalidAccountStatus)
	})
}
