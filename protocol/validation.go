package protocol

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
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

// ParseComponentID extracts and validates a component ID.
func ParseComponentID(componentID string) (string, string, error) {
	badFormat := fmt.Errorf("component ID must be of the form <username>_<component-name>")
	if strings.TrimSpace(componentID) == "" {
		return "", "", badFormat
	}
	lastIndex := strings.LastIndex(componentID, "_")
	if lastIndex == -1 || lastIndex == 0 || lastIndex == len(componentID) {
		return "", "", badFormat
	}
	owner := componentID[0:lastIndex]
	name := componentID[lastIndex+1:]
	if strings.TrimSpace(owner) == "" || strings.TrimSpace(name) == "" {
		return "", "", badFormat
	}
	testOwner := `^[a-z0-9\-]{3,38}$`
	testName := `^[a-z0-9\-]{3,100}$`

	if ok, err := regexp.Match(testOwner, []byte(owner)); !ok || err != nil {
		return "", "", badFormat
	}
	if ok, err := regexp.Match(testName, []byte(name)); !ok || err != nil {
		return "", "", badFormat
	}
	if string(owner[0]) == "-" || string(owner[len(owner)-1]) == "-" || strings.Contains(owner, "--") {
		return "", "", badFormat
	}
	if string(name[0]) == "-" || string(name[len(name)-1]) == "-" || strings.Contains(name, "--") {
		return "", "", badFormat
	}
	return owner, name, nil
}
