/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package importgroups

import (
    "log"

    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/importgroupsmodel"
    . "github.com/magicbutton/magic-files/utils"
)

func ImportGroupsSearch(query string) (*Page[importgroupsmodel.ImportGroups], error) {
    log.Println("Calling ImportGroupssearch")

    return applogic.Search[database.ImportGroups, importgroupsmodel.ImportGroups]("searchindex", query, applogic.MapImportGroupsOutgoing)

}
