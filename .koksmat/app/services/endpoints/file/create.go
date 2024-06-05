/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package file

import (
    "log"
   
    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/filemodel"

)

func FileCreate(item filemodel.File) (*filemodel.File, error) {
    log.Println("Calling Filecreate")

    return applogic.Create[database.File, filemodel.File](item, applogic.MapFileIncoming, applogic.MapFileOutgoing)

}
