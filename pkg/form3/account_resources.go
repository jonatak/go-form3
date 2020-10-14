package form3

import (
	"time"

	"github.com/jonatak/go-form3/pkg/account"
)

// AccountResource define the message structure for account request and response.
type AccountResource struct {
	Type           string           `json:"type"`
	ID             string           `json:"id"`
	OrganisationID string           `json:"organisation_id"`
	Version        int              `json:"version,omitempty"`
	CreatedOn      time.Time        `json:"created_on,omitempty"`
	ModifiedOn     time.Time        `json:"modified_on,omitempty"`
	Attributes     *account.Account `json:"attributes"`
}

// AccountResponse define the structure for account creation response.
type AccountResponse struct {
	Data  AccountResource `json:"data"`
	Links struct {
		Self string `json:"self"`
	}
}

type createAccountRequest struct {
	Data AccountResource `json:"data"`
}

type Accounts struct {
	accountEndpoint *AccountEndpoint
	Data            []AccountResource
}

func (ac *Accounts) Next() {}
