go version

cd ../cmd/$1 && CGO_ENABLED=0 GOOS=$2 GOARCH=amd64 go build -o=../../bin/$3