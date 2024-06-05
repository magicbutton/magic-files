/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package job

import (
    "log"

    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/jobmodel"
    . "github.com/magicbutton/magic-files/utils"
)

func JobSearch(query string) (*Page[jobmodel.Job], error) {
    log.Println("Calling Jobsearch")

    return applogic.Search[database.Job, jobmodel.Job]("searchindex", query, applogic.MapJobOutgoing)

}
