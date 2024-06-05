/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package platform

import (
    "log"
   
    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/platformmodel"

)

func PlatformCreate(item platformmodel.Platform) (*platformmodel.Platform, error) {
    log.Println("Calling Platformcreate")

    return applogic.Create[database.Platform, platformmodel.Platform](item, applogic.MapPlatformIncoming, applogic.MapPlatformOutgoing)

}
