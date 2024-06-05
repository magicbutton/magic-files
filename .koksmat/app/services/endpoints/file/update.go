/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package file

import (
    "log"

    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/filemodel"

)

func FileUpdate(item filemodel.File) (*filemodel.File, error) {
    log.Println("Calling Fileupdate")

    return applogic.Update[database.File, filemodel.File](item.ID,item, applogic.MapFileIncoming, applogic.MapFileOutgoing)

}
