package jenkinsSite

const JENKINS_UPDATE_CENTER_URL = "https://updates.jenkins.io"

// jenkinsSiteInterface represents the methods used from the external JenkinsSite package.
type JenkinsSiteInterface interface {
	GetPluginVersions() (*PluginVersions, error)
	GetUpdateCenter(coreVersion string) (*UpdateCenter, error)
	GetStableCoreVersion() (string, error)
	GetLatestCoreVersion() (string, error)
}

type JenkinsSite struct {
	PluginVersions PluginVersions
	UpdateCenter   UpdateCenter
}

func NewJenkinsSite() JenkinsSite {
	pluginVersions := NewPluginVersions()
	updateCenter := NewUpdateCenter()

	return JenkinsSite{
		PluginVersions: pluginVersions,
		UpdateCenter:   updateCenter,
	}
}

func (js JenkinsSite) GetPluginVersions() (*PluginVersions, error) {
	return js.PluginVersions.Get()
}

func (js JenkinsSite) GetUpdateCenter(coreVersion string) (*UpdateCenter, error) {
	return js.UpdateCenter.Get(coreVersion)
}
