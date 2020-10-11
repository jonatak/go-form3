package form3

import (
	"testing"
)

func TestAccountResourceValidationCountry(t *testing.T) {
	t.Run("AccountResource validation without country should be invalid", func(t *testing.T) {
		ar := &AccountResource{}
		if err := ar.IsValid(); err == nil {
			t.Error()
		}
	})

	t.Run("AccountResource validation without country of 3 char should be invalid", func(t *testing.T) {
		ar := &AccountResource{
			Country: "FRA",
		}

		if err := ar.IsValid(); err == nil {
			t.Error()
		}
	})

	t.Run("AccountResource validation with valid country shouldn't return error", func(t *testing.T) {
		ar := &AccountResource{
			Country: "FR",
		}

		if err := ar.IsValid(); err != nil {
			t.Error()
		}
	})
}
