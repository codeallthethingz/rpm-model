package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/codeallthethingz/rpm-model/model"
	"github.com/logrusorgru/aurora"
)

func main() {
	units := model.PredefinedUnits
	jsOutput := fmt.Sprintf("exports.maxArchiveSizeBytes = %d\n", model.MaxArchiveSizeBytes)
	jsOutput += fmt.Sprintf("exports.modelVersion = '%s'\n", model.ModelVersion)
	jsOutput += "exports.predefinedUnits = ["
	for i, unit := range units {
		jsOutput += "\n  '" + string(unit) + "'"
		if i < len(units)-1 {
			jsOutput += ","
		}
	}
	jsOutput += "\n]\n"

	boundingTypes := model.PredefinedBoundingTypes
	jsOutput += "\nexports.predefinedBoundingTypes = ["
	for i, boundingType := range boundingTypes {
		boundingTypeJSON, err := json.Marshal(boundingType)
		if err != nil {
			panic(err)
		}
		jsOutput += "\n  " + string(boundingTypeJSON)
		if i < len(boundingTypes)-1 {
			jsOutput += ","
		}
	}
	jsOutput += "\n]\n"
	fmt.Printf("%s%s\n", aurora.White("generated"), aurora.Green(" model-js/package-json.js"))
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
