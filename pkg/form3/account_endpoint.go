package form3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jonatak/go-form3/pkg/account"
)

// AccountEndpoint for all account resources methods.
type AccountEndpoint struct {
	*config
}

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

// CreateAccountRequest define the structure for account creation request.
type CreateAccountRequest struct {
	Data AccountResource `json:"data"`
}

// CreateAccountResponse define the structure for account creation response.
type CreateAccountResponse struct {
	Data  AccountResource `json:"data"`
	Links struct {
		Self string `json:"self"`
	}
}

// APIError define error message returned by form3 api.
type APIError struct {
	StatusCode   int
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Err          error
}

func (a *APIError) Error() string {
	if a.Err != nil {
		return a.Err.Error()
	}
	return a.ErrorMessage
}

func (a *APIError) Unwrap() error {
	return a.Err
}

// Create an account.
func (ae *AccountEndpoint) Create(AccountID string, ac *account.Account) (*CreateAccountResponse, error) {
	if err := ac.IsValid(); err != nil {
		return nil, err
	}

	request := &CreateAccountRequest{
		Data: AccountResource{
			Type:           "accounts",
			ID:             AccountID,
			OrganisationID: ae.OrganisationID,
			Version:        0,
			Attributes:     ac,
		},
	}

	jsonRQ, err := json.Marshal(request)
	if err != nil {
		return nil, &APIError{Err: err}
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/v1/organisation/accounts", ae.URL), bytes.NewBuffer(jsonRQ))
	if err != nil {
		return nil, &APIError{Err: err}
	}

	req.Header.Set("Accept", httpAccept)
	req.Header.Set("Content-Type", httpContentType)
	resp, err := ae.client.Do(req)

	if err != nil {
		return nil, &APIError{Err: err}
	}

	defer resp.Body.Close()

	switch resp.StatusCode {

	case 201:
		response := CreateAccountResponse{}
		json.NewDecoder(resp.Body).Decode(&response)
		return &response, nil

	default:
		err := APIError{
			StatusCode: resp.StatusCode,
		}
		json.NewDecoder(resp.Body).Decode(&err)
		return nil, &err
	}
}
