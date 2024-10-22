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

# Test server

```
curl http://localhost:8090/auth -X POST -H "Content-Type: application/json" \
-d '{"email":"admin@example.com","password":"secret"}'
```

# Using cli to generate token pairs for test

For `/home/anvi/PycharmProjects/go-test-course-udemy/webapp/cmd/api/setup_test.go`

```
anvi@anvi-thinkpad-x1:~/PycharmProjects/go-test-course-udemy/webapp$ go run ./cmd/cli -action=expired
EXPIRED Token:
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiYXVkIjoiZXhhbXBsZS5jb20iLCJleHAiOjE3Mjg4NzcwOTksImlzcyI6ImV4YW1wbGUuY29tIiwibmFtZSI6IkpvaG4gRG9lIiwic3ViIjoiMSJ9.l_qsNXkSp6pQuxa5GvFkK2xtkNDYSCtwM3xgSSMLjt4
```
