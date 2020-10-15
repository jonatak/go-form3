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

// AccountPageResponse represent a paginate result of the list endpoints.
type AccountPageResponse struct {
	Data  []AccountResource
	Links struct {
		Self  string `json:"self"`
		First string `json:"first"`
		Last  string `json:"last"`
		Next  string `json:"next"`
		Prev  string `json:"prev"`
	}
}
