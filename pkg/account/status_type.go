package account

import "errors"

// Status is a custom data type to define allowed account_status
type Status string

// Allowed AccountStatus
const (
	PendingAccount   Status = "pending"
	ConfirmedAccount Status = "confirmed"
	FailedAccount    Status = "failed"
)

// ErrInvalidAccountStatus is a specific error for accountStatus validation
var ErrInvalidAccountStatus = errors.New("Invalid AccountStatus")

// IsValid check that an AccountStatus is valid.
func (ac Status) IsValid() error {
	switch ac {
	case PendingAccount, ConfirmedAccount, FailedAccount:
		return nil
	}
	return ErrInvalidAccountStatus
}
