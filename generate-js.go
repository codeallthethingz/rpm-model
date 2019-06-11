package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/codeallthethingz/rpm-model/model"
)

func main() {
	units := model.PredefinedUnits
	createDirIfNotExist("model-js")
	jsOutput := "exports.predefinedUnits = ["
	for i, unit := range units {
		jsOutput += "\n  '" + string(unit) + "'"
		if i < len(units)-1 {
			jsOutput += ","
		}
	}
	jsOutput += "\n]\n"

	boundingTypes := model.PredefinedBoundingTypes
	jsOutput += "\nexports.predefinedBoundingTypes = {"
	for i, boundingType := range boundingTypes {
		jsOutput += "\n  '" + string(boundingType.Name) + "'"
		boundingTypeJSON, err := json.Marshal(boundingType)
		if err != nil {
			panic(err)
		}
		jsOutput += ": " + string(boundingTypeJSON)
		if i < len(boundingTypes)-1 {
			jsOutput += ","
		}
	}
	jsOutput += "\n}\n"
	fmt.Println(jsOutput)
	ioutil.WriteFile("model-js/package-json.js", []byte(jsOutput), 0644)
}

func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}
