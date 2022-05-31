.PHONY: build run upx test shbuild docker swag db

# Default Dev Env
os = darwin
mode = dev
config= ../configs/

app_os = $(app)_$(os)


build:
	@echo "Build $(app)"
	cd cmd/$(app) && CGO_ENABLED=0 GOOS=$(os) GOARCH=amd64 go build -o=../../bin/$(app_os)

run:
	@echo "Build $(app) and Run"
	cd cmd/$(app) && go build -o=../../bin/${app}_${mode} && \
	cd ../../bin && ./$(app)_${mode} server --mode=$(mode) --app=$(app) --config=$(config)

upx:
	@echo 'Build $(app) command:'
	cd cmd/$(app) && \
	CGO_ENABLED=0 GOOS=$(os) GOARCH=amd64 go build -gcflags="-m -l" -ldflags="-w -s" -o=../../bin  && \
    cd ../../bin && rm -f $(app_os) && \
	upx -$(level) $(app) -o $(app_os) && ls -lh  && \
	./$(app_os) server --mode=$(mode) --app=$(app) --config=$(config)


shbuild:
	pwd && cd ./scripts/ && chmod +x build.sh && ./build.sh $(app) $(os) $(app_os)


docker:
	pwd && chmod +x ./scripts/docker.sh && ./scripts/docker.sh $(app) $(mode)


swag:
	swag init -g cmd/admin/main.go --output ./api/admin --exclude ./internal/api && \
	swag init -g cmd/api/main.go --output ./api/api --exclude ./internal/admin

db:
	pwd && chmod +x ./scripts/database.sh && ./scripts/database.sh $(h) $(P) $(u) $(p) $(db)


job:
	@echo "Build $(app) and Run"
	cd cmd/$(app) && go build -o=../../bin/${app}_${mode} && \
	cd ../../bin && ./$(app)_${mode} job --mode=$(mode) --app=$(app) --config=$(config)