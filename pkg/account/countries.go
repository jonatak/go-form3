package account

import "errors"

// Country define a specific type for country for further validation.
type Country string

// Const of all Accepted countries
const (
	GB Country = "GB"
	AU         = "AU"
	BE         = "BE"
	CA         = "CA"
	FR         = "FR"
	GR         = "GR"
	HK         = "HK"
	IT         = "IT"
	LU         = "LU"
	NL         = "NL"
	PL         = "PL"
	PT         = "PT"
	ES         = "ES"
	CH         = "CH"
	US         = "US"
)

// ErrInvalidCountry is a specific error returned if a country isn't valid.
var ErrInvalidCountry = errors.New("Submitted country is invalid")

// IsValid validate that a country is in the range of supported code.
func (c Country) IsValid() error {
	switch c {
	case GB, AU, BE, CA, FR, GR, HK, IT, LU, NL, PL, PT, ES, CH, US:
		return nil
	}
	return ErrInvalidCountry
}
