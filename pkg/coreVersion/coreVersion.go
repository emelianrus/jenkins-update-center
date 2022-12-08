package coreVersion

// we don't have cache for coreVersion module

import (
	"errors"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

// Jenkins update center root page url
const JENKINS_UPDATE_CENTER_URL = "https://updates.jenkins.io"

const URL_LOCATION = "latestCore.txt"

const (
	STABLE_URL = JENKINS_UPDATE_CENTER_URL + "/stable/" + URL_LOCATION
	LATEST_URL = JENKINS_UPDATE_CENTER_URL + "/current/" + URL_LOCATION
)

// Get latest core version of jenkins from update center
// https://github.com/jenkins-infra/update-center2/blob/master/site/LAYOUT.md#latest-core-file
func GetLatestCoreVersion() (string, error) {
	return getPage(LATEST_URL)
}

// Get stable core version of jenkins from update center
// https://github.com/jenkins-infra/update-center2/blob/master/site/LAYOUT.md#latest-core-file
func GetStableCoreVersion() (string, error) {
	return getPage(STABLE_URL)
}

// Download page and read content
func getPage(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		logrus.Errorln(err)
		return "", err
	}
	content, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		logrus.Errorln(err)
		return "", err
	}

	if response.StatusCode != 200 {
		return "", errors.New("status code is not 200")
	}
	return string(content), nil
}
