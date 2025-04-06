package customHTML

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func GetHTML(rawURL string) (string, error) {
	baseURL, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("error parsing the raw url: %v", err)
	}

	res, err := http.Get(baseURL.String())
	if err != nil {
		return "", fmt.Errorf("error in fetching the html: %v", err)
	}
	if res.StatusCode >= 400 {
		return "", fmt.Errorf("%v status code for %v", res.StatusCode, err)
	}
	if !strings.Contains(res.Header["Content-Type"][0], "text/html") {
		return "", fmt.Errorf("text/html content-type type required, fetched content-type: %v", res.Header["Content-Type"][0])
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	htmlString := string(body)

	return htmlString, nil
}
