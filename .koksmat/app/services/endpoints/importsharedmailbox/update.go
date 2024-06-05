/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package importsharedmailbox

import (
    "log"

    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/importsharedmailboxmodel"

)

func ImportSharedMailboxUpdate(item importsharedmailboxmodel.ImportSharedMailbox) (*importsharedmailboxmodel.ImportSharedMailbox, error) {
    log.Println("Calling ImportSharedMailboxupdate")

    return applogic.Update[database.ImportSharedMailbox, importsharedmailboxmodel.ImportSharedMailbox](item.ID,item, applogic.MapImportSharedMailboxIncoming, applogic.MapImportSharedMailboxOutgoing)

}
