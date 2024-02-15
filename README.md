# Shipping calculator MVP

Application that can calculate the number of packs we need to ship to the customer.
The API is written in Golang & be usable by a HTTP API and has unit tests.

### Testing
go test
go test -bench=.

You can also use usage.http with VSCode REST Client extension to test the API

### Build and run tests container
docker build -t shipping-test .
docker run shipping-test

### TODO (but out of scope)
- protect packSizes with RWMutex
- store packSizes in database
- input checks in UI
