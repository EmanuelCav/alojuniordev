FROM golang:1.22-bullseye

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

EXPOSE 4000

CMD ["./alojuniordev"]
