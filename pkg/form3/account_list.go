package form3

import (
	"fmt"
)

func (ae *AccountEndpoint) getPage(page string, ap *AccountPageResponse) (*AccountPageResponse, error) {
	if page == "" {
		return nil, nil
	}
	resp, err := ae.list(fmt.Sprintf("%s/%s", ae.URL, page), 0)
	return resp, err
}

// List accounts.
// This will return (nil, nil) in case the resources isn't found.
func (ae *AccountEndpoint) List(pageSize int) (*AccountPageResponse, error) {
	resp, err := ae.list(fmt.Sprintf("%s/v1/organisation/accounts", ae.URL), pageSize)
	return resp, err
}

// Next return next page.
func (ae *AccountEndpoint) Next(ap *AccountPageResponse) (*AccountPageResponse, error) {
	return ae.getPage(ap.Links.Next, ap)
}

// Prev return next page.
func (ae *AccountEndpoint) Prev(ap *AccountPageResponse) (*AccountPageResponse, error) {
	return ae.getPage(ap.Links.Prev, ap)
}

// First return first page.
func (ae *AccountEndpoint) First(ap *AccountPageResponse) (*AccountPageResponse, error) {
	return ae.getPage(ap.Links.First, ap)
}

// Last return Last page.
func (ae *AccountEndpoint) Last(ap *AccountPageResponse) (*AccountPageResponse, error) {
	return ae.getPage(ap.Links.Last, ap)
}
