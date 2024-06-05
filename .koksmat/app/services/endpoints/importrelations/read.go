/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package importrelations

import (
    "log"
    "strconv"
    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/importrelationsmodel"

)

func ImportRelationsRead(arg0 string) (*importrelationsmodel.ImportRelations, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling ImportRelationsread")

    return applogic.Read[database.ImportRelations, importrelationsmodel.ImportRelations](id, applogic.MapImportRelationsOutgoing)

}
