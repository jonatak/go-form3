package form3

import "errors"

// AccountStatus is a custom data type to define allowed account_status
type AccountStatus string

// Allowed AccountStatus
const (
	PendingAccount   AccountStatus = "pending"
	ConfirmedAccount AccountStatus = "confirmed"
	FailedAccount    AccountStatus = "failed"
)

// IsValid check that an AccountStatus is valid.
func (ac AccountStatus) IsValid() error {
	switch ac {
	case PendingAccount, ConfirmedAccount, FailedAccount:
		return nil
	}
	return errors.New("Invalid AccountClassification type")
}
