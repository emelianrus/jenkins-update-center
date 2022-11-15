package updateCenter

// TODO: its not cmd package
import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	JENKINS_PLUGINS_URL = "https://updates.jenkins.io"
	JSON_FILE_NAME      = "update-center.actual.json"

	UPDATE_CENTER_JSON_LOCATION = JENKINS_PLUGINS_URL + "/" + JSON_FILE_NAME
)

const CACHED_FILE_NAME = "update-center.json"

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

// you can pass empty string to get latest core version package
func Get(coreVersion string) *UpdateCenter {

	if coreVersion == "" {
		logrus.Warnln("[WARN] You didn't pass '--core'. Will use LTS core version")
	} else {
		coreVersion = "?version=stable-" + coreVersion
	}

	// Create the file
	file, err := os.Create(CACHED_FILE_NAME)
	if err != nil {
		logrus.Errorln(err)
		return nil
	}
	defer file.Close()

	URL := UPDATE_CENTER_JSON_LOCATION + coreVersion
	var myClient = &http.Client{Timeout: 10 * time.Second}

	response, err := myClient.Get(URL)
	logrus.Infof("Downloading: %s \n", URL)
	if err != nil {
		logrus.Errorln(err)
		return nil
	}
	defer response.Body.Close()

	// Check server response
	if response.StatusCode != http.StatusOK {
		logrus.Printf("bad status: %s \n", response.Status)
		return nil
	}
	_, err = io.Copy(file, response.Body)
	if err != nil {
		logrus.Errorln(err)
		return nil
	}

	fileContent, err := ioutil.ReadFile(CACHED_FILE_NAME)
	if err != nil {
		logrus.Println(err)
	}

	var updateCenter UpdateCenter
	if err := json.Unmarshal(fileContent, &updateCenter); err != nil {
		logrus.Errorln(err)
		logrus.Errorln("Can not unmarshal JSON")
	}
	return &updateCenter
}
