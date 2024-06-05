/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package container

import (
    "log"
    "strconv"
    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/containermodel"

)

func ContainerRead(arg0 string) (*containermodel.Container, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling Containerread")

    return applogic.Read[database.Container, containermodel.Container](id, applogic.MapContainerOutgoing)

}
