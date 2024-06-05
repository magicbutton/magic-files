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
	"github.com/magicbutton/magic-files/services/models/importgroupsmodel"
   
)


func MapImportGroupsOutgoing(db database.ImportGroups) importgroupsmodel.ImportGroups {
    return importgroupsmodel.ImportGroups{
        ID:        db.ID,
        CreatedAt: db.CreatedAt,
        CreatedBy: db.CreatedBy,
        UpdatedAt: db.UpdatedAt,
        UpdatedBy: db.UpdatedBy,
                Tenant : db.Tenant,
        Name : db.Name,
        Description : db.Description,
        Url : db.Url,
        Groupid : db.Groupid,
        Type : db.Type,

    }
}

func MapImportGroupsIncoming(in importgroupsmodel.ImportGroups) database.ImportGroups {
    return database.ImportGroups{
        ID:        in.ID,
        CreatedAt: in.CreatedAt,
        CreatedBy: in.CreatedBy,
        UpdatedAt: in.UpdatedAt,
        UpdatedBy: in.UpdatedBy,
                Tenant : in.Tenant,
        Name : in.Name,
        Description : in.Description,
        Url : in.Url,
        Groupid : in.Groupid,
        Type : in.Type,
        Searchindex : in.Name,

    }
}
