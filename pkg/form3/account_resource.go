package form3

import "errors"

// Country define a specific type for country for further validation.
type Country string

// BaseCurrency define a specific type for BaseCurrency for further validation.
type BaseCurrency string

// InvalidAccountError represent validation error for AccountResource.
type InvalidAccountError struct {
	Field string
	Err   error
}

func (i *InvalidAccountError) Error() string {
	return i.Field + ": " + i.Err.Error()
}

// AccountResource define an account.
type AccountResource struct {
	Country               Country               `json:"country"`
	BaseCurrency          BaseCurrency          `json:"base_currency"`
	BankID                string                `json:"bank_id"`
	BankIDCode            string                `json:"bank_id_code"`
	AccountNumber         string                `json:"account_number"`
	BIC                   string                `json:"bic"`
	IBAN                  string                `json:"iban"`
	CustomerID            string                `json:"customer_id"`
	Name                  [4]string             `json:"name"`
	AlternativeNames      [3]string             `json:"alternative_names"`
	AccountClassification AccountClassification `json:"account_classification"`
	JointAccount          bool                  `json:"joint_account"`
	AccountMatchingOptOut bool                  `json:"account_matching_opt_out"`
	SecondaryIdentifier   string                `json:"secondary_identifier"`
	Switched              bool                  `json:"switcher"`
	Status                AccountStatus         `json:"status"`
}

// IsValid verify that an account is valid.
func (ac *AccountResource) IsValid() error {
	if ac.Country == "" {
		return &InvalidAccountError{
			Field: "Country",
			Err:   errors.New("Country field is mandatory"),
		}
	}
	if len(ac.Country) != 2 {
		return &InvalidAccountError{
			Field: "Country",
			Err:   errors.New("invalid country code"),
		}
	}
	if ac.BaseCurrency != "" && len(ac.BaseCurrency) != 3 {
		return &InvalidAccountError{
			Field: "BaseCurrency",
			Err:   errors.New("invalid base currency"),
		}
	}
	if ac.BankID != "" && len(ac.BankID) > 11 {
		return &InvalidAccountError{
			Field: "BaseCurrency",
			Err:   errors.New("length of BankID should be 11 char maximum"),
		}
	}
	if ac.BIC != "" && !(len(ac.BIC) == 11 && len(ac.BIC) == 8) {
		return &InvalidAccountError{
			Field: "BIC",
			Err:   errors.New("length of BIC should be 11 or 8 char"),
		}
	}
	if err := ac.AccountClassification.IsValid(); ac.AccountClassification != "" && err != nil {
		return &InvalidAccountError{
			Field: "AccountClassification",
			Err:   err,
		}
	}
	if err := ac.Status.IsValid(); ac.AccountClassification != "" && err != nil {
		return &InvalidAccountError{
			Field: "Status",
			Err:   err,
		}
	}
	return nil
}
