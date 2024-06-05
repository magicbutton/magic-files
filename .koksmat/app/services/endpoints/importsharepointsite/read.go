/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package importsharepointsite

import (
    "log"
    "strconv"
    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/importsharepointsitemodel"

)

func ImportSharePointSiteRead(arg0 string) (*importsharepointsitemodel.ImportSharePointSite, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling ImportSharePointSiteread")

    return applogic.Read[database.ImportSharePointSite, importsharepointsitemodel.ImportSharePointSite](id, applogic.MapImportSharePointSiteOutgoing)

}
