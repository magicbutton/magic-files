/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package job

import (
    "log"
    "strconv"
    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/jobmodel"

)

func JobRead(arg0 string) (*jobmodel.Job, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling Jobread")

    return applogic.Read[database.Job, jobmodel.Job](id, applogic.MapJobOutgoing)

}
