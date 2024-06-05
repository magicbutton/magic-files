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
	"github.com/magicbutton/magic-files/services/models/importrelationsmodel"
   
)


func MapImportRelationsOutgoing(db database.ImportRelations) importrelationsmodel.ImportRelations {
    return importrelationsmodel.ImportRelations{
        ID:        db.ID,
        CreatedAt: db.CreatedAt,
        CreatedBy: db.CreatedBy,
        UpdatedAt: db.UpdatedAt,
        UpdatedBy: db.UpdatedBy,
                Tenant : db.Tenant,
        Name : db.Name,
        Description : db.Description,
        Url : db.Url,
                User_id : db.User_id,
        Type : db.Type,
                Importgroup_Id_id : db.Importgroup_Id_id,

    }
}

func MapImportRelationsIncoming(in importrelationsmodel.ImportRelations) database.ImportRelations {
    return database.ImportRelations{
        ID:        in.ID,
        CreatedAt: in.CreatedAt,
        CreatedBy: in.CreatedBy,
        UpdatedAt: in.UpdatedAt,
        UpdatedBy: in.UpdatedBy,
                Tenant : in.Tenant,
        Name : in.Name,
        Description : in.Description,
        Url : in.Url,
                User_id : in.User_id,
        Type : in.Type,
                Importgroup_Id_id : in.Importgroup_Id_id,
        Searchindex : in.Name,

    }
}
