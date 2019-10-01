BUILD_VERSION   := $(shell cat version)
BUILD_DATE      := $(shell date "+%F %T")
COMMIT_SHA1     := $(shell git rev-parse HEAD)

release:
	gox -osarch="darwin/amd64 linux/386 linux/amd64" \
        -output="dist/{{.Dir}}_{{.OS}}_{{.Arch}}" \
    	-ldflags "-w -s \
				-X 'github.com/mritd/mmh/cmd.Version=${BUILD_VERSION}' \
				-X 'github.com/mritd/mmh/cmd.BuildDate=${BUILD_DATE}' \
				-X 'github.com/mritd/mmh/cmd.CommitID=${COMMIT_SHA1}'"

clean:
	rm -rf dist

install:
	go install -ldflags "-w -s \
						-X 'github.com/mritd/mmh/cmd.Version=${BUILD_VERSION}' \
						-X 'github.com/mritd/mmh/cmd.BuildDate=${BUILD_DATE}' \
						-X 'github.com/mritd/mmh/cmd.CommitID=${COMMIT_SHA1}'"

.PHONY : all release clean install

.EXPORT_ALL_VARIABLES:

GO111MODULE = on
GOPROXY = https://goproxy.cn,direct
GOSUMDB = sum.golang.google.cn