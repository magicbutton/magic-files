/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package importsharedmailbox

import (
    "log"
   
    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/importsharedmailboxmodel"

)

func ImportSharedMailboxCreate(item importsharedmailboxmodel.ImportSharedMailbox) (*importsharedmailboxmodel.ImportSharedMailbox, error) {
    log.Println("Calling ImportSharedMailboxcreate")

    return applogic.Create[database.ImportSharedMailbox, importsharedmailboxmodel.ImportSharedMailbox](item, applogic.MapImportSharedMailboxIncoming, applogic.MapImportSharedMailboxOutgoing)

}
