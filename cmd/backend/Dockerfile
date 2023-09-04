FROM golang:1.21.0-alpine AS build

WORKDIR /app

# COPY go.mod go.sum ./
COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN go build -o /backend cmd/backend/main.go

CMD ["/backend"]