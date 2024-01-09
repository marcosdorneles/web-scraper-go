FROM golang:1.21.5

RUN mkdir /app
ADD . /app
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY .env ./.env
RUN go mod download

COPY *.go ./

RUN go build -o main .

EXPOSE 8081

CMD ["/app/main"]
