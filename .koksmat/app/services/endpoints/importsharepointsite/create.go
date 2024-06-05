/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package importsharepointsite

import (
    "log"
   
    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/importsharepointsitemodel"

)

func ImportSharePointSiteCreate(item importsharepointsitemodel.ImportSharePointSite) (*importsharepointsitemodel.ImportSharePointSite, error) {
    log.Println("Calling ImportSharePointSitecreate")

    return applogic.Create[database.ImportSharePointSite, importsharepointsitemodel.ImportSharePointSite](item, applogic.MapImportSharePointSiteIncoming, applogic.MapImportSharePointSiteOutgoing)

}
