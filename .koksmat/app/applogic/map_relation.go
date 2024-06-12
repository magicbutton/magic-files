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
        CreatedBy: db.CreatedBy,
        UpdatedAt: db.UpdatedAt,
        UpdatedBy: db.UpdatedBy,
                Name : db.Name,
        Description : db.Description,
                User_id : db.User_id,
        Type : db.Type,
        Batchname : db.Batchname,

    }
}

func MapRelationIncoming(in relationmodel.Relation) database.Relation {
    return database.Relation{
        ID:        in.ID,
        CreatedAt: in.CreatedAt,
        CreatedBy: in.CreatedBy,
        UpdatedAt: in.UpdatedAt,
        UpdatedBy: in.UpdatedBy,
                Name : in.Name,
        Description : in.Description,
                User_id : in.User_id,
        Type : in.Type,
        Batchname : in.Batchname,
        Searchindex : in.Name,

    }
}
