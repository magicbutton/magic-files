/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package platformcategory

import (
    "log"
    "strconv"
    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/platformcategorymodel"

)

func PlatformCategoryRead(arg0 string) (*platformcategorymodel.PlatformCategory, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling PlatformCategoryread")

    return applogic.Read[database.PlatformCategory, platformcategorymodel.PlatformCategory](id, applogic.MapPlatformCategoryOutgoing)

}
