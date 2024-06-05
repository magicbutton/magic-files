/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package importsharepointsite

import (
    "log"

    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/importsharepointsitemodel"
    . "github.com/magicbutton/magic-files/utils"
)

func ImportSharePointSiteSearch(query string) (*Page[importsharepointsitemodel.ImportSharePointSite], error) {
    log.Println("Calling ImportSharePointSitesearch")

    return applogic.Search[database.ImportSharePointSite, importsharepointsitemodel.ImportSharePointSite]("searchindex", query, applogic.MapImportSharePointSiteOutgoing)

}
