echo "Build Stage"
go build -v -o=./bin/$1_build  ./cmd/$1/main.go