package account

// BaseCurrency define a specific type for BaseCurrency for further validation.
type BaseCurrency string

// Account define an account.
type Account struct {
	Country               Country        `json:"country"`
	BaseCurrency          BaseCurrency   `json:"base_currency"`
	BankID                string         `json:"bank_id"`
	BankIDCode            string         `json:"bank_id_code"`
	AccountNumber         string         `json:"account_number"`
	BIC                   string         `json:"bic"`
	IBAN                  string         `json:"iban"`
	CustomerID            string         `json:"customer_id"`
	Name                  [4]string      `json:"name"`
	AlternativeNames      [3]string      `json:"alternative_names"`
	AccountClassification Classification `json:"account_classification"`
	JointAccount          bool           `json:"joint_account"`
	AccountMatchingOptOut bool           `json:"account_matching_opt_out"`
	SecondaryIdentifier   string         `json:"secondary_identifier"`
	Switched              bool           `json:"switcher"`
	Status                Status         `json:"status"`
}

// IsValid verify that an account is valid.
func (r *Account) IsValid() error {
	if r.Country == "" {
		return &InvalidAccountError{
			Field: "Country",
			Err:   ErrFieldMandatory,
		}
	}
	if err := r.Country.IsValid(); err != nil {
		return &InvalidAccountError{
			Field: "Country",
			Err:   err,
		}
	}
	if r.BaseCurrency != "" && len(r.BaseCurrency) != 3 {
		return &InvalidAccountError{
			Field: "BaseCurrency",
			Err:   ErrFieldInvalid,
		}
	}
	if r.BankID != "" && len(r.BankID) > 11 {
		return &InvalidAccountError{
			Field: "BaseCurrency",
			Err:   ErrFieldBankIDInvalidLength,
		}
	}
	if r.BIC != "" && !(len(r.BIC) == 11 && len(r.BIC) == 8) {
		return &InvalidAccountError{
			Field: "BIC",
			Err:   ErrFieldBICInvalidLength,
		}
	}
	if err := r.AccountClassification.IsValid(); r.AccountClassification != "" && err != nil {
		return &InvalidAccountError{
			Field: "AccountClassification",
			Err:   err,
		}
	}
	if err := r.Status.IsValid(); r.Status != "" && err != nil {
		return &InvalidAccountError{
			Field: "Status",
			Err:   err,
		}
	}
	return nil
}
