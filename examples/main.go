package main

import (
	"fmt"

	"github.com/emelianrus/jenkins-update-center/pkg/jenkinsSite"
)

func main() {
	js := jenkinsSite.NewJenkinsSite()

	latestCoreVersion, _ := js.GetLatestCoreVersion()
	fmt.Println(latestCoreVersion)

	stableCoreVersion, _ := js.GetStableCoreVersion()
	fmt.Println(stableCoreVersion)

	releaseNotes, _ := js.DownloadReleaseNotes("blueocean")
	fmt.Println(releaseNotes)

	pluginVersions, _ := js.GetPluginVersions()
	fmt.Println(pluginVersions)

	updateCenter, _ := js.GetUpdateCenter("2.401.3")
	fmt.Println(updateCenter)
}
