/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package batch

import (
    "log"

    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/batchmodel"
    . "github.com/magicbutton/magic-files/utils"
)

func BatchSearch(query string) (*Page[batchmodel.Batch], error) {
    log.Println("Calling Batchsearch")

    return applogic.Search[database.Batch, batchmodel.Batch]("searchindex", query, applogic.MapBatchOutgoing)

}
