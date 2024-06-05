/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package platform

import (
    "log"

    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/platformmodel"

)

func PlatformUpdate(item platformmodel.Platform) (*platformmodel.Platform, error) {
    log.Println("Calling Platformupdate")

    return applogic.Update[database.Platform, platformmodel.Platform](item.ID,item, applogic.MapPlatformIncoming, applogic.MapPlatformOutgoing)

}
