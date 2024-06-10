/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//GenerateGoModel v2
package batchmodel
import (
	"encoding/json"
	"time"
    // "github.com/magicbutton/magic-files/database/databasetypes"
)

func UnmarshalBatch(data []byte) (Batch, error) {
	var r Batch
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Batch) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Batch struct {
    ID        int    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    CreatedBy string `json:"created_by"`
    UpdatedAt time.Time `json:"updated_at"`
    UpdatedBy string `json:"updated_by"`
        Tenant string `json:"tenant"`
    Name string `json:"name"`
    Description string `json:"description"`
    Url string `json:"url"`
    Status string `json:"status"`
    Scheduledstart time.Time `json:"scheduledstart"`
    Type string `json:"type"`
    Job_id int `json:"job_id"`

}

