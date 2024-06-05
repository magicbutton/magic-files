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
	"github.com/magicbutton/magic-files/services/models/jobmodel"
   
)


func MapJobOutgoing(db database.Job) jobmodel.Job {
    return jobmodel.Job{
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
        Prefix : db.Prefix,
        Type : db.Type,
        Data : db.Data,

    }
}

func MapJobIncoming(in jobmodel.Job) database.Job {
    return database.Job{
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
        Prefix : in.Prefix,
        Type : in.Type,
        Data : in.Data,
        Searchindex : in.Name,

    }
}
