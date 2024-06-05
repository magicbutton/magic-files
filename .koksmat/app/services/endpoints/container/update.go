/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package container

import (
    "log"

    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/containermodel"

)

func ContainerUpdate(item containermodel.Container) (*containermodel.Container, error) {
    log.Println("Calling Containerupdate")

    return applogic.Update[database.Container, containermodel.Container](item.ID,item, applogic.MapContainerIncoming, applogic.MapContainerOutgoing)

}
