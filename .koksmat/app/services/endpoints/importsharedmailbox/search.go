/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package importsharedmailbox

import (
    "log"

    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/importsharedmailboxmodel"
    . "github.com/magicbutton/magic-files/utils"
)

func ImportSharedMailboxSearch(query string) (*Page[importsharedmailboxmodel.ImportSharedMailbox], error) {
    log.Println("Calling ImportSharedMailboxsearch")

    return applogic.Search[database.ImportSharedMailbox, importsharedmailboxmodel.ImportSharedMailbox]("searchindex", query, applogic.MapImportSharedMailboxOutgoing)

}
