package account

import (
	"errors"

	"log"
)

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
// Note for reviewer: I assumed that the list of countries listed in the documentation
// are the only supported one, I guess that having an extra library here to validate the country
// code on the client side could be a nice to have, didn't include any here to match the specs of the
// take home exercise.
func (c Country) IsValid() error {
	switch c {
	case GB, AU, BE, CA, FR, GR, HK, IT, LU, NL, PL, PT, ES, CH, US:
		return nil
	default:
		if len(c) == 2 { // Allow other country code to not tied client library and server too much.
			log.Printf("Country code %s isn't in the list of supported countries", c)
			return nil
		}
	}
	return ErrInvalidCountry
}
