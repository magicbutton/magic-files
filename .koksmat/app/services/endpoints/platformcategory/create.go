/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package platformcategory

import (
    "log"
   
    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/platformcategorymodel"

)

func PlatformCategoryCreate(item platformcategorymodel.PlatformCategory) (*platformcategorymodel.PlatformCategory, error) {
    log.Println("Calling PlatformCategorycreate")

    return applogic.Create[database.PlatformCategory, platformcategorymodel.PlatformCategory](item, applogic.MapPlatformCategoryIncoming, applogic.MapPlatformCategoryOutgoing)

}
