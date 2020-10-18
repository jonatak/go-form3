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
}

func (a *APIError) Error() string {
	return a.ErrorMessage
}

func (ae *AccountEndpoint) doRequest(req *http.Request) (*http.Response, error) {
	req.Header.Set("Accept", httpAccept)
	req.Header.Set("Content-Type", httpContentType)
	return ae.client.Do(req)
}

func (ae *AccountEndpoint) list(url string, pageSize int) (*AccountPageResponse, error) {

	req, err := http.NewRequest(
		"GET", url, nil,
	)

	if err != nil {
		return nil, err
	}

	if pageSize != 0 {
		q := req.URL.Query()
		q.Add("page[size]", fmt.Sprintf("%d", pageSize))
		req.URL.RawQuery = q.Encode()
	}

	resp, err := ae.doRequest(req)

	defer resp.Body.Close()

	response := AccountPageResponse{}
	switch resp.StatusCode {

	case 200:
		err := json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			return nil, err
		}
		return &response, nil
	case 404:
		return nil, &APIError{
			StatusCode: http.StatusNotFound,
		}
	default:
		apiErr := &APIError{
			StatusCode: resp.StatusCode,
		}
		err := json.NewDecoder(resp.Body).Decode(&apiErr)
		if err != nil {
			return nil, err
		}
		return nil, apiErr
	}
}

// Create an account.
func (ae *AccountEndpoint) Create(accountID string, ac *account.Account) (*AccountResponse, error) {
	if err := ac.IsValid(); err != nil {
		return nil, err
	}

	request := &createAccountRequest{
		Data: AccountResource{
			Type:           "accounts",
			ID:             accountID,
			OrganisationID: ae.OrganisationID,
			Version:        0,
			Attributes:     ac,
		},
	}

	jsonRQ, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/v1/organisation/accounts", ae.URL), bytes.NewBuffer(jsonRQ))
	if err != nil {
		return nil, err
	}

	resp, err := ae.doRequest(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	switch resp.StatusCode {

	case 201:
		response := AccountResponse{}
		err := json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			return nil, err
		}
		return &response, nil

	default:
		apiErr := &APIError{
			StatusCode: resp.StatusCode,
		}
		err := json.NewDecoder(resp.Body).Decode(&apiErr)
		if err != nil {
			return nil, err
		}
		return nil, err
	}
}

// Fetch an account.
// This will return (nil, nil) in case the resources isn't found.
func (ae *AccountEndpoint) Fetch(accountID string) (*AccountResponse, error) {

	resp, err := http.Get(fmt.Sprintf("%s/v1/organisation/accounts/%s", ae.URL, accountID))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	response := AccountResponse{}

	switch resp.StatusCode {
	case 200:
		err := json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			return nil, err
		}
		return &response, nil
	case 404:
		return nil, &APIError{
			StatusCode: http.StatusNotFound,
		}
	default:
		apiErr := &APIError{
			StatusCode: resp.StatusCode,
		}
		err := json.NewDecoder(resp.Body).Decode(&apiErr)
		if err != nil {
			return nil, err
		}
		return nil, err
	}
}

// Delete an account.
func (ae *AccountEndpoint) Delete(accountID string, version int) (int, error) {

	req, err := http.NewRequest(
		"DELETE", fmt.Sprintf("%s/v1/organisation/accounts/%s",
			ae.URL, accountID),
		nil,
	)
	if err != nil {
		return 0, err
	}

	q := req.URL.Query()
	q.Add("version", fmt.Sprintf("%d", version))
	req.URL.RawQuery = q.Encode()
	resp, err := ae.doRequest(req)

	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	return resp.StatusCode, nil
}
