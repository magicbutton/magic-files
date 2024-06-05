/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package batch

import (
    "log"

    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/batchmodel"

)

func BatchUpdate(item batchmodel.Batch) (*batchmodel.Batch, error) {
    log.Println("Calling Batchupdate")

    return applogic.Update[database.Batch, batchmodel.Batch](item.ID,item, applogic.MapBatchIncoming, applogic.MapBatchOutgoing)

}
