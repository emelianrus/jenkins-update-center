package jenkinsSite

import (
	"io"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// Do request with passed url and return content
func DoRequest(url string) ([]byte, error) {
	var client = &http.Client{Timeout: 10 * time.Second}

	response, err := client.Get(url)
	logrus.Infof("Downloading: %s \n", url)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		logrus.Errorf("bad status: %s \n", response.Status)
		return nil, err
	}

	content, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return content, nil
}
