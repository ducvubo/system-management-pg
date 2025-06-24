# FROM golang:alpine AS builder

# WORKDIR /build

# COPY . .

# RUN go mod download

# RUN go build -o system.management.pg.com ./cmd/server

# FROM scratch

# COPY ./config /config

# COPY --from=builder /build/system.management.pg.com /

# EXPOSE 13000

# ENTRYPOINT [ "/system.management.pg.com", "config/local.yaml" ]



FROM golang:1.22-alpine AS builder

RUN apk add --no-cache git

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -trimpath -ldflags="-s -w" -o system.management.pg.com ./cmd/server


FROM alpine:latest
COPY system.management.pg.com /
COPY ./config /config
EXPOSE 13000
ENTRYPOINT ["/system.management.pg.com", "config/local.yaml"]
