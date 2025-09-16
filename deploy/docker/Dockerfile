FROM golang:1.25.0 AS builder

WORKDIR /souph

COPY . .
COPY go.mod go.sum ./

RUN go mod download

COPY ./deploy/docker/.env /souph/.env
COPY ./souph.json /souph/souph.json

ENV GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0

RUN go build -o /souph/main ./cmd

FROM alpine:latest

WORKDIR /souph

COPY --from=builder /souph/main .
COPY --from=builder /souph/souph.json souph.json

CMD ["./main"]