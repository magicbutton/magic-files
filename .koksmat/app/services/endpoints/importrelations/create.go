/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package importrelations

import (
    "log"
   
    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/importrelationsmodel"

)

func ImportRelationsCreate(item importrelationsmodel.ImportRelations) (*importrelationsmodel.ImportRelations, error) {
    log.Println("Calling ImportRelationscreate")

    return applogic.Create[database.ImportRelations, importrelationsmodel.ImportRelations](item, applogic.MapImportRelationsIncoming, applogic.MapImportRelationsOutgoing)

}
