package reqbuilder

import (
	"io"
	"net/http"
	"strings"

	"github.com/bisonai/huobi_futures/sdk/log"
)

func HttpGet(url string) (string, error) {
	logger := log.GetInstance()
	logger.Start()

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	logger.StopAndLog("GET", url)

	return string(result), err
}

func HttpPost(url string, body string) (string, error) {
	logger := log.GetInstance()
	logger.Start()

	resp, err := http.Post(url, "application/json", strings.NewReader(body))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	logger.StopAndLog("POST", url)

	return string(result), err
}
