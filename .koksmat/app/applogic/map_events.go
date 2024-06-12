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
	"github.com/magicbutton/magic-files/services/models/eventsmodel"
   
)


func MapEventsOutgoing(db database.Events) eventsmodel.Events {
    return eventsmodel.Events{
        ID:        db.ID,
        CreatedAt: db.CreatedAt,
        CreatedBy: db.CreatedBy,
        UpdatedAt: db.UpdatedAt,
        UpdatedBy: db.UpdatedBy,
                Tenant : db.Tenant,
        Name : db.Name,
        Description : db.Description,
        Url : db.Url,
        Type : db.Type,
                User_id : db.User_id,
        Operation : db.Operation,
        Batchname : db.Batchname,
                File_id : db.File_id,

    }
}

func MapEventsIncoming(in eventsmodel.Events) database.Events {
    return database.Events{
        ID:        in.ID,
        CreatedAt: in.CreatedAt,
        CreatedBy: in.CreatedBy,
        UpdatedAt: in.UpdatedAt,
        UpdatedBy: in.UpdatedBy,
                Tenant : in.Tenant,
        Name : in.Name,
        Description : in.Description,
        Url : in.Url,
        Type : in.Type,
                User_id : in.User_id,
        Operation : in.Operation,
        Batchname : in.Batchname,
                File_id : in.File_id,
        Searchindex : in.Name,

    }
}
