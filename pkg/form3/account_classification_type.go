package form3

import "errors"

// AccountClassification is a custom datatype to validate account_classification.
type AccountClassification string

// Allowed AccountClassification.
const (
	PersonalAccount AccountClassification = "Personal"
	BusinessAccount AccountClassification = "Business"
)

// IsValid verify that an AccountClassification is valid.
func (ac AccountClassification) IsValid() error {
	switch ac {
	case PersonalAccount, BusinessAccount:
		return nil
	}
	return errors.New("Invalid AccountClassification type")
}
