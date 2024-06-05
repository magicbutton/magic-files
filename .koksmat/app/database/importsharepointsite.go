/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//version: pølsevogn2
package database

import (
	"time"
    
	"github.com/uptrace/bun"
)

type ImportSharePointSite struct {
	bun.BaseModel `bun:"table:importsharepointsite,alias:importsharepointsite"`

	ID             int     `bun:"id,pk,autoincrement"`
	CreatedAt      time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	CreatedBy      string `bun:"created_by,"`
	UpdatedAt      time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedBy      string `bun:"updated_by,"`
	DeletedAt      time.Time `bun:",soft_delete,nullzero"`
        Tenant string `bun:"tenant"`
    Searchindex string `bun:"searchindex"`
    Name string `bun:"name"`
    Description string `bun:"description"`
    Isteamsconnected bool `bun:"isteamsconnected"`
    Localeid string `bun:"localeid"`
    Hubsiteid string `bun:"hubsiteid"`
    Status string `bun:"status"`
    Url string `bun:"url"`
    Title string `bun:"title"`
    Lastcontentmodified time.Time `bun:"lastcontentmodified"`
    Releatedgroupid string `bun:"releatedgroupid"`
    Groupid string `bun:"groupid"`
    Isteamschannelconnected bool `bun:"isteamschannelconnected"`
    Ishubsite bool `bun:"ishubsite"`
    Sharingcapability string `bun:"sharingcapability"`
    Owner string `bun:"owner"`

}

