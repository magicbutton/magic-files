/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
//GenerateMapModel v1.1
package applogic
import (
	//"encoding/json"
	//"time"
	"github.com/magicbutton/magic-files/database"
	"github.com/magicbutton/magic-files/services/models/containermodel"
   
)


func MapContainerOutgoing(db database.Container) containermodel.Container {
    return containermodel.Container{
        ID:        db.ID,
        CreatedAt: db.CreatedAt,
        UpdatedAt: db.UpdatedAt,
                Tenant : db.Tenant,
        Name : db.Name,
        Description : db.Description,
        Url : db.Url,
        Type : db.Type,
        Level1 : db.Level1,
        Level2 : db.Level2,
        Level3 : db.Level3,
        Level4 : db.Level4,
        Level5 : db.Level5,
        Level6 : db.Level6,
        Level7 : db.Level7,
        Level8 : db.Level8,
        Level9 : db.Level9,
        Level10 : db.Level10,
        Level11 : db.Level11,
        Level12 : db.Level12,
        Level13 : db.Level13,
        Level14 : db.Level14,
        Level15 : db.Level15,
        Level16 : db.Level16,

    }
}

func MapContainerIncoming(in containermodel.Container) database.Container {
    return database.Container{
        ID:        in.ID,
        CreatedAt: in.CreatedAt,
        UpdatedAt: in.UpdatedAt,
                Tenant : in.Tenant,
        Name : in.Name,
        Description : in.Description,
        Url : in.Url,
        Type : in.Type,
        Level1 : in.Level1,
        Level2 : in.Level2,
        Level3 : in.Level3,
        Level4 : in.Level4,
        Level5 : in.Level5,
        Level6 : in.Level6,
        Level7 : in.Level7,
        Level8 : in.Level8,
        Level9 : in.Level9,
        Level10 : in.Level10,
        Level11 : in.Level11,
        Level12 : in.Level12,
        Level13 : in.Level13,
        Level14 : in.Level14,
        Level15 : in.Level15,
        Level16 : in.Level16,

    }
}
