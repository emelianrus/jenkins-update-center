package jenkinsSite

type JenkinsSite struct {
	PluginVersions PluginVersions
	UpdateCenter   UpdateCenter
}

const JENKINS_UPDATE_CENTER_URL = "https://updates.jenkins.io"

func NewJenkinsSite() JenkinsSite {
	pluginVersions := NewPluginVersions()
	updateCenter := NewUpdateCenter()

	return JenkinsSite{
		PluginVersions: pluginVersions,
		UpdateCenter:   updateCenter,
	}
}
