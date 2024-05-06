/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: true
---
*/
//generator:  noma3
package importdata

// noma2
import (
	"log"

	"github.com/magicbutton/magic-files/applogic"
	"github.com/magicbutton/magic-files/database"
	"github.com/magicbutton/magic-files/services/models/importdatamodel"
)

func ImportDataCreate(item importdatamodel.ImportData) (*importdatamodel.ImportData, error) {
	log.Println("Calling ImportDataCreate")

	return applogic.Create[database.ImportData, importdatamodel.ImportData](item, applogic.MapImportDataIncoming, applogic.MapImportDataOutgoing)

}
