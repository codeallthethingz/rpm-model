package protocol

import (
	"fmt"
	"regexp"
	"strconv"
)

// ParseVersion extracts and validates version number.
func ParseVersion(version string) (int, int, int, string, error) {
	badVersion := fmt.Errorf("version number must match semantic versioning: https://semver.com")
	test := `([0-9]+)\.([0-9]+)\.([0-9]+)(.*)`
	testTrailing := `^[A-Z]+$`
	re := regexp.MustCompile(test)
	if ok, err := regexp.Match(test, []byte(version)); !ok || err != nil {
		return 0, 0, 0, "", badVersion
	}
	matches := re.FindSubmatch([]byte(version))
	major, err := strconv.Atoi(string(matches[1]))
	if err != nil {
		return 0, 0, 0, "", err
	}
	minor, err := strconv.Atoi(string(matches[2]))
	if err != nil {
		return 0, 0, 0, "", err
	}
	patch, err := strconv.Atoi(string(matches[3]))
	if err != nil {
		return 0, 0, 0, "", err
	}
	trailing := ""
	if string(matches[4]) != "" {
		trailing = string(matches[4])[1:]
		if ok, err := regexp.Match(testTrailing, []byte(trailing)); !ok || err != nil {
			return 0, 0, 0, "", badVersion
		}
	}
	if major == 0 && minor == 0 && patch == 0 {
		return 0, 0, 0, "", badVersion
	}
	return major, minor, patch, trailing, nil
}
