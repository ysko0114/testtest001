FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o main main.go

EXPOSE 9000

CMD ["./main"]

