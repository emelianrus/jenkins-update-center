package request

import (
	"io"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

func GetFileName(url string) string {
	return strings.Split(path.Base(url), "?")[0]
}

// request
func Do(url string) ([]byte, error) {
	fileName := GetFileName(url)
	// Create the file
	file, err := os.Create(fileName)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	defer file.Close()

	var myClient = &http.Client{Timeout: 10 * time.Second}

	response, err := myClient.Get(url)
	logrus.Infof("Downloading: %s \n", url)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	defer response.Body.Close()

	// Check server response
	if response.StatusCode != http.StatusOK {
		logrus.Printf("bad status: %s \n", response.Status)
		return nil, err
	}
	_, err = io.Copy(file, response.Body)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}

	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		logrus.Println(err)
		return nil, err
	}

	return fileContent, nil
}
