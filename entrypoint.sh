#/bin/sh
env GOOS=linux GOARCH=amd64 go build -o ../bin/chkdns-linux
env GOOS=windows GOARCH=amd64 go build -o ../bin/chkdns-win.exe
env GOOS=linux GOARCH=arm64 go build -o ../bin/chkdns-arm64
chown $1:$2 ../bin/*
