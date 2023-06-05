#!/bin/sh

export PATH="$PATH:$(go env GOPATH)/bin"

echo "Generating go proto sources..."
cd proto

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./god.proto

cd ..
echo "Done"
