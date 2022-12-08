package pluginVersions

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/emelianrus/jenkins-update-center/pkg/request"
	"github.com/sirupsen/logrus"
)

// Core jenkins url for update-center
const JENKINS_UPDATE_CENTER_URL = "https://updates.jenkins.io"

const (
	URL_LOCATION = "plugin-versions.json"
	URL          = JENKINS_UPDATE_CENTER_URL + "/" + "current/" + URL_LOCATION
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

// TODO: DRY, make common
// Returns *PluginVersions type with data
func Get() *PluginVersions {
	var fileContent []byte

	CACHE_FILE_NAME := URL_LOCATION

	if _, err := os.Stat(CACHE_FILE_NAME); errors.Is(err, os.ErrNotExist) {
		logrus.Infoln("cache miss")
		fileContent, err = request.Do(URL)
		if err != nil {
			logrus.Println(err)
		}
	} else {
		logrus.Infoln("cache hit")
		fileContent, err = os.ReadFile(CACHE_FILE_NAME)
		if err != nil {
			logrus.Println(err)
		}
	}

	var pluginVersions PluginVersions
	if err := json.Unmarshal(fileContent, &pluginVersions); err != nil {
		logrus.Errorln(err)
		logrus.Errorln("Can not unmarshal JSON")
	}
	return &pluginVersions
}
