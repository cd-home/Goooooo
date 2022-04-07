docker --version
localtime=`date +%Y%m%d%H%M%S`
version=$1:${localtime}
service=$1_$2

docker build -t ${version} -f ./build/$1/Dockerfile .
docker run -d --name ${service} -p 8080:8080  ${version} -app=$1 -mode=$2