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
	"github.com/magicbutton/magic-files/services/models/relationmodel"
   
)


func MapRelationOutgoing(db database.Relation) relationmodel.Relation {
    return relationmodel.Relation{
        ID:        db.ID,
        CreatedAt: db.CreatedAt,
        UpdatedAt: db.UpdatedAt,
                Tenant : db.Tenant,
        Name : db.Name,
        Description : db.Description,
                User_id : db.User_id,
        Type : db.Type,

    }
}

func MapRelationIncoming(in relationmodel.Relation) database.Relation {
    return database.Relation{
        ID:        in.ID,
        CreatedAt: in.CreatedAt,
        UpdatedAt: in.UpdatedAt,
                Tenant : in.Tenant,
        Name : in.Name,
        Description : in.Description,
                User_id : in.User_id,
        Type : in.Type,

    }
}
