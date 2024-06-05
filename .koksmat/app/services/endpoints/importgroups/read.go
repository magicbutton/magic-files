/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package importgroups

import (
    "log"
    "strconv"
    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/importgroupsmodel"

)

func ImportGroupsRead(arg0 string) (*importgroupsmodel.ImportGroups, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling ImportGroupsread")

    return applogic.Read[database.ImportGroups, importgroupsmodel.ImportGroups](id, applogic.MapImportGroupsOutgoing)

}
