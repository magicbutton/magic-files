/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package events

import (
    "log"

    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/eventsmodel"
    . "github.com/magicbutton/magic-files/utils"
)

func EventsSearch(query string) (*Page[eventsmodel.Events], error) {
    log.Println("Calling Eventssearch")

    return applogic.Search[database.Events, eventsmodel.Events]("searchindex", query, applogic.MapEventsOutgoing)

}
