/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package importgroups

import (
    "log"

    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/importgroupsmodel"

)

func ImportGroupsUpdate(item importgroupsmodel.ImportGroups) (*importgroupsmodel.ImportGroups, error) {
    log.Println("Calling ImportGroupsupdate")

    return applogic.Update[database.ImportGroups, importgroupsmodel.ImportGroups](item.ID,item, applogic.MapImportGroupsIncoming, applogic.MapImportGroupsOutgoing)

}
