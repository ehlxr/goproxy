BUILD_VERSION   := $(shell cat version)
BUILD_TIME	:= $(shell date "+%F %T")
COMMIT_SHA1     := $(shell git rev-parse HEAD)

VERSION_PATH	   	= $(shell awk '/^module/{print $$2}' $(shell go env GOMOD))/util
LD_GIT_COMMIT      	= -X '$(VERSION_PATH).GitCommit=$(COMMIT_SHA1)'
LD_BUILD_TIME      	= -X '$(VERSION_PATH).BuildTime=$(BUILD_TIME)'
LD_GO_VERSION      	= -X '$(VERSION_PATH).GoVersion=`go version`'
LD_GATEWAY_VERSION	= -X '$(VERSION_PATH).Version=$(BUILD_VERSION)'
LD_FLAGS           	= "$(LD_GIT_COMMIT) $(LD_BUILD_TIME) $(LD_GO_VERSION) $(LD_GATEWAY_VERSION) -w -s"

build:
	@$(if $(findstring 0,$(shell type gox >/dev/null 2>&1;echo $$?)),,\
		echo "Can't find gox command, will start installation...";\
		GO111MODULE=off go get -v -u github.com/mitchellh/gox)

	@gox -ldflags $(LD_FLAGS) -osarch="darwin/amd64 linux/386 linux/amd64 windows/amd64" \
		-output="dist/{{.Dir}}_{{.OS}}_{{.Arch}}"

clean:
	@rm -rf dist

install:
	@go install -ldflags $(LD_FLAGS)

# 压缩。需要安装 https://github.com/upx/upx
upx:
	@upx dist/**

release: build
	@$(if $(findstring 0,$(shell type ghr >/dev/null 2>&1;echo $$?)),,\
		echo "Can't find ghr command, will start installation...";\
		GO111MODULE=off go get -v -u github.com/tcnksm/ghr)

	@ghr -u ehlxr -t $(GITHUB_RELEASE_TOKEN) -replace -delete --debug ${BUILD_VERSION} dist

.PHONY : build release clean install upx

# this tells 'make' to export all variables to child processes by default.
.EXPORT_ALL_VARIABLES:

GO111MODULE = on
GOPROXY = https://goproxy.cn,direct
GOSUMDB = sum.golang.google.cn
