FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /usr/src/app

COPY . .

RUN go mod download

RUN go build -o bin/api api/main.go

FROM golang:alpine

WORKDIR /app

COPY --from=builder /usr/src/app/bin .

CMD ["./api"]