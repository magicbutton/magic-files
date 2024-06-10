/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//GenerateGoModel v2
package jobmodel
import (
	"encoding/json"
	"time"
    // "github.com/magicbutton/magic-files/database/databasetypes"
)

func UnmarshalJob(data []byte) (Job, error) {
	var r Job
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Job) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Job struct {
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
    Prefix string `json:"prefix"`
    Type string `json:"type"`
    Connection_id int `json:"connection_id"`
    Data interface{} `json:"data"`

}

