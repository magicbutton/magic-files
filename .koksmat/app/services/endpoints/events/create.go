/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package events

import (
    "log"
   
    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/eventsmodel"

)

func EventsCreate(item eventsmodel.Events) (*eventsmodel.Events, error) {
    log.Println("Calling Eventscreate")

    return applogic.Create[database.Events, eventsmodel.Events](item, applogic.MapEventsIncoming, applogic.MapEventsOutgoing)

}
