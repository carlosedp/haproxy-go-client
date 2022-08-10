#!/bin/bash

# Set the versions
SWAGGER_VER=0.29.0
CLIENT_NATIVE_VER=4.0.0

# Clone the client-native repository with models and APIs
if [[ ! -d "client-native-repo" ]]; then
    git clone --quiet "https://github.com/haproxytech/client-native" "client-native-repo" "$@"
else
    pushd "client-native-repo" >/dev/null || return
    git checkout master
    git pull --autostash --quiet
    popd >/dev/null || return
fi

pushd client-native-repo || exit
git checkout v${CLIENT_NATIVE_VER}

# Generate the API spec
go run ./specification/build/build.go -file ./specification/haproxy-spec.yaml > haproxy_spec.yaml

# Install the swagger tool
curl -Ls -o swagger "https://github.com/go-swagger/go-swagger/releases/download/v${SWAGGER_VER}/swagger_$(go env GOOS)_$(go env GOARCH)"
chmod +x swagger

# Generate the client
mkdir -p "$GOPATH/src/github.com/carlosedp/haproxy-go-client"

MAJOR_VER=$(echo "4.0.0" |cut -d. -f1)
./swagger generate client -f haproxy_spec.yaml -A "Data Plane" -t "$GOPATH/src/github.com/carlosedp/haproxy-go-client" --existing-models "github.com/haproxytech/client-native/v${MAJOR_VER}/models" -r ../header.txt
cp -R "$GOPATH"/src/github.com/carlosedp/haproxy-go-client/* ../
# rm -rf "$GOPATH/src/github.com/carlosedp/haproxy-go-client"
popd  || exit

go mod tidy
