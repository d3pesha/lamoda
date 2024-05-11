FROM golang:latest

WORKDIR /lamoda

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build cmd/main.go

CMD ["./main"]