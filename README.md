# go-form3

Go library for form3 api.

## Tech decision

I decided to stick with the standard library for unit and integration testing, I only included
testify for assert.

I choose to use the `-short` option to separate unit and integration test (instead of build flag).

There isn't any uuid generation, I let the user decide how he will generate it.

I tried to keep the code really simple as I only have to implement one endpoint.
Also for the pagination I used simple method to reflect the api, but one possible improvement could be to use a custom iterator:

```golang
iter := client.Account.List(limit)
for iter.Next() {
  account := iter.Account()
  doSomething(account)
}
```

The downside of an iterator would be that you would have to create one for each endpoint or resort to using `interface{}`.

I created the account package because I assumed that we might want extra validation on the client side before submitting data to the API.

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

  account := client.Account.List(limit)

  // Loop over accounts
  for accountPage, err := client.Account.List(limit); accountPage != nil; account, err = client.Account.Next(accountPage) {
    for account := range accountPage.Data {
      doSomething(account)
    }
  }
}
```
