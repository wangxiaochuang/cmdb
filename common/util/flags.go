package util

import (
    goflag "flag"
    "os"
    "strings"

    "github.com/wxc/cmdb/common/auth"
    "github.com/wxc/cmdb/common/blog"
    "github.com/wxc/cmdb/common/version"

    "github.com/spf13/pflag"
)

func WordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
    if strings.Contains(name, "_") {
            return pflag.NormalizedName(strings.Replace(name, "_", "-", -1))
    }
    return pflag.NormalizedName(name)
}

// AddCommonFlags add common flags that is needed by all modules
func AddCommonFlags(cmdline *pflag.FlagSet) *bool {
    version := cmdline.Bool("version", false, "show version information")
    return version
}

func InitFlags() {
    pflag.CommandLine.SetNormalizeFunc(WordSepNormalizeFunc)
    pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
    ver := AddCommonFlags(pflag.CommandLine)
    pflag.Parse()

    // add handler if flag include --version/-v
    if *ver {
            version.ShowVersion()
            os.Exit(0)
    }

    blog.Infof("[auth] enableAuth: %v", auth.EnableAuthorize())
}
