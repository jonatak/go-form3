# go-form3

Go library for form3 api.

## Presentation

- Submitter: __Jonathan Billaud__
- Linkedin: https://www.linkedin.com/in/jonathanbillaud/

## Go Experience

I am new to golang. So far, I've done 4 projects with golang (5 counting this one):
- A wrapper around python pip
- A kafka message republisher with a RestApi
- Plugins for a slack bot
- A project using kafka goka to monitor kafka messages

## Tech decision

I decided to stick with the standard library for unit and integration testing, I only included
testify for assert.

I choose to put the integration test in its own package, I don't know what the best practice within the Go community,
but I though that having a separate package made more sense as I want to be able to test the library the same way a user will use it.

I decided to use build flags to separate unit and integration test.

There isn't any uuid generation, I let the user decide how he will generate it.

## Run Unit Test

```bash
$ make unit-test
```

## Run Integration Test

The integration test use docker and docker-compose.
It will run integration and unit test.
```bash
$ make tests
```

## Usage

```golang
import (
    "github.com/jonatak/go-form3/pkg/form3"
)

func main() {
    form3OrdID := "your-form3-organisation-uuid"
    form3Endpoint := "https://api.form3.tech"
    client := form3.New(form3OrdID, form3Endpoint)

    // Create an account

    accountId := "ad27e265-9605-4b4b-a0e5-3003ea9cc4dd" // you should use a valid uuid as an id.
    account := &account.Account{
		Country:      "GB",
		BaseCurrency: "GBP",
		BankID:       "400300",
		BankIDCode:   "GBDSC",
		BIC:          "NWBKGB22",
    }

    response, err := client.Account.Create(accountId, account)

    // Fetch Account
    response, err := client.Account.Fetch(accountID)

    // Delete an account
    responseCode, err := client.Account.Delete(accountID, response.Data.Version)

    // List accounts
    limit := 10
    accounts, err := client.Account.List(limit)

    // Retrieve Next page
    accounts, err = client.Account.Next(accounts)

    // Retrieve Previous page
    accounts, err = client.Account.Prev(accounts)

    // Retrieve Last page
    accounts, err = client.Account.Last(accounts)

    // Retrieve First page
    accounts, err = client.Account.First(accounts)
}
```
