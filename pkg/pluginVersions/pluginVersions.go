package pluginVersions

import (
	"encoding/json"

	"github.com/emelianrus/jenkins-update-center/pkg/request"
	"github.com/sirupsen/logrus"
)

const REPO = "current"

// Core jenkins url for update-center
const JENKINS_UPDATE_CENTER_URL = "https://updates.jenkins.io"

// Endpoint file name
const URL_LOCATION = "plugin-versions.json"

const URL = JENKINS_UPDATE_CENTER_URL + "/" + REPO + "/" + URL_LOCATION

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

// TODO: DRY, make common
// Returns *PluginVersions, error type with data
func Get() (*PluginVersions, error) {

	content, err := request.DoRequestWithCache(URL)
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
