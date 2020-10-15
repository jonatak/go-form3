# go-form3

Go library for form3 api.

## Presentation

- Submitter: __Jonathan Billaud__
- Linkedin: https://www.linkedin.com/in/jonathanbillaud/

## Go Experience

I am new to golang, until now, I did 4 projects with golang (5 counting this one):
- A wrapper around python pip
- A kafka message republisher with a RestApi
- Plugins for a slack bot
- A project using kafka goka to monitore kafka messages

## Run Unit Test

```bash
$ make unit-test
```

## Run Integration Test

The integration test use docker and docker-compose.
```bash
$ make integration-test
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

    accountId := "ad27e265-9605-4b4b-a0e5-3003ea9cc4dd"
    account := &account.Account{
		Country:      "GB",
		BaseCurrency: "GBP",
		BankID:       "400300",
		BankIDCode:   "GBDSC",
		BIC:          "NWBKGB22",
    }

    response, err := client.Account.Create("ad27e265-9605-4b4b-a0e5-3003ea9cc4dd", )

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
