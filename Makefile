BUILD_DATE = $(shell date +"%Y-%m-%d:T%H:%M:%S")
GIT_COMMIT := $(shell git rev-parse --short HEAD || echo 'dev')
VERSION = $(shell git describe --tags --abbrev=0 | tr -d '\n')
linker_flags = '-s -X go-tricks/version.buildDate=${BUILD_DATE} -X go-tricks/version.gitCommit=${GIT_COMMIT} -X go-tricks/version.gitVersion=${VERSION}'

build:
	@go build -ldflags=${linker_flags} -o=./main ./main.go
