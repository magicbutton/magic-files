/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package user

import (
    "log"
   
    "github.com/magicbutton/magic-files/applogic"
    "github.com/magicbutton/magic-files/database"
    "github.com/magicbutton/magic-files/services/models/usermodel"

)

func UserCreate(item usermodel.User) (*usermodel.User, error) {
    log.Println("Calling Usercreate")

    return applogic.Create[database.User, usermodel.User](item, applogic.MapUserIncoming, applogic.MapUserOutgoing)

}
