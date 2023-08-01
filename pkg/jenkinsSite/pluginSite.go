package jenkinsSite

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

type PluginSiteReleases struct {
	Releases []struct {
		TagName     string `json:"tagName"`
		Name        string `json:"name"`
		PublishedAt string `json:"publishedAt"`
		HTMLURL     string `json:"htmlURL"`
		BodyHTML    string `json:"bodyHTML"`
	} `json:"releases"`
}

type ReleaseNote struct {
	Name      string
	Tag       string
	BodyHTML  string
	HTMLURL   string
	CreatedAt string
}

// Will download last 10 release notes in html string format for plugin
func (ps *JenkinsSite) DownloadReleaseNotes(projectName string) ([]ReleaseNote, error) {
	logrus.Infoln("[PluginSiteReleases][Download]")
	releaseNotes := []ReleaseNote{}

	url := fmt.Sprintf("https://plugin-site-issues.jenkins.io/api/plugin/%s/releases", projectName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logrus.Errorln("Error creating request:", err)
		return nil, nil
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Errorln("Error making request:", err)
		return nil, errors.New("error making request")
	}

	var releases PluginSiteReleases
	err = json.NewDecoder(resp.Body).Decode(&releases)
	if err != nil {
		logrus.Errorln("can not decode plugin site releases")
		return nil, nil
	}

	for _, release := range releases.Releases {
		releaseNotes = append(releaseNotes, ReleaseNote{
			Name:      release.Name,
			Tag:       release.TagName,
			BodyHTML:  release.BodyHTML,
			HTMLURL:   release.HTMLURL,
			CreatedAt: release.PublishedAt,
		})
	}

	return releaseNotes, nil
}
