FROM golang:alpine AS builder

WORKDIR /go/src/github.com/drieiro/comprobadns

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o comprobadns .

FROM alpine

COPY --from=builder /go/src/github.com/drieiro/comprobadns/ .

CMD ./comprobadns
