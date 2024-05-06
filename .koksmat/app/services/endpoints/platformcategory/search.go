/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package platformcategory

import (
    "log"

    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/platformcategorymodel"
    . "github.com/magicbutton/magic-files/utils"
)

func PlatformCategorySearch(query string) (*Page[platformcategorymodel.PlatformCategory], error) {
    log.Println("Calling PlatformCategorysearch")

    return applogic.Search[database.PlatformCategory, platformcategorymodel.PlatformCategory]("name", query, applogic.MapPlatformCategoryOutgoing)

}
