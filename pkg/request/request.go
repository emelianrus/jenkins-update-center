package request

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// Get file name from url, used to create cache file
func getFileNameFromUrl(url string) string {
	return strings.Split(path.Base(url), "?")[0]
}

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

// Do request with passed url and return content, and create/check cached file on file system
func DoRequestWithCache(url string, cacheFileName string) ([]byte, error) {
	var fileContent []byte

	// check is file exist as "cached"
	if _, err := os.Stat(cacheFileName); errors.Is(err, os.ErrNotExist) {
		logrus.Infoln("cache miss")

		fileContent, err = DoRequest(url)
		if err != nil {
			logrus.Errorln(err)
			return nil, err
		}
		// create file
		file, err := os.Create(cacheFileName)
		if err != nil {
			logrus.Errorln(err)
			return nil, err
		}
		// write content to file
		b, err := file.Write(fileContent)
		if err != nil {
			logrus.Errorln(err)
			return nil, err
		}
		logrus.Debugf("wrote %d bytes to %s\n", b, cacheFileName)
		defer file.Close()
	} else {
		logrus.Infoln("cache hit")
		fileContent, err = os.ReadFile(cacheFileName)
		if err != nil {
			logrus.Errorln(err)
			return nil, err
		}
	}
	return fileContent, nil
}
