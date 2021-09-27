package version

import (
	"fmt"
)

var (
	CCVersion       = "17.03.28"
	CCTag           = "2017-03-28 Release"
	CCBRANCH        = ""
	CCBuildTime     = "2017-03-28 19:50:00"
	CCGitHash       = "unknown"
	CCRunMode       = "product"   // product, test, dev
	CCDistro        = "community" // enterprise, community
	CCDistroVersion = "9999.9999.9999"
	ServiceName     = "unknown"
)

// CCRunMode enumeration
var (
	CCRunModeProduct = "product"
	CCRunModeTest    = "test"
	CCRunModeDev     = "dev"
)

var (
	CCDistrEnterprise = "enterprise"
	CCDistrCommunity  = "community"
)

var (
	CanCreateSetModuleWithoutTemplate = true
)

func ShowVersion() {
	fmt.Printf("%s", GetVersion())
}

// GetVersion return the version info
func GetVersion() string {
	version := fmt.Sprintf(`Version     : %s
Tag         : %s
BuildTime   : %s
GitHash     : %s
RunMode     : %s
Distribution: %s
ServiceName : %s
`, CCVersion, CCTag, CCBuildTime, CCGitHash, CCRunMode, CCDistro, ServiceName)
	return version
}
