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
	"github.com/magicbutton/magic-files/services/models/importsharepointsitemodel"
   
)


func MapImportSharePointSiteOutgoing(db database.ImportSharePointSite) importsharepointsitemodel.ImportSharePointSite {
    return importsharepointsitemodel.ImportSharePointSite{
        ID:        db.ID,
        CreatedAt: db.CreatedAt,
        CreatedBy: db.CreatedBy,
        UpdatedAt: db.UpdatedAt,
        UpdatedBy: db.UpdatedBy,
                Name : db.Name,
        Description : db.Description,
        Localeid : db.Localeid,
        Hubsiteid : db.Hubsiteid,
        Status : db.Status,
        Url : db.Url,
        Title : db.Title,
        Releatedgroupid : db.Releatedgroupid,
        Groupid : db.Groupid,
        Sharingcapability : db.Sharingcapability,
        Owner : db.Owner,

    }
}

func MapImportSharePointSiteIncoming(in importsharepointsitemodel.ImportSharePointSite) database.ImportSharePointSite {
    return database.ImportSharePointSite{
        ID:        in.ID,
        CreatedAt: in.CreatedAt,
        CreatedBy: in.CreatedBy,
        UpdatedAt: in.UpdatedAt,
        UpdatedBy: in.UpdatedBy,
                Name : in.Name,
        Description : in.Description,
        Localeid : in.Localeid,
        Hubsiteid : in.Hubsiteid,
        Status : in.Status,
        Url : in.Url,
        Title : in.Title,
        Releatedgroupid : in.Releatedgroupid,
        Groupid : in.Groupid,
        Sharingcapability : in.Sharingcapability,
        Owner : in.Owner,
        Searchindex : in.Name,

    }
}
