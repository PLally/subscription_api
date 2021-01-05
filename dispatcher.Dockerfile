FROM golang:1.15.4-alpine

WORKDIR /src

COPY . .

RUN go build -o /bin/dispatcher ./cmd/dispatcher

CMD ["/bin/dispatcher"]