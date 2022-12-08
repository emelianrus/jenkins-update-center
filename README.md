
# Jenkins update center parser go

Golang package to parse jenkins update center json file/endpoint


`https://updates.jenkins.io`

```
updateCenter := updateCenter.Get(jenkins core version)

pluginVersions := pluginVersions.Get()

stableCoreVersion, _ := coreVersion.GetStableCoreVersion()
fmt.Println(stableCoreVersion)

latestCoreVersion, _ := coreVersion.GetLatestCoreVersion()
fmt.Println(latestCoreVersion)
```


## TODO:

* support repos
```
stable
current
experimental
```

* Latest core endpoint https://github.com/jenkins-infra/update-center2/blob/master/site/LAYOUT.md#latest-core-file
* Release history https://github.com/jenkins-infra/update-center2/blob/master/site/LAYOUT.md#release-history-json-file
* Plugin documentation https://github.com/jenkins-infra/update-center2/blob/master/site/LAYOUT.md#plugin-documentation-urls-json-file

* Make `Get` function as common for endpoints. currently we duplicate code