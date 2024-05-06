/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package container

import (
    "log"

    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/containermodel"
    . "github.com/magicbutton/magic-files/utils"
)

func ContainerSearch(query string) (*Page[containermodel.Container], error) {
    log.Println("Calling Containersearch")

    return applogic.Search[database.Container, containermodel.Container]("name", query, applogic.MapContainerOutgoing)

}
