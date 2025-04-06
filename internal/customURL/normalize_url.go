package customURL

import (
	"errors"
	"net/url"
	"strings"
)

func NormalizeURL(rawUrl string) (string, error) {
	parsedURL, err := url.Parse(rawUrl)
	if err != nil {
		return "", errors.New("error in parsing the raw url")
	}

	path, _ := strings.CutSuffix(parsedURL.Path, "/")
	return strings.ToLower(parsedURL.Host + path), nil
}
