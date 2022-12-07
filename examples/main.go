package main

import (
	"fmt"

	"github.com/emelianrus/jenkins-update-center/pkg/pluginVersions"
	"github.com/emelianrus/jenkins-update-center/pkg/updateCenter"
)

func main() {

	updateCenter := updateCenter.Get("2.232.3")
	// updateCenter := updateCenter.Get("") // will use latest core version

	// all items you can find here
	// https://github.com/emelianrus/jenkins-update-center/blob/master/pkg/updateCenter/updateCenter.go#L24-L108
	fmt.Println(updateCenter.Plugins["blueocean"].Labels)

	// OR

	pluginVersions := pluginVersions.Get()
	//fmt.Println(pluginVersions.Plugins)
	fmt.Println("")
	//fmt.Println(pluginVersions.Plugins["blueocean"])

	fmt.Println("")
	fmt.Println(pluginVersions.Plugins["blueocean"]["1.25.3"])
}
