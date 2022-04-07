FROM golang:1.17 AS Builder
LABEL stage=Build
WORKDIR /app/build/

ENV GOOS=linux CGO_ENABLED=0 GOARCH=amd64 GOPROXY=https://goproxy.cn,direct

# Cache
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .

RUN chmod +x ./scripts/build.sh && ./scripts/build.sh api linux api


FROM scratch
LABEL stage=Run
WORKDIR /app/release/
COPY --from=Builder /app/build/ /app/release/
EXPOSE 8000


ENTRYPOINT ["./bin/api"]