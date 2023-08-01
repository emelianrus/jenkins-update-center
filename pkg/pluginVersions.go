package jenkinsSite

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

// PluginVersions type
// https://github.com/jenkins-infra/update-center2/blob/master/site/LAYOUT.md#plugin-versions-json-file
type PluginVersions struct {
	GenerationTimestamp string `json:"generationTimestamp"`

	Plugins map[string]map[string]struct {
		BuildData    string `json:"buildDate"`
		Dependencies []struct {
			Name     string `json:"name"`
			Optional bool   `json:"optional"`
			Version  string `json:"version"`
		} `json:"dependencies"`

		Name string `json:"name"`

		ReleaseTimestamp string `json:"releaseTimestamp"`
		RequiredCore     string `json:"requiredCore"`
		Sha1             string `json:"sha1"`
		Sha256           string `json:"sha256"`
		Url              string `json:"url"`
		Version          string `json:"version"`
	} `json:"plugins"`

	Signature struct {
		Certificates        []string `json:"certificates"`
		CorrectDigest       string   `json:"correct_digest"`
		CorrectDigest512    string   `json:"correct_digest512"`
		CorrectSignature    string   `json:"correct_signature"`
		CorrectSignature512 string   `json:"correct_signature512"`
	} `json:"signature"`

	UpdateCenterVersion string `json:"updateCenterVersion"`
}

// https://updates.jenkins.io/current/plugin-versions.json

func NewPluginVersions() PluginVersions {
	return PluginVersions{}
}

func (pv PluginVersions) Get() (*PluginVersions, error) {
	logrus.Debugln("loading pluginVersions")
	url := JENKINS_UPDATE_CENTER_URL + "/" + "current" + "/" + "plugin-versions.json"

	content, err := DoRequest(url)
	if err != nil {
		return nil, err
	}
	var pluginVersions PluginVersions
	if err := json.Unmarshal(content, &pluginVersions); err != nil {
		logrus.Errorln("Can not unmarshal JSON")
		logrus.Errorln(err)
		return nil, err
	}
	return &pluginVersions, nil
}
