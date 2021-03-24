FROM golang:1.15.4-alpine

WORKDIR /src

COPY . .

RUN go build -o /bin/rpc ./cmd/rpc

CMD ["/bin/rpc"]