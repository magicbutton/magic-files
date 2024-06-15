/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package sql

import (
    "log"

    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/sqlmodel"
    . "github.com/magicbutton/magic-files/utils"
)

func SQLSearch(query string) (*Page[sqlmodel.SQL], error) {
    log.Println("Calling SQLsearch")

    return applogic.Search[database.SQL, sqlmodel.SQL]("searchindex", query, applogic.MapSQLOutgoing)

}
