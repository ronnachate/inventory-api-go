## Description

Inventory Rest api go version with gin and gorm, logging with zerolog

## Running the app

```bash
# development
$ go run app/main.go
```

### Test
```bash
# Run all tests
$ go test ./...

# test coverage
$ go test ./... -coverprofile [OUT_FILE]

# print out coverage in browser
$ go tool cover -html=[OUT_FILE]
```