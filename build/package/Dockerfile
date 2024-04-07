###################################################################################
################################ Stage 1: #########################################
###################################################################################
FROM golang:1.22-alpine AS builder

# Set destination for COPY
WORKDIR /app

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

# Download Go modules
COPY ./go.mod ./go.sum ./
COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build cmd/web/main.go

###################################################################################
################################ Stage 2: #########################################
###################################################################################
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
CMD ["./main"]
