/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//GenerateGoModel v1
package importgroupsmodel
import (
	"encoding/json"
	"time"
    // 
)

func UnmarshalImportGroups(data []byte) (ImportGroups, error) {
	var r ImportGroups
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ImportGroups) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ImportGroups struct {
    ID        int    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    CreatedBy string `json:"created_by"`
    UpdatedAt time.Time `json:"updated_at"`
    UpdatedBy string `json:"updated_by"`
        Tenant string `json:"tenant"`
    Name string `json:"name"`
    Description string `json:"description"`
    Url string `json:"url"`
    Groupid string `json:"groupid"`
    Type string `json:"type"`

}

