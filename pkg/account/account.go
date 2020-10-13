package account

import "time"

// BaseCurrency define a specific type for BaseCurrency for further validation.
type BaseCurrency string

// Account define an account.
type Account struct {
	Country               Country        `json:"country"`
	BaseCurrency          BaseCurrency   `json:"base_currency,omitempty"`
	BankID                string         `json:"bank_id,omitempty"`
	BankIDCode            string         `json:"bank_id_code,omitempty"`
	AccountNumber         string         `json:"account_number,omitempty"`
	BIC                   string         `json:"bic,omitempty"`
	IBAN                  string         `json:"iban,omitempty"`
	CustomerID            string         `json:"customer_id,omitempty"`
	Name                  [4]string      `json:"name,omitempty"`
	AlternativeNames      [3]string      `json:"alternative_names,omitempty"`
	AccountClassification Classification `json:"account_classification,omitempty"`
	JointAccount          bool           `json:"joint_account"`
	AccountMatchingOptOut bool           `json:"account_matching_opt_out"`
	SecondaryIdentifier   string         `json:"secondary_identifier,omitempty"`
	Switched              bool           `json:"switcher"`
	Status                Status         `json:"status,omitempty"`
}

// Resource define the message structure for account request and response.
type Resource struct {
	Type           string    `json:"type"`
	ID             string    `json:"id"`
	OrganisationID string    `json:"organisation_id"`
	Version        int       `json:"version,omitempty"`
	CreatedOn      time.Time `json:"created_on,omitempty"`
	ModifiedOn     time.Time `json:"modified_on,omitempty"`
	Attributes     *Account  `json:"attributes"`
}

// CreateRequest define the structure for account creation request.
type CreateRequest struct {
	Data Resource `json:"data"`
}

// Response define the structure for account creation response.
type Response struct {
	Data  Resource `json:"data"`
	Links struct {
		Self string `json:"self"`
	}
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
	if r.BIC != "" && !(len(r.BIC) == 11 || len(r.BIC) == 8) {
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
