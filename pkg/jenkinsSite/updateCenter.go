package jenkinsSite

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

type UpdateCenter struct {
	ConnectionCheckUrl string `json:"connectionCheckUrl"`
	Core               struct {
		BuildData string `json:"buildDate"`
		Name      string `json:"name"`
		Sha1      string `json:"sha1"`
		Sha256    string `json:"sha256"`
		Size      int    `json:"size"`
		Url       string `json:"url"`
		Version   string `json:"version"`
	} `json:"core"`

	Deprecations map[string]struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"deprecations"`

	GenerationTimestamp string `json:"generationTimestamp"`
	Id                  string `json:"id"`

	Plugins map[string]struct {
		BuildData     string `json:"buildDate"`
		DefaultBranch string `json:"defaultBranch"`

		Dependencies []struct {
			Name     string `json:"name"`
			Optional bool   `json:"optional"`
			Version  string `json:"version"`
		} `json:"dependencies"`

		Developers []struct {
			DeveloperId string `json:"developerId"`
			Name        string `json:"name"`
		} `json:"developers"`

		Excerpt string `json:"excerpt"`

		Gav string `json:"gav"`

		IssueTrackers []struct {
			ReportUrl string `json:"reportUrl"`
			Type      string `json:"type"`
			ViewUrl   string `json:"viewUrl"`
		} `json:"issueTrackers"`

		Labels            []string `json:"labels"`
		Name              string   `json:"name"`
		Popularity        int      `json:"popularity"`
		PreviousTimestamp string   `json:"previousTimestamp"`
		PreviousVersion   string   `json:"previousVersion"`
		ReleaseTimestamp  string   `json:"releaseTimestamp"`
		RequiredCore      string   `json:"requiredCore"`
		Scm               string   `json:"scm"`
		Sha1              string   `json:"sha1"`
		Sha256            string   `json:"sha256"`
		Size              int      `json:"size"`
		Title             string   `json:"title"`
		Url               string   `json:"url"`
		Version           string   `json:"version"`
		// MinimumJavaVersion string `json:"minimumJavaVersion"`
		Wiki string `json:"wiki"`
	} `json:"plugins"`

	Signature struct {
		Certificates        []string `json:"certificates"`
		CorrectDigest       string   `json:"correct_digest"`
		CorrectDigest512    string   `json:"correct_digest512"`
		CorrectSignature    string   `json:"correct_signature"`
		CorrectSignature512 string   `json:"correct_signature512"`
	} `json:"signature"`

	UpdateCenterVersion string `json:"updateCenterVersion"`

	Warnings []struct {
		Id       string `json:"id"`
		Message  string `json:"message"`
		Name     string `json:"name"`
		Type     string `json:"type"`
		Url      string `json:"url"`
		Versions []struct {
			LastVersion string `json:"lastVersion"`
			Pattern     string `json:"pattern"`
		} `json:"versions"`
	} `json:"warnings"`
}

// https://updates.jenkins.io/stable/update-center.actual.json

// TODO: make this configurable
// Version of jenkins update center repo
// https://github.com/jenkins-infra/update-center2/blob/master/site/LAYOUT.md#regular-tiered-update-sites-lts-and-weekly

func NewUpdateCenter() UpdateCenter {
	return UpdateCenter{}
}

// Returns *UpdateCenter, error type with data
// you can pass empty string to get latest core version package or specific version of jenkins core
func (uc UpdateCenter) Get(coreVersion string) (*UpdateCenter, error) {
	logrus.Debugln("loading updateCenter")
	var urlParam string
	// get update center for specific jenkins core, should be arg for URL
	if coreVersion == "" {
		logrus.Warnln("[WARN] You didn't pass '--core'")
	}
	// would be good to have param, but all tools requires version only
	// like give me update center for "1.255.3" without stable prefix or so
	// stable here for workaround for now, mb will be changed in future
	urlParam = "?version=" + coreVersion

	url := JENKINS_UPDATE_CENTER_URL + "/" + "stable" + "/" + "update-center.actual.json"

	content, err := DoRequest(url + urlParam)
	if err != nil {
		return nil, err
	}

	var updateCenter UpdateCenter
	if err := json.Unmarshal(content, &updateCenter); err != nil {
		logrus.Errorln(err)
		logrus.Errorln("Can not unmarshal JSON")
		return nil, err
	}
	return &updateCenter, nil
}
