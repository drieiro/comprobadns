#!/bin/bash

set -eu

IMAGE="golang"
TARGET="$(dirname "$0" | xargs realpath)"
VERSION="v1.2"

while getopts "v:i:h" opt
do
    case "$opt" in
        v)
            VERSION="$OPTARG"
            ;;
        i)
            IMAGE="$OPTARG"
            ;;
        h)
            echo "Usage: $0 [-i image] [-v version]"
            exit 0
            ;;
        *)
            exit 1
            ;;
    esac
done

main() {
    [ -d "$TARGET/bin" ] || mkdir "$TARGET/bin"
    docker run --rm --name chkdns-build-$$ \
                    --volume "$TARGET/bin:/go/bin" \
                    --volume "$TARGET:/go/src" \
                    --workdir /go/src \
                    --user root "$IMAGE" \
                    sh entrypoint.sh "$(id -u)" "$(id -g)"
}

main
