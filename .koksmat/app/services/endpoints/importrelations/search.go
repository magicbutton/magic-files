/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package importrelations

import (
    "log"

    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/importrelationsmodel"
    . "github.com/magicbutton/magic-files/utils"
)

func ImportRelationsSearch(query string) (*Page[importrelationsmodel.ImportRelations], error) {
    log.Println("Calling ImportRelationssearch")

    return applogic.Search[database.ImportRelations, importrelationsmodel.ImportRelations]("searchindex", query, applogic.MapImportRelationsOutgoing)

}
