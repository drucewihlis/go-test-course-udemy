# Primeapp-cli

# Webapp
Init
```
cd webapp/
go mod init webapp
```

Run app

```
go run ./cmd/web
http://localhost:8080/
```

Test

`go test -v ./...`

Generate cov report:

```
cd cmd/web
go test -v -coverprofile=coverage.out . && go tool cover -html=coverage.out -o coverage.html
```
