
# Jenkins update center parser go

Golang package to parse jenkins update center json file/endpoint

original repo:
https://github.com/jenkins-infra/update-center2/blob/master/site/LAYOUT.md

parses:
`https://updates.jenkins.io`


## Supported endpoints

* Root update center page aka `update-center.json`
* Plugin versions aka `plugin-versions.json`
* Get stable(LTS) jenkins core `stable/latestCore.txt`
* Get latest jenkins core aka `current/latestCore.txt`
* Get plugin release notes `https://plugin-site-issues.jenkins.io/api/plugin/<pluginName>/releases`


example:

```
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
```



## TODO:

* support switch between repos
```
stable
current
experimental
```
currently only hard coded ones

* Release history https://github.com/jenkins-infra/update-center2/blob/master/site/LAYOUT.md#release-history-json-file
* Plugin documentation https://github.com/jenkins-infra/update-center2/blob/master/site/LAYOUT.md#plugin-documentation-urls-json-file
