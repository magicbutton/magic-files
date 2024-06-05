/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package importgroups

import (
    "log"
   
    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/importgroupsmodel"

)

func ImportGroupsCreate(item importgroupsmodel.ImportGroups) (*importgroupsmodel.ImportGroups, error) {
    log.Println("Calling ImportGroupscreate")

    return applogic.Create[database.ImportGroups, importgroupsmodel.ImportGroups](item, applogic.MapImportGroupsIncoming, applogic.MapImportGroupsOutgoing)

}
