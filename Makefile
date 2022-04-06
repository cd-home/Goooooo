.PHONY: build run upx

os = darwin


app_os = $(app)_$(os)


build:
	@echo "Build $(app)"
	cd cmd/$(app) && CGO_ENABLED=0 GOOS=$(os) GOARCH=amd64 go build -o=../../bin/$(app_os)

run:
	@echo "Build $(app) and Run"
	cd cmd/$(app) && go build -o=../../bin && \
	cd ../../bin && ./$(app) -mode=dev -config=../configs/


upx:
	@echo 'Build $(app) command:'
	cd cmd/$(app) && \
	CGO_ENABLED=0 GOOS=$(os) GOARCH=amd64 go build -gcflags="-m -l" -ldflags="-w -s" -o=../../bin  && \
    cd ../../bin && rm -f $(app_os) && \
	upx -$(level) $(app) -o $(app_os) && ls -lh  && \
	./$(app_os) -mode=dev -config=../configs/