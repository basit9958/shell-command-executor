FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

# Copy the go source
COPY main.go main.go
COPY models/ models/
COPY routers/ routers/

# Build
RUN go build -o app

EXPOSE 3000

CMD ["./app"]

