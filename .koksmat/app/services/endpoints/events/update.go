/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package events

import (
    "log"

    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/eventsmodel"

)

func EventsUpdate(item eventsmodel.Events) (*eventsmodel.Events, error) {
    log.Println("Calling Eventsupdate")

    return applogic.Update[database.Events, eventsmodel.Events](item.ID,item, applogic.MapEventsIncoming, applogic.MapEventsOutgoing)

}
