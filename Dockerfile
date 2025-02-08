FROM golang:1.23.6-alpine3.21 AS builder
WORKDIR /build
COPY . .
RUN GOOS=linux CGO_ENABLED=0 GOFLAGS=-mod=vendor go build -buildvcs=false -o /build/pupa -ldflags "-s -w" .

FROM alpine:3.21
COPY --from=builder /build/pupa /pupa
RUN chmod +x /pupa
EXPOSE 8080
CMD ["/pupa"]