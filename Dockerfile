FROM golang:latest

LABEL maintainer="Prem <prem.svmm@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

ENTRYPOINT ["sh", "entrypoint.sh"]