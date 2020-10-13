package form3

import "net/http"

const (
	httpAccept      string = "vnd.api+json"
	httpContentType        = "application/vnd.api+json"
)

type config struct {
	URL            string
	OrganisationID string
	client         *http.Client
}

// Client define the resources for form3 Api requests
type Client struct {
	Account AccountEndpoint
}

// New create an API client to form3 api.
func New(OrganisationID string, BaseURL string) *Client {
	config := config{
		URL:            BaseURL,
		OrganisationID: OrganisationID,
		client:         &http.Client{},
	}
	return &Client{
		Account: AccountEndpoint{
			config: &config,
		},
	}
}
