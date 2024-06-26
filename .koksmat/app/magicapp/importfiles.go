package magicapp

import (
	"fmt"
	"os"

	"github.com/magicbutton/magic-files/services/endpoints/importdata"
	"github.com/magicbutton/magic-files/services/models/importdatamodel"
)

func ImportFiles(filepath string) (*importdatamodel.ImportData, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {

		return nil, fmt.Errorf("could not read file: %w", err)
	}

	OpenDatabase()

	importRecord := importdatamodel.ImportData{
		Name:        filepath,
		Description: "Imported file",

		Data: data,
	}
	return importdata.ImportDataCreate(importRecord)

}
