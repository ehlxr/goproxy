package metadata

var BannerBase64 = "DQogIF9fXyAgIF9fICBfX19fICBfX19fICAgX18gIF8gIF8gIF8gIF8gDQogLyBfXykgLyAgXCggIF8gXCggIF8gXCAvICBcKCBcLyApKCBcLyApDQooIChfIFwoICBPICkpIF9fLyApICAgLyggIE8gKSkgICggICkgIC8gDQogXF9fXy8gXF9fLyhfXykgIChfX1xfKSBcX18vKF8vXF8pKF9fLyAgDQo="
var VersionTpl = `%s
Name: goproxy
Version: %s
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
