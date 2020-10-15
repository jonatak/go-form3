package account

import "errors"

// Classification is a custom datatype to validate account classification.
type Classification string

// Allowed AccountClassification.
const (
	PersonalAccount Classification = "Personal"
	BusinessAccount Classification = "Business"
)

// ErrInvalidAccountClassificationType is a specific error for AccountClassification validation.
var ErrInvalidAccountClassificationType = errors.New("Invalid AccountClassification type")

// IsValid verify that an AccountClassification is valid.
func (ac Classification) IsValid() error {
	switch ac {
	case PersonalAccount, BusinessAccount:
		return nil
	}
	return ErrInvalidAccountClassificationType
}
