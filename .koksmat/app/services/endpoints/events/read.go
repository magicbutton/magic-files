/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package events

import (
    "log"
    "strconv"
    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/eventsmodel"

)

func EventsRead(arg0 string) (*eventsmodel.Events, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling Eventsread")

    return applogic.Read[database.Events, eventsmodel.Events](id, applogic.MapEventsOutgoing)

}
