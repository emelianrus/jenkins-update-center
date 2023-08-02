package main

import (
	"fmt"

	"github.com/emelianrus/jenkins-update-center/v2/pkg/jenkinsSite"
)

func main() {
	js := jenkinsSite.NewJenkinsSite()

	latestCoreVersion, _ := js.GetLatestCoreVersion()
	fmt.Println(latestCoreVersion)

	stableCoreVersion, _ := js.GetStableCoreVersion()
	fmt.Println(stableCoreVersion)

	releaseNotes, _ := js.DownloadReleaseNotes("blueocean")
	fmt.Println(releaseNotes)

	pluginVersions, _ := js.PluginVersions.Get()
	fmt.Println(pluginVersions)

	updateCenter, _ := js.UpdateCenter.Get("2.401.3")
	fmt.Println(updateCenter)
}
