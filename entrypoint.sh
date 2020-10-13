#!/bin/bash
export PROJECT_NAME="go-form3"
export PKG="github.com/jonatak/$PROJECT_NAME"
export PKG_LIST=`go list $PKG/... | grep -v /vendor/`

# Wait until API is ready before running test.
until curl -f $FORM3_ENTRYPOINT/v1/organisation/accounts; do
  >&2 echo "API is unavailable - sleeping"
  sleep 1
done

go test -v -cover --tags=integration ./...
