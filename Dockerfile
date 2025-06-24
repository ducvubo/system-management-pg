FROM golang:alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o system.management.pg.com ./cmd/server

FROM scratch

COPY ./config /config

COPY --from=builder /build/system.management.pg.com /

EXPOSE 13000

ENTRYPOINT [ "/system.management.pg.com", "config/local.yaml" ]


