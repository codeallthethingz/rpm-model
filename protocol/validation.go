package protocol

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/codeallthethingz/rpm-model/model"
)

// ParseVersion extracts and validates version number.
func ParseVersion(version string) (int, int, int, string, error) {
	if version == "latest" {
		return 0, 0, 0, "", nil
	}
	badVersion := fmt.Errorf("version number must match X.X.X-AAAAA")
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
func ParseComponentID(componentID string) (string, string, string, error) {
	badFormat := fmt.Errorf("component ID must be of the form <username>_<component-name> or  <username>_<component-name>@version")
	if strings.TrimSpace(componentID) == "" {
		return "", "", "", badFormat
	}
	lastIndex := strings.LastIndex(componentID, "_")
	if lastIndex == -1 || lastIndex == 0 || lastIndex == len(componentID) {
		return "", "", "", badFormat
	}
	owner := componentID[0:lastIndex]
	name := componentID[lastIndex+1:]
	indexOfAt := strings.LastIndex(componentID, "@")
	if indexOfAt != -1 {
		name = componentID[lastIndex+1 : indexOfAt]
	}
	if strings.TrimSpace(owner) == "" || strings.TrimSpace(name) == "" {
		return "", "", "", badFormat
	}
	testOwner := `^[a-z0-9\-]{3,38}$`
	testName := `^[a-z0-9\-]{3,100}$`
	version := ""
	if indexOfAt != -1 {
		version = componentID[indexOfAt+1:]
		_, _, _, _, err := ParseVersion(version)
		if err != nil {
			return "", "", "", badFormat
		}
	}
	if ok, err := regexp.Match(testOwner, []byte(owner)); !ok || err != nil {
		return "", "", "", badFormat
	}
	if ok, err := regexp.Match(testName, []byte(name)); !ok || err != nil {
		return "", "", "", badFormat
	}
	if string(owner[0]) == "-" || string(owner[len(owner)-1]) == "-" || strings.Contains(owner, "--") {
		return "", "", "", badFormat
	}
	if string(name[0]) == "-" || string(name[len(name)-1]) == "-" || strings.Contains(name, "--") {
		return "", "", "", badFormat
	}

	return owner, name, version, nil
}

// ValidateConstructJSON makes sure the construct json is all looking good
func ValidateConstructJSON(constructJSON *model.ConstructJSON) *UserFacingMessage {
	var missingFields []string
	errors := ""
	if isBlank(constructJSON.Name) {
		missingFields = append(missingFields, "name")
	}
	if isBlank(constructJSON.Version) {
		missingFields = append(missingFields, "version")
	}
	if isBlank(constructJSON.License) {
		missingFields = append(missingFields, "license")
	}
	if isBlank(string(constructJSON.Units)) {
		missingFields = append(missingFields, "units")
	}
	if constructJSON.Bounds == nil || len(constructJSON.Bounds) == 0 {
		missingFields = append(missingFields, "bounds")
	} else {
		errors = validateFirstBounds(constructJSON)
	}
	if len(missingFields) > 0 {
		errors = "construct.json must include "
		errors += makeCommaAndString(missingFields)
	}

	if !isBlank(errors) {
		return &UserFacingMessage{
			Message:    errors,
			StatusCode: 400,
		}
	}
	return nil
}

func isBlank(element string) bool {
	if strings.TrimSpace(element) == "" {
		return true
	}
	return false
}

func validateFirstBounds(constructJSON *model.ConstructJSON) string {
	firstBound := constructJSON.Bounds[0]
	if firstBound.Name != "total-area" {
		return "first bound must be called total-area"
	}
	allowed := ""
	for _, bt := range model.PredefinedBoundingTypes {
		allowed += bt.Name + ", "
		if bt.Name == firstBound.BoundingType.Name {
			return validateMeasurements(bt, firstBound.BoundingType)
		}
	}
	return "bounding type name \"" + firstBound.BoundingType.Name + "\" is not in the allowable list for total-area.  Valid options are " + allowed
}

func validateMeasurements(template model.BoundingType, pjBoundingType model.BoundingType) string {
	errors := ""
	missingFields := []string{}
	badFields := []string{}
	for k := range template.Measurements {
		value, ok := pjBoundingType.Measurements[k]
		if !ok {
			missingFields = append(missingFields, k)
		} else {
			if isBlank(value) {
				missingFields = append(missingFields, k)
			} else {
				_, err := strconv.ParseFloat(value, 64)
				if err != nil {
					badFields = append(badFields, k)
				}
			}
		}
	}
	sort.Strings(missingFields)
	sort.Strings(badFields)
	if len(missingFields) > 0 {
		errors = "bounding type \"" + template.Name + "\" is missing metrics "
		errors += makeCommaAndString(missingFields)
	}
	if len(badFields) > 0 {
		if !isBlank(errors) {
			errors += " also "
		}
		errors += makeCommaAndString(badFields) + " must be decimal"
		if len(badFields) > 1 {
			errors += "s"
		}
	}
	return errors
}

func makeCommaAndString(items []string) string {
	errors := ""
	for i, field := range items {
		if i == len(items)-1 && len(items) > 1 {
			errors += " and "
		}
		errors += field
		if i < len(items)-2 {
			errors += ", "
		}
	}
	return errors
}
