.PHONY: build run upx test shbuild docker swag db buildx

# Default Dev Env
config= ../configs/
app_os_arch = $(app)_$(os)_$(arch)

build:
	@echo "Build $(app)"
	cd cmd/$(app) && CGO_ENABLED=0 go build -gcflags="-m -l" -ldflags="-w -s" -o=../../bin/$(app)

buildx:
	@echo "Build $(app)"
	cd cmd/$(app) && CGO_ENABLED=0 GOOS=$(os) GOARCH=$(arch) go build -gcflags="-m -l" -ldflags="-w -s" -o=../../bin/$(app_os_arch)

run:
	@echo "Build $(app) and Run"
	cd cmd/$(app) && go build -o=../../bin/${app}_${mode} && \
	cd ../../bin && ./$(app)_${mode} server --mode=$(mode) --app=$(app) --config=$(config)

upx:
	@echo 'Build $(app) command:'
	cd cmd/$(app) && \
	CGO_ENABLED=0 go build -gcflags="-m -l" -ldflags="-w -s" -o=../../bin/${app}_${mode} && \
    cd ../../bin && \
	upx -$(level) ${app}_${mode} -o $(app)_upx_${mode} && ls -lh  && \
	./$(app)_$(mode) server --mode=$(mode) --app=$(app) --config=$(config)

docker:
	pwd && chmod +x ./scripts/docker.sh && ./scripts/docker.sh $(app) $(mode)

swag:
	swag init -g cmd/admin/main.go --output ./api/admin --exclude ./internal/api && \
	swag init -g cmd/api/main.go --output ./api/api --exclude ./internal/admin

# make db h=127.0.0.1 P=3306 u=root p=root@123456 db=admin_dev
db:
	pwd && chmod +x ./scripts/database.sh && ./scripts/database.sh $(h) $(P) $(u) $(p) $(db)

job:
	@echo "Build $(app) and Run"
	cd cmd/$(app) && go build -o=../../bin/${app}_${mode} && \
	cd ../../bin && ./$(app)_${mode} job --mode=$(mode) --app=$(app) --config=$(config)