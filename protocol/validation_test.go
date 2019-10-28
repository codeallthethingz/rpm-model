package protocol

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/codeallthethingz/rpm-model/model"
	"github.com/stretchr/testify/require"
)

func TestBadConstructJSON(t *testing.T) {
	err := ValidateConstructJSON(&model.ConstructJSON{})
	require.NotNil(t, err)
	require.Equal(t, "construct.json must include name, version, license, units and bounds", err.Error())
	require.NotNil(t, ValidateConstructJSON(&model.ConstructJSON{Bounds: []model.Bound{model.Bound{}}}))
	require.NotNil(t, ValidateConstructJSON(&model.ConstructJSON{Bounds: []model.Bound{model.Bound{Name: "total-area"}}}))
	require.NotNil(t, ValidateConstructJSON(&model.ConstructJSON{Bounds: []model.Bound{model.Bound{Name: "total-area", BoundingType: model.BoundingType{Name: "cuboid", Measurements: map[string]string{"height": ""}}}}}))
	require.NotNil(t, ValidateConstructJSON(&model.ConstructJSON{Bounds: []model.Bound{model.Bound{Name: "total-area", BoundingType: model.BoundingType{Name: "cuboid", Measurements: map[string]string{"height": "", "width": "N", "length": "aoeu"}}}}}))
}
func TestGoodConstructJSON(t *testing.T) {
	contents, err := ioutil.ReadFile("../model/examples/user_bolt-round-20mm/construct.json")
	require.Nil(t, err)
	p := &model.ConstructJSON{}
	require.Nil(t, json.Unmarshal(contents, p))
	require.Nil(t, ValidateConstructJSON(p))
}
func TestParseComponentID(t *testing.T) {
	owner, name, version, err := ParseComponentID("0codeall-thethingz_hex-nut-particle1")
	require.Nil(t, err)
	require.Equal(t, "", version)
	require.Equal(t, "0codeall-thethingz", owner)
	require.Equal(t, "hex-nut-particle1", name)

	owner, name, version, _ = ParseComponentID("0codeall-thethingz_hex-nut-particle1@1.2.3")
	require.Equal(t, "1.2.3", version)
	require.Equal(t, "0codeall-thethingz", owner)
	require.Equal(t, "hex-nut-particle1", name)

	owner, name, version, _ = ParseComponentID("0codeall-thethingz_hex-nut-particle1@latest")
	require.Equal(t, "latest", version)
	require.Equal(t, "0codeall-thethingz", owner)
	require.Equal(t, "hex-nut-particle1", name)

	badComponentID(t, "0codeall-thethingz_hex-nut-particle1@l")
	badComponentID(t, "0codeall-thethingz_hex-nut-particle1@1.2")
	badComponentID(t, "0codeall-thethingz_hex-nut-particle1@@1.2.3")
	badComponentID(t, "0codeall-thethingz_hex-nut-particle1@1.2.2.3")
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
	badComponentID(t, " ")
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
	badVersion(t, "1.X.3")
	badVersion(t, "1.2.X")
	badVersion(t, "0.0.0")
	badVersion(t, "AOEU")
	badVersion(t, "1234")
	badVersion(t, "X.X.X")
	badVersionSpecificError(t, "1341234123412341234123412341234123412341234.1.6", "value out of range")
	badVersionSpecificError(t, "1.1341234123412341234123412341234123412341234.6", "value out of range")
	badVersionSpecificError(t, "1.2.1341234123412341234123412341234123412341234", "value out of range")
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
	badVersionSpecificError(t, version, "version number must match X.X.X-AAAAA")
}
func badComponentID(t *testing.T, componentID string) {
	_, _, _, err := ParseComponentID(componentID)
	require.NotNil(t, err)
	require.Contains(t, err.Error(), "component ID must be of the form")
}

func filled() *model.ConstructJSON {
	return &model.ConstructJSON{
		Name:    "test thing",
		Version: "0.0.1",
		License: "MIT",
		Units:   "mm",
	}
}

func TestValidateConstructJSON(t *testing.T) {
	err := ValidateConstructJSON(&model.ConstructJSON{
		Version: "0.0.1",
	})
	require.Contains(t, err.Error(), "construct.json must include name, ")
}

func TestValidateConstructBadBounds(t *testing.T) {
	pj := filled()
	pj.Bounds = []model.Bound{
		model.Bound{
			Name: "total-aquarium",
		},
	}
	err := ValidateConstructJSON(pj)
	require.Contains(t, err.Error(), "first bound must be called")
}

func TestValidateConstructBadBoundsBadBoundingType(t *testing.T) {
	pj := filled()
	pj.Bounds = []model.Bound{
		model.Bound{
			Name: "total-area",
			BoundingType: model.BoundingType{
				Name: "cube",
			}}}
	err := ValidateConstructJSON(pj)
	require.Error(t, err)
	require.Contains(t, err.Error(), "\"cube\" is not in the allowable list")
}

func TestValidateConstructBadBoundsMissingMetric(t *testing.T) {
	pj := filled()
	pj.Bounds = []model.Bound{
		model.Bound{
			Name: "total-area",
			BoundingType: model.BoundingType{
				Name: "cuboid",
			}}}
	err := ValidateConstructJSON(pj)
	require.NotNil(t, err)
	require.Contains(t, err.Error(), "missing metrics")
}

func TestValidateConstructBadBoundsBadMetric(t *testing.T) {
	pj := filled()
	pj.Bounds = []model.Bound{
		model.Bound{
			Name: "total-area",
			BoundingType: model.BoundingType{
				Name: "cuboid",
				Measurements: map[string]string{
					"width":  "k",
					"height": "",
					"length": "5h",
				},
			}}}
	err := ValidateConstructJSON(pj)
	require.NotNil(t, err)
	require.Contains(t, err.Error(), "missing metrics")
	require.Contains(t, err.Error(), "length and width must be decimals")
}
