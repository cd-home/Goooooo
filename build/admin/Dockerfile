FROM golang:1.17 AS Builder
LABEL stage=Build
WORKDIR /app/build/

# Build Variables
ARG APP = "admin"
ARG GOOS  = "linux"
ARG GOARCH = "amd64"
ARG TARGET = ${APP}

# Build Variables And Runtime Variables
ENV GOOS=${GOOS} CGO_ENABLED=0 GOARCH=${GOARCH} GOPROXY=https://goproxy.cn,direct

# Cache
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .

RUN cd ./cmd/${APP} && CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build -o=../../bin/${TARGET} .


FROM alpine:3.14
LABEL stage=Run
WORKDIR /app/release/
COPY --from=Builder /app/build/ /app/release/
EXPOSE 8080
RUN apk update && \
	apk add --no-cache \
	vim && \
	curl && \
	ca-certificates && \
	bash && \
	tzdata && \
	ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
	echo Asia/Shanghai > /etc/timezone

ENTRYPOINT ["./bin/admin", "server"]