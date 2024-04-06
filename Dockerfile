FROM golang:1.22-alpine

# Set destination for COPY
WORKDIR /app

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

# Download Go modules
COPY go.mod go.sum ./

COPY ./ ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build cmd/web/main.go

EXPOSE 8088
RUN ls
CMD ["./main"]
