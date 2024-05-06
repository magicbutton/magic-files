/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//GenerateGoModel v1
package importmodel
import (
	"encoding/json"
	"time"
    // "github.com/magicbutton/magic-files/database/databasetypes"
)

func UnmarshalImport(data []byte) (Import, error) {
	var r Import
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Import) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Import struct {
    ID        int    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
        Tenant string `json:"tenant"`
    Name string `json:"name"`
    Description string `json:"description"`
    User_id int `json:"user_id"`
    Data interface{} `json:"data"`

}

