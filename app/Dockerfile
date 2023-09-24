FROM golang:1.21

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o app cmd/fleeter/main.go

EXPOSE 3000

CMD ["./app"]
