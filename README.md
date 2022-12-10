
# Jenkins update center parser go

Golang package to parse jenkins update center json file/endpoint

original repo:
https://github.com/jenkins-infra/update-center2/blob/master/site/LAYOUT.md

parses:
`https://updates.jenkins.io`

examples:

```
	fmt.Println("----------- Update center ----------->>")
	updateCenter, _ := updateCenter.Get("2.232.3")
	// updateCenter, _ := updateCenter.Get("") // will use latest core version

	// all items you can find here
	// https://github.com/emelianrus/jenkins-update-center/blob/master/pkg/updateCenter/updateCenter.go#L24-L108
	fmt.Println(updateCenter.Plugins["blueocean"].Labels)
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
```


## TODO:

* support repos
```
stable
current
experimental
```

* Release history https://github.com/jenkins-infra/update-center2/blob/master/site/LAYOUT.md#release-history-json-file
* Plugin documentation https://github.com/jenkins-infra/update-center2/blob/master/site/LAYOUT.md#plugin-documentation-urls-json-file
