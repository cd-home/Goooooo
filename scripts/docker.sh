docker --version
localtime=`date +%Y%m%d%H%M%S`
version=$1:${localtime}
service=$1_$2
echo "Docker Build"
docker build -t ${version} -f ./build/$1/Dockerfile .
echo "Docker Run"
docker run -d --name ${service} -p 8080:8080  ${version} --app=$1 --mode=$2
# docker run -d --name ${service} -p 8080:8080 -e APP_NAME=$1 -e APP_MODE=$2 ${version}