.PHONY: default
default: build

APP=sbanken
VERSION=1.2.0

## build: build binaries and generate checksums
build:
	mkdir dist/${VERSION}
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -ldflags '-s' -installsuffix cgo -o ${APP} cmd/sbanken/main.go
	tar -czvf ${APP}_${VERSION}_linux_amd64.tar.gz ${APP}
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -a -ldflags '-s' -installsuffix cgo -o ${APP} cmd/sbanken/main.go
	tar -czvf ${APP}_${VERSION}_darwin_amd64.tar.gz ${APP}
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -a -ldflags '-s' -installsuffix cgo -o ${APP}.exe cmd/sbanken/main.go
	tar -czvf ${APP}_${VERSION}_windows_amd64.tar.gz ${APP}.exe
	md5sum ${APP}_${VERSION}_linux_amd64.tar.gz > ${APP}_${VERSION}_checksums.txt
	md5sum ${APP}_${VERSION}_darwin_amd64.tar.gz >> ${APP}_${VERSION}_checksums.txt
	md5sum ${APP}_${VERSION}_windows_amd64.tar.gz >> ${APP}_${VERSION}_checksums.txt
	mv ${APP}_${VERSION}_linux_amd64.tar.gz dist/${VERSION}/
	mv ${APP}_${VERSION}_darwin_amd64.tar.gz dist/${VERSION}/
	mv ${APP}_${VERSION}_windows_amd64.tar.gz dist/${VERSION}/
	mv ${APP}_${VERSION}_checksums.txt dist/${VERSION}/
	rm ${APP} ${APP}.exe

.PHONY: run
## run: Run sbanken (set SBANKEN_CONFIG=path/to/config.yaml)
run:
	go run -race cmd/sbanken/main.go --config=$(SBANKEN_CONFIG)

.PHONY: test
## test: Run the tests
test:
	go test -race -v ./...

.PHONY: help
## help: Print this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
