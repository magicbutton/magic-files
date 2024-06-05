/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package batch

import (
    "log"
    "strconv"
    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/batchmodel"

)

func BatchRead(arg0 string) (*batchmodel.Batch, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling Batchread")

    return applogic.Read[database.Batch, batchmodel.Batch](id, applogic.MapBatchOutgoing)

}
