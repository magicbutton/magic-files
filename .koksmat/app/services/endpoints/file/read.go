/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package file

import (
    "log"
    "strconv"
    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/filemodel"

)

func FileRead(arg0 string) (*filemodel.File, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling Fileread")

    return applogic.Read[database.File, filemodel.File](id, applogic.MapFileOutgoing)

}
