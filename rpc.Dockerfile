FROM golang:1.15.4-alpine

WORKDIR /src

COPY . .

RUN go build -o /bin/rpc ./cmd/rpc

WORKDIR /

CMD ["/bin/rpc"]