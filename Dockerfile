FROM golang:1.19-alpine

WORKDIR /app
COPY . /app

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]