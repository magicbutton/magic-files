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

type File struct {
	bun.BaseModel `bun:"table:file,alias:file"`

	ID             int     `bun:"id,pk,autoincrement"`
	CreatedAt      time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	CreatedBy      string `bun:"created_by,"`
	UpdatedAt      time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedBy      string `bun:"updated_by,"`
	DeletedAt      time.Time `bun:",soft_delete,nullzero"`
        Tenant string `bun:"tenant"`
    Name string `bun:"name"`
    Description string `bun:"description"`
    Url string `bun:"url"`
    Searchindex string `bun:"searchindex"`
    Type string `bun:"type"`
    Size int `bun:"size"`
    Extension string `bun:"extension"`
    Createdby string `bun:"createdby"`
    Created time.Time `bun:"created"`
    Modified time.Time `bun:"modified"`
    Modifiedby string `bun:"modifiedby"`
    Version string `bun:"version"`
    Published bool `bun:"published"`
    Level1 string `bun:"level1"`
    Level2 string `bun:"level2"`
    Level3 string `bun:"level3"`
    Level4 string `bun:"level4"`
    Level5 string `bun:"level5"`
    Level6 string `bun:"level6"`
    Level7 string `bun:"level7"`
    Level8 string `bun:"level8"`
    Level9 string `bun:"level9"`
    Level10 string `bun:"level10"`
    Level11 string `bun:"level11"`
    Level12 string `bun:"level12"`
    Level13 string `bun:"level13"`
    Level14 string `bun:"level14"`
    Level15 string `bun:"level15"`
    Level16 string `bun:"level16"`
    Batchname string `bun:"batchname"`

}

