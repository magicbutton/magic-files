/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package sql

import (
    "log"
    "strconv"
    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/sqlmodel"

)

func SQLRead(arg0 string) (*sqlmodel.SQL, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling SQLread")

    return applogic.Read[database.SQL, sqlmodel.SQL](id, applogic.MapSQLOutgoing)

}
