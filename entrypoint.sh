#/bin/sh
env GOOS=linux GOARCH=amd64 go build -buildvcs=false -o ../bin/comprobadns-linux-amd64
env GOOS=windows GOARCH=amd64 go build -buildvcs=false -o ../bin/comprobadns-win-amd64.exe
env GOOS=linux GOARCH=arm64 go build -buildvcs=false -o ../bin/comprobadns-linux-arm64
chown $1:$2 ../bin/*
