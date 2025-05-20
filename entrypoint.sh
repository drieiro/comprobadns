#/bin/sh
env GOOS=linux GOARCH=amd64 go build -buildvcs=false -o ../bin/comprobadns-linux
env GOOS=windows GOARCH=amd64 go build -buildvcs=false -o ../bin/comprobadns-win.exe
env GOOS=linux GOARCH=arm64 go build -buildvcs=false -o ../bin/comprobadns-arm64
chown $1:$2 ../bin/*
