PATH=$(pwd)/bin:$PATH

./bin/protoc/bin/protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. ./proto/dynamicequity/service.proto