/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package importrelations

import (
    "log"

    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/importrelationsmodel"

)

func ImportRelationsUpdate(item importrelationsmodel.ImportRelations) (*importrelationsmodel.ImportRelations, error) {
    log.Println("Calling ImportRelationsupdate")

    return applogic.Update[database.ImportRelations, importrelationsmodel.ImportRelations](item.ID,item, applogic.MapImportRelationsIncoming, applogic.MapImportRelationsOutgoing)

}
