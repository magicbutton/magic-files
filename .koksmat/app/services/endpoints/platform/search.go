/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package platform

import (
    "log"

    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/platformmodel"
    . "github.com/magicbutton/magic-files/utils"
)

func PlatformSearch(query string) (*Page[platformmodel.Platform], error) {
    log.Println("Calling Platformsearch")

    return applogic.Search[database.Platform, platformmodel.Platform]("name", query, applogic.MapPlatformOutgoing)

}
