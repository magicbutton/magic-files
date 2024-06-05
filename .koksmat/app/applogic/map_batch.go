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
	"github.com/magicbutton/magic-files/services/models/batchmodel"
   
)


func MapBatchOutgoing(db database.Batch) batchmodel.Batch {
    return batchmodel.Batch{
        ID:        db.ID,
        CreatedAt: db.CreatedAt,
        CreatedBy: db.CreatedBy,
        UpdatedAt: db.UpdatedAt,
        UpdatedBy: db.UpdatedBy,
                Tenant : db.Tenant,
        Name : db.Name,
        Description : db.Description,
        Url : db.Url,
        Status : db.Status,
        Type : db.Type,
                Job_id : db.Job_id,

    }
}

func MapBatchIncoming(in batchmodel.Batch) database.Batch {
    return database.Batch{
        ID:        in.ID,
        CreatedAt: in.CreatedAt,
        CreatedBy: in.CreatedBy,
        UpdatedAt: in.UpdatedAt,
        UpdatedBy: in.UpdatedBy,
                Tenant : in.Tenant,
        Name : in.Name,
        Description : in.Description,
        Url : in.Url,
        Status : in.Status,
        Type : in.Type,
                Job_id : in.Job_id,
        Searchindex : in.Name,

    }
}
