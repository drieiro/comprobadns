#/bin/sh
env GOOS=linux GOARCH=amd64 go build -o ../bin
env GOOS=windows GOARCH=amd64 go build -o ../bin
chown $1:$2 ../bin/*
