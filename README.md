# Primeapp-cli

# Webapp
Init
```
cd webapp/
go mod init webapp
```

Run app

```
docker-compose up -d  # spins up db
go run ./cmd/web  # runs app itself
http://localhost:8080/
```

Test

`go test -v ./...`

Generate cov report:

```
go test ./cmd/web -v -coverprofile=coverage.out . && go tool cover -html=coverage.out -o coverage.html
```

Run integration tests (Spin up thier own postgres via docker image, in the end destroy it)

```
go test -v -tags=integration . 
```

# Using Postgres db

`docker-compose up` to start db

Correct user is 
```
admin@example.com
secret
```
