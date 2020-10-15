package form3

import "fmt"

// List accounts.
// This will return (nil, nil) in case the resources isn't found.
func (ae *AccountEndpoint) List(pageSize int) (*AccountPageResponse, error) {
	resp, err := ae.list(fmt.Sprintf("%s/v1/organisation/accounts", ae.URL), pageSize)
	return resp, err
}

// Next return next page.
func (ae *AccountEndpoint) Next(ap *AccountPageResponse) (*AccountPageResponse, error) {
	if ap.Links.Next == "" {
		return nil, nil
	}
	resp, err := ae.list(fmt.Sprintf("%s/%s", ae.URL, ap.Links.Next), 0)
	return resp, err
}

// Prev return next page.
func (ae *AccountEndpoint) Prev(ap *AccountPageResponse) (*AccountPageResponse, error) {
	if ap.Links.Prev == "" {
		return nil, nil
	}
	resp, err := ae.list(fmt.Sprintf("%s/%s", ae.URL, ap.Links.Prev), 0)
	return resp, err
}

// First return first page.
func (ae *AccountEndpoint) First(ap *AccountPageResponse) (*AccountPageResponse, error) {
	if ap.Links.First == "" {
		return nil, nil
	}
	resp, err := ae.list(fmt.Sprintf("%s/%s", ae.URL, ap.Links.First), 0)
	return resp, err
}

// Last return Last page.
func (ae *AccountEndpoint) Last(ap *AccountPageResponse) (*AccountPageResponse, error) {
	if ap.Links.Last == "" {
		return nil, nil
	}
	resp, err := ae.list(fmt.Sprintf("%s/%s", ae.URL, ap.Links.Last), 0)
	return resp, err
}
