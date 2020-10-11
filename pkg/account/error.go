package account

import "errors"

// InvalidAccountError represent validation error for AccountResource.
type InvalidAccountError struct {
	Field string
	Err   error
}

// ErrFieldMandatory an error for mandatory field.
var ErrFieldMandatory = errors.New("Field is mandatory")

// ErrFieldInvalid error for invalid field.
var ErrFieldInvalid = errors.New("Field is invalid")

// ErrFieldBankIDInvalidLength error for invalid field length.
var ErrFieldBankIDInvalidLength = errors.New("length of BankID should be 11 char maximum")

// ErrFieldBICInvalidLength error for invalid field length.
var ErrFieldBICInvalidLength = errors.New("length of BIC should be 11 or 8 char")

func (i *InvalidAccountError) Error() string {
	return i.Field + ": " + i.Err.Error()
}

func (i *InvalidAccountError) Unwrap() error {
	return i.Err
}
