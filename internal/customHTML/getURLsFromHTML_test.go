package customHTML

import (
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func TestURLsFromHTML(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		inputHTML     string
		expected      []string
		errorContains string
	}{
		{
			name:     "absolute url",
			inputURL: "https://blog.boot.dev",
			inputHTML: `<html>
							<body>
								<a href="https://blog.boot.dev">
									<span>Boot.dev</span>
								</a>
							</body>
						</html>`,
			expected:      []string{"https://blog.boot.dev"},
			errorContains: "",
		},
		{
			name:     "relative url",
			inputURL: "https://blog.boot.dev",
			inputHTML: `<html>
							<body>
								<a href="/path/one">
									<span>Boot.dev</span>
								</a>
							</body>
						</html>`,
			expected:      []string{"https://blog.boot.dev/path/one"},
			errorContains: "",
		},
		{
			name:     "relative and absolute url",
			inputURL: "https://blog.boot.dev",
			inputHTML: `<html>
							<body>
								<a href="https://example.com/one">
									<span>Boot.dev</span>
								</a>
								<a href="/path/one">
									<span>Boot.dev</span>
								</a>
							</body>
						</html>`,
			expected:      []string{"https://example.com/one", "https://blog.boot.dev/path/one"},
			errorContains: "",
		},
		{
			name:     "invalid base url",
			inputURL: ":\\invalidURL",
			inputHTML: `<html>
							<body>
								<a href="https://example.com/one">
									<span>Boot.dev</span>
								</a>
								<a href="/path/one">
									<span>Boot.dev</span>
								</a>
							</body>
						</html>`,
			expected:      nil,
			errorContains: "error in parsing base url",
		},
		{
			name:     "invalid href url",
			inputURL: "https://blog.boot.dev",
			inputHTML: `<html>
							<body>
								<a href=":\\invalidURL">
									<span>Boot.dev</span>
								</a>
								<a href="/path/one">
									<span>Boot.dev</span>
								</a>
							</body>
						</html>`,
			expected:      []string{"https://blog.boot.dev/path/one"},
			errorContains: "",
		},
		{
			name:     "nested anchor tags",
			inputURL: "https://blog.boot.dev",
			inputHTML: `<html>
							<body>
								<a href="/path">
									<span>Boot.dev</span>
								</a>
								<a href="/path/one">
										<span>Boot.dev</span>
								</a>
								<div>
									<a href="/path/two">
										<span>Boot.dev</span>
									</a>
								</div>
							</body>
						</html>`,
			expected:      []string{"https://blog.boot.dev/path", "https://blog.boot.dev/path/one", "https://blog.boot.dev/path/two"},
			errorContains: "",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			baseURL, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: couldn't parse input URL: %v", i, tc.name, err)
				return
			}
			actual, err := GetURLsFromHTML(tc.inputHTML, baseURL)
			if err != nil && !strings.Contains(err.Error(), tc.errorContains) {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if err != nil && tc.errorContains == "" {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if err == nil && tc.errorContains != "" {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
				return
			}
		})
	}
}
