FROM golang:1.15.4-alpine

WORKDIR /src

COPY . .

RUN go build -o /bin/http_api ./cmd/http_api

CMD ["/bin/http_api"]