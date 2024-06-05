/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package importsharedmailbox

import (
    "log"
    "strconv"
    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/importsharedmailboxmodel"

)

func ImportSharedMailboxRead(arg0 string) (*importsharedmailboxmodel.ImportSharedMailbox, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling ImportSharedMailboxread")

    return applogic.Read[database.ImportSharedMailbox, importsharedmailboxmodel.ImportSharedMailbox](id, applogic.MapImportSharedMailboxOutgoing)

}
