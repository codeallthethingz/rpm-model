package protocol

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseComponentID(t *testing.T) {
	owner, name, err := ParseComponentID("0codeall-thethingz_hex-nut-particle1")
	require.Nil(t, err)
	require.Equal(t, "0codeall-thethingz", owner)
	require.Equal(t, "hex-nut-particle1", name)

	badComponentID(t, "_")
	badComponentID(t, " _ ")
	badComponentID(t, "missing")
	badComponentID(t, "Case_Sensitive")
	badComponentID(t, "Case_sensitive")
	badComponentID(t, "case_Sensitive")
	badComponentID(t, "too_many_underscores")
	badComponentID(t, "~_~")
	badComponentID(t, "___")
	badComponentID(t, "-_-")
	badComponentID(t, "a_  ")
	badComponentID(t, "a_")
	badComponentID(t, " _a")
	badComponentID(t, "_a")
	badComponentID(t, "usernameiswaytooloooooooooooooooooooooooooooooooooooooooooooooooong_nut")
	badComponentID(t, "user_componentidiswaytooloooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooong")
	badComponentID(t, "usernameendswith-_nut")
	badComponentID(t, "-usernamestartswithhyphen_nut")
	badComponentID(t, "user_-componentidstartswithhyphen")
	badComponentID(t, "user_-componentidendswith-")
	badComponentID(t, "consecutive--hyphens_componentid")
	badComponentID(t, "user_consecutive--hyphens")
}

func TestParseVersion(t *testing.T) {
	major, minor, patch, trailing, err := ParseVersion("1.2.13")
	require.Nil(t, err)
	require.Equal(t, 1, major)
	require.Equal(t, 2, minor)
	require.Equal(t, 13, patch)
	require.Equal(t, "", trailing)

	major, minor, patch, trailing, _ = ParseVersion("3.2.1-SNAPSHOT")
	require.Equal(t, 3, major)
	require.Equal(t, 2, minor)
	require.Equal(t, 1, patch)
	require.Equal(t, "SNAPSHOT", trailing)

	badVersion(t, "3..2.1-SNAPSHOT")
	badVersion(t, "3..2.1-SNAPSHOT")
	badVersion(t, "3.2..1-SNAPSHOT")
	badVersion(t, "3..2.1--SNAPSHOT")
	badVersion(t, ".2.1-SNAPSHOT")
	badVersion(t, "1..1-SNAPSHOT")
	badVersion(t, "1.2.-SNAPSHOT")
	badVersion(t, "1.2.9-snapshot")
	badVersion(t, "1.2.9-SNAP1")
	badVersion(t, "1.2.-10")
	badVersion(t, "0.0.0")
	badVersion(t, "AOEU")
	badVersion(t, "1234")
	badVersion(t, "X.X.X")
	badVersionSpecificError(t, "1341234123412341234123412341234123412341234.1.6", "value out of range")
	badVersion(t, "3.2.1-SNAPSHOT-aoeu")
	badVersion(t, "3.2.1--")
	badVersion(t, "3.2.1-")
	badVersion(t, "3.2.1- ")
	badVersion(t, " ")
	badVersion(t, "")
}

func badVersionSpecificError(t *testing.T, version string, msg string) {
	_, _, _, _, err := ParseVersion(version)
	require.NotNil(t, err)
	require.Contains(t, err.Error(), msg)
}
func badVersion(t *testing.T, version string) {
	badVersionSpecificError(t, version, "version number must match semantic versioning: https://semver.com")
}
func badComponentID(t *testing.T, componentID string) {
	_, _, err := ParseComponentID(componentID)
	require.NotNil(t, err)
	require.Contains(t, err.Error(), "component ID must be of the form")
}
