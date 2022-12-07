package pluginVersions

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/emelianrus/jenkins-update-center/pkg/request"
	"github.com/sirupsen/logrus"
)

const JENKINS_PLUGINS_URL = "https://updates.jenkins.io"

const (
	CACHED_FILE_NAME = "plugin-versions.json"
	URL_LOCATION     = JENKINS_PLUGINS_URL + "/" + "current/plugin-versions.json"
)

type pluginVersions struct {
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
// you can pass empty string to get latest core version package
func Get() *pluginVersions {
	var fileContent []byte

	if _, err := os.Stat(request.GetFileName(URL_LOCATION)); errors.Is(err, os.ErrNotExist) {
		logrus.Infoln("cache miss")
		fileContent, err = request.Do(URL_LOCATION)
		if err != nil {
			logrus.Println(err)
		}
	} else {
		logrus.Infoln("cache hit")
		fileContent, err = os.ReadFile(request.GetFileName(URL_LOCATION))
		if err != nil {
			logrus.Println(err)
		}
	}

	var pluginVersions pluginVersions
	if err := json.Unmarshal(fileContent, &pluginVersions); err != nil {
		logrus.Errorln(err)
		logrus.Errorln("Can not unmarshal JSON")
	}
	return &pluginVersions
}
