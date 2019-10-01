BUILD_VERSION   := $(shell cat version)
BUILD_TIME		:= $(shell date "+%F %T")
COMMIT_SHA1     := $(shell git rev-parse HEAD)

VERSION_PATH	   	= $(shell awk '/^module/{print $$2}' $(shell go env GOMOD))/util
LD_GIT_COMMIT      	= -X '$(VERSION_PATH).GitCommit=$(COMMIT_SHA1)'
LD_BUILD_TIME      	= -X '$(VERSION_PATH).BuildTime=$(BUILD_TIME)'
LD_GO_VERSION      	= -X '$(VERSION_PATH).GoVersion=`go version`'
LD_GATEWAY_VERSION	= -X '$(VERSION_PATH).Version=$(BUILD_VERSION)'
LD_FLAGS           	= "$(LD_GIT_COMMIT) $(LD_BUILD_TIME) $(LD_GO_VERSION) $(LD_GATEWAY_VERSION) -w -s"

release:
	gox -osarch="darwin/amd64 linux/386 linux/amd64" \
        -output="dist/{{.Dir}}_{{.OS}}_{{.Arch}}" \
    	-ldflags $(LD_FLAGS)

clean:
	rm -rf dist

install:
	go install -ldflags $(LD_FLAGS)

.PHONY : release clean install

.EXPORT_ALL_VARIABLES:

GO111MODULE = on
GOPROXY = https://goproxy.cn,direct
GOSUMDB = sum.golang.google.cn