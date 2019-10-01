package util

import (
	"encoding/base64"
	"fmt"
	"runtime"
)

var bannerBase64 = "DQogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgDQogICBfX19fICAgX19fXyBfX19fX19fX19fX19fICBfX19fX19fICBfX19fX18uX18uDQogIC8gX19fXCAvICBfIFxcX19fXyBcXyAgX18gXC8gIF8gXCAgXC8gICZsdDsgICB8ICB8DQogLyAvXy8gICZndDsgICZsdDtfJmd0OyApICB8XyZndDsgJmd0OyAgfCBcKCAgJmx0O18mZ3Q7ICZndDsgICAgJmx0OyBcX19fICB8DQogXF9fXyAgLyBcX19fXy98ICAgX18vfF9ffCAgIFxfX19fL19fL1xfIFwvIF9fX198DQovX19fX18vICAgICAgICB8X198ICAgICAgICAgICAgICAgICAgICAgXC9cLyAgICAgDQo="
var versionTpl = `%s

Name: goproxy
Version: %s
Arch: %s
BuildTime: %s
GitCommit: %s
GoVersion: %s
`

var (
	Version   string
	BuildTime string
	GitCommit string
	GoVersion string
)

// PrintVersion Print out version information
func PrintVersion() {
	banner, _ := base64.StdEncoding.DecodeString(bannerBase64)
	fmt.Printf(versionTpl, banner, Version, runtime.GOOS+"/"+runtime.GOARCH, BuildTime, GitCommit, GoVersion)
}
