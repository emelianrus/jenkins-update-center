package jenkinsSite

func (c JenkinsSite) GetLatestCoreVersion() (string, error) {
	latestVersionUrl := JENKINS_UPDATE_CENTER_URL + "/current/" + "latestCore.txt"
	content, err := DoRequest(latestVersionUrl)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (c JenkinsSite) GetStableCoreVersion() (string, error) {
	STABLE_URL := JENKINS_UPDATE_CENTER_URL + "/stable/" + "latestCore.txt"

	content, err := DoRequest(STABLE_URL)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
