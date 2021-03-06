FROM golang:latest

LABEL maintainer="beliz çimen sevinç"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 8080

RUN go build

CMD [""]
