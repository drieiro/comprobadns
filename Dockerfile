FROM golang:alpine AS builder

WORKDIR /go/src/github.com/drieiro/chkdns

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o chkdns .

FROM alpine

COPY --from=builder /go/src/github.com/drieiro/chkdns/ .

CMD ./chkdns
