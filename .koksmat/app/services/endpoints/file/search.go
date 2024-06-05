/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package file

import (
    "log"

    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/filemodel"
    . "github.com/magicbutton/magic-files/utils"
)

func FileSearch(query string) (*Page[filemodel.File], error) {
    log.Println("Calling Filesearch")

    return applogic.Search[database.File, filemodel.File]("searchindex", query, applogic.MapFileOutgoing)

}
