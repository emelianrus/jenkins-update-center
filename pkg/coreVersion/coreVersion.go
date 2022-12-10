package coreVersion

// we don't have cache for coreVersion module

import (
	"github.com/emelianrus/jenkins-update-center/pkg/request"
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
	content, err := request.DoRequest(LATEST_URL)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// Get stable core version of jenkins from update center
// https://github.com/jenkins-infra/update-center2/blob/master/site/LAYOUT.md#latest-core-file
func GetStableCoreVersion() (string, error) {
	content, err := request.DoRequest(STABLE_URL)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
