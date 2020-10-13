package form3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jonatak/go-form3/pkg/account"
)

// AccountEndpoint for all account resources methods.
type AccountEndpoint struct {
	*config
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

func (ae *AccountEndpoint) doRequest(req *http.Request) (*http.Response, error) {
	req.Header.Set("Accept", httpAccept)
	req.Header.Set("Content-Type", httpContentType)
	return ae.client.Do(req)
}

// Create an account.
func (ae *AccountEndpoint) Create(AccountID string, ac *account.Account) (*account.CreateResponse, error) {
	if err := ac.IsValid(); err != nil {
		return nil, err
	}

	request := &account.CreateRequest{
		Data: account.Resource{
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

	resp, err := ae.doRequest(req)

	if err != nil {
		return nil, &APIError{Err: err}
	}

	defer resp.Body.Close()

	switch resp.StatusCode {

	case 201:
		response := account.CreateResponse{}
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
