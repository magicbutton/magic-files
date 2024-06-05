/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package container

import (
    "log"
   
    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/containermodel"

)

func ContainerCreate(item containermodel.Container) (*containermodel.Container, error) {
    log.Println("Calling Containercreate")

    return applogic.Create[database.Container, containermodel.Container](item, applogic.MapContainerIncoming, applogic.MapContainerOutgoing)

}
