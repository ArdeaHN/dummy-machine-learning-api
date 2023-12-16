FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main

EXPOSE 8000

CMD ["./main"]
