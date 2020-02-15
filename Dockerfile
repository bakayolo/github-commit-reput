FROM golang:1.13-alpine as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o github-commit-reput cmd/main.go

FROM alpine:3.9.3

WORKDIR /app

COPY --from=builder /app/github-commit-reput github-commit-reput

ENTRYPOINT ["./github-commit-reput"]