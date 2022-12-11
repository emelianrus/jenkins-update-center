package main

import (
	"fmt"

	"github.com/emelianrus/jenkins-update-center/pkg/coreVersion"
	"github.com/emelianrus/jenkins-update-center/pkg/pluginVersions"
	"github.com/emelianrus/jenkins-update-center/pkg/updateCenter"
)

func main() {

	fmt.Println("----------- Update center ----------->>")
	updateCenterByVersion, _ := updateCenter.Get("2.232.3")
	fmt.Println(updateCenterByVersion.Core.Version)
	updateCenter, _ := updateCenter.Get("") // will use latest core version
	// all items you can find here
	// https://github.com/emelianrus/jenkins-update-center/blob/master/pkg/updateCenter/updateCenter.go#L24-L108
	fmt.Println(updateCenter.Core.Version)
	fmt.Println("<<----------- Update center -----------")

	fmt.Println("----------- PluginVersions ----------->>")
	pluginVersions, _ := pluginVersions.Get()
	fmt.Println(pluginVersions.Plugins["blueocean"]["1.25.3"])
	fmt.Println("<<----------- PluginVersions -----------")

	fmt.Println("----------- StableCoreVersion ----------->>")
	stableCoreVersion, _ := coreVersion.GetStableCoreVersion()
	fmt.Println(stableCoreVersion)
	fmt.Println("<<----------- StableCoreVersion -----------")

	fmt.Println("----------- LatestCoreVersion ----------->>")
	latestCoreVersion, _ := coreVersion.GetLatestCoreVersion()
	fmt.Println(latestCoreVersion)
	fmt.Println("<<----------- LatestCoreVersion -----------")
}
